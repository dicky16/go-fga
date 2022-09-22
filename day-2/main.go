package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//conditional statement
	// ConditionalStatement()
	//loop
	// Loop()
	//array
	// Array()
	MiniQuiz()
}

func MiniQuiz() {
	type Int = int
	arrInt := []Int{1, 2, 3, 4, 1, 2, 3, 1, 2, 3, 4, 1, 2, 3, 5}
	// fmt.Println(len(arrInt))
	result := make(map[int]Int)
	for _, v := range arrInt {
		result[v] += 1
		// fmt.Println(v)
	}
	result[6] += 1
	fmt.Println(result)
	for key, v := range result {
		fmt.Println(key, v)
	}
}

func Array() {
	//array
	arrString := []string{"o", "ok", "oko"}
	arrString = append(arrString, "ini tambahan")
	// fmt.Println(arrString[0])
	// fmt.Println(arrString[0:2])
	//array multidimensi
	arrMulti := [3][3][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	fmt.Println(arrMulti[1][0][0])
	//struct
	// type Struct1 struct {
	// 	arrStr := []string,

	// }
	//map
	mapString := make(map[string]int)
	mapString["k1"] = 2
	mapString["k2"] = 3
	val, ok := mapString["key3"]
	fmt.Println(val, ok)
}

func Loop() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	//infinite loop
	h := 0
	for {
		h += 10
		fmt.Println(h)
		if h == 10000 {
			break
		}
	}
	//for loop
	for i := 0; i <= 6; i++ {
		fmt.Println()
		for k := 0; k <= 6; k++ {
			if i > k {
				fmt.Print("*")
			}
		}
	}
	fmt.Println()
	for i := 0; i <= 6; i++ {
		fmt.Println()
		for k := 0; k < 5; k++ {
			fmt.Print("*")
		}
	}
}

func ConditionalStatement() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		//using if else
		cuaca := scanner.Text()
		if cuaca == "cerah" {
			fmt.Println("Yuk kita dolan")
		} else if cuaca == "mendung" {
			fmt.Println("dirumah aja, mau hujan nih")
		} else {
			fmt.Println("bebas sih mau ngapain")
		}
		//using switch case
		switch cuaca {
		case "cerah":
			fmt.Println("Yuk kita dolan")
		case "mendung":
			fmt.Println("dirumah aja, mau hujan nih")
		default:
			fmt.Println("bebas sih mau ngapain")
		}
	}
}
