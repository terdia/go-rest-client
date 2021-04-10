package example

import (
	"encoding/json"
	"errors"
	rest "github.com/terdia/go-rest-client/src"
	"io/ioutil"
	"net/http"
)

const (
	aniFactApiUrl      = "https://cat-fact.herokuapp.com"
	jsonPlaceholderUrl = "https://jsonplaceholder.typicode.com/posts"
)

type animalFact struct {
	Id   string `json:"_id"`
	Text string `json:"text"`
}

func CreatePost(request *CreatePostRequest, headers http.Header) (*CreatePostResponse, error) {
	res, err := rest.Post(jsonPlaceholderUrl, request, headers)
	if err != nil {
		return nil, err
	}

	if res == nil || res.Body == nil {
		return nil, errors.New("timeout or server not reachable")
	}

	if res.StatusCode > 299 {
		return nil, errors.New("bad request")
	}

	body, _ := ioutil.ReadAll(res.Body)

	var post = CreatePostResponse{}
	if marshalErr := json.Unmarshal(body, &post); marshalErr != nil {
		return nil, errors.New("invalid api response")
	}

	return &post, nil
}

func GetAnimalFacts(factId string) (*animalFact, error) {

	res, err := rest.Get(aniFactApiUrl+"/facts/"+factId, nil)
	if err != nil {
		return nil, err
	}

	if res == nil || res.Body == nil {
		return nil, errors.New("timeout or server not reachable")
	}

	if res.StatusCode > 299 {
		return nil, errors.New("something bad")
	}

	body, _ := ioutil.ReadAll(res.Body)

	var fact = animalFact{}
	if marshalErr := json.Unmarshal(body, &fact); marshalErr != nil {
		return nil, errors.New("invalid api response")
	}

	return &fact, nil
}
