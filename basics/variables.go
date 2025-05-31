package basics

import "fmt"

var naam = "Super"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(naam)

	doeNaam()

	fmt.Println(naam)
}

func doeNaam() {
	naam := "Sub"
	fmt.Println(naam)

	{
		naam := "SubSub"
		fmt.Println(naam)
	}

	fmt.Println(naam)
}