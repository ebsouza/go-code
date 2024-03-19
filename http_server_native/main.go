package main

import (
	"encoding/json"
	"example/http-server/note"
	"fmt"
	"net/http"
	"strconv"
)

var db = note.NewDB()

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /notes", func(w http.ResponseWriter, r *http.Request) {
		var n note.Note
		err := json.NewDecoder(r.Body).Decode(&n)

		if err != nil {
			fmt.Println(err)
			return
		}

		db.AddNote(n)

		w.WriteHeader(http.StatusCreated)
	})

	router.HandleFunc("GET /notes", func(w http.ResponseWriter, r *http.Request) {
		notes := db.GetNotes()
		j, _ := json.Marshal(notes)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	router.HandleFunc("GET /notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.PathValue("id"))

		n, ok := db.GetNote(id)

		if ok == false {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		j, _ := json.Marshal(n)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	router.HandleFunc("PUT /notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println("update a note by id", id)
	})

	router.HandleFunc("DELETE /notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.PathValue("id"))

		db.RemoveNote(id)

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Starting simple http server... \nReady!")
	http.ListenAndServe(":8080", router)
}
