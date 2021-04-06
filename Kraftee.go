package main

type Kraftee struct {
	First               string
	Last                string
	Nickname            string
	Year                string
	RefreshTokenEnvName string
	StravaId            string
	StravaAccessToken   string
}

func (k Kraftee) fullName() string {
	return k.First + " " + k.Last
}
