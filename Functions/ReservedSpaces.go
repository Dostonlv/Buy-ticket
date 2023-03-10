package Functions

import "fmt"

func ReservedSpaces(places []uint, ticketCount []uint) []uint {
	GetCustomPlaces(places, ticketCount)
	var orderedPlace []uint
	var orderPlace uint
	//var ticketCount uint
	var expr uint
	var a = true
	for a {
		fmt.Println("Joy buyurtma qilamizmi unda 1 ni yozing yoki to'xtatamizmi? unda 2 ni yozing: ")
		fmt.Scan(&expr)
		switch expr {
		case 1:
			fmt.Scan(&orderPlace)
			orderedPlace = append(orderedPlace, orderPlace)
			//places = [50]uint(RemoveIndex(places, orderPlace))
		case 2:
			a = false
		}

	}
	return orderedPlace

}
