package main

import (
	"fmt"
	"log"
	"net/http"
	"store-product-restapi/router"
)

func main(){
	r := router.Router()
	
	fmt.Println("Starting server on the port 8080....")
	log.Fatal(http.ListenAndServe(":8800", r))
}