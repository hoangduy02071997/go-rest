package common

import "fmt"

type Media struct {
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"`
	From   string `json:"cloud_name"`
	Ext    string `json:"ext"`
}

func (m *Media) FullFill(domain string) {
	m.Url = fmt.Sprintf("%s/%s", domain, m.Url)
}
