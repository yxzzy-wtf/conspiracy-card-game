package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Card struct {
	Name        map[string]string
	Description map[string]string
	Copies      int
}

type Action struct {
	Card
	Instant    map[string]string
	Persistent bool
}

type Mission struct {
	Card
	Points int
}

type Stance struct {
	Card
	Type            string
	AgainstBetrayal int
	AgainstBargain  int
}

type Final struct {
	Card
	Priority int
}

func main() {
	lang := os.Args[1]

	a, _ := os.Open("cards/actions.json")
	defer a.Close()
	dec := json.NewDecoder(a)
	actions := []Action{}
	if err := dec.Decode(&actions); err != nil {
		panic(err)
	}

	fmt.Print("\n% ACTIONS\n")

	for _, act := range actions {
		name, ex := act.Name[lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Name for card %v", lang, act))
		}

		desc, ex := act.Description[lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Description for action card %v", lang, act))
		}

		var inst string
		if act.Instant != nil {
			inst, ex = act.Instant[lang]
			if !ex {
				panic(fmt.Sprintf("No '%v' language Instant for action card %v", lang, act))
			}
		}

		var pers string
		if act.Persistent {
			pers = "true"
		}

		for i := 0; i < act.Copies; i++ {
			fmt.Printf("\\action{%v}{%v}{%v}{%v}\n", name, desc, inst, pers)
		}
	}

	s, _ := os.Open("cards/stances.json")
	defer s.Close()
	dec = json.NewDecoder(s)
	stances := []Stance{}
	if err := dec.Decode(&stances); err != nil {
		panic(err)
	}

	fmt.Print("\n% STANCES\n")
	for _, stn := range stances {
		name, ex := stn.Name[lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Name for stance card %v", lang, stn))
		}

		desc, ex := stn.Description[lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Description for stance card %v", lang, stn))
		}

		stancetype := strings.ToUpper(stn.Type)

		for i := 0; i < stn.Copies; i++ {
			fmt.Printf("\\stance{%v}{%v}{%v}{%v}{%v}\n", name, desc, stancetype, stn.AgainstBargain, stn.AgainstBetrayal)
		}
	}

	f, _ := os.Open("cards/finals.json")
	defer f.Close()
	dec = json.NewDecoder(f)
	finals := []Final{}
	if err := dec.Decode(&finals); err != nil {
		panic(err)
	}

	fmt.Print("\n% FINALS\n")
	for _, fin := range finals {
		name, ex := fin.Name[lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Name for final card %v", lang, fin))
		}

		desc, ex := fin.Description[lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Description for final card %v", lang, fin))
		}

		for i := 0; i < fin.Copies; i++ {
			fmt.Printf("\\final{%v}{%v}{%v}\n", name, desc, fin.Priority)
		}
	}

	//TODO
	m, _ := os.Open("cards/missions.json")
	defer f.Close()
	dec = json.NewDecoder(m)
	missions := []Mission{}
	if err := dec.Decode(&missions); err != nil {
		panic(err)
	}
}
