package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"dalle/entity"
)

const maxConnection = 10

type Service struct {
	client *http.Client
	apiKey string
}

func New(apiKey string) *Service {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = maxConnection
	t.MaxConnsPerHost = maxConnection
	t.MaxIdleConnsPerHost = maxConnection
	return &Service{
		client: &http.Client{
			Transport: t,
		},
		apiKey: apiKey,
	}
}

func (s *Service) GenerateImage(ctx context.Context, prompt string) (entity.Image, error) {
	req := imageRequest{
		Model:  "dall-e-3",
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
	}
	var res imageResponse
	if err := s.doRequest(ctx, imagePath, req, &res); err != nil {
		return entity.Image{}, err
	}
	if len(res.Data) == 0 {
		return entity.Image{}, fmt.Errorf("no image data")
	}
	return res.Data[0], nil
}

func (s *Service) doRequest(ctx context.Context, path string, apiRequest, apiResponse any) error {
	request, err := json.Marshal(apiRequest)
	if err != nil {
		return err
	}

	req, err := s.buildRequest(ctx, path, request)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("server error, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return json.Unmarshal(body, apiResponse)
}

func (s *Service) buildRequest(ctx context.Context, path string, request []byte) (*http.Request, error) {
	u, err := url.JoinPath(host, path)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
