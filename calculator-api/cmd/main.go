package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

type RequestBody struct {
	A int `json:"a"`
	B int `json:"b"`
}
type ResponseBody struct {
	Result int `json:"result"`
}
type DivisionResponseBody struct {
	Result    int `json:"result"`
	Remainder int `json:"remainder"`
}
type SumRequestBody struct {
	Items []int `json:"items"`
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := http.NewServeMux()

	router.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			logger.Error("Error decoding request body: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sum := body.A + body.B

		response, err := json.Marshal(&ResponseBody{Result: sum})
		if err != nil {
			logger.Error("Error encoding response: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	router.HandleFunc("POST /subtract", func(w http.ResponseWriter, r *http.Request) {
		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			logger.Error("Error decoding request body: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		difference := body.A - body.B

		response, err := json.Marshal(&ResponseBody{Result: difference})
		if err != nil {
			logger.Error("Error encoding response: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	router.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			logger.Error("Error decoding request body: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product := body.A * body.B

		response, err := json.Marshal(&ResponseBody{Result: product})
		if err != nil {
			logger.Error("Error encoding response: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	router.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			logger.Error("Error decoding request body: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if body.B == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Cannot divide by zero"))
			return
		}

		quotient := body.A / body.B
		remainder := body.A % body.B

		response, err := json.Marshal(&DivisionResponseBody{Result: quotient, Remainder: remainder})
		if err != nil {
			logger.Error("Error encoding response: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	router.HandleFunc("POST /sum", func(w http.ResponseWriter, r *http.Request) {
		var body SumRequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			logger.Error("Error decoding request body: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sum := 0
		for _, item := range body.Items {
			sum += item
		}

		response, err := json.Marshal(&ResponseBody{Result: sum})
		if err != nil {
			logger.Error("Error encoding response: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info("Starting server on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Error starting server: " + err.Error())
	}
}
