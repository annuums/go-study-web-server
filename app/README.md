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

- 그런데 우리 코드는, `NewHandler() http.Handler`를 직접 정의하여 http에서 실행하고 있어요.
- 그 이유는 코드를 분리하여 더 쉽게 관리하기 위함입니다.
- 또한 `http.NewServeMux()`를 통해 각 모듈에서 담당하는 라우팅을 관리할 수 있어요. 이는 User에 대한 라우팅은 UserMux, Post에 대한 라우팅은 PostMux 등을 만들어 할당 할 수 있게 됩니다.

  - 헷갈릴 수 있어요!

    - 결국 모두 `Handler`기 때문에, Mux를 사용하지 않고 Global Default Mux를 이용해 `http.Handle`등으로 곧바로 연결할 수 있어요.
    - 하지만 단일 프로젝트가 아닌, 혹은 큰 프로젝트에서 http.Handle을 통해 모든 핸들러를 등록한다면 이는 관리적인 측면 에서 활용성, 용이성이 떨어질 수 있답니다.

    ```go
    func main() {
        rootMux := root.NewHandler()
        homeMux := handlers.HomeHandler()

        //* rootMux에서 /home 으로 시작하는 uri는 homeMux가 처리하도록 합니다.
        //* StripPrefix? 요청 URL의 Path 중, 설정된 접두사를 제거하고 하위 내용을 핸들러에 전달하는 역할을 합니다.
        //* 즉, /home/hello가 요청으로 들어오면, /hello 부분만 homeMux에 전달돼요.
        rootMux.Handle("/home/", http.StripPrefix("/home/", homeMux))
        http.ListenAndServe(":5000", rootMux)
    }
    ```

    - 해당 내용은 직접 코드를 작성해 테스트하고, 차이점을 알아보세요!

#### 톺아보기

```go
func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", handlers.HomeHandler)

	return mux
}
```

- 우리 코드는 `ServeMux`를 만들고, `/home`을 처리하도록 명시했어요.
- 이 때 `Handler`로는 `handlers/home.go` 아래에 있는 `HomeHandler`를 이용합니다.

#### 종합하면...

- `/{상위_URI}/home`으로 명시되는 요청을 처리합니다. 즉, Routing을 담당하고 있어요.
