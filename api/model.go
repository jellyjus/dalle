package api

type generateImageRequest struct {
	HD     bool   `json:"hd"`
	Prompt string `json:"prompt"`
}
