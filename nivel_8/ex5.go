package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type sortIdade []user

func (p sortIdade) Len() int           { return len(p) }
func (p sortIdade) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p sortIdade) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type sortSobrenome []user

func (p sortSobrenome) Len() int           { return len(p) }
func (p sortSobrenome) Less(i, j int) bool { return p[i].Last < p[j].Last }
func (p sortSobrenome) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	// Normal
	users := []user{u1, u2, u3}
	printUsers(users, "Sem Ordenação: ")

	// Ordenado por idade
	sort.Sort(sortIdade(users))
	printUsers(users, "Ordenação por idade: ")

	// Ordenando sayings
	for _, v := range users {
		sort.Strings(v.Sayings)
	}
	printUsers(users, "Sayings Ordenadas: ")

	// Ordenando sobrenomes
	sort.Sort(sortSobrenome(users))
	printUsers(users, "Ordenação por sobrenome: ")
}

func printUsers(users []user, mensagem string) {
	fmt.Println(mensagem)
	for i, v := range users {
		fmt.Printf("\t%d. Nome: %s\n", i+1, v.First)
		fmt.Printf("\t%d. Sobrenome: %s\n", i+1, v.Last)
		fmt.Printf("\t%d. Idade: %d\n", i+1, v.Age)
		fmt.Printf("\t%d. Sayings: %v\n", i+1, v.Sayings)
		fmt.Println("")
	}
}
