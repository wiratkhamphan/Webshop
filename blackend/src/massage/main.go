package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type PostID int

type Post struct {
	ID    PostID `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Grade1 struct {
	ID    PostID `json:"id"`
	Grade int    `json:"grade"`
}

// Function to calculate the grade letter based on the score
func Grade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func PostGrade(c *fiber.Ctx) error {

	var grade Grade1
	if err := c.BodyParser(&grade); err != nil {
		return err
	}

	gradeLetter := Grade(grade.Grade)

	response := struct {
		ID    PostID `json:"id"`
		Grade string `json:"grade"`
	}{
		ID:    grade.ID,
		Grade: gradeLetter,
	}

	return c.JSON(response)
}

func PostGradeTest(c *fiber.Ctx) error {
	var TestGrade Grade1

	if err := c.BodyParser(&TestGrade); err != nil {
		return err

	}
	Test1Grad := Grade(TestGrade.Grade)

	resJsonTestGrade := struct {
		ID    PostID `json:"id"`
		Grade string `json:"grade"`
	}{
		ID:    TestGrade.ID,
		Grade: Test1Grad,
	}

	return c.JSON(resJsonTestGrade)
}

func main() {
	app := fiber.New()

	app.Get("/Get", Getdata)
	app.Get("/Get/:ID", GetPostByID)
	app.Post("/Post1", CreatePost)
	app.Post("Grade", PostGrade)
	app.Post("/Grade/test", PostGradeTest)

	app.Listen(":8080")
}

func Conmysql() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "shoplek"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		return nil, err
	}

	// Ensure the database connection is valid
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Getdata(c *fiber.Ctx) error {
	// Assuming Conmysql() returns *sql.DB and error
	db, err := Conmysql()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close() // Close the database connection when done

	var Posts []Post
	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var Port Post
		err := rows.Scan(&Port.ID, &Port.Title, &Port.Body) // Assuming Post has fields ID, Title, and Content
		if err != nil {
			log.Fatal(err)
			return err
		}
		Posts = append(Posts, Port)
	}

	// Do something with Posts if needed

	return c.JSON(Posts)
}
func GetPostByID(c *fiber.Ctx) error {
	// Retrieve the post ID from the URL parameters
	postID := c.Params("ID")

	// Assuming Conmysql() returns *sql.DB and error
	db, err := Conmysql()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer db.Close()

	// Query the database to get the post with the specified ID
	var post Post
	err = db.QueryRow("SELECT * FROM post WHERE ID = ?", postID).Scan(&post.ID, &post.Title, &post.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no post is found with the given ID, return a 404 Not Found response
			return c.Status(fiber.StatusNotFound).SendString("Post not found")
		}
		// For other errors, return a 500 Internal Server Error response
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return the retrieved post as JSON
	return c.JSON(post)
}

func CreatePost(c *fiber.Ctx) error {
	// Parse the request body to extract post data
	var newPost Post
	if err := c.BodyParser(&newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Assuming Conmysql() returns *sql.DB and error
	db, err := Conmysql()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer db.Close()

	// Insert the new post into the database
	_, err = db.Exec("INSERT INTO post (Title, Body) VALUES (?, ?)", newPost.Title, newPost.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return a 201 Created response with the created post
	return c.SendString("update : mysql post")
}
