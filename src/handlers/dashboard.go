package handlers

import (
	"src/config"
	"src/models"

	"github.com/gofiber/fiber/v2"
)

func DashboardPage(c *fiber.Ctx) error {
	var jmlsUsers int64
	var jmlProducts int64

	config.DB.Model(&models.Users{}).Count(&jmlsUsers)
	config.DB.Model(&models.Products{}).Count(&jmlProducts)
	return c.Render("dashboard", fiber.Map{
		"Title":       "Dashboard",
		"JmlUsers":    jmlsUsers,
		"JmlProducts": jmlProducts,
	})
}
