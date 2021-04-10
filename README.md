# Go Rest Client

A Simple HTTP client for Golang with mock compatibility, take a look at
the [example](https://github.com/terdia/go-rest-client/blob/main/src/example) folder for usage and sample tests

```go
package main

import (
"fmt"
"github.com/terdia/go-rest-client/src/example"
//gorest "github.com/terdia/go-rest-client/src"
//"net/http"
//"time"
)

func main() {

//You can provide custom input to the HTTP client like so, in most cases the default is okay.
//gorest.Client = &http.Client{
//Transport:     nil,
//CheckRedirect: nil,
//Jar:           nil,
//Timeout:       1000 * time.Millisecond,
//}
	
result, err := example.GetAnimalFacts("590b9d90229d260020af0b06")
if err != nil {
fmt.Println(err)
}
fmt.Println(result)

	header := map[string][]string{
		"Content-type": {"application/json; charset=UTF-8"},
	}

	post, err := example.CreatePost(&example.CreatePostRequest{
		UserId: 1,
		Title:  "test",
		Body:   "data",
	}, header)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(post)
}

```

## Contributing

Thank you for considering contributing to Go Rest Client! pull request all welcomed

## License

Go Rest Client is open-sourced software licensed under
the [MIT license](https://github.com/terdia/go-rest-client/blob/main/LICENSE).

