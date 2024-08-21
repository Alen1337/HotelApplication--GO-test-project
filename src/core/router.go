package core

import (
	api_v1 "HotelApplication/src/api/v1"
	"net/http"
)

func RouterInit() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/api/v1/person", api_v1.PersonHandler)
	http.HandleFunc("/api/v1/room", api_v1.RoomHandler)
}
