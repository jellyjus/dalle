package api

import (
	"dalle/extapi/openai"
	"dalle/repository"

	"github.com/gofiber/fiber/v3"
)

type router struct {
	openai *openai.Service
	repo   *repository.Repository
}

func ApplyRoutes(app *fiber.App, oai *openai.Service, repo *repository.Repository) {
	r := router{
		openai: oai,
		repo:   repo,
	}

	api := app.Group("/api")
	image := api.Group("/image")
	image.Post("/generate", r.generateImage)
}

func (r *router) generateImage(c fiber.Ctx) error {
	var req generateImageRequest
	if err := c.Bind().JSON(&req); err != nil {
		return err
	}
	if req.Prompt == "" {
		return fiber.ErrBadRequest
	}

	image, err := r.openai.GenerateImage(c.Context(), req.Prompt)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := r.repo.AddImage(c.Context(), image); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(req)
}
