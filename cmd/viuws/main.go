package main

import (
	"github.com/viuws/viuws-desktop/internal/nmhost"
)

func main() {
	go nmhost.Run()
}
