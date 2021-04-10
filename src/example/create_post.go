package example

type CreatePostRequest struct {
	UserId int64  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type CreatePostResponse struct {
	PostId int64  `json:"id"`
	UserId int64  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
