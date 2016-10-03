package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoDoc struct {
	ID bson.ObjectId `bson:"_id"`
}

// Company - target company to work at
type Company struct {
	mongoDoc
	Name         string
	Website      string
	Applications []Application
}

// Application is a job application
type Application struct {
	mongoDoc
	Title       string
	Link        string
	Status      string
	Date        time.Time
	Since       time.Time
	Description string
}

func dialInfo() (*mgo.DialInfo, error) {
	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		return nil, fmt.Errorf("'MONGODB_URL' cannot be empty")
	}
	return mgo.ParseURL(mongoURI)
}
