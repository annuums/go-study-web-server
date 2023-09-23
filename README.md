# annuums

- 본 Repository는 annuums에서 제작한 Become a DevOps/SRE Engineer 프로젝트 중 [Golang - 웹 서버 정보하기] 코드 저장소 입니다.

---

## Golang - 웹 서버 정복하기 Repository

### 학습 목표

- `Home Router`를 분리하여 만들어 `/home`으로 발생하는 요청을 처리할 수 있다.
- Router, Handler를 분리하여 작성하는 이유를 이해한다.

#### starter.go 톺아보기

```go
func main() {
  err := log.Println("Server is running on :3000...")
  if err != nil {
    log.Fatal(err)
  }
  http.ListenAndServe(":3000", app.NewHandler())
}
```

- 우리의 웹 서버는 `app` 패키지에서 `NewHandler`를 바탕으로 웹 서버를 실행하고 있어요.
- 주소는 `localhost:3000`이랍니다.
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
