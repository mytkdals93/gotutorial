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

## Go 테스트 
- Go는 빌트인으로 테스팅을 지원한다. (네이밍 컨벤션, testing 패키지, gotest 커맨드로 빠르게 작성하고 실행시킬 수 있음)
- Ending a file's name with _test.go tells the go test command that this file contains test functions.
- _test.go로 이름이 끝나면 해당 파일에 test function이 포함되어 있음을 알 수 있다.
