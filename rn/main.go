package main 

import (
	"fmt"
	"os"
)

func giveMetheEquivalentoRomanBase(s, index int) string {
	thedigitsInIndexZero := [10]string{"I","I+I", "I+I+I", "V-I", "V", "V+I", "V+I+I", "V+I+I+I", "X-I"}
	thedigitsInIndexOne := [10]string{"X","X+X", "X+X+X", "L-X", "L", "L+X", "L+X+X", "L+X+X+X", "C-X"}
	thedigitsInIndexTwo := [10]string{"C","C+C", "C+C+C", "D-C", "D", "D+C", "D+C+C", "D+C+C+C", "M-C"}
	thedigitsInIndexThree := [10]string{"M","M+M", "M+MM"}

	if index == 0 {
		for i:=1; i<= 9; i++ {
			if i == s {
				return thedigitsInIndexZero[i-1]
			}
		}
	}
	if index == 1 {
		for i:=1; i<= 9; i++ {
			if i == s {
				return thedigitsInIndexOne[i-1]
			}
		}
	}
	if index == 2 {
		for i:=1; i<= 9; i++ {
			if i == s {
				return thedigitsInIndexTwo[i-1]
			}
		}
	}
	if index == 3 {
		for i:=1; i<= 3; i++ {
			if i == s {
				return thedigitsInIndexThree[i-1]
			}
		}
	}
	return "Cannot Convert this digit"
}

func Atoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if 0 <= int(byte(s[i]))%48 && int(byte(s[i]))%48 <= 9 {
			result = int(byte(s[i]))%48 + result*10
		} else {
			return 0
		}
	}
	return result
}

func main()  {
	arguments := os.Args[1:]
	var digitsOfTheNumbers []int
	word := ""
	bakar := ""
	if len(arguments) != 1 {
		fmt.Print("Cannot Convert this digit")
		return
	}
	nb := Atoi(arguments[0])
	if len(arguments) == 1 {
		for nb != 0 {
			digitsOfTheNumbers = append(digitsOfTheNumbers, nb%10)
			nb = nb / 10 
		}
	}
	for i := len(digitsOfTheNumbers)-1; i >= 0; i-- {
		word += giveMetheEquivalentoRomanBase(digitsOfTheNumbers[i], i)
		bakar += giveMetheEquivalentoRomanBase(digitsOfTheNumbers[i], i)
		if i == 0{
			break
		}
		bakar += "+"
	}
	fmt.Println(bakar)
	fmt.Println(word)
}
