package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type product struct {
	Name string `json:"name"`
	Price int	`json:"price"`
	Published bool `json:"published"`
}

func main()  {
	p := product{
		Name: "MacBook Pro",
		Price: 1500,
		Published: true,
	}
	
					//retorna um json
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

}