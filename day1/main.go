package main

import (
	"fmt"
)

type Company struct {
	Name     string
	Employee int
	Revenue  float64
}

func main() {

	fmt.Println("--- ex0 ---")
	ex0()

	fmt.Println("--- ex1 ---")
	ex1()

	fmt.Println("--- ex1.2 ---")
	ex1_2(20, 2)

	fmt.Println("--- ex2 ---")
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	ex2(x)

	fmt.Println("--- ex3 ---")
	ex3()

	fmt.Println("--- ex3.1 ---")
	ex3_1(1, 1000, 9)

	fmt.Println("--- ex4 ---")
	ex4()

	fmt.Println("--- ex4.1 ---")
	ex4_1("ine t")

	fmt.Println("--- ex5 ---")
	ex5()

	fmt.Println("--- ex6 ---")
	ex6()

	fmt.Println("--- extar ---")
	extar()

}

// 0
func ex0() {

	i := 2

	fmt.Println("if condition")

	if i == 0 {
		fmt.Println("zero")
	} else if i == 1 {
		fmt.Println("one")
	} else if i == 2 {
		fmt.Println("two")
	} else if i == 3 {
		fmt.Println("three")
	} else {
		fmt.Println("Error")
	}
}

// 1
func ex1() {

	result := 0

	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Print(i , " ")
			result++
		}
	}

	fmt.Printf("ans = %d \n", result)

}

// 1.2
func ex1_2(num1 int, num2 int) {
	result := 1

	for i := 0; i < num2; i++ {
		result *= num1
	}

	fmt.Printf("result is %d\n", result)
}

// 2
func ex2(arr []int) {
	if len(arr) == 0 {
		return
	}

	min := arr[0]
	max := arr[0]

	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	fmt.Printf("Min is : %d || Max is : %d\n", min, max)
}

// 3
func ex3() {

	count := 0

	for i := 1; i <= 1000; i++ {
		num := i
		for num > 0 {
			if num%10 == 9 {
				count++
			}
			num /= 10
		}
	}

	fmt.Printf(" มีทั้งหมด: %d ตัว\n", count)
}

//3.1

func ex3_1(start int, end int, target int) {
	count := 0

	for i := start; i <= end; i++ {
		num := i
		for num > 0 {
			if num%10 == target {
				count++
			}
			num /= 10
		}
	}

	fmt.Printf(" มีทั้งหมด: %d ตัว\n", count)
}

// 4
func ex4() {
	result := ""
	text := "AW SOME GO!"

	for _, char := range text {
		if char != ' ' {
			result += string(char)
		}
	}

	fmt.Printf("Result: %s\n", result)
}

// 4.1
func ex4_1(text string) {
	result := ""

	for _, char := range text {
		if char != ' ' {
			result += string(char)
		}
	}

	fmt.Printf("Result: %s\n", result)
}

// 5
func ex5() {
	users := map[string]map[string]string{

		"user1": {
			"name":    "Safe",
			"age":     "23",
			"address": "chaingmai",
		},

		"user2": {
			"name":    "New",
			"age":     "23",
			"address": "chaingmai",
		},

		"user3": {
			"name":    "Bank",
			"age":     "25",
			"address": "chaingmai",
		},
	}

	fmt.Println("Result -:")
	for _, data := range users {
		age := data["age"]

		fmt.Printf("Name -: %s || Age: %s\n", data["name"], age)
		fmt.Printf("Address -: %s\n\n", data["address"])
	}
}

// 6
func ex6() {

	var c1 Company
	c1.Name = "Google"
	c1.Employee = 350000
	c1.Revenue = 100000000.50

	c2 := new(Company)
	c2.Name = "Microsoft"
	c2.Employee = 220000
	c2.Revenue = 100000000.90

	c3 := Company{
		Name:     "apple",
		Employee: 500000,
		Revenue:  100000000.50,
	}

	fmt.Printf("Name: %s || Employee: %d || unit Revenue: %.2f dollars \n", c1.Name, c1.Employee, c1.Revenue)
	fmt.Printf("Name: %s || Employee: %d || unit Revenue: %.2f dollars \n", c2.Name, c2.Employee, c2.Revenue)
	fmt.Printf("Name: %s || Employee: %d || unit Revenue: %.2f dollars \n", c3.Name, c3.Employee, c3.Revenue)

}

// extar
func extar() {
	for i := 1; i <= 6; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
