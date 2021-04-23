package entity

type Notification struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Extras      map[string]string `json:"extras"`
}
