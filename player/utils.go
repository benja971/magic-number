package player

func FindPlayer(slice *[]Player, name string) (*Player, bool) {
	for i := range *slice {
		if (*slice)[i].Name == name {
			return &(*slice)[i], true
		}
	}

	return nil, false
}
