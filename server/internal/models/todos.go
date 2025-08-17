package models

import "go.mongodb.org/mongo-driver/v2/bson"

type QueryTodo struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`
}

type TODO struct {
	ID      bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
}

type Todos struct {
	Todos []TODO `json:"todos"`
}
