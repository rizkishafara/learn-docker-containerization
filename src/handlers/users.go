package handlers

import (
	"fmt"
	"src/config"
	"src/models"

	"github.com/gofiber/fiber/v2"
)

func UsersPage(c *fiber.Ctx) error {
	var users []models.Users
	config.DB.Find(&users)

	return c.Render("users", fiber.Map{
		"Title": "Kelola User",
		"Users": users,
	})
}

func UserCreate(c *fiber.Ctx) error {
	user := new(models.Users)
	
	user.NamaUser = c.FormValue("nama_user")
	user.Email = c.FormValue("email")

	fmt.Println("nama: ", user.NamaUser)

	result := config.DB.Create(user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Redirect("/users")
}
func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.Users

	if err := config.DB.First(&user, "id_user = ?", id).Error; err != nil {
		return c.Status(404).SendString("User tidak ditemukan")
	}

	user.NamaUser = c.FormValue("nama_user")
	user.Email = c.FormValue("email")

	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(500).SendString("Gagal update: " + err.Error())
	}
	
	return c.Redirect("/users")
}

func UserDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Users{}, "id_user = ?", id)
	return c.Redirect("/users")
}
