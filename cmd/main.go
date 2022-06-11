package main

import (
	"PagerDutyMachineCode/internal/server"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	//initalise my server
	//Rest enpoints

	wg := sync.WaitGroup{}
	wg.Add(1)

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.GetRoot)
	mux.HandleFunc("/teams/create", server.CreateTeam)

	err := http.ListenAndServe(":3333", mux)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	wg.Wait()

	//dsn := "user:Password@123@tcp(127.0.0.1:3306)/pagerduty?charset=utf8mb4&parseTime=True&loc=Local"
	//_, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println(err.Error())
	//	panic("failed to connect database")
	//}

}
