package main

import (
	"io"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"fmt"

)

type Arduino struct {
	port io.ReadWriteCloser
	portName  string
	baudRate uint
}

func NewArduino(portName string, baudRate uint) *Arduino {
	port := openSerialPort(portName, baudRate)
	a := &Arduino{port,portName, baudRate}
	return a
}

func openSerialPort(portName string, baudRate uint) (port io.ReadWriteCloser) {
	// Set up options.
	options := serial.OpenOptions{
		PortName: portName,
		BaudRate: baudRate,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
	}

	// Open the serial port.
	port, err := serial.Open(options)

	// If any error occurs
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	return port
}

func (a Arduino) readBytes() (returnBytes []byte) {
	for {
		buf := make([]byte, 32)
		n, err := a.port.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from serial port: ", err)
			}
		} else {
			buf = buf[:n]
			return buf
		}
	}
}

func (a Arduino) writeBytes(data []byte) {
	// Write array bytes to serial port
	n, err := a.port.Write(data)

	// If any error occurs
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	} else {
		// Show success message
		fmt.Print("Wroted data: ")
		fmt.Println(data)
		fmt.Println("Data size:", n ,"bytes")
	}
}

func (a Arduino) printlnBytes() {
	for {
		buf := make([]byte, 32)
		n, err := a.port.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from serial port: ", err)
			}
		} else {
			buf = buf[:n]
			fmt.Println("Rx: ", buf)
			//fmt.Println("Rx: ", hex.EncodeToString(buf))
			//time.Sleep(10 * time.Millisecond)
		}
	}
}
