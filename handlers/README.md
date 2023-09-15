# annuums

- 본 Repository는 annuums에서 제작한 Become a DevOps/SRE Engineer 프로젝트 중 [Golang - 웹 서버 정보하기] 코드 저장소 입니다.

---

## Golang - 웹 서버 정복하기 Repository

### handlers

- Golang의 `net/http`는 HTTP 요청을 처리하기 위해 `Handler`라는 인터페이스를 구현해요. 인터페이스로 추상화 되어 있기 때문에, 요청을 처리하는 함수, 객체 등을 쉽게 만들 수 있어요.

- 해당 인터페이스의 메서드를 구현하여 요청을 처리해 볼까요?

```go
package handlers

import (
	"fmt"
	"net/http"
)

type HomeHandler struct {}

/**
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
func (home *HomeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    //* 실제로 어떤 URL이 요청되는지 확인해 보세요.
	fmt.Printf("Requested %s\n", req.URL.Path)

	switch req.URL.Path {
	case "/":
		indexHandler(res, req)
	case "/test":
		testHandler(res, req)
	default:
		http.NotFound(res, req)
	}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getIndex(res, req)
	case http.MethodPost:
		postIndex(res, req)
	}
}

func getIndex(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Hello, This is Get Handler! You can [GET, POST] to /home")
}

func postIndex(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Helo, This is Post Handler! You can [GET, POST] to /home")
}

...
```

- 위 코드는 `ch1`과는 다르게 `Handler`를 구현하여 사용하고 있어요. Router 형식으로 생성하여, 최초 요청을 전달받은 뒤 URI 기준으로 분기하고 있어요.
- 상위 URI인 `/home/*`의 요청에 대해 `/home` 뒷부분의 URI Segments 중, `/`과 `/test`만 처리하는 것을 확인할 수 있어요.
  - 특히 각 URL에 등록되는 기능이 많아질수록, `/home`과 `/home/test` 또한 라우터를 분기시켜 따로 관리할 수 있어요.
- 우리는 URL 구분에서 그치지 않고, 각 Handler 아래에서 요청된 `Http Methods`에 따라 처리를 나누고 있어요. 우리는 그 중 `GET`과 `POST`만 처리하고 있답니다.

- 실제로 `localhost:3000/home`과 `localhost:3000/home/test`에 GET, POST 요청을 각각 보내면,
  `Hello, This is {GET|POST} Handler! You can [GET, POST] to {/home|/home/test}` 응답을 확인할 수 있어요.
- 이처럼 라우터 패턴과 MVC 패턴을 조합한다면, 별도의 웹 서버 프레임워크 없이 웹 서버를 구현할 수 있답니다.

- `req.URL.Path`는 `req *http.Request`로 전달된 요청의 `URL`을 확인할 수 있어요.
- `req.Method`는 `req *http.Request`로 전달된 요청의 `Http Method`를 확인할 수 있어요.

#### Handler 인터페이스 다시보기

- Golang의 `net/http`는 HTTP 요청을 처리하기 위해 `Handler`라는 인터페이스를 구현해요. 인터페이스로 추상화 되어 있기 때문에, 요청을 처리하는 함수, 객체 등을 쉽게 만들 수 있어요.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

- 구조체를 정의하여 발생하는 요청들을 처리하려면, `ServeHTTP(ResponseWRiter, *Request)`를 구현하여 할당하면 된답니다.

```go
type homeHandler struct {}

func (home *homeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    ...
}

func main() {
    ...
    http.Handle("/foo", &homeHandler{})
    ...
}
```
