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


func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "Hello, Home!")
}
```

- 위 코드에서 `func HomeHandler(...)`는 함수 명은 다르지만, 실제로 `Handler` 인터페이스의 `ServeHTTP`와 똑같이 생겼어요. 즉, 해당 함수를 통해 요청을 처리할 수 있어요.

- `handlers`는 실제로 요청을 처리하는 함수를 담당하는 것이지, 어떤 `경로`를 어떤 `Http Method`로 요청 받았을 때, ... 등은 신경쓰지 않아요. 다만, 이러한 요청을 `어떻게 처리` 할 거야! 를 담당하는 것이죠.

#### 한 줄씩 톺아보기

```go
func HomeHandler(res http.ResponseWriter, req *http.Request) {
    ...
```

- 해당 구문은 하나의 `handler` 함수를 정의한답니다. 일반적으로 웹 통신은 다음과 같이 동작해요
  1. Client가 Server로 `요청`을 보낸다.
  2. Server는 해당 `요청`이 유효한지 확인한다.
  3. Server는 `요청`을 처리하는 Logic을 실행한다.
  4. Server는 `응답`을 Client에게 반환한다.
- 여기서 `요청`은 `req *http.Request` 인자로, `응답`은 `res http.ResponseWriter` 인자를 통해 처리해요.

```go
...
    res.WriteHeader(http.StatusOK)
    ...
```

- 해당 구문은 `응답`을 처리합니다. 여기서는 `응답 코드`를 `http.StatusOK`, 즉 `200`으로 설정하여 반환해요.
  - [Http Response Code 자세히 보기](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

```go
    ...
    fmt.Fprintln(res, "Hello, Home!")
    ...
```

- 해당 구문은 `응답`을 처리합니다. 여기서는 `응답 메시지`를 작성하고 있어요. `HomeHandler`를 통해 처리되는 요청은, `Hello, Home!`이라는 내용을 받을거에요

#### 종합하면...

- `HomeHandler`로 처리되는 요청은 `200` 응답 코드와 `Hello, Home!` 메시지를 반환 받을 거에요.

#### Handler 인터페이스

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
