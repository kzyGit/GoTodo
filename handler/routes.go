package handler

import (
    "net/http"

    "../db"
)

// Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.
// A powerful URL router and dispatcher for golang. 

func SetUpRouting() *http.ServeMux {
    todoHandler := &todoHandler{
        samples:  &db.Sample{},
    }

    mux := http.NewServeMux()
	mux.HandleFunc("/samples", todoHandler.GetSamples)
	mux.HandleFunc("/todo" func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
        case http.MethodGet:
            todoHandler.getAllTodo(w, r)
        case http.MethodPost:
            todoHandler.saveTodo(w, r)
        case http.MethodDelete:
            todoHandler.deleteTodo(w, r)
        default:
            responseError(w, http.StatusNotFound, "")
        }
	})

    return mux
}