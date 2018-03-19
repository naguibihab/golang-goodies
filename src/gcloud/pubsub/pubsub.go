package main

import "cloud.google.com/go/pubsub"
import "golang.org/x/net/context"

import "fmt"
import "log"

func main() {
    fmt.Printf("Running...\n")

    ctx := context.Background()
    projectID := "my-project-id"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	const topic = "my-topic"
	const msg = "Hello from the API, I must have called a thousand times"

	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
		return
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}

func publish(client *pubsub.Client, topic, msg string) error {
	ctx := context.Background()
	// [START publish]
	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	// [END publish]
	return nil
}
