package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Card struct {
	Name        map[string]string
	Description map[string]string
	Copies      int
	Introduced  string
	Updated     string
}

type Sigil struct {
	Card
	Initials map[string]string
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
	var lang = flag.String("lang", "en", "Language strings to use in the output")
	var version = flag.String("printvers", "", "Specific version to print out")
	flag.Bool("snap", false, "")

	flag.Parse()

	si, _ := os.Open("cards/sigils.json")
	defer si.Close()
	dec := json.NewDecoder(si)
	sigils := []Sigil{}
	if err := dec.Decode(&sigils); err != nil {
		panic(err)
	}

	fmt.Print("\n% SIGILS\n")

	for _, sig := range sigils {
		if *version != "" && sig.Updated != *version {
			continue
		}

		name, ex := sig.Name[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Name for card %v", lang, sig))
		}

		desc, ex := sig.Description[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Description for action card %v", lang, sig))
		}

		init, ex := sig.Initials[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Initials for action card %v", lang, sig))
		}

		for i := 0; i < sig.Copies; i++ {
			fmt.Printf("\\sigil{%v}{%v}{%v}\n", name, desc, init)
		}
	}

	a, _ := os.Open("cards/actions.json")
	defer a.Close()
	dec = json.NewDecoder(a)
	actions := []Action{}
	if err := dec.Decode(&actions); err != nil {
		panic(err)
	}

	fmt.Print("\n% ACTIONS\n")

	for _, act := range actions {
		if *version != "" && act.Updated != *version {
			continue
		}

		name, ex := act.Name[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Name for card %v", lang, act))
		}

		desc, ex := act.Description[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Description for action card %v", lang, act))
		}

		var inst string
		if act.Instant != nil {
			inst, ex = act.Instant[*lang]
			if !ex {
				panic(fmt.Sprintf("No '%v' language Instant for action card %v", lang, act))
			}
		}

		var pers string
		if act.Persistent {
			pers = "true"
		}

		for i := 0; i < act.Copies; i++ {
			fmt.Printf("\\action{%v}{%v}{%v}{%v}{%v}{%v}\n", name, desc, inst, pers,
				act.Introduced, act.Updated)
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
		if *version != "" && stn.Updated != *version {
			continue
		}

		name, ex := stn.Name[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Name for stance card %v", lang, stn))
		}

		desc, ex := stn.Description[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Description for stance card %v", lang, stn))
		}

		stancetype := strings.ToUpper(stn.Type)

		for i := 0; i < stn.Copies; i++ {
			fmt.Printf("\\stance{%v}{%v}{%v}{%v}{%v}{%v}{%v}\n", name, desc, stancetype,
				stn.AgainstBargain, stn.AgainstBetrayal, stn.Introduced, stn.Updated)
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
		if *version != "" && fin.Updated != *version {
			continue
		}

		name, ex := fin.Name[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Name for final card %v", lang, fin))
		}

		desc, ex := fin.Description[*lang]
		if !ex {
			panic(fmt.Sprintf("No '%v' language Description for final card %v", lang, fin))
		}

		for i := 0; i < fin.Copies; i++ {
			fmt.Printf("\\final{%v}{%v}{%v}{%v}{%v}\n", name, desc, fin.Priority,
				fin.Introduced, fin.Updated)
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
