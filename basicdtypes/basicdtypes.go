package basicdtypes

import (
	"fmt"
)

func Basicdtypes() {
	var intNum0 int = 4
	var floatNum float64 = 3835348732916297342669289.3748466566
	fmt.Printf("Int is %d \n", intNum0)
	fmt.Printf("Float is %f \n", floatNum)

	var intNum1 int = 3

	fmt.Println(intNum0 / intNum1)

	fmt.Println(float64(intNum0 / intNum1))
	fmt.Println(float64(intNum0) / float64(intNum1))

	//type inference using := symbol

	name := "Piyush"
	print(name)

}
