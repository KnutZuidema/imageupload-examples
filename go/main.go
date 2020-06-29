package main

import (
	"encoding/json"
	"errors"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// maxFileSize is the maximum size for uploaded files
	// = 20 MB
	maxFileSize = 20 << 20
	formFileKey = "image"
	address     = "localhost:3579"
)

type Error struct {
	Message string `json:"message"`
}

func main() {
	// router holds all known endpoints, applies middleware and implements the http.Listener interface
	router := mux.NewRouter()

	// add Content-Type header for all responses
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			h.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(maxFileSize); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(Error{Message: "file size is too large"}); err != nil {
				log.Printf("ERROR: write error: %v\n", err)
			}
			return
		}
		file, header, err := r.FormFile(formFileKey)
		if err != nil && errors.Is(err, http.ErrMissingFile) {
			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(Error{Message: "missing file"}); err != nil {
				log.Printf("ERROR: write error: %v\n", err)
			}
		} else if err != nil {
			log.Printf("get form file: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(Error{Message: "something went wrong"}); err != nil {
				log.Printf("ERROR: write error: %v\n", err)
			}
			return
		}
		// handle file in a separate goroutine so the user does not have to wait
		go handleFile(file, header)
		w.WriteHeader(http.StatusNoContent)
	}).Methods(http.MethodPost)

	log.Printf("INFO: listening on http://%s\n", address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Printf("ERROR: listen: %v\n", err)
	}
}

func handleFile(file multipart.File, header *multipart.FileHeader) {
	// do something with the file
}
