package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/xml"
	"strconv"
)

type RoomInfo struct {
	TotalCount  string  `xml:"total_count,attr,omitempty"`
}

func (twitter *Twitter) ParseRoomFile(RoomFile string) error {
	file, err := ioutil.ReadFile(RoomFile)
	RoomList := strings.Split(string(file), "\n")
	if err != nil {
		return err
	}
	for _, room := range RoomList {
		twitter.Rooms = append(twitter.Rooms, room)
	}

	return nil
}

func (twitter *Twitter) TweetRoom(room string, users string) error {
	_, err := twitter.Api.PostTweet(fmt.Sprintf("There are currently %s people chatting in http://tinychat.com/%s", users, room), nil)
	if err != nil {
		return err
	}
	return nil
}

func (twitter *Twitter) RunCrawler() error {
	err := twitter.ParseRoomFile("rooms.txt")
	if err != nil {
		return err
	}

	for {
		for _, Room := range twitter.Rooms {
			response, err := http.Get(fmt.Sprintf("http://api.tinychat.com/%s.xml", Room))
			if err != nil {
				continue
			}
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				continue
			}

			var roominfo RoomInfo
			err = xml.Unmarshal([]byte(body), &roominfo)
			if err != nil {
				continue
			}

			if len(roominfo.TotalCount) == 0 {
				continue
			}
			TotalCount, err := strconv.Atoi(roominfo.TotalCount)
			if err != nil {
				continue
			}
			if TotalCount >= 2 {
				err = twitter.TweetRoom(Room, roominfo.TotalCount)
				if err != nil {
					fmt.Print("[WARN] %s\n", err)
					continue
				}
				fmt.Printf("[INFO] Tweeted room %s\n", Room)
			}
		}
		time.Sleep(30 * time.Minute)
	}

	return nil
}
