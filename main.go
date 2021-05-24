package main

import "fmt"

func main() {
	priceInfo, err := CrawPage("http://prod.danawa.com/info/?pcode=3148363")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", priceInfo)
}
