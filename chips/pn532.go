// Copyright 2014 jet tsang. All rights reserved.
// license that can be found in the LICENSE file.

package pn532

import (
	"github.com/jetsanix/golang-pcduino"
	"io"
)

const (
	PN532_PREAMBLE   byte = 0x00
	PN532_STARTCODE1 byte = 0x00
	PN532_STARTCODE2 byte = 0xFF
	PN532_POSTAMBLE  byte = 0x00

	PN532_HOSTTOPN532 byte = 0xD4

	PN532_FIRMWAREVERSION     byte = 0x02
	PN532_GETGENERALSTATUS    byte = 0x04
	PN532_SAMCONFIGURATION    byte = 0x14
	PN532_INLISTPASSIVETARGET byte = 0x4A
	PN532_INDATAEXCHANGE      byte = 0x40
	PN532_MIFARE_READ         byte = 0x30
	PN532_MIFARE_WRITE        byte = 0xA0

	PN532_AUTH_WITH_KEYA byte = 0x60
	PN532_AUTH_WITH_KEYB byte = 0x61

	PN532_WAKEUP byte = 0x55

	PN532_MIFARE_ISO14443A byte = 0x0

	KEY_A byte = 1
	KEY_B byte = 2
)

type NFCDrive struct {
	NFCHandle io.ReadWriteCloser
}

func OpenPN532() (nd *NFCDrive, err error) {
	uartDevice, err := pcduino.SetupUART(115200)
	return &NFCDrive{
		NFCHandle: uartDevice,
	}, err
}

func (nd *NFCDrive) PN532UartWakeup() []byte {
	var wakeupBuffer = []byte{0x55, 0x55, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x03, 0xfd, 0xd4, 0x14, 0x01, 0x17, 0x00} //wake up NFC module

	nd.NFCHandle.Write(wakeupBuffer)
	buf := make([]byte, 15)
	nd.NFCHandle.Read(buf)
	return buf
}
