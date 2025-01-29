package shared

import (
	"net/http"

	"hiyoko-fiber/pkg/logging/file"

	"github.com/gofiber/fiber/v2"
)

func ResponseOK(c *fiber.Ctx, data interface{}) error {
	if err := c.SendStatus(http.StatusOK); err != nil {
		log.Error("OK response Set status failed", "error", err, "data", data)
		return err
	}
	if err := c.JSON(data); err != nil {
		log.Error("OK response JSON conversion failed", "error", err, "data", data)
		return err
	}
	return nil
}

func ResponseCreate(c *fiber.Ctx, data interface{}) error {
	if err := c.SendStatus(http.StatusCreated); err != nil {
		log.Error("Created response JSON conversion failed", "error", err, "data", data)
		return err
	}
	if err := c.JSON(data); err != nil {
		log.Error("Created response JSON conversion failed", "error", err, "data", data)
		return err
	}
	return nil
}

func ResponseNoContent(c *fiber.Ctx) error {
	if err := c.SendStatus(http.StatusNoContent); err != nil {
		log.Error("NoContent response JSON conversion failed", "error", err)
		return err
	}
	return nil
}

func ResponseBadRequest(c *fiber.Ctx, code string) error {
	if err := c.SendStatus(http.StatusBadRequest); err != nil {
		log.Error("BadRequest response Set status failed", "error", err, "code", code)
		return err
	}
	// Bad request always returns none code for security reasons
	if err := c.JSON(ErrorResponse{
		Code:    NoneCode,
		Message: GetErrorMessage(http.StatusBadRequest),
	}); err != nil {
		log.Error("BadRequest response JSON conversion failed", "error", err, "code", code)
		return err
	}
	return nil
}

func ResponseUnauthorized(c *fiber.Ctx) error {
	if err := c.SendStatus(http.StatusUnauthorized); err != nil {
		log.Error("Unauthorized response Set status failed", "error", err)
		return err
	}
	if err := c.JSON(ErrorResponse{
		Code:    NoneCode,
		Message: GetErrorMessage(http.StatusUnauthorized),
	}); err != nil {
		log.Error("Unauthorized response JSON conversion failed", "error", err)
		return err
	}
	return nil
}

func ResponseForbidden(c *fiber.Ctx, code string) error {
	if err := c.SendStatus(http.StatusForbidden); err != nil {
		log.Error("Forbidden response Set status failed", "error", err)
		return err
	}
	if err := c.JSON(ErrorResponse{
		Code:    code,
		Message: GetErrorMessage(http.StatusForbidden),
	}); err != nil {
		log.Error("Forbidden response JSON conversion failed", "error", err)
		return err
	}
	return nil
}

func ResponseNotFound(c *fiber.Ctx, code string) error {
	if err := c.SendStatus(http.StatusNotFound); err != nil {
		log.Error("NotFound response Set status failed", "error", err, "code", code)
		return err
	}
	if err := c.JSON(ErrorResponse{
		Code:    code,
		Message: GetErrorMessage(http.StatusNotFound),
	}); err != nil {
		log.Error("NotFound response JSON conversion failed", "error", err, "code", code)
		return err
	}
	return nil
}

func ResponseRequestTimeout(c *fiber.Ctx, code string) error {
	if err := c.SendStatus(http.StatusRequestTimeout); err != nil {
		log.Error("RequestTimeout response Set status failed", "error", err, "code", code)
		return err
	}
	if err := c.JSON(ErrorResponse{
		Code:    code,
		Message: GetErrorMessage(http.StatusRequestTimeout),
	}); err != nil {
		log.Error("RequestTimeout response JSON conversion failed", "error", err, "code", code)
		return err
	}
	return nil
}
