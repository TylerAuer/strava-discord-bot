package main

type Kraftee struct {
	First               string
	Last                string
	RefreshTokenEnvName string
	StravaId            string
	StravaAccessToken   string
	daysBeforeNag       int
}

func (k Kraftee) FullName() string {
	return k.First + " " + k.Last
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
