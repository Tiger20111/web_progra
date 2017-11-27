package main

import "fmt"
import "unicode"

func RemoveEven(input []int) []int {
	var output []int
 	for _, i:= range input{
		if i % 2 == 1 {
			output= append(output, i);
		}
	}
 return output
}
func PowerGenerator(agrument int) func() int {
	degree:=1;
	return func() int{
		degree = degree * agrument;
		return degree;
	}
}
func DifferentWordsCount(input string) int {
	set:=make(map[string]bool)
	input = input + " ";
	word:= "";
	count := 0;
	for _, symb := range(input) {
		if unicode.IsLetter(symb) {
			word = word + string(unicode.ToLower(symb));
		} else {
			if word !="" {
				if set[word] == false {
					count += 1;
				}
				set[word] = true;
				word = "";
			}
		}
	}
return count;
}

