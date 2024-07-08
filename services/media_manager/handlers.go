package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

func uploadImage(s *storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20) // 10 MB

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Failed to receive the file", http.StatusInternalServerError)
			log.Printf("Failed to receive the file: %s\n", err)
			return
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to read the file", http.StatusInternalServerError)
			log.Printf("Failed to read the file: %s\n", err)
			return
		}

		name := generateObjectName()
		fileType := filepath.Ext(header.Filename)

		if err := s.Save(r.Context(), name, fileType, header.Size, data); err != nil {
			http.Error(w, "Failed to save the file", http.StatusInternalServerError)
			log.Printf("Failed to save the file: %s\n", err)
		}

		link := os.Getenv("HOST") + "/img/" + name

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(link))
	}
}

func receiveImage(s *storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")

		data, ext, err := s.Load(r.Context(), name)
		if err != nil {
			http.Error(w, fmt.Sprintf("File with name %s is not found", name), http.StatusNotFound)
			log.Printf("File is not found: %s\n", err)
		}

		w.Header().Set("Content-Type", fmt.Sprintf("image/%s", ext))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func generateObjectName() string {
	const n = 32
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
