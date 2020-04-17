package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Hint maps a question to it's correct answer and hint.
type Hint struct {
	P survey.Prompt
	A string
	H string
}

var pick = survey.Question{
	Prompt: &survey.Select{
		Message: "Hvilket egg trenger du hint for?",
		Options: []string{"Egg 1", "Egg 2"},
	},
}

var egg1 = []Hint{
	{
		P: &survey.Input{
			Message: "Stav navnet på skogen der Mikkel Rev og Klatremus bor.",
		},
		A: "HAKKEBAKKESKOGEN",
		H: "EGG I EGG",
	},
	{
		P: &survey.Input{
			Message: "Hva er 2 * 10?",
		},
		A: "20",
		H: "LILLA",
	},
}

var egg2 = []Hint{
	{
		P: &survey.Select{
			Message: "Hva er 20 på kinesisk?",
			Options: []string{"ER SHÍ", "SHÍ ER", "SAN SHÍ"},
		},
		A: "ER SHÍ",
		H: "TEPPE",
	},
	{
		P: &survey.Select{
			Message: "Hva er påskefargen?",
			Options: []string{"BLÅ", "GUL", "ROSA", "ORANGE", "RØD", "BRUN"},
		},
		A: "GUL",
		H: "UGLE",
	},
}

func main() {
	var i1, i2 int

LOOP:
	for {

		var eggName string
		survey.AskOne(
			&survey.Select{
				Message: "Hvilket egg trenger du hint for?",
				Options: []string{"EGG 1", "EGG 2"},
			},
			&eggName,
		)

		var ix *int
		var eggx []Hint
		switch eggName {
		case "EGG 1":
			ix = &i1
			eggx = egg1
		case "EGG 2":
			ix = &i2
			eggx = egg2
		default:
			break LOOP
		}

		for i := 0; i < *ix; i++ {
			fmt.Println(eggx[i].H)
		}
		if *ix >= len(eggx) {
			fmt.Println("Ikke flere hint")
			continue LOOP
		}
		h := eggx[*ix]

		var a string
		survey.AskOne(h.P, &a)
		if strings.ToUpper(a) == h.A {
			fmt.Println(h.H)
			*ix++
		} else {
			fmt.Println("Feil svar!")
		}
	}
}
