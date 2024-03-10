package clickupservice

import (
	"context"
	"fmt"
)

func GetTeams() {
	client := getClient()

	// teamid: 6913006

	teams, _, err := client.Teams.GetTeams(context.Background())

	if err != nil {
		panic(err)
	}

	for _, team := range teams {
		fmt.Printf("%+v\n", team.ID)
	}
}
