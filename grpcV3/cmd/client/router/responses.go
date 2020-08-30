package router

type List struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"statusCode"`
}
