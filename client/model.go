package client

type Result struct {
	HtmlFormattedUrl string `json:"htmlFormattedUrl,omitempty"`

	HtmlTitle string `json:"htmlTitle,omitempty"`

	HtmlSnippet string `json:"htmlSnippet,omitempty"`

	FormattedUrl string `json:"formattedUrl,omitempty"`

	Title string `json:"title,omitempty"`

	Snippet string `json:"snippet,omitempty"`

	Icon Icon `json:"icon,omitempty"`
}

type Icon struct {
	Src    string `json:"src,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}
