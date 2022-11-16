package main

import (
	"context"
	"log"
	"os"

	"go.temporal.io/sdk/client"

	"temporal-simple-demo/internal/worker"
	"temporal-simple-demo/internal/workflow"
)

const (
	temporalHostPortEnvironmentVariableKey = "TEMPORAL_HOST_PORT"
	userName                               = "Stefano"
)

func main() {
	ctx := context.Background()

	temporalHostPort := os.Getenv(temporalHostPortEnvironmentVariableKey)

	// The client and worker are heavyweight objects that should be created once per process.
	temporalClient, err := client.Dial(client.Options{HostPort: temporalHostPort})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()

	wfRun, err := temporalClient.ExecuteWorkflow(ctx,
		client.StartWorkflowOptions{
			TaskQueue: worker.TaskQueue,
		},
		workflow.SendReminderEmail,
		userName,
	)
	if err != nil {
		log.Fatalln("Executing the workflow", err)
	}

	if err := wfRun.Get(ctx, nil); err != nil {
		log.Fatalln("Workflow failed", err)
	}

	os.Exit(0)
}
