package game

import (
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
