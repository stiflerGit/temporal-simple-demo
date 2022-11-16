package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"temporal-simple-demo/internal/activities"
)

// MyWorkflowArgument ...
type MyWorkflowArgument struct{}

// MyWorkflow ...
func MyWorkflow(ctx workflow.Context, _ MyWorkflowArgument) error {
	var a activities.Activities

	logger := workflow.GetLogger(ctx)

	ctx = workflow.WithActivityOptions(ctx,
		workflow.ActivityOptions{
			ScheduleToCloseTimeout: 24 * time.Hour,
		},
	)

	logger.Info("workflow started")

	err := workflow.ExecuteActivity(ctx, a.Fail, time.Second).Get(ctx, nil)
	if err != nil {
		logger.Error("MyWorkflow activity execution failed: a.FailAfter", "error", err)
		return err
	}

	logger.Info("workflow completed")

	return nil
}

// SendReminderEmail ...
func SendReminderEmail(ctx workflow.Context, user string) error {
	var a activities.Activities

	ctx = workflow.WithActivityOptions(ctx,
		workflow.ActivityOptions{
			ScheduleToCloseTimeout: time.Hour,
		},
	)

	logger := workflow.GetLogger(ctx)

	for {
		var hasSignedUp bool
		err := workflow.ExecuteActivity(ctx, a.HasUserSignedUp, user).Get(ctx, &hasSignedUp)
		if err != nil {
			logger.Error("executing activity HasUserSignedUp", "error", err)
			return err
		}

		if hasSignedUp {
			logger.Info("workflow completed")
			return nil
		}

		err = workflow.ExecuteActivity(ctx, a.SendReminderEmail, user).Get(ctx, nil)
		if err != nil {
			logger.Error("executing activity SendReminderEmail", "error", err)
			return err
		}

		if err := workflow.Sleep(ctx, 30*time.Second); err != nil {
			logger.Error("error while sleeping", "error", err)
			return err
		}
	}
}

func BPMNWorkflow(ctx workflow.Context, user string) error {
	var a activities.Activities

	ctx = workflow.WithActivityOptions(ctx,
		workflow.ActivityOptions{
			ScheduleToCloseTimeout: time.Hour,
		},
	)

	logger := workflow.GetLogger(ctx)

	userSignedUpChannel := workflow.GetSignalChannel(ctx, "USER_SIGNED_UP")
	timer := workflow.NewTimer(ctx, 30*(24*time.Hour))

	selector := workflow.NewSelector(ctx)
	selector.AddFuture(timer, func(f workflow.Future) {
		err := workflow.ExecuteActivity(ctx, a.SendReminderEmail, user).Get(ctx, nil)
		if err != nil {
			logger.Error("executing activity SendReminderEmail", "error", err)
		}
	})

	selector.AddReceive(userSignedUpChannel, func(ch workflow.ReceiveChannel, more bool) {
		err := workflow.ExecuteActivity(ctx, a.SendWelcomeEmail, user).Get(ctx, nil)
		if err != nil {
			logger.Error("executing activity SendWelcomeEmail", "error", err)
		}
	})

	for {
		selector.Select(ctx)
	}
}
