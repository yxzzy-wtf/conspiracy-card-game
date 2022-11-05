package main

import (
	"encoding/json"
	"os"
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
	a, _ := os.Open("cards/actions.json")
	defer a.Close()
	dec := json.NewDecoder(a)
	actions := []Action{}
	if err := dec.Decode(&actions); err != nil {
		panic(err)
	}

	s, _ := os.Open("cards/stances.json")
	defer s.Close()
	dec = json.NewDecoder(s)
	stances := []Stance{}
	if err := dec.Decode(&stances); err != nil {
		panic(err)
	}

	f, _ := os.Open("cards/finals.json")
	defer f.Close()
	dec = json.NewDecoder(f)
	finals := []Final{}
	if err := dec.Decode(&finals); err != nil {
		panic(err)
	}

	m, _ := os.Open("cards/missions.json")
	defer f.Close()
	dec = json.NewDecoder(m)
	missions := []Mission{}
	if err := dec.Decode(&missions); err != nil {
		panic(err)
	}
}
