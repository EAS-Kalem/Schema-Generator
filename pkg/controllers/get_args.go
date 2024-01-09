package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/models"
)

func Args(Microservice string, Handlers string) string {
	var Args models.Args
	token := "REQUIRED"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/Enterprise-Automation/"+Microservice+"/contents/handlers/"+Handlers, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "token "+token)
	argsReq, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)

	}

	defer argsReq.Body.Close()
	body, _ := io.ReadAll(argsReq.Body)
	if err := json.Unmarshal(body, &Args); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	decoded, err := base64.StdEncoding.DecodeString(Args.Content)
	if err != nil {
		fmt.Println("Error decoding string:", err)
	}
	_, after, _ := strings.Cut(string(decoded), "//Args:{")
	before, _, _ := strings.Cut(string(after), "}")

	return before
}
