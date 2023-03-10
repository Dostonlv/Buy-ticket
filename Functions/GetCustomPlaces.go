package Functions

import (
	"fmt"
	"github.com/TwiN/go-color"
)

func GetCustomPlaces(places []uint, ticketCount []uint) {
	s := len(ticketCount)
	if s == 0 {
		fmt.Println(color.Ize(color.Yellow, "Qolgan joy"), len(places))
		for i := 0; i < len(places); i++ {
			if places[i]%10 == 0 {
				fmt.Print(color.Ize(color.Green, places[i]))
				print(" ")
				println()
			} else {
				fmt.Print(color.Ize(color.Green, places[i]))
				print(" ")
			}

		}
	} else {
		fmt.Println("Qolgan joy", len(places)-len(ticketCount))
		for i := 0; i < len(places); i++ {
			if places[i]%10 == 0 {
				fmt.Print(color.Ize(color.Green, places[i]))
				print(" ")
				println()
			} else {
				var s uint
				s = uint(len(ticketCount))
				var j uint
				for j = 0; j < s; j++ {
					if places[i] != ticketCount[j] {

						fmt.Print(color.Ize(color.Green, places[i]))
						print(" ")
					} else {
						//j++
						fmt.Print(color.Ize(color.Red, ticketCount[j]))
						print(" ")
					}
				}
			}

		}
	}
}
