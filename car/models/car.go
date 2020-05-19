package models

import "strconv"

// ICar 車インタフェース
type ICar interface {
	run(int) string
	stop() string
}

// Car 車構造
type Car struct {
	Name  string
	speed int
}

// Run 走る
func (car *Car) Run(speed int) string {
	car.speed = speed
	return strconv.Itoa(speed) + "kmで走行"
}

// Stop 止まる
func (car *Car) Stop() string {
	car.speed = 0
	return "停止"
}
