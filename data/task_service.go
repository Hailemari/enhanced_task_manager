package data

import (
	"log"
	"errors"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Hailemari/enhanced_task_manager/models"
)



var taskCollection *mongo.Collection

func ConnectDB(mongoURI string) error {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Ping the database to ensure the connection is established
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")
	taskCollection = client.Database("taskDB").Collection("tasks")
	return nil
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	// Find all documents in the collection
	cursor, err := taskCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Printf("Cannot decode a task %v", err)
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskByID(id string) (models.Task, bool) {
	var task models.Task
	
	// Find a single document by ID
	err := taskCollection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task, false
		}
		log.Printf("Error finding task by ID %v", err)
		return task, false
	}
	return task, true
}
func DeleteTask(id string) error {
	// Delete the task by ID
    result, err := taskCollection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
    if err != nil {
        log.Printf("Error deleting task with ID %s: %v\n", id, err)
        return err
    }

	// Check if a task was deleted
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func UpdateTask(id string, updatedTask models.Task) error {
	// Find the task and update it with the new data
	result := taskCollection.FindOneAndUpdate(context.TODO(), bson.D{{Key:  "id", Value: id}}, bson.D{{Key: "$set", Value: updatedTask}})

	// Check if the task was successfully updated
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return errors.New("task not found")
		}
		return result.Err()
	}
	return nil
}

func AddTask(newTask models.Task) error {
	// Insert the new task directly into the collection
	_, err := taskCollection.InsertOne(context.TODO(), newTask)

	if err != nil {
		return err
	}

	return nil
}
