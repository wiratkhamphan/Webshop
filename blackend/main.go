package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// User struct สำหรับเก็บข้อมูลผู้ใช้
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "shoplek"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ทดสอบการเชื่อมต่อ
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL database")

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
	app.Post("/login", func(c *fiber.Ctx) error {
		// รับข้อมูลการเข้าสู่ระบบจากผู้ใช้
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&loginData); err != nil {
			return err
		}

		// ค้นหาผู้ใช้ในฐานข้อมูล
		var user User
		err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", loginData.Username).Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
			}
			return err
		}

		// // เปรียบเทียบรหัสผ่าน
		// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
		// if err != nil {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
		// }

		// ส่งข้อความอนุญาต
		return c.SendStatus(fiber.StatusOK)
	})

	// เริ่มต้นเซิร์ฟเวอร์
	log.Fatal(app.Listen(":8080"))
}
