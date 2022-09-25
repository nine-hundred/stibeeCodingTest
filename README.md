# stibeeCodingTest

## Description
스티비 백엔드 엔지니어 채용 과제 - “창과 방패”

## Usage
1. Go 서버 실행
````
// go 설치 후,
$ go run main.go
````

2. 스테이지 시작
````
// n번째 스테이지 시작
// ex: 1번째 스테이지를 시작하려면, 
// $ make start-stage1
$ make start-stage(n)
````

## What I did (on Stage3)
 1. Stage1과 Stage2는 쉽게 통과
 2. Stage3를 시작하면 서버가 죽길래 dmesg로 확인하니 OOM 발생
    - swap 메모리 추가로 해결
 3. Stage3를 n번 시도하면 어쩌다 한번 pass고 대부분 timeout 발생(대략 40 ~ 50초)
    - Stage3의 input 크기와 사용되는 문자들 확인
    - Stage3에선 input의 종류가 다양하지 않아 map과 slice로 아스키 값 +1 을 대치하려 하였으나, 시간 차이가 크지 않아 관둠
    - 슬라이스의 사이즈가 동적으로 추가되는 것에 리소스가 낭비한다고 판단 후, 슬라이스 사이즈 고정 후 index에 따라 문자 입력
    - bufio(buffer)와 string.Builer를 사용하여 입출력 시간 줄이려고 함
    - 8 ~ 30초로 줄임(10초 안에 안정적으로 들어오진 않음)
