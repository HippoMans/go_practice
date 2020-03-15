##Go 언어 설치
###1.Install Go Language
	$ sudo apt-get update
	$ wget http://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
	$ sudo tar -xvzf go1.11.4.linux-amd64.tar.gz
	$ sudo mv go /usr/local

###2. Setup Go Environment
	$ export GOROOT=/root/local/go
	$ export GOPATH=$HOME/go
	$ export PATH=PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
	$ vim ~/.profile

###3. Update current shell session and verify
	$ source ~/.profile
	$ go version			// go 버전 확인
	$ go env				// go 환경 변수 확인
	$ echo $GOPATH
	$ ls $HOME/go

##GO 프로그램 빌드와 컴파일
###Go 디렉토리 
	1. bin : 소스 파일 컴파일해서 실행 파일이 생성되는 디렉토리
	2. pkg : 패키지 컴파일하여 라이브러리 파일이 생성되는 디렉토리, 아키텍처 단위
	3. src : 작성한 소스 파일과 인터넷에서 자동으로 받아온 소스 파일이 저장되는 디렉토리

###Go 파일 빌드 후 실행
	1. go build 파일이름
	2. ./파일이름

###Go 파일 컴파일
	1. go run 파일이름 

###소스 파일 내용 정리 표준 출력으로 보여준
	1. gofmt 파일이름

##GO 프로그램 읽는 방법
###패키지 선언(package)
	1. 모든 Go 프르그램은 반드시 패키지 선언으로 시작
	2. 패키지는 Go언어에서 코드 조직화하고 재사용 수단
	3. Go 프로그램은 두 가지 유형 : 실행 프로그램과 라이브러리
	4. 실행 가능한 애플리케이션은 터미널에서 바로 실행할 수 있는 종류의 프로그램
	5. 라이브러리는 다른 프로그램에서 사용할 수 있게 패키징된 코드의 모음

###importt "fmt"
	1. import : 다른 패키지에 포함된 코드 이용
	2. fmt 패키지(형식화[format]의 축약형) : 입력과 출력에 대한 형식화
	3. fmt : 큰 따옴표 사용하는 것은 "문자열 리터럴", "수식(expression)"에 해당
	4. Go 문자열은 길이가 정해진 문자(글자, 숫자, 기호 등) 나열

###주석
	1. //주석은 //뒤의 모든 텍스트를 주석으로 간주
	2. /* */ 주석은 두 * 사이의 내용을 모두 주석으로 간주

###main함수 선언
func main() {
	fmt.Println("Hello World")
}
	1. 함수(function) : Go 프로그램의 기본 구성 요소
	2. 모든 함수는 func 키워드로 시작
	3. main은 함수의 이름
	4. 괄호() 안은 0개 이상의 매개변수가 포함한다.
	5. 매개 변수 뒤에는 반환형이 올 수 있다. 이때 반환형은 return 값이 있으면 그에 맞는 자료형이 들어가고, return 값이 없으면 필요가 없다.
	6. 중괄호 {} 안에는 함수를 호출하였을 때 수행한 내용을 선언한다.
	7. main 함수는 프로그램을 실행할 때 가장 먼저 호출되는 특별한 함수  

##데이터 타입
값이 저장되는 방식을 정의한 것을 데이터 타입이라고 한다.

###정수(숫자)
	1. Go 정수 타입 : uint8, uint16, uint32, uint64, int8, int16, int32, int64
	2. 8, 16, 32, 64는 각 타입이 사용하는 비트 수
	3. uint는 "부호가 없는 정수(unsigned integer)", int는 "부호가 있는 정수(signed integer)"
	4. 부호가 없는 정수(unsigned integer)는 양수, 0만 값으로 들어간다.
	5. byte는 uint8을 별칭으로 한다. byte는 다른 타입을 정의하는데 자주 사용된다. 
	6. rune은 uint32을 별칭으로 한다.
	7. uint, int, uintptr는 CPU와 OS 등에 의존적인 정수 타입이다. 즉, CPU의 기본 비트(32bit, 64bit)와 OS(linux, window, mac)에 따라 다르다.
	8. 일반적으로 정수를 사용할 때는  int 타입이  사용된다. uint를 사용하려면 선언을 따로 선언을 해 주어야 한다.
 
###부동 소수점(숫자)
	1. 소수부가 포함된 숫자(예: 1.234, 123.4, 0.00001234, 12340000)
	2. float32와 float64라는 두가지 "부동 소수점 타입"과 complex64와 complex128이라고 하는 "복소수 타입"이 존재
	3. 일반적으로 부동소수점을 사용할 때는 float64 타입이 사용된다. 
	4. 부동 소수점 수는 부정확하다. 즉, 부정확하는 것은 계산을 했을 때 원하는 결과와 실제 결과가 다를 수 있다는 뜻이다. 
 	5. 예를 들어 1.01 - 0.99를 계산하면 0.020000000000000018으로 출력된다. 0.02의 결과값을 예상했지만 실제값은 달리 나온다.
	6. 정확도를 향상시키지 위해서는 큰 부동 소수점을 사용할 수 있는 float64를 사용하는 것이 좋다.
	7. float32는 소숫점 8자리까지 값이 나온다면 float64는 소숫점 15자리까지 값이 나옴으로써 float64가 더 정확하다.








