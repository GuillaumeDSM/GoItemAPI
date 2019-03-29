package main

import "time"

type Item struct {
	Text   string    `json:"text"`
	Rating int       `json:"rating"`
	Date   time.Time `json:"date"`
}

type Items []Item
