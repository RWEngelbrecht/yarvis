package helper

import (
	"log"
	"os"

	c "github.com/RWEngelbrecht/yarvis/colours"
)

func OutputError(errText string) {
	l := log.New(os.Stderr, "", 0)
	l.Println(string(c.Red) + errText)
	// panic(string(c.Red), errors.New(errText).Error())
}
