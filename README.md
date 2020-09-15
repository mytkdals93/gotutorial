# gotutorial
GO 공식 튜토리얼
## 모듈 만들기
- go mod init <경로지정>
- path 안적으면 디렉토리 경로로 됨
- go module은 디펜던시를 자동으로 입력시켜주는데 만약에 배포된 패키지가 아니라 로컬에 있는 패키지라면 replace를 지정하여 알려줘야함 hello/go.mod참고

## 에러 핸들
- GO에서는 함수에 에러 발생시 error를 리턴, 없으면 nil 리턴
- 해당 함수를 콜한 곳에서 err처리
- error.New("message") => error리턴. error.Error() => message리턴
- greeting/greetings.go 참고

## [log 패키지](https://golang.org/src/log/log.go?s=3306:3345#L60)
- 로그 설정, prefix flag등
- - log.SetPrefix("greeting: ")
- -	log.SetFlags(log.LstdFlags | log.Lmsgprefix | log.Lshortfile)
- - flag는  참조

## [math/rand 패키지](https://golang.org/pkg/math/rand/)
- 랜덤 숫자를 뽑을 때 이용.
- - 시드 지정 rand.Seed(time.Now().UnixNano()) (프로그램 상 완전한 랜덤은 있을 수 없기 떄문에 타임을 이용해 늘 변경되는 값을 지정)
- - 
