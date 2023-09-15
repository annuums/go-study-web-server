# annuums

- 본 Repository는 annuums에서 제작한 Become a DevOps/SRE Engineer 프로젝트 중 [Golang - 웹 서버 정보하기] 코드 저장소 입니다.

---

## Golang - 웹 서버 정복하기 Repository

### TDD

- 기능을 만들 때, 해당 기능이 잘 동작하는지 확인하는 것은 매우 중요해요.
- 사전에 테스트를 수행함으로써 Production에 올라가 있는 코드가 적어도 내가 의도한대로 잘 동작하는지 확인할 수 있어요.
- TDD는 기본적으로 어떤 서비스에 대한 기능적 요구사항을 분석하고, 해당하는 기능이 잘 동작하도록 근간이 되는 코드를 작성한 뒤, 이것이 잘 동작한다면 Business Logic을 작성하는 방법이에요.
  - 즉, 테스트에 통과하지 않는다면 Business Logic을 작성하지 않아요.
- 하지만 분명 코드를 작성할 때, 외부의 기능이 필요할 때가 있어요.
  - 데이터베이스를 연동한다든지, 함께 일하는 동료의 코드를 사용한다든지, 심지어 외부 패키지(라이브러리)등을 이용할 때가 있어요.
- 그런데 이러한 외부의 모든 것을 코드에 포함한다는 것은 쉽지 않답니다. 이를 해결하기 위해 우리는 `Mock Data` (쉽게 목업이라는 단어를 들어봤죠?)를 만들어 해당하는 기능들이 `**동작한다고 가정**`하고 테스트 한답니다.

#### 톺아보기

```go
func TestHomeIndex(t *testing.T) {
	assert := assert.New(t)

	homeURI := "/home"
	//* NewHandler가 명시하는 Mock 서버를 생성해요.
	ts := httptest.NewServer(app.NewHandler())
    //* 서버는 종료해줘야 해요.
    //* defer 지시자를 통해 `TestHomeIndex`함수가 종료될 때 Close()하도록 지시합니다.
	defer ts.Close()

	//* 서버 주소에 테스트할 homeURI를 붙여요
	testURL := fmt.Sprintf("%s%s", ts.URL, homeURI)

    //* 단순히 http.Get(ts.URL + homeURI)로 해도 된답니다.
	res, err := http.Get(testURL)

    //* 현재 우리의 실제 HomeHandler는 에러가 없기 때문에,
    //* NoError인지 확인해요.
	assert.NoError(err)
    //* 이후 http.StatusOK가 맞는지 확인한답니다.
	assert.Equal(http.StatusOK, res.StatusCode)

    //* 실제 요청을 보내고, 응답을 가져와요.
    //* 변수를 할당하지 않은 부분은 error를 반환받는 부분이랍니다.
	data, _ := io.ReadAll(res.Body)

    //* 이후 응답이 내가 기대한 값과 같은지 비교한답니다.
    //* 만약 다르다면, 해당 Test Case는 실패하고 말아요.
	assert.Equal("Hello, Home!\n", string(data))
}
```

- 위 코드는 `/home` 주소로 http 요청을 보냈을 때, 기대한 값이 나오는지 확인하는 코드랍니다.
- 실제로는 많은 것을 `Mocking`해야 하지만, 여기서는 단순히 요청을 보내면 그 결과가 나오는지 보는것을 정의했기 때문에, 간단하게 만들었어요.
- 자세한 내용은 코드에 주석을 달았기 때문에, 그것으로 대체할게요.
