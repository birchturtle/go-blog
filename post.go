package main

import (
	"time"
)

type Post struct {
	id        int
	author    User
	published time.Time
	tags      []string
	content   string
}
