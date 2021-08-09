package main

import (
	"fmt"
)

func main (){
	name:= "Myra"
	age:= 20

	fmt.Println("Hello, my name is ",name, "and my age is ",age)
    	TodayIsMyBirthday(&age)
	fmt.Println("After party my age is ",age)
}

func TodayIsMyBirthday(age *int) int{
	*age = 21

	fmt.Println("Today is my ",*age," birthday and I am going to party")
	return *age
}
