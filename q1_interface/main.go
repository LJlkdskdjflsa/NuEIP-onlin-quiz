/*
一、物件導向繼承介面
今有車輛「汽車」和「機車」,請使用物件的繼承介面描述二者相同與差異,及二物件的執行程式碼
繼承界面規定了需要存在的function
但允許方法實作依照子類別需求變化

*/
package main

import (
	"fmt"
)

type Vehicle interface {
	Run()
}

type VehicleType int

const (
	Car VehicleType = iota
	Scooter
)

type CarClass struct{}

func NewCar() *CarClass {
	fmt.Printf("Starting a new Car\n")
	return new(CarClass)
}

func (d *CarClass) Run() {
	fmt.Println("car run")
}

type ScooterClass struct{}

func NewScooter() *ScooterClass {

	fmt.Printf("Starting a new Scooter\n")
	return new(ScooterClass)
}

func (d *ScooterClass) Run() {
	fmt.Println("scooter run")
}

func New(t VehicleType) Vehicle {
	switch t { //依據不同傳入,調用不同初始化函釋
	case Car:
		return NewCar()
	case Scooter:
		return NewScooter()
	default:
		panic("Unknown type")
	}
}

func main() {
	vehicles := make([]Vehicle, 0)

	car := New(Car)
	scooter := New(Scooter)

	vehicles = append(vehicles, car, scooter)

	for _, a := range vehicles {
		a.Run()
	}
}
