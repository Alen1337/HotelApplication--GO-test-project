package core

import (
	api_v1 "HotelApplication/src/api/v1"
	"fmt"
	"html/template"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	isAuthorized, authKey := api_v1.CheckAuthorization(w, r)

	if !isAuthorized {
		fmt.Fprintf(w, "Authorization code missing!")
		return
	}

	t, _ := template.ParseFiles("src/public/index.html")

	t.Execute(w, authKey)
}
