//Задана строка S из малых латинских букв, требуется узнать длину наибольшей подстроки
//состоящей не более чем из двух различных символов.

//Примеры:
//"abcdef" -> 2
//"aaaaa" -> 5
//"abaccab" -> 4

package main

import "fmt"

func MaxSubstringLength(str string, k int) int {

	dict := make(map[string]int)
	length := 0
	start := 0

	// проходим скользящим окном
	for i, c := range str{

		// проверка ключей + заполнение словаря
		//fmt.Println(i ,string(c))
		_, ok := dict[string(c)]
		if !ok {
			dict[string(c)] = 1
		} else {
			dict[string(c)] += 1
		}

		// если элементов в слоаре > 2, то удаляем ключ
		for len(dict) > k {
			c := str[start]
			start += 1
			dict[string(c)] -= 1
			if dict[string(c)] == 0 {
				delete(dict, string(c))
			}
		}

		if (i - start + 1) > length {
			length = i - start + 1
		}

	}
	return length
}


func main() {
	fmt.Println(MaxSubstringLength("abcdef", 2))
	fmt.Println(MaxSubstringLength("aaaaaa", 2))
	fmt.Println(MaxSubstringLength("abaccab", 2))
	fmt.Println(MaxSubstringLength("abcdacbdefefefefef", 2))
}
