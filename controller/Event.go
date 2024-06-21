package controller

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	var events []models.Event

	data := database.DB.Find(&events)
	if data.Error != nil {
		panic(data.Error)
	}

	result := models.Result{
		Status:  200,
		Message: "OK",
		Data:    events,
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func Create(ctx *fiber.Ctx) error {
	var event models.Event

	if err := ctx.BodyParser(&event); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Result{
			Status:  400,
			Message: "Error Bad Request",
			Data:    []models.Event{},
		})
	}

	if err := database.DB.Create(&event).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Result{
			Status:  500,
			Message: "Error Internal Server",
			Data:    []models.Event{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.Result{
		Status:  200,
		Message: "OK",
		Data:    []models.Event{event},
	})

}

func Search(ctx *fiber.Ctx) error {
	title := ctx.Params("title")

	var events []models.Event
	if err := database.DB.Where("LOWER(title) LIKE ?", "%"+title+"%").Find(&events).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Result{
			Status:  404,
			Message: "Event Not Found",
			Data:    []models.Event{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.Result{
		Status:  200,
		Message: "OK",
		Data:    events,
	})
}

func Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var event models.Event

	if err := ctx.BodyParser(&event); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Result{
			Status:  400,
			Message: "Error Bad Request",
			Data:    []models.Event{},
		})
	}

	if row := database.DB.Where("id = ? ", id).Updates(&event).RowsAffected; row == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.Result{
			Status:  500,
			Message: "Error Bad Request",
			Data:    []models.Event{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.Result{
		Status:  200,
		Message: "OK",
		Data:    []models.Event{event},
	})
}

func Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := database.DB.Delete(&models.Event{}, id).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Result{
			Status:  400,
			Message: "Error Bad Request",
			Data:    []models.Event{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.Result{
		Status:  200,
		Message: "OK",
		Data:    []models.Event{},
	})
}
