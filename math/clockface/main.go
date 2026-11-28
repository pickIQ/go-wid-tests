package main

import (
	"os"
	"time"

	clockface "go-wid-tests/math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
