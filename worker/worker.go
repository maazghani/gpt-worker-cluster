package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

// Task is a struct that represents a task that the worker should perform
type Task struct {
	Text string `json:"text"`
}

func main() {
	// Connect to the NATS server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to the task queue
	_, err = nc.Subscribe("tasks", func(msg *nats.Msg) {
		// Parse the task from the message
		var task Task
		err := json.Unmarshal(msg.Data, &task)
		if err != nil {
			log.Println("Error parsing task:", err)
			return
		}

		// Call the OpenAI API to generate text based on the task
		result, err := generateText(task.Text)
		if err != nil {
			log.Println("Error generating text:", err)
			return
		}

		// Log the result
		log.Println("Result:", result)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Wait for messages to arrive
	select {}
}

func generateText(text string) (string, error) {
	// Replace this with a call to the OpenAI API to generate text based on the input text
	return "Generated text based on " + text, nil
}

