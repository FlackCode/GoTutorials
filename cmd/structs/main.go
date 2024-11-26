package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	//ownerInfo owner
	//carType
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

//type owner struct {
//	name string
//}
//
//type carType struct {
//	name string
//}

type engine interface {
	milesLeft() uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func canMakeIt(e engine, miles uint8) { //can take anything as long as object has milesLeft method that returns uint8
	if miles <=e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("Fuel up!")
	}
}

//assigning function to gasEngine

func main () {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15 }
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in tank: %v \n", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
}
