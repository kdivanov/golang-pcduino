// Copyright 2014 jet tsang. All rights reserved.
// license that can be found in the LICENSE file.

package pcduino_test

import (
	"github.com/jetsanix/golang-pcduino"
	"testing"
	"time"
)

func TestPin(t *testing.T) {
	gd, err := pcduino.OpenPin(pcduino.GPIO2)
	if err != nil {
		t.Fail()
	}
	gd.SetMode(pcduino.OUT)
	gd.SetLevel(pcduino.HIGH)
	time.Sleep(1000 * time.Millisecond)
	gd.SetLevel(pcduino.LOW)
	time.Sleep(1000 * time.Millisecond)
	gd.SetLevel(pcduino.HIGH)
	time.Sleep(1000 * time.Millisecond)
	gd.SetLevel(pcduino.LOW)
	gd.Close()
}
