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

func krafteeList() []Kraftee {
	ta := Kraftee{"Tyler", "Auer", "Ugly Stick", "2007", "TYLER", "20419783", ""}

	kraftees := []Kraftee{ta}

	return kraftees
}
