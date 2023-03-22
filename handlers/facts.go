package handlers

import (
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	// Get the fact ID from the request parameters
	factID := c.Params("id")

	// Check if the fact exists in the database
	var fact models.Fact
	if err := database.DB.Db.First(&fact, factID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fact not found",
		})
	}

	// Parse the request body into a new Fact object
	newFact := new(models.Fact)
	if err := c.BodyParser(newFact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Update the fact in the database
	fact.Question = newFact.Question
	fact.Answer = newFact.Answer
	database.DB.Db.Save(&fact)

	return c.Status(fiber.StatusOK).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	// Get the fact ID from the request parameters
	factID := c.Params("id")

	// Check if the fact exists in the database
	var fact models.Fact
	if err := database.DB.Db.First(&fact, factID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fact not found",
		})
	}

	// Delete the fact from the database
	database.DB.Db.Delete(&fact)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fact deleted successfully",
	})
}
