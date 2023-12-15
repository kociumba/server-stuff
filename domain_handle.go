package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type Data struct {
	Greetings string `json:"greetings"`
	Example   struct {
		Name string `json:"name"`
		Info string `json:"info"`
	} `json:"example"`
	Balls float64 `json:"balls"`
}

var is_balls chan bool

// helloHandler handles the HTTP requests for the /hello endpoint.
//
// It accepts GET and POST requests. For GET requests, it reads the contents
// of the "hello.json" file and returns it as the response. For POST requests,
// it decodes the request body into a Request object, constructs a Response
// object with a greeting message, and returns it as the response.
//
// Parameters:
// - w: the HTTP response writer that will be used to write the response.
// - r: the HTTP request object that contains the request data.
//
// Return type: None
func DomainHandler(w http.ResponseWriter, r *http.Request) {

	is_balls = make(chan bool)

	is_balls <- false

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		data := Data{
			Greetings: "hello x",
			Example: struct {
				Name string `json:"name"`
				Info string `json:"info"`
			}{
				Name: "balls",
				Info: "use only the name field to get a response using the post method",
			},
			Balls: checkBalls(),
		}

		jsonData, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			panic(err)
		}

		w.Write(jsonData)

	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")

		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := Response{
			Message: "",
		}

		if req.Name == "balls" {
			response.Message = "https://www.youtube.com/watch?v=Ke0Li-babo4 " + "balls aquierd"
			is_balls <- true
		} else {
			response.Message = "hello " + req.Name + ", where are your balls ???"
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)

	}
}
