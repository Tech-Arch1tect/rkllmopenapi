/*
 * OpenAI API
 *
 * APIs for sampling from and fine-tuning language models
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"

	// WARNING!
	// Pass --git-repo-id and --git-user-id properties when generating the code
	//
	"github.com/Tech-Arch1tect/rkllmopenapi/config"
	sw "github.com/Tech-Arch1tect/rkllmopenapi/go"
)

func main() {
	config.LoadConfig()

	routes := sw.ApiHandleFunctions{}

	log.Printf("Server started")

	router := sw.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}
