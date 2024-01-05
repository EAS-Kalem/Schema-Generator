package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/models"
	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/controllers"
)
var fileName = "motorbikes.json"
var app = fiber.New()
func main() {
    
	
	newFileName := controllers.GetSchema(fileName)
	controllers.VerifySyntax(newFileName)
	// controllers.FileContents(Service, Action)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Thank god it works 🙏")
    })

    app.Post("/", controllers.ValidateSchema, func(c *fiber.Ctx) error {
        body := new(models.Schema)
        c.BodyParser(&body)
        return c.Status(fiber.StatusOK).JSON(body)
    })

    app.Listen(":3000")
}



var Validator = validator.New()



  
