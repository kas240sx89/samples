package main

import (
	service2 "github.com/kas240sx89/samples/profileService/internal/service"
	"io"
	"log"
	"net/http"
)

type HttpEndpoints struct {
	svc *service2.Service
}

//NewServer creates a new http server and listens for requests
func NewServer(service *service2.Service) {

	s := HttpEndpoints{service}

	http.HandleFunc("/get_profile", s.GetProfile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//GetProfile retrieves the profile attached to the passed in "id"
func (h *HttpEndpoints) GetProfile(w http.ResponseWriter, req *http.Request) {

	id := req.FormValue("id")
	if id == "" {
		io.WriteString(w, "id is empty\n")
		w.WriteHeader(400)
		return
	}

	profile, err := h.svc.GetProfile(id)
	// TODO with better error, better error codes
	if err != nil {
		w.WriteHeader(404)
		return
	}

	profileJSON, err := profile.ToJSON()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(profileJSON)
	w.WriteHeader(200)
	return
}
