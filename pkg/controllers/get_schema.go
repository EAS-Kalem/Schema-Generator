package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Enterprise-Automation/lms-schema-generator-app/pkg/models"
	"io"
	"net/http"
	"os"
	"log"
)

func GetSchema(fileName string) string {
	var FileContains models.FileContains

	resp, err := http.Get("https://api.github.com/repos/" + "EAS-Kalem" + "/" + "test-repo" + "/contents/" + fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Get Schema:", resp.Status)

	if err := json.Unmarshal(body, &FileContains); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	decoded, err := base64.StdEncoding.DecodeString(FileContains.Content)
	if err != nil {
		fmt.Println("Error decoding string:", err)
	}
	
	file, err := json.MarshalIndent(string(decoded), "", " ")
	if err != nil {
		log.Println("Unable to create json file")
	}
	newFileName := "Schema-"+ fileName
	_ = os.WriteFile(newFileName, file, 0644)
    return string(newFileName)
}