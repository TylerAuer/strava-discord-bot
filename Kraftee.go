package main

type Kraftee struct {
	first               string
	Last                string
	RefreshTokenEnvName string
	StravaId            string
	StravaAccessToken   string
	daysBeforeNag       int
}

func (k Kraftee) FullName() string {
	return k.first + " " + k.Last
}

// SafeFirstName returns the first name of a Kraftee, but uses substitutions for
// Kraftees that have the same name. SafeFirstName should always be used in place of Kraftee.First
func (k Kraftee) SafeFirstName() string {
	switch k.StravaId {
	case "80996402":
		return "Quella"
	case "89420051":
		return "Sweeney"
	default:
		return k.first
	}
}

func (k Kraftee) GetStravaAccessToken() string {
	if k.StravaAccessToken != "" {
		return k.StravaAccessToken
	} else {
		token := fetchStravaAccessToken(k)
		k.StravaAccessToken = token
		return token
	}
}
