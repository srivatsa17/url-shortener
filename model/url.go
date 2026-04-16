package model

type URLShortenRequest struct {
	URL string `json:"url"`
}

type URLShortenResponse struct {
	LongURL    string `json:"long_url"`
	ShortURL   string `json:"short_url"`
	CreatedAt  string `json:"created_at"`
	ClickCount int64  `json:"click_count"`
	Error      string `json:"error,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
