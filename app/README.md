# annuums

- 본 Repository는 annuums에서 제작한 Become a DevOps/SRE Engineer 프로젝트 중 [Golang - 웹 서버 정보하기] 코드 저장소 입니다.

---

## Golang - 웹 서버 정복하기 Repository

### Http Serve - 웹 서버 구동

- 웹 서버는 `net/http` 패키지 아래의 `http`를 이용해 구동할 수 있어요.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        res.WriteHeader(http.StatusOK)
        fmt.Fprint(res, "Hello, World!")
    })

    fmt.Println("Server is running on port 5000...")
    err := http.ListenAndServe(":5000", nil)

    if err != nil {
        log.Fatal(err)
    }
}
```

- 위 코드에서 `http.HandleFunc(...)` 함수는 지정된 URI로 들어오는 요청에 대해 어떻게 처리할 것인지를 명시한답니다. 자세한 내용은, `handlers/README.md`를 읽어주세요.

- 처리할 요청들을 모두 정의한 뒤, `http.ListenAndServe("주소:포트", Handler)`을 통해 서버를 실행할 수 있답니다.

#### 톺아보기

```go
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/home/", http.StripPrefix("/home", &handlers.HomeHandler{}))
    //* 이 코드를 통해 위에 둥록되지 않은 요청은 모두 404 Not Found 응답을 반환한답니다.
	mux.Handle("/", http.NotFoundHandler())

	return mux
}
```

- 우리 코드는 `ServeMux`를 만들고, `/home/`으로 시작하는 요청을 모두 `handlers` 패키지 아래의 `HomeHandler` 구조체를 이용해 처리하도록 명시했어요.
- 그리고 제일 마지막에 `/`로 시작하는 요청을 모두 `http.NotFoundHandler()`를 통해 처리하게 했어요. 이는 등록되지 않은 요청들은 모두 404 Not Found를 응답한다는 걸 뜻해요.
- 이처럼 기능에 따라 `Handler`를 여러개 생성하고, 라우터를 만들어 각자 처리하게 한다면 코드를 관리하기 쉬워지겠죠?
