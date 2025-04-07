package game

import (
	"example/guess_number/player"
	"slices"

	"github.com/google/uuid"
)

func FindGameById(slice []Game, id uuid.UUID) (*Game, bool) {
	for i := range slice {
		if slice[i].ID == id {
			return &slice[i], true
		}
	}

	return nil, false
}

func DeleteGame(slice []Game, id uuid.UUID) ([]Game, bool) {
	slice = slices.DeleteFunc(slice, func(g Game) bool {
		return g.ID == id
	})
	return slice, true
}

func GetPlayerNames(players []player.Player) []string {
	names := make([]string, len(players))
	for i, player := range players {
		names[i] = player.Name
	}
	return names
}
