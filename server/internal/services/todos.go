package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/pm-cloudify/multi-container-app/server/internal/db"
	"github.com/pm-cloudify/multi-container-app/server/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func getTodosCollection() *mongo.Collection {
	return db.GetCollection("test", "todos")
}

func GetAllTodos() (*models.Todos, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := getTodosCollection().Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := models.Todos{
		Todos: make([]models.TODO, 0),
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.TODO
		if err := cur.Decode(&result); err != nil {
			log.Println(err)
			return nil, err
		}

		res.Todos = append(res.Todos, result)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

func GetTodo(q *models.QueryTodo) (*models.TODO, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := getTodosCollection().Find(ctx, q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := models.Todos{
		Todos: make([]models.TODO, 0),
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.TODO
		if err := cur.Decode(&result); err != nil {
			log.Println(err)
			return nil, err
		}

		res.Todos = append(res.Todos, result)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	if len(res.Todos) == 0 {
		return nil, nil
	}

	return &res.Todos[0], nil
}

func CreateNewTODO(data *models.NewTodoData) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := getTodosCollection().InsertOne(ctx, data)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("created: ", res.InsertedID)

	return nil
}

func UpdateOneTodo(q *models.QueryTodo, data *models.NewTodoData) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := getTodosCollection().UpdateOne(ctx, q, data)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("modified: ", res.ModifiedCount)

	if res.ModifiedCount == 0 {
		return errors.New("not modified")
	}

	return nil
}
