package puppy

import "fmt"

const Category = "pomerian"

func Barks() string {
	return "woo woooo"
}

func TagVer1() {
	fmt.Println("This is tagged version 1.0.0")
}

func TagVer2() {
	fmt.Println("This is tagged version 1.2.0")
}

func TagVer3() string {
	return "This is tagged version 1.3.0"
}

func TagVer4() string {
	return "This is tagged version 1.3.0"
}
