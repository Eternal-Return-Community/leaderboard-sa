package structs

type Pathnotes struct {
	Articles []Notes `json:"articles"`
}

type Notes struct {
	Thumbnail_url string   `json:"thumbnail_url"`
	I18ns         Location `json:"i18ns"`
}

type Location struct {
	En_Us English `json:"en_Us"`
}

type English struct {
	Title       string `json:"title"`
	Description string `json:"summary"`
	Link        string `json:"content_link"`
	Time        string `json:"created_at_for_humans"`
}
