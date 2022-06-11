package server

import (
	"PagerDutyMachineCode/internal/developers"
	"PagerDutyMachineCode/internal/teams"
	"fmt"
	"io"
	"net/http"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / createTeam req\n")
	var devs = []developers.Developer{
		{
			Name:        "Abhinav",
			PhoneNumber: "000000010",
		},
	}

	team, err := teams.GetCore().Create("team1", devs)

	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	respString := fmt.Sprintf("%#v", team)
	io.WriteString(w, respString)
}

func Status() {
	fmt.Println("I'm alive")
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
