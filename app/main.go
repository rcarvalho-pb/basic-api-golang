package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.LoadTemplates()
	router := router.Generate()

	fmt.Println("Starting webapp in Port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}