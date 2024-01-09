package api

import (
	"dalle/api/templates"
	"dalle/extapi/openai"
	"dalle/repository"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/filesystem"
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

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Use("/admin", filesystem.New(filesystem.Config{
		Root: templates.FS,
	}))

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

	image, err := r.openai.GenerateImage(c.Context(), req.Prompt, req.HD)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := r.repo.AddImage(c.Context(), image); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(image)
}
