package main

import (
	"bufio"
	"fmt"
	"github.com/Dostonlv/Buy-ticket.git/Functions"
	"os"
	"strings"
)

type UserInfo struct {
	firstName   string
	lastName    string
	email       string
	ticketCount uint
}

func main() {

	// TwoSumm()
	//var allTicks int = 50
	//var zakaz int
	//action := true
	//
	//allPlaces := 0
	//
	//var name string
	//var last string
	//var email string
	//
	//var sum []string
	//
	//for action {
	//	var sub []string
	//
	//	fmt.Scanln(&name)
	//	fmt.Scanln(&last)
	//	fmt.Scanln(&email)
	//
	//	sub = append(sub, name, last, email)
	//	sum = append(sum, sub...)
	//
	//	fmt.Print("nechta joy :")
	//	fmt.Scanln(&zakaz)
	//
	//	if zakaz > 0 {
	//		allPlaces += zakaz
	//	}
	//
	//	if allPlaces >= 50 {
	//		fmt.Println("joy qomadi")
	//		// fmt.Println(sum)
	//		for i := 0; i < len(sum); i += 3 {
	//
	//			fmt.Println(sum[i])
	//			fmt.Println(sum[i+1])
	//			fmt.Println(sum[i+2])
	//			fmt.Println()
	//
	//		}
	//		return
	//	}
	//
	//	allTicks -= zakaz
	//	fmt.Println("qogan joy : ", allTicks)
	//
	//	if zakaz == 101 {
	//		action = false
	//	}
	//
	//}
	//
	//fmt.Println(allPlaces)

	var ticketCount []uint
	places := []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	fmt.Println("Welcome to Buy tisket CLI")
	fmt.Println("1. Get Custom Places")
	fmt.Println("2. Reserved Spaces Place")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	var a string
	for len(places) > 0 {
		fmt.Scan(&a)
		switch choice {
		case "1":
			Functions.GetCustomPlaces(places, ticketCount)
		case "2":
			ticketCount = Functions.ReservedSpaces(places, ticketCount)
		}
	}
	//fmt.Println(ticketCount)
}
