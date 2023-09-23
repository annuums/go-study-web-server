# annuums

- 본 Repository는 annuums에서 제작한 Become a DevOps/SRE Engineer 프로젝트 중 [Golang - 웹 서버 정보하기] 코드 저장소 입니다.

---

## Golang - 웹 서버 정복하기 Repository

### 웹 서버 처리 Mux 만들ㄱ

```go
package app

import (
	"net/http"

	routers "github.com/annuums/go-study-web-server/routes"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/home/", http.StripPrefix("/home", &routers.HomeRouter{}))
	mux.Handle("/", http.NotFoundHandler())

	return mux
}

```

- 설명 쓰기
