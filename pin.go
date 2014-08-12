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

type PINNUM int

const (
	GPIO0  PINNUM = iota
	GPIO1  PINNUM = iota
	GPIO2  PINNUM = iota
	GPIO3  PINNUM = iota
	GPIO4  PINNUM = iota
	GPIO5  PINNUM = iota
	GPIO6  PINNUM = iota
	GPIO7  PINNUM = iota
	GPIO8  PINNUM = iota
	GPIO9  PINNUM = iota
	GPIO10 PINNUM = iota
	GPIO11 PINNUM = iota
	GPIO12 PINNUM = iota
	GPIO13 PINNUM = iota
	GPIO14 PINNUM = iota
	GPIO15 PINNUM = iota
	GPIO16 PINNUM = iota
	GPIO17 PINNUM = iota
	GPIO18 PINNUM = iota
	GPIO19 PINNUM = iota
	GPIO20 PINNUM = iota
	GPIO21 PINNUM = iota
	GPIO22 PINNUM = iota
	GPIO23 PINNUM = iota
)

const (
	SPI_CS     PINNUM = GPIO10
	SPI_MOSI   PINNUM = GPIO11
	SPI_MISO   PINNUM = GPIO12
	SPI_CLK    PINNUM = GPIO13
	SPIEX_CS   PINNUM = GPIO20
	SPIEX_MOSI PINNUM = GPIO21
	SPIEX_MISO PINNUM = GPIO22
	SPIEX_CLK  PINNUM = GPIO23
)

const (
	MAX_GPIO_NUM      PINNUM = 23
	MAX_GPIO_MODE_NUM PINNUM = 8
	MAX_PWM_NUM       PINNUM = 5
	MAX_ADC_NUM       PINNUM = 11
)

type gpioDrive struct {
	pinNum     PINNUM
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
func OpenPin(pinNum PINNUM) (gd *gpioDrive, err error) {
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
	return &gpioDrive{
		pinNum:     pinNum,
		modeHandle: modeHandle,
		pinHandle:  pinHandle,
	}, nil
}

//
func (gd *gpioDrive) SetMode(pm PINMODE) (n int, err error) {
	defer func() {
		gd.modeHandle.Close()
	}()
	return gd.modeHandle.Write([]byte(pm))
}

func (gd *gpioDrive) SetLevel(il IOLEVEL) (n int, err error) {
	return gd.pinHandle.Write([]byte(il))
}

func (gd *gpioDrive) Write(p []byte) (n int, err error) {
	return gd.pinHandle.Write(p)
}

func (gd *gpioDrive) Read(p []byte) (n int, err error) {
	return gd.pinHandle.Read(p)

}
func (gd *gpioDrive) Close() error {
	err := gd.pinHandle.Close()
	if err != nil {
		return err
	}
	return nil
}
