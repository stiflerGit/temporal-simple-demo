package main

import (
	"log"
	"os"
	"time"

	sdkClient "go.temporal.io/sdk/client"
	sdkWorker "go.temporal.io/sdk/worker"

	"temporal-simple-demo/internal/activities"
	"temporal-simple-demo/internal/store"
	"temporal-simple-demo/internal/worker"
	"temporal-simple-demo/internal/workflow"
	"temporal-simple-demo/pkg/emailSender"
)

const (
	temporalHostPortEnvironmentVariableKey = "TEMPORAL_HOST_PORT"
	activityDurationEnvironmentVariableKey = "ACTIVITY_DURATION"
)

func main() {
	temporalHostPort := os.Getenv(temporalHostPortEnvironmentVariableKey)
	activityDurationEnvVar := os.Getenv(activityDurationEnvironmentVariableKey)

	activityDuration, err := time.ParseDuration(activityDurationEnvVar)
	if err != nil {
		log.Fatalln("parsing activity duration", err)
	}

	// The client and worker are heavyweight objects that should be created once per process.
	temporalClient, err := sdkClient.Dial(sdkClient.Options{HostPort: temporalHostPort})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()

	w := sdkWorker.New(temporalClient, worker.TaskQueue, sdkWorker.Options{})

	w.RegisterWorkflow(workflow.MyWorkflow)
	w.RegisterWorkflow(workflow.SendReminderEmail)

	activities := activities.New(activityDuration, &store.Storage{}, &emailSender.EmailSender{})
	w.RegisterActivity(activities)

	err = w.Run(sdkWorker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
