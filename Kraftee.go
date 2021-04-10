package main

type Kraftee struct {
	First               string
	Last                string
	RefreshTokenEnvName string
	StravaId            string
	StravaAccessToken   string
}

func (k Kraftee) FullName() string {
	return k.First + " " + k.Last
}

func (k Kraftee) GetStravaAccessToken() string {
	if k.StravaAccessToken != "" {
		return k.StravaAccessToken
	} else {
		token := getStravaAccessToken(k)
		k.StravaAccessToken = token
		return token
	}
}
