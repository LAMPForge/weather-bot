package main

import (
	"fmt"
	"log"
)

func (c *Controller) set_cron(cronExpr string, cronFunc func()) {
	_, err := c.cronScheduler.AddFunc(cronExpr, cronFunc)
	if err != nil {
		log.Fatalf("Failed to schedule cron job: %v", err)
	}

	c.cronScheduler.Start()
	fmt.Printf("Cron job added with expression '%s'.\n", cronExpr)
}
