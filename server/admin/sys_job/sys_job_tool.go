package sys_job

import (
	"context"
	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/quartz"
	"log"
	"x_admin/util"
)

var sched quartz.Scheduler

func Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create scheduler
	sched = quartz.NewStdScheduler()

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
	loadFromDb()
	sched.Wait(ctx)

}

func loadFromDb() {
	list, e := Service.ListAll()
	if e != nil {
		log.Println(e)
		return
	}
	for _, v := range list {
		if v.Status == "1" {
			continue
		}
		AddJob(v.Name, v.CronExpression, v.InvokeTarget)
	}
}

func JobExists(jobKey *quartz.JobKey) bool {
	_, err := sched.GetScheduledJob(jobKey)
	return err == nil
}

func AddJob(jobName string, cron string, invokeTarget string) {
	jobKey := quartz.NewJobKey(jobName)
	if JobExists(jobKey) {
		sched.DeleteJob(jobKey)
	}

	cronTrigger, _ := quartz.NewCronTrigger(cron)
	functionJob := job.NewFunctionJob(func(_ context.Context) (int, error) {
		log.Printf("invoke target: %s\n", invokeTarget)
		util.Get(invokeTarget)
		return 42, nil
	})
	sched.ScheduleJob(quartz.NewJobDetail(functionJob, jobKey), cronTrigger)
}

func DelJob(jobName string) {
	jobKey := quartz.NewJobKey(jobName)
	if JobExists(jobKey) {
		sched.DeleteJob(jobKey)
	}
}

func TriggerJob(jobName string, invokeTarget string) {
	util.Get(invokeTarget)
}
