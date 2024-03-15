package restful

import "github.com/gofiber/fiber/v2"

var (
	e *fiber.App
)

func setup() {
	e = NewRouter()
}
