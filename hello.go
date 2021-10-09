package main

import (
    "fmt"
    "log"
    "net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
    
)

type Post struct {
    Id uint
    Caption  string   
    Imageurl string   
    Timestamp string
}

type Users struct {
    Id uint
    Name string
    Email string
    Password string
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Print("HEY")
}

func main() {    
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Nishma:root@cluster0.g32ph.mongodb.net/Cluster0?retryWrites=true&w=majority"))
	if err != nil { return err }
    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil { return err }

    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
