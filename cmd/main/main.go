package main

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/controllers"
	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/models"
)

func main() {
	Microservices1 := controllers.Microservices()
	var structInstances []models.RespR // Slice 
	for _, microsvc := range Microservices1 {
		Handlers1 := controllers.Handlers(microsvc.Name)
		for _, hndlr := range Handlers1 {
			re := regexp.MustCompile(`_`)
			resultString := re.ReplaceAllString(hndlr.Name, " ")
			camelCase := UpperCase(resultString)
			filter := strings.ReplaceAll(camelCase, " ", "")
			filter1 := strings.ReplaceAll(filter, ".go", "")
			filter2 := strings.ReplaceAll(filter1, ".Go", "")
			Args1 := controllers.Args(microsvc.Name, hndlr.Name)
			stringSlice := strings.Split(Args1, ",")

			myStructInstance := models.RespR{
				Service: microsvc.Name,
				Handlers: models.HandlersR{
					FileName:    hndlr.Name,
					ServiceName: filter2,
					Args:        stringSlice,
				},
			}
			structInstances = append(structInstances, myStructInstance)
		}
	}
	fmt.Println(structInstances)
}

func UpperCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
}

