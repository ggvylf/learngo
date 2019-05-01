package main

import "fmt"

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Usb interface {
	Start()
	Stop()
}

func (p Phone) Start() {
	fmt.Println(p.Name, "starting")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "stoping")
}

//该方法只有Phone类型有
func (p Phone) Call() {
	fmt.Println(p.Name, "calling")
}

func (c Camera) Start() {
	fmt.Println(c.Name, "starting")
}

func (c Camera) Stop() {
	fmt.Println(c.Name, "stoping")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()

}

func main() {

	var usbArr [3]Usb

	usbArr[0] = Phone{"apple"}
	usbArr[1] = Camera{"sony"}
	usbArr[2] = Phone{"xiaomi"}

	var c Computer

	for _, v := range usbArr {
		c.Working(v)
		fmt.Println()
	}

}
