package main

import (
	"fmt"

	"github.com/soe-j/go-practice/car/models"
)

func main() {
	myCar := models.NewCar("マイカー")
	fmt.Println(myCar.GetName(), myCar.Run(50))
	fmt.Println(myCar.GetName(), myCar.Stop())

	aisha := models.NewCar("愛車")
	fmt.Println(aisha.GetName(), aisha.Run(30))
	fmt.Println(aisha.GetName(), aisha.Run(100))
	fmt.Println(aisha.GetName(), aisha.Stop())

	truck := models.NewBigCar("社用車")
	fmt.Println(truck.GetName(), truck.Run(40))
	fmt.Println(truck.GetName(), truck.Stop())
}
