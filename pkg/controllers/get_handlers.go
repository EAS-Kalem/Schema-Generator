package controllers

import (
	// "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/models"
)

// var arr2 string
func Handlers(Microservices string) models.Handlers {
	var Handlers models.Handlers
	token := "REQUIRED"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+"Enterprise-Automation"+"/"+Microservices+"/contents/handlers", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	handlersReq, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}

	defer handlersReq.Body.Close()
	body, _ := io.ReadAll(handlersReq.Body)

	if err := json.Unmarshal(body, &Handlers); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return Handlers
}
