package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

type Car struct {
	factory string
	model   string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor:%s, model:%s", p.vendor, p.model)
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "qinlin"
	h.lastName = "yang"
	interfaces = append(interfaces, h)
	c := &Car{}
	c.factory = "benz"
	c.model = "s"
	p := Plane{}
	p.vendor = "vendor"
	p.model = "model"
	interfaces = append(interfaces, c, p)
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}

}
