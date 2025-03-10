package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	EndPoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirsName string `json:"first_name"`
		LastName string `json:"last_name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}

	ErrorRes struct {
		Error string `json:"error"`
	}
)

func MakeEndpoints(s Service) EndPoints {

	return EndPoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}

}

func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"Invalid request format"})
			return
		}
		if req.FirsName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"First name is required"})
			return
		}
		if req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"last name is required"})
			return
		}
		err := s.Create(req.FirsName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{err.Error()})
			return
		}
		fmt.Println(req)
		json.NewEncoder(w).Encode(req)
	}
}

func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"okG": true})
	}
}

func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Make GetAll endpoint")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Make update endpoint")
		json.NewEncoder(w).Encode(map[string]bool{"okU": true})
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Make delete endpoint")
		json.NewEncoder(w).Encode(map[string]bool{"okD": true})
	}
}
