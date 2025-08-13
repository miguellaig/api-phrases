package responses

type PhraseResponse struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Original    string `json:"original"`
	Translation string `json:"translation"`
	Language    string `json:"language"`
}
