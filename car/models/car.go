package models

import "strconv"

// Car 車インタフェース
type Car interface {
	GetName() string
	Run(int) string
	Stop() string
}

// car 車構造
type car struct {
	name  string
	speed int
}

// NewCar 新しい車を作る
func NewCar(name string) Car {
	return &car{name: name, speed: 0}
}

// GetName 名前を取得する
func (car *car) GetName() string {
	return car.name
}

// Run 走る
func (car *car) Run(speed int) string {
	car.speed = speed
	return strconv.Itoa(speed) + "kmで走行"
}

// Stop 止まる
func (car *car) Stop() string {
	car.speed = 0
	return "停止"
}
