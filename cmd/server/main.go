package main

import (
	"encoding/json"
	"fmt"

	"github.com/fhbkoller/workshop-go/domain"
)

func main() {
	fmt.Println("Olá mundo!")

	relatives := map[string]interface{}{
		"father":   "Laerto",
		"mother":   "Marilei",
		"siblings": []string{"Eduardo"},
	}
	user := domain.NewUser("1", "Koller", 27, []interface{}{"12345677", "12347890"}, relatives)

	fmt.Printf("User: %s\n", user)
	fmt.Printf("User: %v\n", user)
	fmt.Printf("User: %+v\n", user)

	for i, v := range user.GetPhones() {
		fmt.Printf("Phones[%d]: %s\n", i, v)
	}

	for k, v := range user.GetRelatives() {
		fmt.Printf("[%s]: %s\n", k, v)
	}

	j, _ := json.Marshal(user)

	fmt.Printf("%s\n", string(j))

	user2 := domain.NewUser("1", "Fernando", 0, []interface{}{}, nil)

	j2, _ := json.Marshal(user2)

	fmt.Printf("%s\n", string(j2))
}
