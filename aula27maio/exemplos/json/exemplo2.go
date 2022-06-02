package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type products struct {
	Name string `json:"name"`
	Price int	`json:"price"`
	Published bool `json:"published"`
}

func main()  {
	jsonData := `{"Name": "MacBook Pro", "Price": 1500, "Published":true}`

	var p products
			  // transforma em um array de bytes e joga na variável p
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Price:", p.Price)
	fmt.Println("Published:", p.Published)
}