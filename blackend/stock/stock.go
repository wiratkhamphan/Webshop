package stock

import (
	"log"

	"github.com/gofiber/fiber/v2"
	condb "github.com/wiratkhamphan/Webshop.git/condb"
)

type Box struct {
	ID     int    `json:"id"`
	BoxID  string `json:"box_id"`
	Status string `json:"status"`
	Time   string `json:"time"`
}

func Stock(c *fiber.Ctx) error {
	// Establish a connection to the database
	db, err := condb.NewDB()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}

	// Perform a query to fetch stock information
	rows, err := db.Query("SELECT * FROM boxkk")
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error querying the database"})
	}
	defer rows.Close()

	// Iterate over the rows and scan data into Box structs
	var boxes []Box
	for rows.Next() {
		var box Box
		err := rows.Scan(&box.ID, &box.BoxID, &box.Status, &box.Time)
		if err != nil {
			log.Println(err)
			continue // Skip this row and continue with the next one
		}
		boxes = append(boxes, box)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error iterating over query results"})
	}

	// Return the retrieved data as JSON
	return c.JSON(boxes)
}
