package main

import (
	"github.com/ChimeraCoder/anaconda"
	"fmt"
)

type Twitter struct {
	Api    *anaconda.TwitterApi
	Rooms  []string
}

func main() {
	fmt.Printf("[INFO] Starting TinyChat Crawler\n")
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	api := anaconda.NewTwitterApi("", "")

	twitter := &Twitter{Api: api}

	err := twitter.RunCrawler()
	if err != nil {
		fmt.Println(err)
	}
}
