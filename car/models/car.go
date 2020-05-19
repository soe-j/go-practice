package models

import "strconv"

// ICar 車インタフェース
type ICar interface {
	GetName() string
	Run(int) string
	Stop() string
}

// Car 車構造
type Car struct {
	name  string
	speed int
}

// NewCar 新しい車を作る
func NewCar(name string) ICar {
	return &Car{name: name, speed: 0}
}

// GetName 名前を取得する
func (car *Car) GetName() string {
	return car.name
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
