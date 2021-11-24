package main

import (
	"fmt"
	"log"
	"projects/review-finder/configs"
	"projects/review-finder/tools/email"
	"sync"
)

func main() {
	fmt.Println("\n  ----------------------------- Starting Short Term Rental Finder --------------------------------")

	basicConfig, err := configs.Init()
	if err != nil {
		log.Fatal("failed to init config, err: " + err.Error())
	}

	emailServer := email.NewEmailService(basicConfig.EmailSettings)
	jobWG := sync.WaitGroup{}
	for _, alert := range basicConfig.Alerts {
		jobWG.Add(1)

		a := alert
		go a.CreateNewJob(emailServer)
	}
	jobWG.Wait()

	fmt.Println("\n  ----------------------------- Ending STR Finder ----------------------------------")

}
