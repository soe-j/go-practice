package main

import (
	"fmt"

	"github.com/soe-j/go-practice/car/models"
)

func main() {
	myCar := &models.Car{Name: "マイカー"}
	fmt.Println(myCar.Name, myCar.Run(50))
	fmt.Println(myCar.Name, myCar.Stop())

	aisha := &models.Car{Name: "愛車"}
	fmt.Println(aisha.Name, aisha.Run(30))
	fmt.Println(aisha.Name, aisha.Run(100))
	fmt.Println(aisha.Name, aisha.Stop())
}
