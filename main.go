
package main

func main() {

	arduino := NewArduino("/dev/ttyACM0", 9600)

	testWrite(arduino)
	testRead(arduino)

	arduino.port.Close()
}

func testWrite(arduino *Arduino) {
	data := []byte{'1'}
	arduino.writeBytes(data)
}

func testRead(arduino *Arduino) {
	arduino.printlnBytes()
}







