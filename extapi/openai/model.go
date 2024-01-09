package openai

import "dalle/entity"

const host = "https://api.openai.com/v1"

const imagePath = "/images/generations"

type imageRequest struct {
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	N       int    `json:"n"`
	Size    string `json:"size"`
	Quality string `json:"quality,omitempty"`
}

type imageResponse struct {
	Data []entity.Image `json:"data"`
}
