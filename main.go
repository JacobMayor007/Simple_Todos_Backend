package main

import (
	"log"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todos struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"createdAt"`
}

var (
	TodosItems = []Todos{}
	mutex      = &sync.Mutex{}
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Accept, Content-Type, Origins",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	app.Get("/todos", AllTodos)
	app.Post("/todos", CreateTodos)
	app.Patch("/todos", UpdateTodos)
	app.Delete("/todos", DeleteTodos)

	log.Fatal(app.Listen(":3201"))

}

func AllTodos(c *fiber.Ctx) error {
	return c.JSON(TodosItems)
}

func CreateTodos(c *fiber.Ctx) error {
	var data Todos
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	mutex.Lock()
	data.Id = len(TodosItems) + 1
	TodosItems = append(TodosItems, data)
	mutex.Unlock()

	return c.JSON(data)
}

func UpdateTodos(c *fiber.Ctx) error {
	found := false

	var data Todos

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i := 0; i < len(TodosItems); i++ {
		if TodosItems[i].Id == data.Id {
			TodosItems[i].Title = data.Title
			TodosItems[i].Completed = data.Completed
			found = true
			break
		}
	}

	if !found {
		return fiber.ErrNotFound // 404 if todo not found
	}
	method := c.Method()
	log.Printf("%s", method)

	return c.JSON(data)
}
func DeleteTodos(c *fiber.Ctx) error {
	var data Todos
	if err := c.BodyParser(&data); err != nil {
		return fiber.ErrBadRequest
	}

	mutex.Lock()
	defer mutex.Unlock()

	found := false
	for i := 0; i < len(TodosItems); i++ {
		if TodosItems[i].Id == data.Id {

			TodosItems = append(TodosItems[:i], TodosItems[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fiber.ErrNotFound
	}

	log.Printf("%s", c.Method())
	return c.JSON(fiber.Map{
		"message": "Todo deleted successfully",
		"id":      data.Id,
	})
}
