package response

type (
	ConfluenceContent struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		Status string `json:"status"`
		Title  string `json:"title"`
		Body   Body   `json:"body"`
	}
	Body struct {
		Storage Storage `json:"storage"`
	}
	Storage struct {
		Value string `json:"value"`
	}
)
