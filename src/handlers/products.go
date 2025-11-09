package handlers

import (
	"src/config"
	"src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ProductsPage(c *fiber.Ctx) error {
	var products []models.Products
	config.DB.Find(&products)

	return c.Render("products", fiber.Map{
		"Title":  "Kelola Produk",
		"Produk": products,
	})
}
func ProductCreate(c *fiber.Ctx) error {
	product := new(models.Products)
	// if err := c.BodyParser(product); err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	// }

	product.NamaProduct = c.FormValue("nama_product")

	// parse harga (string -> uint)
	hargaStr := c.FormValue("harga")
	if hargaStr != "" {
		harga64, err := strconv.ParseUint(hargaStr, 10, 64)
		if err != nil {
			return c.Status(400).SendString("invalid harga")
		}
		product.Harga = uint(harga64)
	}

	// parse stok (string -> uint)
	stokStr := c.FormValue("stok")
	if stokStr != "" {
		stok64, err := strconv.ParseUint(stokStr, 10, 64)
		if err != nil {
			return c.Status(400).SendString("invalid stok")
		}
		product.Stok = uint(stok64)
	}

	result := config.DB.Create(product)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Redirect("/products")
}

func ProductUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Products

	if err := config.DB.First(&product, "id_product = ?", id).Error; err != nil {
		return c.Status(404).SendString("Produk tidak ditemukan")
	}

	product.NamaProduct = c.FormValue("nama_product")

	// parse harga (string -> uint)
	hargaStr := c.FormValue("harga")
	if hargaStr != "" {
		harga64, err := strconv.ParseUint(hargaStr, 10, 64)
		if err != nil {
			return c.Status(400).SendString("invalid harga")
		}
		product.Harga = uint(harga64)
	}

	// parse stok (string -> uint)
	stokStr := c.FormValue("stok")
	if stokStr != "" {
		stok64, err := strconv.ParseUint(stokStr, 10, 64)
		if err != nil {
			return c.Status(400).SendString("invalid stok")
		}
		product.Stok = uint(stok64)
	}

	if err := config.DB.Save(&product).Error; err != nil {
		return c.Status(500).SendString("Gagal update: " + err.Error())
	}
	return c.Redirect("/products")
}

func ProductDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Products{}, "id_product = ?", id)

	return c.Redirect("/products")
}
