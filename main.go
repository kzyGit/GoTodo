package main

import (
    "fmt"
    "log"
    "net/http"

	"./handler"
	"time"
	"./db"
)

func main() {
	var postgres *db.Postgres
	var err error

	for i:=0; i<10; i++ {
		time.sleep(3 * time.second)
		postgres, err = db.ConnectPostgres()
	}
	if err != nil {
        panic(err)
    } else if postgres == nil {
        panic("postgres is nil")
    }

    mux := handler.SetUpRouting()

    fmt.Println("http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}