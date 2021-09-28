package main

import "fmt"

func AppendAnimals(animals []string) {
	animals = append(animals, "老虎", "大象")
}

func AppendAnimalsPointer(animals *[]string) {
	*animals = append(*animals, "老虎", "大象")
}

func UpdateAnimals(animals []string) {
	for i := range animals {
		animals[i] = "兔子"
	}
}

func main() {
	input := []string{"猴子"}
	AppendAnimals(input)
	fmt.Println(input)

	AppendAnimalsPointer(&input)
	fmt.Println(input)

	updateInput := []string{"老虎", "大象"}
	UpdateAnimals(updateInput)
	fmt.Println(updateInput)

}
