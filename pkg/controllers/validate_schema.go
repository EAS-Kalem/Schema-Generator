package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/models"
	// "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    // "strconv"
	
	

)

func VerifySyntax(newFileName string) {
	jsonFile, err := os.Open(newFileName)
    if err != nil {
        fmt.Println(err)
    } else {
		fmt.Println("Successfully Opened "+newFileName)
	}
	
   
    defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println("Edit for full os.Open resp -", string(byteValue[3]))
}

func ValidateSchema(c *fiber.Ctx) error {
	var Validator = validator.New()
    var errors []*models.IError
    body := new(models.Schema)
    c.BodyParser(&body)

    err := Validator.Struct(body)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var el models.IError
            el.Field = err.Field()
            el.Tag = err.Tag()
            el.Value = err.Param()
			
            errors = append(errors, &el)
        }
        return c.Status(fiber.StatusBadRequest).JSON(errors)
    }
    return c.Next()
}