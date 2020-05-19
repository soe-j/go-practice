package main

import (
	"fmt"
	"strconv"
)

// ICar 車インタフェース
type ICar interface {
	run(int) string
	stop() string
}

// Car 車構造
type Car struct {
	name  string
	speed int
}

func (car *Car) run(speed int) string {
	car.speed = speed
	return strconv.Itoa(speed) + "kmで走行"
}
func (car *Car) stop() string {
	car.speed = 0
	return "停止"
}

func main() {
	myCar := &Car{name: "マイカー", speed: 0}
	fmt.Println(myCar.name, myCar.stop())

	aisha := &Car{name: "愛車", speed: 0}
	fmt.Println(aisha.name, aisha.run(30))
	fmt.Println(aisha.name, aisha.run(100))
	fmt.Println(myCar.name, myCar.run(50))
	fmt.Println(aisha.name, aisha.stop())
}
