# annuums

- 본 Repository는 annuums에서 제작한 Become a DevOps/SRE Engineer 프로젝트 중 [Golang - 웹 서버 정보하기] 코드 저장소 입니다.

---

## Golang - 웹 서버 정복하기 Repository

### routers

- `Router`는 들어오는 HTTP 요청에 대해 `URI`를 기준으로 요청을 처리할 함수, 객체 등을 구분해요. 예를 들어서, `/home/address`는 `집의 주소와 관련한 정보를 보여주고`, `/home/people`은 `집에 살고있는 가족 구성원 정보를 얻는` Logic을 수행할 수 있어요.

- 즉, 요청이 들어오면 이를 처리할 Logic을 할당할 수 있어요.

- 뿐만 아니라 Router마다 `유저 인증, 권한 인가, 로그` 등 `Middle Logic, Middleware`를 다르게 할당할 수 있어요.

```go
package routers

import (
	"fmt"
	"net/http"

	"github.com/annuums/go-study-web-server/handlers"
)

type HomeRouter struct{}

/*
*

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/
func (home *HomeRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Requested %s\n", req.URL.Path)

	handler := &handlers.HomeHandler{}
	switch req.URL.Path {
	case "/":
		handler.IndexHandler(res, req)
	default:
		http.NotFound(res, req)
	}
}
```

- 위 코드는 `package routers`를 새롭게 만들고, 그 안에 `HomeRouter` 구조체를 만들었어요. 아참, 우리는 `Router, Handler`를 분리하여 관리함으로써 관리 용이성을 높이고, 재사용성을 높일 수 있어요. 이에 따라 해당 함수는 `routes/` 디렉토리 아래에 `home.routes.go`라는 파일에 작성됐어요.

  - 즉, 또 다른 `UserRouter`가 필요하다면, `routes/user.routes.go` 파일에 작성하면 되겠죠?

- 이외의 것은 `ch2`에서 확인한 것과 같아요. 단지 `HomeHandler` -> `HomeRouter`로 바뀌었고, `ServeHTTP` 아래에 `Handler`들이 모두 `handler` 객체의 함수로 바뀌었어요.
