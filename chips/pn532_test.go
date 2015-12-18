package pn532_test

import (
	"fmt"
	"github.com/jetsanix/golang-pcduino/chips"
	"testing"
)

func TestPN532(t *testing.T) {
	pndrive, err := pn532.OpenPN532()
	if err != nil {
		t.Fail()
	}
	fmt.Println("begin wake up")
	buf := pndrive.PN532UartWakeup()
	fmt.Println(buf)
}
