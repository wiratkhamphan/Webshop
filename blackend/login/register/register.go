package register

import (
	"log"

	"github.com/gofiber/fiber/v2"
	condb "github.com/wiratkhamphan/Webshop.git/condb"
)

type Ruser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Register(c *fiber.Ctx) error {
	var userData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	db, err := condb.NewDB()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}
	defer db.Close()

	// Insert user data into the database
	_, err = db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
		userData.Username, userData.Password, userData.Email)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.SendStatus(fiber.StatusCreated)
}
