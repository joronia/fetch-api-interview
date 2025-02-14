package responses

type ID struct {
	Id string `json:"id"`
}

type Points struct {
	Points int64 `json:"points"`
}

type ErrorExceptionMessage struct {
	Description string `json:"description"`
}
