package main

import "fmt"

func main() {

	// 타입有
	const age int = 10
	const name string = "sky"
	fmt.Println("타입有 : ", age, name)
	/* 컴파일 에러
	const score int				// 대입값이 없으면 상수로 사용 불가
	age = 20					// 타입이 없다면 const를 붙여줘야 한다.
	name = "Hippo"				// 타입이 없다면 const를 붙여줘야 한다.
	*/

	//타입無
	const height = 190      // 타입이 없어도 const 선언한 후 값을 대입하면 대입값에 따라 타입이 선언된다. (int)
	const nickname = "Mans" // 타입이 없어도 const 선언한 후 값을 대입하면 대입값에 따라 타입이 선언된다. (string)

	/*컴파일 에러
	const address 				// 타입도 없고 값도 없으면 컴파일 에러 발생
	*/

	fmt.Println("타입無 : ", height, nickname)

	/*	복수의 const 값 대입
		cosnt 상수명1, 상수명2 자료형 = 초기값1, 초기값2
		const 상수명1, 상수명2 = 초기값1, 초기값2
	*/
	const x, y int = 30, 50
	const email, phone = "hello@naver", 010

	fmt.Println("(1) double const : ", x, y)
	fmt.Println("(2) double const : ", email, phone)
}
