package models

type Fans struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
	Data    People `json:"data"`
}

type People struct {
	ReVersion int64 `json:"re_version"`
	Total     int64 `json:"total"`
	List      []Up  `json:"list"`
}

type Up struct {
	Attribute      int64       `json:"attribute"`
	Face           string      `json:"face"`
	Mid            int64       `json:"mid"`
	Mtime          string      `json:"mtime"`
	Sign           string      `json:"sign"`
	Tag            interface{} `json:"tag"`
	Special        int64       `json:"special"`
	Uname          string      `json:"uname"`
	OfficialVerify struct {
		Desc string `json:"desc"`
		Type int64  `json:"type"`
	} `json:"official_verify"`
}
