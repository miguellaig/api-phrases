package requests

type PhrasesRequest struct {
	Original    string `json:"original"`
	Translation string `json:"translation"`
	Language    string `json:"language"`
}
