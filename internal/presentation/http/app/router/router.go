package router

import (
	"hiyoko-fiber/internal/pkg/auth/v1"
	"hiyoko-fiber/internal/presentation/http/app/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(f *fiber.App, h handler.AppHandler) {
	v1 := f.Group("/v1")

	v1.Post("/signup", h.Signup)
	v1.Post("/signin", h.Signin)

	v1Guard := v1.Use(auth.Auth())
	v1Guard.Get("/users/me", h.GetMe)
}
