package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	login "github.com/wiratkhamphan/Webshop.git/login"
	register "github.com/wiratkhamphan/Webshop.git/login/register"
)

func main() {

	// สร้าง Fiber app
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Set("Access-Control-Allow-Credentials", "true")

		if c.Method() == "OPTIONS" {
			c.SendStatus(fiber.StatusNoContent)
			return nil
		}

		return c.Next()
	})

	// เส้นทาง API สำหรับการเข้าสู่ระบบ
	app.Post("/login", login.Login)
	app.Post("/register", register.Register)

	// เริ่มต้นเซิร์ฟเวอร์
	log.Fatal(app.Listen(":8080"))
}
