// 컴파일러의 진입 파일 이름
// 보통은 main.go 파일로 만든다
package testing

import (
	"fmt"
	"strings"
)

// package 뒤에 있는 이름의 함수를 실행시켜야 하기 때문에
// 반드시 같은 이름의 함수를 생성해야한다.

func testFunction(name string) (length int, uppercase string) {
	// defer: 함수가 끝났을떄 추가적인 작업을 시킨다 -> next() ??
	defer fmt.Println("I'm done")

	length = 101010101 // 변수 수정
	uppercase = strings.ToUpper(name)

	// 리턴만 작성해도 선언식에서 리턴되는 값을 자동으로 리턴한다.
	// 하지만 정확하게 무엇이 리턴되는지 작성하는 것이 좋다
	fmt.Println(length, uppercase)
	return length, uppercase
}

func constAndVar() (string, string, string) {
	const name string = "const 상수 선언"

	var varName = "var 변수 선언"

	// :=는 go가 타입을 찾는다
	// 축약형은 함수 내부에서만 사용 가능.
	shortcutVar := "shortcut var 선언"

	return name, varName, shortcutVar
}

func testDefaultReturnType(arg string) (length int, upperCase string) {
	length = 1001001
	upperCase = strings.ToUpper(arg)

	// return length, upperCase
	return
	// 변수 리턴을 안해도 함수 리턴에 선언된 값들은 리턴된다.
	// 하지만 명확하게 어떤 값을 리턴해주는지 명시하는 것이 좋다
}

func deferTesting() (a string, b string) {
	defer fmt.Println("I'm third")

	a = "I'm first"

	b = "I'm second"

	fmt.Println("deferTesting Function End")

	return a, b
}

func loopTesting(numbers ...int) (int, int) {
	// fmt.Println(numbers) // [1 2 3 4 5]
	countRange := 0
	countOriginForLoop := 0

	// range로 for 선언 시 index, number를 준다
	for _, number := range numbers {
		countRange += number
	}

	// JS 방식 for loop
	for i := 0; i < len(numbers)-1; i++ {
		countOriginForLoop += numbers[i]
	}

	return countRange, countOriginForLoop
}

func ifElseTesting(age int) bool {
	// 1. 조건식을 ()없이 작성

	// 2. variable expresstion
	// 조건식 이전에 뱐수 생성 가능
	// 조건식에서만 사용되는 변수로 선언하기 때문에 가독성에 좋다
	// 변수의 브라켓이 조건식 내부로 들어간다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		fmt.Println(koreanAge)
	}
	// fmt.Println(koreanAge) // 없음

	return true
}

func LowLevelProgramming() {
	// & = 주소 보기
	// * = 주소의 값 보기 Pointer

	a := 2
	b := &a // a의 주소를 b에 저장

	// 값이 저잔된 주소 보기
	fmt.Println(b)  // a의 주소
	fmt.Println(&b) // b의 주소
	fmt.Println(*b) // b의 값(주소)이 가리키는 a의 값(2)

	a = 10 // a의 값 변경

	// a의 주소는 동일
	fmt.Println(&b) // b가 가진 a의 주소는 변경되지 않는다
	fmt.Println(*b) // 10 출력

	// a의 주소를 가진 b를 통해 a의 값 변경
	*b = 20
	fmt.Println(a) // 20 출력

	return
}

func goArrayTesting() {
	// 길이가 정적인 Array 선언
	limitedArray := [3]string{"Name1", "Name2"} // 4 = length
	fmt.Println(limitedArray)
	// Array 요소 추가
	limitedArray[2] = "추가" // 2 = index

	// --------------------------------------- //

	// 길이가 동적인 Array 선언(slice)
	unLimitedArray := []string{"slice1", "slice2"}

	// slice에서 요소 추가하기 append(array, args)
	// append 메서드는 기존값을 수정하지 않고 새로운 값을 리턴한다.
	newUnlimitedArray := append(unLimitedArray, "added elements slice3")
	fmt.Println(unLimitedArray)
	fmt.Println(newUnlimitedArray)

	return
}

func mapTesting() {

	// go에서는 생성자 함수가 없기 떄문에 수동으로 생성해줘야 한다.
	type person struct {
		name  string
		age   int
		phone string
		food  []string
	}

	foods := []string{"a", "b"}
	chanyangClone := person{name: "chanyang", age: 12, phone: "fd", food: foods}

	chanyang := map[string]string{"name": "chanyang", "age": "12"}
	// map[keyType] valueType{"key":"value", ...}
	fmt.Println(chanyang)
	fmt.Println(chanyangClone)

	// map의 반복문
	for key, value := range chanyang {
		fmt.Println(key, value)
	}

}

func testing() {
	// 변수 선언 방법
	constName, varName, shortcutVar := constAndVar()
	fmt.Println(constName, varName, shortcutVar)

	// 리턴 타입 설정 및 nicked return 연습
	arg := "lowercase"
	argLength, upperCase := testDefaultReturnType(arg)
	fmt.Println(argLength, upperCase)

	// defer => 함수가 끝난 후에 동작을 설정한다.
	a, b := deferTesting()
	fmt.Println(a, b)

	// golang의 반복문은 for 하나다
	countRange, countOriginForLoop := loopTesting(1, 2, 3, 4, 5)
	fmt.Println(countRange, countOriginForLoop)

	// if else 조건문 확이
	age := 17
	passAge := ifElseTesting(age)
	fmt.Println(passAge)

	// switch 확인

	// Go만 가지고 있는 특징 Low-Level Programming
	LowLevelProgramming()

	//go의 Array
	goArrayTesting()

	// maps
	mapTesting()
}
