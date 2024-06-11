package common

type Media struct {
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"`
	From   string `json:"cloud_name"`
	Ext    string `json:"ext"`
}
