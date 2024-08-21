package api_v1

import (
	api_v1_models "HotelApplication/src/api/v1/models"
	"fmt"
	"net/http"
	"strconv"
)

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	isAuthorized, _ := CheckAuthorization(w, r)

	if !isAuthorized {
		fmt.Fprintf(w, "Authorization code missing!")
		return
	}

	switch r.Method {
	case "GET":
		if searchRoomByID(w, r) {
			return
		}
	case "POST":

	default:
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
	}
}

func searchRoomByID(w http.ResponseWriter, r *http.Request) bool {
	roomString := r.URL.Query().Get("room")

	if roomString == "" {
		return false
	}

	roomInt, err := strconv.Atoi(roomString)
	if err != nil {
		fmt.Println("Room has to be an Integer! Details: ")
	}

	// TODO: Search in the database for the room details
	room := api_v1_models.Room{ID: roomInt, Status: api_v1_models.Clean}

	fmt.Fprintf(w, "Room: %d, Status: %d", room.ID, room.Status)

	return true
}
