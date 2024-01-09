package controllers

import (
	// "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	// "strings"
	"regexp"

	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/models"
)

var arr string

func Microservices() models.Microservices {
	var Microservices models.Microservices
	var matchedMicroservices models.Microservices
	token := "REQUIRED"
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.github.com/orgs/"+"Enterprise-Automation"+"/"+"repos?per_page=100", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	reposReq, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)

	}
	defer reposReq.Body.Close()

	body, _ := io.ReadAll(reposReq.Body)
	if err := json.Unmarshal(body, &Microservices); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for _, microsvc := range Microservices {
		regexPattern := `^lms-check-.*-ms$`
		regex := regexp.MustCompile(regexPattern)
		if regex.MatchString(microsvc.Name) {
			fmt.Printf("Match: %s\n", microsvc.Name)
			matchedMicroservices = append(matchedMicroservices, microsvc)
		} else {
			// fmt.Println("No Match", microsvc.Name)
		}
	}
	return matchedMicroservices
}
