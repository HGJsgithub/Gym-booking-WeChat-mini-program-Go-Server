package getVenueState

func checkIfVenueTypeLegal(venueType string) bool {
	switch venueType {
	case "badminton":
		return true
	case "tennis":
		return true
	case "tableTennis":
		return true
	}
	return false
}
