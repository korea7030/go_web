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
