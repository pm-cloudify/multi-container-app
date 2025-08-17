package models

import "go.mongodb.org/mongo-driver/v2/bson"

type QueryTodo struct {
	ID bson.ObjectID `bson:"_id,omitempty" json:"id"`
}

type TODO struct {
	ID      bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title   string        `json:"title" xml:"title" form:"title"`
	Content string        `json:"content" xml:"content" form:"content"`
}

type NewTodoData struct {
	Title   string `json:"title" xml:"title" form:"title" bson:"title"`
	Content string `json:"content" xml:"content" form:"content" bson:"content"`
}

type Todos struct {
	Todos []TODO `json:"todos"`
}
