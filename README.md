### go testing
#### GoConvey
https://github.com/smartystreets/goconvey

#### install
```
$ go get github.com/smartystreets/goconvey
```

#### quick start browser
- project path에서 "goconvey" 명령어 수행
```
$ goconvey
```
- localhost:8080 으로 접속하여 test 수행

### go testify assert
https://github.com/stretchr/testify

Go lang에서 test를 쉽게 해주기 위한 기능
#### install
```
$ go get github/com/stretchr/testify/assert
```

#### 사용법
```go
    assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"aaa", "last_name":"bbb", "email":"ccc"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("aaa", user.FirstName)
	assert.Equal("bbb", user.LastName)
```

### gorilla mux
url path router 지정  
https://github.com/gorilla/mux
#### install
```bash
$ go get -u github.com/gorilla/mux
```
#### examples
```go
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/products", ProductsHandler)
    r.HandleFunc("/articles", ArticlesHandler)
    http.Handle("/", r)
}
```
### gorilla pat
RESTful api를 간단하게 구현하게 하는 package  
https://github.com/gorilla/pat

### unrolled render
JSON, XML, text, binary data, HTML templates의 rendering을 쉽게 해줌  
https://github.com/unrolled/render

### negroni
Go의 웹 미들웨어 기능을 모아놓은 package  
Routing, Static, Logging 등의 기능 제공  
https://github.com/urfave/negroni/blob/master/translations/README_ko_KR.md#staticv

### antage eventsource
http server를 위한 server-sent eventsource 기능 제공 pkackage  
https://github.com/antage/eventsource