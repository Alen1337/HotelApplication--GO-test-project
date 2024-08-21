package api_v1

import (
	api_v1_models "HotelApplication/src/api/v1/models"
	"fmt"
	"net/http"
	"strconv"
)

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	isAuthorized, _ := CheckAuthorization(w, r)

	if !isAuthorized {
		fmt.Fprintf(w, "Authorization code missing!")
		return
	}

	switch r.Method {
	case "GET":
		if searchByRoom(w, r) {
			return
		}
		if searchByFirstName(w, r) {
			return
		}
		if searchByLastName(w, r) {
			return
		}
	case "POST":
		if addPerson(w, r) {
			return
		}
	default:
		http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
	}
}

func searchByRoom(w http.ResponseWriter, r *http.Request) bool {
	room := r.URL.Query().Get("room")

	if room == "" {
		return false
	}

	roomInt, err := strconv.Atoi(room)
	if err != nil {
		fmt.Println("Room has to be an Integer! Details: ")
	}

	// TODO: Search in the database for the right person(s) by the room
	person := api_v1_models.Person{FirstName: "John", LastName: "Doe", Room: roomInt}

	fmt.Fprintf(w, "First Name: %s, Last Name: %s, Room: %d", person.FirstName, person.LastName, person.Room)

	return true
}

func searchByFirstName(w http.ResponseWriter, r *http.Request) bool {
	firstName := r.URL.Query().Get("firstname")

	if firstName == "" {
		return false
	}

	// TODO: Search in the database for the right person(s) by the first name
	person := api_v1_models.Person{FirstName: firstName, LastName: "Doe", Room: 231}

	fmt.Fprintf(w, "First Name: %s, Last Name: %s, Room: %d", person.FirstName, person.LastName, person.Room)

	return true
}

func searchByLastName(w http.ResponseWriter, r *http.Request) bool {
	lastName := r.URL.Query().Get("lastname")

	if lastName == "" {
		return false
	}

	// TODO: Search in the database for the right person(s) by the last name
	person := api_v1_models.Person{FirstName: "Tracy", LastName: lastName, Room: 376}

	fmt.Fprintf(w, "First Name: %s, Last Name: %s, Room: %d", person.FirstName, person.LastName, person.Room)

	return true
}

func addPerson(w http.ResponseWriter, r *http.Request) bool {
	firstName := r.URL.Query().Get("firstname")
	lastName := r.URL.Query().Get("lastname")
	room := r.URL.Query().Get("room")

	// TODO: more validation
	roomInt, err := strconv.Atoi(room)
	if err != nil {
		fmt.Println("Room has to be an Integer! Details: ")
	}

	if firstName == "" || lastName == "" || room == "" {
		fmt.Fprintf(w, "Person details missing! \nFirst Name: %s, Last Name: %s, Room: %s", firstName, lastName, room)
		return false
	}
	

	// TODO: add new person to the database
	person := api_v1_models.Person{FirstName: firstName, LastName: lastName, Room: roomInt}

	fmt.Fprintf(w, "Person successfully added to the database! \nFirst Name: %s, Last Name: %s, Room: %d", person.FirstName, person.LastName, person.Room)

	return true
}
