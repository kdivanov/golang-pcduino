// Copyright 2014 jet tsang. All rights reserved.
// license that can be found in the LICENSE file.

package pcduino

import (
	"io"
	"os"
	"strconv"
	"syscall"
)

type PINMODE string

// 5 mode in GPIO pins
const (
	IN       PINMODE = "0"
	OUT      PINMODE = "1"
	IO_SPI   PINMODE = "2"
	IO_SPIEX PINMODE = "3"
	IO_UART  PINMODE = "3"
)

type IOLEVEL string

const (
	// Low represents 0.
	LOW IOLEVEL = "0"

	// High represents 1.
	HIGH IOLEVEL = "1"
)

// GPIO Path of pcDuino
const (
	GPIO_PIN_PATH  string = "/sys/devices/virtual/misc/gpio/pin/"
	GPIO_MODE_PATH string = "/sys/devices/virtual/misc/gpio/mode/"
	GPIO_PREFIX    string = "gpio"
)

const (
	GPIO0  int = iota
	GPIO1  int = iota
	GPIO2  int = iota
	GPIO3  int = iota
	GPIO4  int = iota
	GPIO5  int = iota
	GPIO6  int = iota
	GPIO7  int = iota
	GPIO8  int = iota
	GPIO9  int = iota
	GPIO10 int = iota
	GPIO11 int = iota
	GPIO12 int = iota
	GPIO13 int = iota
	GPIO14 int = iota
	GPIO15 int = iota
	GPIO16 int = iota
	GPIO17 int = iota
	GPIO18 int = iota
	GPIO19 int = iota
	GPIO20 int = iota
	GPIO21 int = iota
	GPIO22 int = iota
	GPIO23 int = iota
)

const (
	SPI_CS     int = GPIO10
	SPI_MOSI   int = GPIO11
	SPI_MISO   int = GPIO12
	SPI_CLK    int = GPIO13
	SPIEX_CS   int = GPIO20
	SPIEX_MOSI int = GPIO21
	SPIEX_MISO int = GPIO22
	SPIEX_CLK  int = GPIO23
)

const (
	MAX_GPIO_NUM      int = 23
	MAX_GPIO_MODE_NUM int = 8
	MAX_PWM_NUM       int = 5
	MAX_ADC_NUM       int = 11
)

type GpioDrive struct {
	pinNum     int
	modeHandle io.ReadWriteCloser
	pinHandle  io.ReadWriteCloser
}

type gpio interface {
	Write(p []byte)
	Read(p []byte)
	SetMode(pm PINMODE)
	SetLevel(il IOLEVEL)
	Close()
}

// Open a GPIO Port with pinNum
func OpenPin(pinNum int) (gd *GpioDrive, err error) {
	modeHandle, modErr := os.OpenFile(GPIO_MODE_PATH+GPIO_PREFIX+strconv.Itoa(int(pinNum)), syscall.O_RDWR, 0777)
	pinHandle, pinErr := os.OpenFile(GPIO_PIN_PATH+GPIO_PREFIX+strconv.Itoa(int(pinNum)), syscall.O_RDWR, 0777)
	if modErr != nil {
		return nil, modErr
	}
	if pinErr != nil {
		return nil, pinErr
	}
	defer func() {
		if modErr != nil && modeHandle != nil {
			modeHandle.Close()
		}
		if pinErr != nil && pinHandle != nil {
			pinHandle.Close()
		}
	}()
	return &GpioDrive{
		pinNum:     pinNum,
		modeHandle: modeHandle,
		pinHandle:  pinHandle,
	}, nil
}

//
func (gd *GpioDrive) SetMode(pm PINMODE) (n int, err error) {
	defer func() {
		gd.modeHandle.Close()
	}()
	return gd.modeHandle.Write([]byte(pm))
}

func (gd *GpioDrive) SetLevel(il IOLEVEL) (n int, err error) {
	return gd.pinHandle.Write([]byte(il))
}

func (gd *GpioDrive) Write(p []byte) (n int, err error) {
	return gd.pinHandle.Write(p)
}

func (gd *GpioDrive) Read(p []byte) (n int, err error) {
	return gd.pinHandle.Read(p)

}
func (gd *GpioDrive) Close() error {
	err := gd.pinHandle.Close()
	if err != nil {
		return err
	}
	return nil
}
