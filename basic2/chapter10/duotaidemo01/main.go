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
}

func (p *Phone) Start() {
	fmt.Println(p.Name, "starting")
}

func (c *Camera) Start() {
	fmt.Println(c.Name, "starting")
}

func main() {

	var usbArr [3]Usb

	usbArr[0] = &Phone{"apple"}
	usbArr[1] = &Camera{"sony"}
	usbArr[2] = &Phone{"xiaomi"}

	for _, v := range usbArr {
		fmt.Printf("%v\n", v)
	}
}
