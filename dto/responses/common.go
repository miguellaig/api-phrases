package responses

type ErrorResponse struct {
	Error string `json:"error"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type Message struct {
	Message string `json:"message"`
}

type APIResponse struct {
	Data any `json:"data"`
}
