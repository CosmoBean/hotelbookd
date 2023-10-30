package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/CosmoBean/hotelbookd/db"
	"github.com/CosmoBean/hotelbookd/models"
	"github.com/CosmoBean/hotelbookd/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http/httptest"
	"testing"
)

const testDatabase = "hotel-reservation-test"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	var (
		mongoHost = utils.GetEnvDefault("MONGO_HOST", "localhost")
		mongoPort = utils.GetEnvDefault("MONGO_PORT", "27017")
	)
	var dbUri = fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Panic("error while connecting to the database ", err)
	}
	return &testdb{db.NewMongoUserStore(client, testDatabase)}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.CreateNewUser)
	params := models.CreateUserRequest{
		Email:     "some@foo.com",
		FirstName: "foo",
		LastName:  "bar",
		Password:  "passfoobar",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("content-type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var user models.User
	json.NewDecoder(resp.Body).Decode(&user)
	if len(user.Id) == 0 {
		t.Errorf("expecting a user id, got nothing")
	}
	if len(user.EncryptedPass) > 0 {
		t.Errorf("expecting encrypted password to be absent in response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("Expected first name : %s, but got : %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("Expected last name : %s, but got : %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("Expected email : %s, but got : %s", params.Email, user.Email)
	}
}
