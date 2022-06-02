package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type produtos struct {
	Name string `json:"name"`
	Price int	`json:"price"`
	Published bool `json:"published"`
}

func main()  {
	jsonData := `[
	{"Name": "MacBook Air", "Price": 1500, "Published":true}, 
	{"Name": "MacBook Pro", "Price": 2500, "Published":true}
	]`

	var p []produtos

			  // transforma em um array de bytes e joga na variável p
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(p))

	for _, pd := range p {
		fmt.Println("Name:", pd.Name)
		fmt.Println("Price:", pd.Price)
		fmt.Println("Published:", pd.Published)
	}
}