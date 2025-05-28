# annuums

- 본 Repository는 annuums에서 제작한 Become a DevOps/SRE Engineer 프로젝트 중 [Golang - 웹 서버 정보하기] 코드 저장소 입니다.

---

## Golang - 웹 서버 정복하기 Repository

### 학습 목표

- 웹 요청에서 필요한 정보를 획득할 수 있다.
- 학생 관리를 위한 간단한 서비스를 만들고, 이를 웹 서버에서 사용할 수 있다.
- 서비스를 설계할 때 구조에 대해 생각해 볼 수 있다.

### 학생 관리 요구 사항

- 학생 등록, 조회, 삭제가 가능해야 한다.

### 학습 순서

- 학생 관리 요구 사항에 대한 인터페이스를 정의한다.
- 학생 관리 인터페이스를 구현한다.
- 학생 관리 인터페이스를 웹 서버에서 사용할 수 있도록 의존성을 관리한다.

#### starter.go 톺아보기

```go
func main() {

  log.Println("Server is running on :5050...")
  err := http.ListenAndServe(":5050", app.NewHandler())
  if err != nil {

    log.Fatalf("Failed to start server: %v", err)
  }
}
```

- 우리의 웹 서버는 `app` 패키지에서 `NewHandler`를 바탕으로 웹 서버를 실행하고 있어요.
- 주소는 `localhost:5050`이랍니다.
- 만약 서버 실행에 실패한다면, 에러를 출력하고 종료돼요.

#### 서버 실행

- 빌드 후 실행하기

```shell
$ go build
$ ./go-study-web-server
```

- 그냥 실행하기

```shell
$ go run .
```

- 테스트 서버 실행하기

```shell
# $GOPATH/bin 환경변수 등록 했다면
$ goconvey
# 안했다면
$ $GOPATH/bin/goconvey
```

### 설치 패키지

- [goconvey](https://github.com/smartystreets/goconvey)
  - `go get github.com/smartystreets/goconvey`
- [testify](https://github.com/stretchr/testify)
  - `go get github.com/stretchr/testify`
- [godotenv](github.com/joho/godotenv)
  - `go get github.com/joho/godotenv`

### net/http

- [net/http](https://pkg.go.dev/net/http)

#### Written by

- dev.whoan(싹난 감자) in Annuums
  - [Github](https://github.com/dev-whoan)
  - dev.whoan@gmail.com
