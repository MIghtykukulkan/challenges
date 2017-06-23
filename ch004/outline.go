package main

import (
	"fmt"
	"sync"
)

func main() {
	var option int32
	var wg sync.WaitGroup

	// you can also add these one at
	// a time if you need to

	wg.Add(1)

	go func() {
		defer wg.Done()
		details := make(map[string]string)
		details["where"] = "ITC Infotech, Bangalore"
		details["when"] = "When: 22nd & 23rd July, 2017"
		details["what"] = "Cash prizes, on-spot internships & job offers, partnerships"

		fmt.Println("Tell us who you are??")
		fmt.Println("[1]Techie, [2]Engineer, [3]Developer, [4]Startup")
		fmt.Printf("1 - 4 : ")
		fmt.Scanf("%v", &option)

		if option < 1 || option > 4 {
			fmt.Println("Sorry not a valid option")
		} else {
			fmt.Println("Activities you will be involved : Coding, prototype, demonstrate")
			fmt.Println("Interested to register for iTech 2017 y/n?:")
			var yesNo string
			fmt.Scanf("%v", &yesNo)
			fmt.Scanf("%v", &yesNo)
			fmt.Println("you choosed", yesNo)
			if yesNo == "y" {
				fmt.Println("Whats awaiting !!! .. ")
				fmt.Println("What :", details["what"])
				fmt.Println("When:", details["when"])
				fmt.Println("Where:", details["where"])
			} else {
				fmt.Printf("Ok !! May be today is not the day")
			}

		}
	}()
	wg.Wait()

}
