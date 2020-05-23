package models

type bigCar struct {
	weight int
	car
}

// NewBigCar 新しいデカい車を作る
func NewBigCar(name string) Car {
	// bigcar := bigCar{weight: 10000}
	// bigcar.name = name
	// bigcar.speed = 0
	// return &bigcar

	// or

	return &bigCar{weight: 10000, car: car{name: name, speed: 0}}
	//                                 ^同じpackage内なら呼べる
}
