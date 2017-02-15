package main

import "io/ioutil"
import "fmt"

func main() {

	filecontent, err := ioutil.ReadFile("readme.txt")

	if err != nil {
		panic(err)
	}

	filestr := string(filecontent)

	fmt.Println(filestr)
	fmt.Println(len(filestr))

}
