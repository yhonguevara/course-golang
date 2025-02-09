package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)

	countries["Venezuela"] = "Caracas"
	countries["España"] = "Madrid"
	countries["Argentina"] = "Buenos Aires"
	// fmt.Println(countries)
	// fmt.Println(countries["España"])

	champions := map[string]int{
		"Real Madrid": 14,
		"AC Milan": 7,
		"Bayern Múnich": 6,
		"Liverpool": 6,
		"Barcelona": 5,
		"Ajax": 4,
	}

	fmt.Println(champions)

	for team, titles := range(champions) {
		fmt.Printf("%s champions %d \n", team, titles)
	}

	delete(champions, "Barcelona")
	fmt.Println(champions)

	points, exists := champions["Atletico de Madrid"]
	fmt.Printf("Los titulos: %d, y el equipo existe %t \n", points, exists)
}