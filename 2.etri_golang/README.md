## Go 언어 설치
### 1.Install Go Language
	$ sudo apt-get update
	$ wget http://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
	$ sudo tar -xvzf go1.11.4.linux-amd64.tar.gz
	$ sudo mv go /usr/local

### 2. Setup Go Environment
	$ export GOROOT=/root/local/go
	$ export GOPATH=$HOME/go
	$ export PATH=PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
	$ vim ~/.profile

### 3. Update current shell session and verify
	$ source ~/.profile
	$ go version			// go 버전 확인
	$ go env				// go 환경 변수 확인
	$ echo $GOPATH
	$ ls $HOME/go

## GO 프로그램 빌드와 컴파일
### Go 디렉토리 
	1. bin : 소스 파일 컴파일해서 실행 파일이 생성되는 디렉토리
	2. pkg : 패키지 컴파일하여 라이브러리 파일이 생성되는 디렉토리, 아키텍처 단위
	3. src : 작성한 소스 파일과 인터넷에서 자동으로 받아온 소스 파일이 저장되는 디렉토리

### Go 파일 빌드 후 실행
	1. go build 파일이름
	2. ./파일이름

### Go 파일 컴파일
	1. go run 파일이름 

### 소스 파일 내용 정리 표준 출력으로 보여준
	1. gofmt 파일이름

## GO 프로그램 읽는 방법
### 패키지 선언(package)
	1. 모든 Go 프르그램은 반드시 패키지 선언으로 시작
	2. 패키지는 Go언어에서 코드 조직화하고 재사용 수단
	3. Go 프로그램은 두 가지 유형 : 실행 프로그램과 라이브러리
	4. 실행 가능한 애플리케이션은 터미널에서 바로 실행할 수 있는 종류의 프로그램
	5. 라이브러리는 다른 프로그램에서 사용할 수 있게 패키징된 코드의 모음

### importt "fmt"
	1. import : 다른 패키지에 포함된 코드 이용
	2. fmt 패키지(형식화[format]의 축약형) : 입력과 출력에 대한 형식화
	3. fmt : 큰 따옴표 사용하는 것은 "문자열 리터럴", "수식(expression)"에 해당
	4. Go 문자열은 길이가 정해진 문자(글자, 숫자, 기호 등) 나열

### 주석
	1. //주석은 //뒤의 모든 텍스트를 주석으로 간주
	2. /* */ 주석은 두 * 사이의 내용을 모두 주석으로 간주

### main함수 선언
```
func main() {
	fmt.Println("Hello World")
}
```
	1. 함수(function) : Go 프로그램의 기본 구성 요소
	2. 모든 함수는 func 키워드로 시작
	3. main은 함수의 이름
	4. 괄호() 안은 0개 이상의 매개변수가 포함한다.
	5. 매개 변수 뒤에는 반환형이 올 수 있다. 이때 반환형은 return 값이 있으면 그에 맞는 자료형이 들어가고, return 값이 없으면 필요가 없다.
	6. 중괄호 {} 안에는 함수를 호출하였을 때 수행한 내용을 선언한다.
	7. main 함수는 프로그램을 실행할 때 가장 먼저 호출되는 특별한 함수  

## 데이터 타입
값이 저장되는 방식을 정의한 것을 데이터 타입이라고 한다.

### 정수(숫자)
	1. Go 정수 타입 : uint8, uint16, uint32, uint64, int8, int16, int32, int64
	2. 8, 16, 32, 64는 각 타입이 사용하는 비트 수
	3. uint는 "부호가 없는 정수(unsigned integer)", int는 "부호가 있는 정수(signed integer)"
	4. 부호가 없는 정수(unsigned integer)는 양수, 0만 값으로 들어간다.
	5. byte는 uint8을 별칭으로 한다. byte는 다른 타입을 정의하는데 자주 사용된다. 
	6. rune은 uint32을 별칭으로 한다.
	7. uint, int, uintptr는 CPU와 OS 등에 의존적인 정수 타입이다. 즉, CPU의 기본 비트(32bit, 64bit)와 OS(linux, window, mac)에 따라 다르다.
	8. 일반적으로 정수를 사용할 때는  int 타입이  사용된다. uint를 사용하려면 선언을 따로 선언을 해 주어야 한다.
	9. 변수 크를 알기 위해서는 unsafe 패키지에서 sizeof() 함수를 사용한다.
 
### 부동 소수점(숫자)
	1. 소수부가 포함된 숫자(예: 1.234, 123.4, 0.00001234, 12340000)
	2. float32와 float64라는 두가지 "부동 소수점 타입"과 complex64와 complex128이라고 하는 "복소수 타입"이 존재
	3. 일반적으로 부동소수점을 사용할 때는 float64 타입이 사용된다. 
	4. 부동 소수점 수는 부정확하다. 즉, 부정확하는 것은 계산을 했을 때 원하는 결과와 실제 결과가 다를 수 있다는 뜻이다. 
 	5. 예를 들어 1.01 - 0.99를 계산하면 0.020000000000000018으로 출력된다. 0.02의 결과값을 예상했지만 실제값은 달리 나온다.
	6. 정확도를 향상시키지 위해서는 큰 부동 소수점을 사용할 수 있는 float64를 사용하는 것이 좋다.
	7. float32는 소숫점 8자리까지 값이 나온다면 float64는 소숫점 15자리까지 값이 나옴으로써 float64가 더 정확하다.

## 문자열
	1. Go 언의 문자열은 개별 바이트로 구성한다.
	2. 보통 문자열의  각 문자는 문자마다 한 바이트를 차지한다.
	3. 문자열 리터럴은 "Hello World"처럼 큰 따옴표를 이용하거나 `Hello World`처럼 역따옴표를 이용한다.
	4. 큰 따옴표로 만든 문자열은 줄바꿈 포함할 수 없고 특별한 이스케이프 문자열을 사용
	5. 즉, Windows에서 enter키를 누르면 발생하는 줄바꿈은 적용이 안되고 \n을 문자열에 적어야 줄바꿈 적용 가능
	6. 문자열 연산으로 문자열의 길이를 구하는 방법 : len("Hello World")
	7. 문자열 내의 각 문자에 접근 방법 : ("Hello World"[1])
	8. 두 문자열 하나로 합치는 방법 : ("Hello " + "World")
	9. 공백도 하나의 문자로 간주한다. 즉, "Hello World"의 문자열 길이는 10이 아니라 11이다.
	10. Go에서 인덱스는 0부터 시작한다. 즉, fmt.Printf("%c",(Hello World[1]))을 하면 'H'가 아니라 'e'가 출력된다.
	11. 맘약 fmt.Println("Hello World[1]")을 하면 'e'가 나오지 않는다. 101이 나온다. fmt.Println()은 문자가 아니라 바이트로 출력값을 표현한다.
	12. 문자열 연결은 덧셈과 같은 기호를 사용한다. "Hello " + "World"를 하면 1개의 문자열 "Hello World"이 만들어진다.

## Boolean
	1. Boolean은 참(true)와 거짓(false)를 나타내는 데 사용되는 특별한 1 바이트 정수 타입이다.
|연산자|의미|
|---|---:|	
|&&|and|
||||or|
|!|not|

	2. and 연산자 동작 방식
|수식|값|
|---|---:|	
|true&&true|true|
|true&&false|false|
|false&&true|false|
|false&&false|false|

	3. or 연산자 동작 방식
|수식|값|
|---|---:|	
|true&&true|true|
|true&&false|true|
|false&&true|true|
|false&&false|false|

	4. ! 연산자 동작 방식
|수식|값|
|---|---:|	
|!true|false|
|!false|true|

## 상수 
	1. 상수는 변수 앞에 const를 붙여서 사용한다.
	2. 상수는 문자 또는 _(밑줄 문자)로 시작해야한다. 숫자로 시작할 수 없다.
	3. 자료형을 생략하면 상수의 자료형은 대입하는 값의 자료형에 따라 결정된다.
	4. 상수 여러 개를 선언하고 초기화할 경우 변수와 값과 ,(콤마)로 구분하여 나열
	5. 상수 선언한 대로 값이 대입되어야 하며 반드시 선언한 상수의 개수와 대입할 값의 개수를 맞추어야 한다.
	6. const 키워드를 붙이지 않으면 변수가 됨으로 주의를 해야한다.

## 변수(Variable)
### 변수의 기본 사용법
```
var x string = "Hello World"
```
	1. 변수는 저장 공간이며, 구체적인 타입과 연관된 이름 소유
	2. Go의 변수는 var 키워드를 이용하여 만든 다음 변수명(x), 타입(string)을 지정한 다음 변수의 값에 "Hello World"를 할당
	3. var x string = "Hello World"와 var x string, x = "Hello World" 두가지 방식으로 사용가능하다.   
	4. var을 사용하지 않고 :=을 사용하여 초기값이 지정된 변수를 생성하는 방법도 있다. x := "Hello World" 방식으로 사용한다.
	5. Go 컴파일러가 변수에 할당하는 리터럴 값 토대로 타입을 추론할 수 있기 때문에 타입은 필요가 없다.
	6. 문자열 리터럴을 할당하고 있으므로 x에는 string 타입 지정
	7. 컴파일러는 var 문장에도 타입 추론 가능하다. var x = "Hello World"방식으로 사용한다.


### 변수 이름 짓는 법
#### 표기법 종류
##### 1. 카멜 표기법(Camel Case)
여러 단어를 연달아 사용할 때 각 단어의 첫 글자를 대문자로 적되, 맨 앞에 오는 글자는 소문자로 표기하는 것이다. 낙타의 등에 있는 혹과 같다고 하여 카멜(Camel) 표기법이라고 부른다. 예로는 camelVariable과 같은 식으로 쓴다. Java의 권장 표기법이다.

##### 2. 파스칼 표기법(Pascal Case)
이 역시 연달아 오는 단어의 모든 앞글자를 대문자로 표기하는 것은 카멜 표기법과 같지만 맨 앞에 오는 문자를 대문자로 표기한다. 카멜 표기법이 단봉낙타라면 파스칼은 쌍봉낙타라고 할 수 있다. 예로는 PowerPoint가 있다. 카멜 표기법과 파스칼 표기법을 적절하게 조합하여, 변수명이나 함수명은 카멜 표기법을 따르고 클래스명은 파스칼 표기법을 따르는 작성 스타일, 혹은 변수명은 카멜로 표기하고 함수와 클래스명은 파스칼로 표기하는 작성 스타일이 대세이다. 전자는 Java, 후자는 C++에서 주로 볼 수 있는 스타일이다. 언어에 따라 전부 카멜, 전부 파스칼로 표기할 것을 권장하기도 한다.

##### 3. 헝가리안 표기법(Hungarian Notation)
접두어에 자료형을 붙이는 것으로 strName, bBusy, szName 등이 있다. 요새는 잘 사용하지 않는 (지양하는) 스타일인데 언어의 종류가 다양한 만큼 자료형도, 문서 데이터도 다양해졌기 때문에 접두어가 의미가 없어졌다. 무엇보다 개발 중간 자료형이 바뀐다면(특히 컨테이너 계열에서 자주 일어난다. ArrayList 계열 컨테이너를 LinkedList나 다른 Set, Map, Hash 등의 컨테이너로 바꿀 때) 모든 변수명을 수정해주어야 하는 난감한 상황이 연출된다. Windows API가 이 표기법을 사용한다.

##### 4. 스네이크 표기법(Snake Case)
단어 사이에 언더바를 넣어서 표기하는 것이다. 허나 한 단어에 언더바를 붙인 _apple 등의 명칭은 C++의 장래 예약어 확장을 위해 지양되고 있다. 자세하게는, 언더바를 사용한 후 바로 대문자로 시작하는(e.g. _Apple, _Banana, _Cucumber) 식별자나 인접한 언더바(_a_apple, _b_Banana), 또는 두 개의 언더바(__Apple, __banana)는 모든 스코프에서 지양된다. 언더바로 시작하는 모든 식별자는 전역 스코프에서만 지양된다. 식별자는 변수명, 함수명, 데이터타입명, 네임스페이스명 등이 포함된다.


### 변수 유효범위
	1. 변수가 사용하도록 허용되는 공간의 범위를 해당 변수의 유효범위(scope)라고 한다.
	2. 변수는 변수는 중첩된 중괄호(블록)을 포함해서 가장 가까운 중괄호 {}(블록) 내에 존재한다.

### 여러 개의 변수 정의
Go에서는 변수 여러 개를 정의해야하는 경우가 있다. Go에서는 축약형 방법을 제공한다.
```
var (
	a = 5
	b = 10
	c = 15
)
```
 
## 제어 구조법 (if, for, switch)

### 아래 기본 for 구문 설명
	1. i <= 10("i는 10보다 작거나 같다")이라는 수식을 평가(실행) 
	2. i <= 10 의 평가가 참이면 블록 안의 문장을 실행
	3. i <= 10의 평가가 거짓이면 블록 이후 문자 실행
	4. 만약 i <= 10의 평가가 참이여서 블록 안의 문장들이 실행되고, 문자의 실행이 종료되면 다시 for 문의 조건으로 돌아간다.
	5. i = i + 1은 조건문이 무한루프에 빠지지 않고 실행될 수 있도록 해준다. 

#### 1. for 표현 방법(1)
```
i := 1
for i <= 10 {
	fmt.Println(i)
	i = i + 1
}
```

#### 2. for 표현 방법(2)
```
for i := 1; i <= 10; i++ {
	fmt.Println(i)
} 
```

### 아래 기본 if 구문 설명
	1. i <= 10 ("i는 10보다 작거나 같다")이라는 수식을 평가(실행)
	2. i <= 10의 평가가 참이면 블록 안의 문자을 실행
	3. i <= 10의 평가가 거짓이면 블록 이후 문자 실행
	4. 만약 i <= 10의 평가가 참이여서 블록 안의 문장들이 실행된다.
	5. if 문은 오직 1번만 진행된다.

#### if 표현 방법
```
i := 1
if i <= 10 {
	fmt.Println(i)
}
```


### goto 구문

#### goto 기본 설명
	1. goto는 선언된 레이블로 이동하도록 한다.
	2. goto문을 남발하여 사용하면 스파게티 코드가 되어 사용하지 않는 것이 좋다.

#### goto 표현 방법
```
a := 1
if a == 1{
	goto Error1
}

Error:
	fmt.Println("Error 1")
	return
```


### switch 구문

#### switch 기본 설명
	1. 여러 값을 비교해야 하는 경우 또는 다수의 조건식을 체크해야 하는 경우 사용한다.
	2. switch문 뒤에 하나의 변수 또는 expression을 조건식으로 지정하여 사용한다.
	3. 조건문의 결과가 case문에 해당 변수와 일치하는 경우 실행하고, 일치하지 않는 경우 다음 case를 비교한다.
	4. 복수 개의 case를 사용할 수 있다.

#### switch 표현 방법
```
i := 6
switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("four")
	case 5: fmt.Println("five")
	default: fmt.Println("Unknown")
}
``` 


## 배열, 슬라이스, 맵

### 배열
배열은 길이가 고정되어 있고, 번호가 매겨진 단일 타입 원소의 묶음이다.

#### 1. 5개의 int 타입으로 구성된 배열
```
var x [5]int
x[4] = 100
fmt.Println(x)
``` 

#### 2. 배열 사용하는 프로그램(float64 타입의 값이 들어있는 5개의 묶음 자료구조)
```
var x [5]float64
x[0] = 98
x[1] = 93
x[2] = 77
x[3] = 82
x[4] = 83

var total float64 = 0
for i:=0; i<5; i++{
	total += x[i]
}
fmt.Println("평균 : ", total/5)
```

#### 3. 배열 생성하는 짧은 방법
```
x := [5]float64 = {98, 93,, 77, 82, 83}
```

#### 4.배열 복사(copy)
배열을 대입하여 배열 전체 복사
이때 발생하는 배열의 복사는 값을 복사하지만 주소를 복사하지 않아 새로운 배열이 생성이 되는 효과를 얻을 수 있다.
주소를 복사하는 reference copy도 있는데 포인터 개념이 필요하기에 포인터 개념에 들어가서 약간의 설명을 덫붙이겠습니다.
```
a := [5]int{1, 2, 3, 4, 5}					// 배열 대입하여 배열 전체 복사
b := a
```

### 슬라이스
	1. 슬라이스(slice)는 길이가 공정되어 있지 않고 크기가 동적으로 할당되는 배열의 형태를 띈 자료구조
	2. 배열과 마찬가지로 슬라이스는 인덱스를 통해 접근이 가능하다. 
	3. 슬라이스는 배열과 달리 길이를 변경할 수 있다.
	4. var []float64(슬라이스) vs var [5]int(배열)
	5. 배열과 슬라이스의 표기법에서 유일하게 차이가 생기는 대괄호([]) 사이의 값이 있고 없고의 차이이다.
	6. 슬라이스 생성 내장 함수는 make 
	7. 슬라이스는 중요한 속성으로 len과 cap이 사용된다.
	8. 슬라이스 예시
	9. x := make([]float64, 5)						// cap = 5인 슬라이스
	10. x := make([]float64, 5, 10) 					// len = 5, cap = 10인 슬라이스
	11. y := []int{1, 2, 3, 4, 5}

### 슬라이스 함수
Go에는 append()와 copy()라는 슬라이스 내장 함수가 있다.

#### 1. append() 함수
기존 슬라이스(첫번째 인자)가져와 그 다음에 이어지는 인자 모두를 덧붙이는 식으로 새 슬라이드 생성
```
slice1 := []int{1,2,3}
slice2 := append(slice, 4, 5)
```

#### 2. copy() 함수
기존 슬라이스 내용을 새 슬라이스에 복사한다. 
```
slice1 := []int{1,2,3}
slice2 := make([]int, 2, 2}
copy(slice1, slice2)
```  


### 맵(map)
	1. 맵(map) : 순서가 없는 키-값(key-value) 쌍의 집합
	2. 다른 언어에서는 map을 연관 배열 또는 해시 테이블(hash table), 딕셔너리(dictionary)으로 표시하고 있다.
	3. 연관 키를 통해서 값을 찾는데 사용한다.

#### 맵을 선언하는 방법
**1. make함수로 키는 string, 값은 int인 map을 만드는 방법**
```
var a map[string]int = make(map[string]int)
```
**2. map을 선언할 때 map 키워드와 자료형을 생략**
```
var b = make(map[string]int)
```
**3. map을 선언할 때 var, map 키워드와 자료형을 생략**
```
c := make(map[string]int)
```

#### var, map 키워드와 자료형 생략한 map
맵(map)을 룩업 테이블(lookup table)이나 디렉토리(dictionary)로 사용하는 굉장히 보편적인 맵(map) 활용법
``` 
elements := make(map[string]string)
```

#### map안에 map 생성하기
```
elements := map[string]map[string]string{
	"one": map[string]string{
    	"name":  "Hippo",
        "number": "1991",
    },
    "two": map[string]string{
        "name":  "Mans",
        "number": "1991",
    },
}
```

#### 기본적인 map 사용하기
```
a := map[string]int{"Hippo":20, "Mans":199}
b := a
```


## 함수
**1. 함수(function) : 0개 이상의 입력 매개변수를 0개 이상의 출력 매개변수로 매핑하는 독립적인 코드 영역**
<br>**2. 함수(다른 말로 프로시저나 서브루틴으로 불린다.)는 블랙박스로 표현된다.**

### main 함수만 사용하는 경우 
```
func main() {
	x := []float64{98, 97,, 77, 82, 83}
	
	total := 0.0
	for _ , v := range x{			// _는 인덱스이고 v는 값
		total += v
		}
	fmt.Println(total / float64(len(x)))
	}
}
```


### average 함수를 사용하는 경우
```
func average(x []float64) float64{
	total := 0.0
	for _, v := range(x){
		total += v
	}
	return total / float64(len(x))
}

func main(){
	x := []float64{98, 98, 77,, 82, 83}
	fmt.Println(average(x))
}

```


### 다중 리턴값
```
func SumAndDiff(a int, b int) (int, int){
	return (a+b), (a-b)
}

func main(){
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(su, diff)
}
```


### 리턴값 생략
```
func hello() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}

func main(){
	a, _, c, _, e := hello()
	fmt.Println(a, c, e)
}
```

## 가변 함수
함수의 매개변수 개수가 정해져 있지 않고 매개변수의 개수를 유동적으로 변하도록 설정하는 함수
```
func sum(n ...int) int{
	total := 0
	for _, value := range(n){
		total += value
	}
}

func main(){
	a := sum(1,2,3,4,5)
	fmt.Println(a)
	b := sum(1,2,3)
	fmt.Println(b)
}
```


## 클로저(Closure)
Go 언어에서 함수는 클로저(Closure)로서 사용될 수도 있다. 클로저(Closure)는 함수 바깥에 있는 변수를 참조하는 함수값(function value)를 일컫는데, 이때의 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다.

### 클로저에서 자주 사용하는 용어
	- 고차함수 : 함수를 인자로 받거나 함수를 반환하는 함수
	- 클로저 : 고차함수가 반환하는 함수
	- 익명함수 : 이름 없이 임시로 쓰는 함수
	- 람다함수(람다식) : 익명함수의 또다른 이름

### 1.익명함수를 클로저로 활용(1)
add에 익명함수 func(int, int)int를 대입하여 활용
```
func main(){
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1,2))
}
```

### 2.익명함수를 클로저로 활용(2)
increment에 익명함수 func() int를 대입하여 활용
```
func main(){
	x := 0
	increment := func() int{
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
```

### 3.고차함수 : 함수를 만드는 함수(1)
outadd에 출력값을 익명함수 func(int)int로 사용하여 활용
```
func main(){
	add := outadd(10)
	fmt.Println(add(3))						//출력값 13
}

func outadd(a int) func(int) int{					// 함수의 반환을 함수로 한다. 
	return func(b int) int{						// 이때 익명함수를  클로저로 반환한다. 
		return a + b
	}
}
```  

### 4. 고차함수 : 함수를 만드는 함수(2)
makeEven에 출력값 func() uint로 사용하여 활용
```
func makeEven() func() uint{
	i := uint(0)
	return func() (value uint) {
		value = i
		i += 2
		return value
	}
}

func main(){
	Even := makeEven()
	fmt.Println(Even())
	fmt.Println(Even())
	fmt.Println(Even())
}
```


## 재귀함수
**1. 함수 자기 자신을 호출하여 프로그램을 실행**<br>
**2. 클로저와 재귀함수는 함수형 프로그래밍으로 알려진 패러다임의 근간이 되는 중요한 기법**
```
func factorial(n uint64) uint64{
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}

func main(){
	fmt.Println(factorial(5))
}
```

## 포인터

### 포인터 기본 사용법
**인자를 받는 함수 호출할 때 해당 인자는 함수로 복사**
```
//zero함수는 value copy를 통해 값을 변화하려고 시도
//실패
func zero(x int){
	x = 0
}

//ten함수는 reference copy를 통해 값을 변화하려고 시도
//성공
func ten(x *int){
	*x = 10
}

func main(){
	x := 5
	zero(x)
	fmt.Println(x)							// 값 변화 X
	ten(&x)
	fmt.Println(x)							// 값 변화 O
}
```

### 포인터의 선언 방식
**포인터형 변수 numberPtr 선언**
int형의 변수를 가리키는 포인터형 변수 numberPtr 선언. 선언시 아무 값도 넣지 않으면 nil로 초기화
```
var numberPtr *int
```

### *와 &연산자
	- *는 포인터 변수를 선언할 때 사용한다. 또한, *는 "역참조(dereference)"를 하는데 사용한다.
	- 포인터 역참조하면 해당 포인터가 가리키는 값에 접근이 가능하다.
	- *numberPtr = 0이라고 쓰면 "int값 0을 numberPtr가 참조하는 메모리 위치에 저장한다."
	- number=0이라고 쓰면 컴파일 에러가 발생한다. number은 int가 아니라 *int를 할당할 수 있다.
	- 변수 주소를 구할 때는 &연산자 사용한다.
	- &number는 *int(int에 대한 포인터)반환, number는 int이다.
	- 포인터를 이용해야 원본 변수 값을 변경 가능하다.

### new - 포인터를 구하는 내장 함수
```
func hundred(x *int){
	*x = 100
}

func main(){
	x := new(int)
	hundred(x)
	fmt.Println(*x)
}
```
#### 위 코드에 대한 분석
	- new : 인자로 타입 하나 받아 해당 타입의 값에 맞는 충분한 메모리를 할당한 후 포인터 반환한다.
	- C++에서는 new와 &가 차이가 있다. 언어마다 차이를 인식하고 세심하게 사용할 것을 주의해야 한다. 
	- Go언어는 garbage collection을 지원하는 언어로 new가 생성한 메모리에 값이 없으면 자동 삭제한다.
	- 포인터는 함수나 구조체, 메서드 등 많은 부분에서 중요하기 때문에 계속적으로 공부하는 것을 추천한다. 

## 구조체, 메서드(Method), 인터페이스(interface)

### 구조체
구조체는 여러 변수를 담을 수 있는 자료형이다.
```
type Color struct{
	red uint64
	blue uint64
	yellow uint64
}

type Person struct{
	name, nickname string
}
```

#### 1. 구조체 초기화
	- 구조체를 다양한 방법으로 인스턴스 생성
	- 기본적으로 0으로 설정된 지역 Color 변수 생성
	- 기본적으로 ""으로 설정한 지역 Person 변수 생성
	- struct의 경우 0은 각 필드에 적절한 기본값(int는 0, float의는 0.0, string은 "", 포인터는 nil) 설정
	- new 함수 사용 가능    예시) green := new(Color), Hippo := new(Person)
	- 모든 필드에 대한 메모리가 할당되고, 각 필드는 0값 또는 ""으로 설정된 후 포인터(*Color, *Person)가 반환
	- 각 필드에 값을 할당하고 싶은 경우   예시) green := Color{red : 0, blue :100, yello :50}
	- 필드가 정의된 순서를 알고 있는 경우 필드명 생략 가능   예시) green := Color{0,100,50}

#### 2. 구조체 필드(Field)
	- .(닷)연산자를 이용하여 필드에 접근이 가능하다.
	- black := Color{red : 100, blue : 100, yellow : 100}
	- black.red, black.blue, black.yellow으로 구조체의 필드에 직접 접근할 수 있다.
	- fmt.Println(black.red, black.blue, black.yello)

### 메서드(Method)
	1. 메서드(method)는 특별한 유형의 함수 형태로 코드를 대폭 축소하는데 중요한 역할을 한다.
	2. func키워드와 함수명 사이에 "수신자(receiver)"추가 - 구조체를 위해 사용되는 함수를  메서드라 한다.
	3. 수신자는 이름과 타입 있다는 점에서 매개변수와 비슷하지만 구조체를 생성한 후 .(닷)을 붙이면 매서드를 이용이 가능하다.
	4. 구조체에 종속하는 함수를 메서드라 하여, 다른 구조체에 같은 함수가 필요한 경우에는 비효율적이다.

### 1. 메서드 정의
```
type Color struct{
	red uint64
	blue uint64
	yellow uint64
}
func (c Color) PrintColor(){
	fmt.Println("red : ", c.red, "blue : ", c.blue, "yellow : ", c.yellow)
}
func (c *Color) ChangeBlack() (uint64, uint64, uint64) {
	c.red = 100
	c.blue = 100
	c.yellow = 100
	return (c.red, c.blue, c.yellow)
}
```

### 2. 구조체 값 복사 vs 구조체 주소 복사
	1.구조체 값 복사는 구조체의 값을 전체로 복사(copy)하여 전달
	2.구조체 주소 복사는 구조체의 주소를 복사(copy)하여 전달
	3.메서드에서 변수의 값을 변화시키기 위해서는 반드시 구조체 주소를 사용

#### 구조체 값 복사와 구조체 주소 복사에 대한 예
```
// 구조체
type Person struct{
	name string
	nickname string
}
//구조체 값 변경 (실패)
func (p Person) changeName(name string){
	p.name = name
}

//구조체 주소로 값 변경(성공)
func (p *Person) changeNickname(nickname string){
	p.nickname = nickname
}
```

### 인터페이스(interface)
	1. 구조체(struct)가 필드들의 집합체라면, 인터페이스(interface)는 메서드들의 집합체이다.
	2. interface는 타입(type)이 구현해야 하는 메서드 원형(prototype)들의 정의이다.
	3. 하나의 사용자 정의 타입이 인터페이스(interface)를 구현하기 위해서는 인터페이스가 갖는 모든 메서드를 구현
	4. 인터페이스(interface)는 구조체(struct)와 마찬가지로 type문으로 사용 정의

### 1. 인터페이스(interface) 구현
```
type Shape interface{
	area() float64
	perimeter() float64
}
type Rect struct{
	width, height float64
}
type Circle struct{
	radius float64
}
//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return w.width * w.height }
func (r Rect) perimeter() float64 {return 2 * (r.width + r.height)}

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 { return math.Pi * c.radius * c.radius }
func (c circle) perimeter() float64 { return 2 * math.Pi * c.radius }
```

### 2. 인터페이스(interface) 사용
	1. 인터페이스 사용하는 일반적인 예로 함수가 파라미터로 인터페이스로 받아들이는 경우 사용
	2. 함수 파라미터가 interface인 경우, 어떤 타입이든 상관없이 해당 인터페이스를 구현하면 모든 입력 파라미터로 사용 가능
	3. 아래의 예제에 showPerimeter()함수는 Shape인터페이스를 파라미터로 받아들여 사용
	4. showPerimeter()함수 내에 인터페이스를 매개변수로 사용하는 메서드들이 있다.

#### 인터페이스(interface) 사용 예시
```
func main(){
	r := Rect{width : 10.2, height : 18.9}
	c := Circle{radius : 3.2}
	showPerimeter(r, c)
}
func showPerimeter(shapes ...Shape){
	// 인터페이스 메서드(area()와 perimeter())를 s에 대입
	for _, s := range shapes {
		a := s.perimeter()
		fmt.Println(a)
	}
}
```

### 3. 인터페이스(interface) 타입
	1. 빈 인터페이스(empty interface)를 인터페이스 타입(interface type)이라고 하여 자주 접하게 된다.
	2. 여러 표준 패키지들의 함수 Prototype에서 빈 인터페이스(empty interface)가 자주 등작한다.
	3. 빈 인터페이스(empty interface)는 interface{}와 같이 표현
	4. 빈 인터페이스(empty interface)는 메서드를 전혀 갖지 않은 빈 인터페이스이다.
	5. Go언어에서 모든 타입들을 나타내기 위해서 빈 인터페이스를 사용한다.
	6. 빈 인터페이스는 어떠한 타입들도 담을 수 있는 컨테이너를 일컫는다.
	7. Go에서 빈 인터페이스는 C/C++에서는 void*이고 자바에서는 object라고 볼 수 있다.
	8. 아래의 예제는 인터페이스 타입 what에 정수 100, 문자열 "Hippo", 함수 number를 넣은 후 실행해 본다.

#### 인터페이스(interface) 타입 예시
```
func main(){
	var what interface{}
	what = 100
	PrintWhat(what)
	what = "Hippo"
	PrintWhat(what)
}

func PrintWhat(v interface{}){
	fmt.Println(v)
}
```

## Goroutine
**1. 고루틴(Goroutine)은 함수를 동시에 실행시킬 수 있도록 하는 기능을 가지고 있다.**<br>
**2. 다른 언어의 스레드 생성 방법보다 간단하게 할 수 있다. 고루틴(Goroutine)은 Go언어의 최대 장점이다**<br>
**3. 고루틴은 운영체제의 쓰레드와는 조금 다른 개념이다. 프로그램 쓰레드를 만들기 때문에 리소스를 적게 사용한다.**<br>

### Goroutine의 사용법
```
func Introduction(){
	fmt.Println("Hello, My name is Nick")

	ByeBye()
}
func ByeBye() {
	fmt.Println("Bye Bye~~~")
}
func main(){
	go Introduction()
	go ByeBye()
}
```

### Closure를 이용하여 Goroutine을 실행
```
func main(){
	runtime.GOMAXPROCS(1)
	s := "Hello, Go World"

	for i := 0; i < 10; i++{
		go func(n int){
			fmt.Println(s, n)
		}(i)
	}
}
```

## Channel

### channel 사용하기

**받기 전용 채널 선언**<br>
```
c <- chan int
```

**보내기 전용 채널 선언**<br>
```
c chan <- int
```

### channel 사용 예시
```
func sumfunction(first int, second int, sum chan int){
	sum <- first + second			// int 형 channel에 두 합 전송
}
func main(){
	ch := make(chan int)			// int형 channel 생성
	go sumfunction(15, 3, ch)		// sum goroutine을 샐행, channel 매개변수로 보냄
	
	sum := <- ch				// channel에 값을 꺼내서 sum에 대입
	fmt.Println(sum)
}
```

### 동기식 goroutine channel 사용법
**1. 송신자가 보낼 값이 들어 올 때까지 대기하고, 수신자는 채널에 값이 들어올 때가지 대기하는 형식**<br>
**2. 동기식channel을 통해 goroutine의 실행 순서를 결정할 수 있다.**
```
func main(){
	channel := make(chan bool)
	count := 10

	go func(){
		for i:=0; i< count; i++{
			channel <- true
			fmt.Println("goroutine : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i:=0; i < count; i++{
		<- channel
		fmt.Println("main func : ", i)
	}
}
```

### close()와 range() 함수를 사용하는 방법
**1. for 반복문에서 range 키워드를 사용하면 channel이 닫힐 때까지 반복하여 값을 꺼낼 수 있다.**<br>
**2. goroutine 안에 channel의 일정 값까지 보낸 뒤 close()를 사용하면, 선언 전까지만 channel에 값이 전송**<br>

```
func main(){
	channel := make(chan int)
	go func() {
		for i:=0; i < 5; i++{
			channel <- i
		
			if i == 3{	
				close(channel)
			}
		}
	}()

	for i := range channel{
		fmt.Println(i)
	}
}
``` 


## Go 출력 함수
fmt 패키지가 제공하는 표준 출력 함수 종류<br>
**func Print(a ...interface{}) (n int, err error) {} : 실행되는 cmd에 바로 출력**<br>
**func Println(a ...interface{}) (n int, err error) {} : 실행되는 cmd에 출력한 후 줄을 띄움(개행)**<br>
**func Printf(format string, a ...interface{}) (n int, err error) {} : 형식 지정자를 이용하여 값 출력**<br>


## Go 입력 함수
fmt 패키지가 제공하는 표준 입력 함수 종류<br>
**func Scan(a ...interface{}) (n int, err error) {} : 콘솔에 공백과 개행문자를 구분하여 입력**<br>
**func Scanln(a ...interface{}) (n int, err error) {} : 콘솔에 공백으로 구분하여 입력**<br>
**func Scanf(format string, a ...interface{}) (n int, err error) {} : 콘솔에 형식지정자를 이용하여 입력**<br>


## 문자열 출력 함수
fmt 패키지가 제공하는 문자열 출력 함수 종류<br>
**func Sprint(a ...interface{}) string {} : 값을 문자열로 만들어서 출력**<br>
**func Sprintln(a ...interface{}) string {} : 값을 문자열로 만든 뒤 끝에 개행 문자(\n)을 붙임**<br>
**func Sprintf(format string, a ...interface{}) string {} : 형식 지정자 이용하여 문자열을 만든 뒤 출력**<br>


## 문자열 입력 함수
fmt 패키지가 제공하는 문자열 입력 함수 종류<br>
**func Sscan(str string, a ...interface{}) (n int, err error) {} : 공백, 개행문자로 구분된 문자열 값을 받음**<br>
**func Sscanln(str string, a ...interface{}) (n int, err error) {} : 공백으로 구분된 문자열 값을 받음**<br>
**func Sscanf(str string, a ...interface{}) (n int, err error) {} : 문자열을 형식지정자로 구분 된 값을 받음**<br>

## TCP 프로토콜 사용 방법
TcpServer와 TcpClient에서 tcp 프로토콜을 이용하여 데이터를 전송하는 방법

### TCP 서버
net 패키지가 제공하는 TCP 함수들의 종류<br>
**func Listen(net, laddr string) (Listener, error) : 프로토콜, IP주소, 포트번호 설정 네트워크 연결 대기**<br>
**func (l *TCPListener) Accept() (Conn, error) : 클라이언트가 연결되면 TCP 연결(커넥션) Conn 값 리턴**<br>
**func (l *TCPListener) Close() error : TCP 연결 대기 닫음**<br>
**func (c *TCPConn) Read(b []byte) (int, error) : 받은 데이터를 읽어온다**<br>
**func (c *TCPConn) Write(b []byte) (int, error) : 데이터를 보냄**<br>
**func (c *TCPConn) Close() error : TCP 연결을 닫는다**<br>
```
func requestHandler(c net.Conn){
	data := make([]byte, 4096)
	for {
		n, err := c.Read(data)
		if err != nil{
			fmt.Println(err)
			return 
		}
		fmt.Println(string(data[:n]))
	
		_, err = c.Write(data[:n])
		if err != nil{
			fmt.Println(err)
			return
		}
	}
}

func main(){
	listen, err := net.Listen("tcp", ":8080:")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		defer conn.Close()
		go requestHandler(conn)
	}
}
```

### TCP 클라이언트
net 패키지가 제공하는 TCP 함수들의 종류<br>
**func Dial(network, address string) (Conn, error) : 프로토콜, IP주소, 포트번호 설정 서버 연결**<br>
***func (c *TCPConn) Close() error : TCP 연결 닫음**<br>
**func (c *TCPConn) Read(b []byte) (int, error) : 받은 데이터를 읽는다.**<br> 
***func (c *TCPConn) Write(b []byte) (int, error) : 데이터를 보냄**<br> 

```
func main(){
	client, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer client.Close()
	go func(c net.Conn){
		data := make([]byte, 4096)
		
		for {
			n, err := c.Read(data)
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(string(data[:n]))
			time.Sleep(2 * time.Second)
		}
	}(client)

	go func(c net.Conn){
		i := 0
		for {
			s := "Golang " + strconv.Itoa(i)
			_ , err := c.Write([]byte(s))
			if err != nil{
				fmt.Println(err)
				return
			}
			i++
			time.Sleep(2 * time.Second)
		}
	}(client)
	fmt.Scanln()
}
```
