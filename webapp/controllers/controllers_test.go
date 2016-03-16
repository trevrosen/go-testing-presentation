package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/trevrosen/go-testing-presentation/webapp/db"
)

func User_Exists(id string) (db.User, error) {
	user := db.User{
		ID:        id,
		FirstName: "Joe",
		LastName:  "Foo",
		Age:       122,
	}
	return user, nil
}

func User_Missing(id string) (db.User, error) {
	user := db.User{}
	err := fmt.Errorf("User not found")
	return user, err
}

func TestShowUserHandler_Exists(t *testing.T) {
	cdb := db.MockDB{
		OnFindUserById: User_Exists,
	}

	server := httptest.NewServer(App(cdb))
	defer server.Close()
	client := new(http.Client)

	req, _ := http.NewRequest("GET", server.URL+"/users/1", nil)
	res, _ := client.Do(req)

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status of 200, found %v", res.Status)
	}
}

func TestShowUserHandler_Missing(t *testing.T) {
	cdb := db.MockDB{
		OnFindUserById: User_Missing,
	}

	server := httptest.NewServer(App(cdb))
	defer server.Close()
	client := new(http.Client)

	req, _ := http.NewRequest("GET", server.URL+"/users/1", nil)
	res, _ := client.Do(req)

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status of 404, found %v", res.Status)
	}

}
