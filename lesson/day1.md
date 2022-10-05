# 클라우드 서버 접속


정보확인
```
$> cat /etc/os-release
$> apt-get update
$> apt-get upgrade
```

# go lang 소개
- https://go.dev/play/
- brace 를 라인 앞에 달면 컴파일 오류남

# top 명령어는 컴퓨터의 상태를 보여줌
```sh
$> top
```
- 1 을 누르면 cpu의 core들의 상태를 보여줌



# GO LANG
- 강타입언어 오토캐스팅(data preserving) 지원안함 (오류남
- 데이터타입 생략가능
- 덕타임핑 인터페이스 
- 슬라이스 : 배열은 배열인데 크기가 정해져있지않은배열. 동적배열. 포인터로 연결된.

# HelloWorld.go
- 컴파일 & 실행

# GO BUILD
```sh
> go run hello.go
> go build hello.go
> ./hello.exe
```

# UBUNTU GO 설치
- 이렇게 했더니 1.13이 설치됨
```sh 
$> sudo apt-get install golang.go
$> sudo apt-get purge golang.go
```

## 최신 버전 설치하는법 
- download [go1.19.2.linux-amd64.tar.gz](https://go.dev/dl/go1.19.2.linux-amd64.tar.gz)
- https://go.dev/doc/install 
```
$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
$ vi ~/.profile
$ mkdir ~/hpc-go/go
```
- .profile 에 환경변수 추가
```sh
export GOROOT=/usr/local/go
export GOPATH=$HOME/hpc-go/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
- vscode 설치
```sh
$ sudo dpkg -i code_1.71.2~~
```

# GIT 설치
```
$> sudo apt-get install git
```

# 요약
- Go는 정적 컴파일 언어
- Go는 강타입 언어
- Go는 가비지컬렉터를 제공
- "12"+12의 결과는?

# 4. 변수
- 카멜케이스 변수명에 여러단어가 이어지면 두번째 단어부터는 대문자로 이어집니다.
- 밑줄은 일반적으로 사용하지 않습니다. 

### 기본형 데이터타입
|자료형|설명|
|-----|--------------------|
|uint8| 1바이트 부호없는 정수|
|uint16||
|uint32||
|uint64||
|int8||
|int16||
|int32||
|int64||
|float32||
|float64||
|complex64|8바이트 복소수(진수,가수)|
|complex128|16바이트 복소수(진수,가수)|
|byte|uint8의 별칭. 1바이트 데이터를 나타낼 떄 사용|
|rune|int32의 별칭. UTF-8로 문자 하나를 나타낼 때 사용|
|int|32비트 컴퓨터에서는 int32. 64 비트 컴퓨터에서는 int64와 같음|
|uint|32비트 컴퓨터에서는 uint32. 64비트 컴퓨터에서는 uint64와 같음|
### 그외 타입
- boolean
- 문자열
- 배열 : 같은 타입의 요소들로 이루어진 연속된 메모리 공간을 나타내는 자료구조입니다.
- 슬라이스 : Go 언어에서 제공하는 가변길이의 배열을 말합니다. 배열은 고정길이로써 한 번 길이가 정해지면 늘리거나 줄일 수 없는 반면 슬라이스는 길이를 늘리거나 줄일 수 있습니다.
- 구조체 :
- 포인터 : 메모리 주소를 값으로 갖는 타입입니다. 포인터를 이용해서 같은 메모리 공간을 가리키는 여러 변수를 만들수 있습니다.
- 함수타입 : 함수를 가리키는 타입입니다. 다른말로 함수 포인터 라고 말합니다. 사용할 함수를 동적 으로 바꿀대 유용합니다.
- 인터페이스 : 메서드 정의의 집합입니다.
- 맵 : 키와 값을 갖는 데이터를 저장해둔 자료구조입니다. 키를 사용해 데이터를 찾는 데 특화된 자료구조입니다. 쉽게 전화번호부나 사전을 생각하시면 됩니다.
- 채널 : 멀티스레드 환경에 특화된 큐 형태 자료구조입니다. 앞으로 각 타입에 대해 자세히 배우게 되니깐 이런게 있구나 정도로알고 넘어가도 됩니다.

### 선언대입문 := 
- 선언대입문이란 말그대로 선언봐 대입을 한꺼번에 하는 구문입니다. var 키워드와 타입을 생략해 변수를 선언 할 수 있습니다.
```go
var b= 3.1415
c:=365
s:="hello world"        //
```

## 타입변환
- type conversion
```go
a:=3    //int
var b float54 = 3.5

var c int = b       //Error - float64 변수를 int 에 대입불가
d := a * b          //Error - 다른 타입인 int 변수와 flaot64 연산 불가

var e int64 = 7 
f := a * e          //Error - a 는 int 타입,  e는 int64 타입으로 같은 정수값이지만 타입이 달라서 연산 불가
```

## 실수 : [IEEE 754](https://ko.wikipedia.org/wiki/IEEE_754)

# 5. fmt  패키지를 이용한 텍스트 입출력
## 5.1 표춘 입출력
### 5.1.1 fmt 패키지
- Print()
- Println()
- Printf() 

## 5.2 표준 입력
- Scan() 표준입력에서 값을 입력받습니다.
- Scanf() 표준입력에서 서식 형태로 값을 입력받습니다.
- Scanln() 표준 입력에서 한줄을 읽어서 값을 입력받습니다.

### 5.2.2 Scan ()
- Scan() 함수는 값을 채워넣을 변수들의 메모리 주소를 인수로 받습니다. 한번에 여러 값을 입력 받을때는 변수 사이를 공란을 두어 구분합니다(enter 키도 공란으로 인식합니다).
```go
func Scan(a ...interface{}) (n int, err error)
```
함수 반환값은 성공적으로 입력한 값 개수와 입력 실패시 에러를 반환합니다.


### 5.2.3 Scanf()
Scanf() 함수는 서식에 맞춘 입력을 받습니다.

### 5.2.4 Scanln()
Scanln() 함수는 한줄을 입력받아서 인수로 드어온 변수메모리 주ㅗ에 갑을 채워줍니다.


# 6. 연산자
## 6.1 산술연산자
|구분|연산자|연산|피연산자 타입|
|-----|--|--|--|
|사칙연산과나머지|+|덧셈|정수,실수,복소수,문자열|
|    |-|뺄셈|정수,실수,복소수|

## 6.2 비교연산자
|연산자|설명|반환값|
|-----|----|-----|
|==|같다|참이면true<br/>거짓이면false|
|!=|다르다||
|<|작다||
