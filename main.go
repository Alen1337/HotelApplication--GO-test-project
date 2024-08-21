package main

import (
	"HotelApplication/src/core"
	"log"
	"net/http"
)

func main() {
	core.RouterInit()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
