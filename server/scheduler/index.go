package scheduler

import (
	"context"
	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/quartz"
	"log"
)

func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create scheduler
	sched := quartz.NewStdScheduler()

	// async start scheduler
	sched.Start(ctx)

	// create jobs
	cronTrigger, _ := quartz.NewCronTrigger("1/5 * * * * *")
	functionJob := job.NewFunctionJob(func(_ context.Context) (int, error) {
		log.Println("hello")
		return 42, nil
	})

	sched.ScheduleJob(quartz.NewJobDetail(functionJob, quartz.NewJobKey("functionJob")), cronTrigger)
	// wait for all workers to exit
	sched.Wait(ctx)
}
