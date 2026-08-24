package main

import "fmt"

func maps_init() {

	// Map Literal

	// teams := map[string][]string{
	// 	"New York Yankees":    {"Aaron Judge", "Juan Soto", "Giancarlo Stanton"},
	// 	"Los Angeles Dodgers": {"Shohei Ohtani", "Mookie Betts", "Freddie Freeman"},
	// 	"Boston Red Sox":      {"Rafael Devers", "Masataka Yoshida", "Masataka Yoshida"},
	// 	"Toronto Blue Jays":   {"Vladimir Guerrero Jr.", "Bo Bichette", "Jose Berrios"},
	// }

	// fmt.Println(teams["New York Yankees"][0])

	ages := make(map[int][]string, 10)
	fmt.Println(ages)
	// Everything not comparable using == or != can't be a key

}
