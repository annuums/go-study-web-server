# annuums

- 본 Repository는 annuums에서 제작한 Become a DevOps/SRE Engineer 프로젝트 중 [Golang - 웹 서버 정보하기] 코드 저장소 입니다.

---

## Golang - 웹 서버 정복하기 Repository

### handlers

- `ch2`와는 다르게 `Handler`가 처리하는 메서드 `Handles`가 `HomeHandler`에 할당되었어요. 이는 `Handler`가 `Router`와 다른 패키지에 존재하고, 객체 또한 구분되기 때문에 외부에 노출되어야 하기 때문이에요.
- 물론 객체에 메서드를 정의하지 않고, `func Handles(res ...)` 등으로 외부에 노출시킬 수 있지만, 이는 코드의 관리를 용이하게 한답니다. 쉽게 객체 지향 프로그래밍에서 메서드는 객체에 존재하는 것을 떠올리면 돼요.
- 반면에, getIndex와 postIndex는 내부적으로 처리되고, 외부에 노출 될 필요가 없기 때문에 숨겨져 있답니다.

```go
package handlers

import (
	"fmt"
	"net/http"
)

type HomeHandler struct{}

func (handler *HomeHandler) Handles(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		handler.getIndex(res, req)
	case http.MethodPost:
		handler.postIndex(res, req)
	}
}

func (handler *HomeHandler) getIndex(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Hello, This is Get Handler! You can [GET, POST] to /home")
}

func (handler *HomeHandler) postIndex(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Helo, This is Post Handler! You can [GET, POST] to /home")
}

```

#### `ch2` Handler 인터페이스 다시보기

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
