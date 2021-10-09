package main

import "time"

type post struct {
	ID              string
	Caption         string
	ImageURL        string
	PostedTimestamp time.Duration
}