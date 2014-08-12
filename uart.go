// Copyright 2014 jet tsang. All rights reserved.
// license that can be found in the LICENSE file.

package pcduino

import (
	"github.com/tarm/goserial"
	"io"
)

const serialName string = "/dev/ttyS1"

func SetupUART(baud int) (io.ReadWriteCloser, error) {
	gd0, _ := OpenPin(GPIO0) //uart_rx
	gd1, _ := OpenPin(GPIO1) //uart_tx
	gd0.SetMode(IO_UART)
	gd1.SetMode(IO_UART)
	uartConf := &serial.Config{Name: serialName, Baud: baud}
	return serial.OpenPort(uartConf)
}
