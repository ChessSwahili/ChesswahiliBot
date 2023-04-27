package main

import (
	"time"

	"github.com/ChessSwahili/ChessSWBot/internal/data"
)

func (sw *SWbot) pollTeam(playersId chan<- []string) {

	ticker := time.NewTicker(time.Minute * 5)

	defer ticker.Stop()

	for range ticker.C {

		playersId <- data.FetchTeamPlayers()

	}
}
