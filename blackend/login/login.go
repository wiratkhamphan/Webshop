package login

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2" // Ensure you're using the correct import path for Fiber v2
	condb "github.com/wiratkhamphan/Webshop.git/condb"
)

// User struct for storing user information
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse login data"})
	}

	db, err := condb.NewDB()
	if err != nil {
		log.Println(err) // Log the error instead of using log.Fatal
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}
	defer db.Close() // Ensure you close the database connection

	var user User
	err = db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", loginData.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Query error"})
	}

	// Assuming you add password hashing verification here
	// // เปรียบเทียบรหัสผ่าน
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	// }

	return c.SendStatus(fiber.StatusOK) // No need to return an error after this
}
