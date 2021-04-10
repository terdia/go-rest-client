package example

import (
	"github.com/stretchr/testify/assert"
	rest "github.com/terdia/go-rest-client/src"
	"github.com/terdia/go-rest-client/src/utils/mocks"
	"testing"
)

const (
	factId string = "591f98803b90f7150a19c229"
)

func init() {
	rest.Client = &mocks.MockHttpClient{}
}

func TestCreatePost(t *testing.T) {
	// arrange
	expected := `{"id": 101, "userId": 1, "title": "the title", "body": "the body"}`
	mocks.NewMockResponse(&mocks.MockResponse{
		StatusCode: 200,
		Body:       expected,
	})
	header := map[string][]string{
		"Content-type": {"application/json; charset=UTF-8"},
	}

	//act
	post, err := CreatePost(&CreatePostRequest{
		UserId: 1,
		Title:  "the title",
		Body:   "the body",
	}, header)

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, post)
	assert.EqualValues(t, 101, post.PostId)
	assert.EqualValues(t, 1, post.UserId)
	assert.EqualValues(t, "the title", post.Title)
	assert.EqualValues(t, "the body", post.Body)
}

func TestCreatePostError(t *testing.T) {
	// arrange
	mocks.NewMockError("http client error")
	header := map[string][]string{
		"Content-type": {"application/json; charset=UTF-8"},
	}

	//act
	post, err := CreatePost(&CreatePostRequest{
		UserId: 1,
		Title:  "the title",
		Body:   "the body",
	}, header)

	//assert
	assert.Nil(t, post)
	assert.NotNil(t, err)
	assert.EqualValues(t, "http client error", err.Error())
}

func TestCreatePostTimeout(t *testing.T) {
	// arrange
	mocks.NewMockResponse(nil)
	header := map[string][]string{
		"Content-type": {"application/json; charset=UTF-8"},
	}

	//act
	post, err := CreatePost(&CreatePostRequest{
		UserId: 1,
		Title:  "the title",
		Body:   "the body",
	}, header)

	//assert
	assert.Nil(t, post)
	assert.NotNil(t, err)
	assert.EqualValues(t, "timeout or server not reachable", err.Error())
}

func TestCreatePostBadStatusCode(t *testing.T) {
	// arrange
	mocks.NewMockResponse(&mocks.MockResponse{
		StatusCode: 400,
		Body:       "some data",
	})
	header := map[string][]string{
		"Content-type": {"application/json; charset=UTF-8"},
	}

	//act
	post, err := CreatePost(&CreatePostRequest{
		UserId: 1,
		Title:  "the title",
		Body:   "the body",
	}, header)

	//assert
	assert.Nil(t, post)
	assert.NotNil(t, err)
	assert.EqualValues(t, "bad request", err.Error())
}

func TestCreatePostInvalidResponseData(t *testing.T) {
	// arrange
	expected := `{"id": "101", "userId": 1, "title": "the title", "body": "the body"}`
	mocks.NewMockResponse(&mocks.MockResponse{
		StatusCode: 200,
		Body:       expected,
	})
	header := map[string][]string{
		"Content-type": {"application/json; charset=UTF-8"},
	}

	//act
	post, err := CreatePost(&CreatePostRequest{
		UserId: 1,
		Title:  "the title",
		Body:   "the body",
	}, header)

	//assert
	assert.Nil(t, post)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid api response", err.Error())
}

func TestGetAnimalFacts(t *testing.T) {
	// arrange
	expected := `{"_id": "591f98803b90f7150a19c229", "text": "In an average year, cat owners in the United States spend over $2 billion on cat food."}`
	mocks.NewMockResponse(&mocks.MockResponse{
		StatusCode: 200,
		Body:       expected,
	})

	//act
	resp, err := GetAnimalFacts(factId)

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.EqualValues(t, "591f98803b90f7150a19c229", resp.Id)
	assert.EqualValues(t, "In an average year, cat owners in the United States spend over $2 billion on cat food.", resp.Text)
}

func TestGetAnimalFactsError(t *testing.T) {
	// arrange
	mocks.NewMockError("http client error")

	//act
	fact, err := GetAnimalFacts(factId)

	//assert
	assert.Nil(t, fact)
	assert.NotNil(t, err)
	assert.EqualValues(t, "http client error", err.Error())
}

func TestGetAnimalFactsTimeout(t *testing.T) {
	// arrange
	mocks.NewMockResponse(nil)

	//act
	fact, err := GetAnimalFacts(factId)

	//assert
	assert.Nil(t, fact)
	assert.NotNil(t, err)
	assert.EqualValues(t, "timeout or server not reachable", err.Error())
}

func TestGetAnimalFactsBadStatusCode(t *testing.T) {
	// arrange
	mocks.NewMockResponse(&mocks.MockResponse{
		StatusCode: 400,
		Body:       "some data",
	})

	//act
	fact, err := GetAnimalFacts(factId)

	//assert
	assert.Nil(t, fact)
	assert.NotNil(t, err)
	assert.EqualValues(t, "something bad", err.Error())
}

func TestGetAnimalFactsInvalidResponseData(t *testing.T) {
	// arrange
	expected := `{"_id": 101, "text": "the text"}`
	mocks.NewMockResponse(&mocks.MockResponse{
		StatusCode: 200,
		Body:       expected,
	})

	//act
	fact, err := GetAnimalFacts(factId)

	//assert
	assert.Nil(t, fact)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid api response", err.Error())
}
