package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trevrosen/go-testing-presentation/webapp/controllers"
	"github.com/trevrosen/go-testing-presentation/webapp/db"
)

func main() {
	cdb, err := db.NewDB()

	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	defer cdb.Gorm.Close()

	portString := ":5000"
	fmt.Printf("[-] Listening on %v\n", portString)
	http.ListenAndServe(portString, controllers.App(cdb))
}
