package models

// 用户信息
type Info struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
	Data    Data   `json:"data"`
}

type Data struct {
	Birthday   string      `json:"birthday"`
	Coins      float64     `json:"coins"` // 这个硬币数量如果没有登录的话都为0，所以这项数据不能获取
	Face       string      `json:"face"`
	FansBadge  bool        `json:"fans_badge"`
	IsFollowed bool        `json:"is_followed"`
	Jointime   int64       `json:"jointime"`
	Level      int         `json:"level"`
	Mid        int64       `json:"mid"`
	Moral      int64       `json:"moral"`
	Name       string      `json:"name"`
	Nameplate  Nameplate   `json:"nameplate"`
	Official   Official    `json:"official"`
	Pendant    Pendant     `json:"pendant"`
	Rank       int64       `json:"rank"`
	Sex        string      `json:"sex"`
	Sign       string      `json:"sign"`
	Silence    int64       `json:"silence"`
	SysNotice  interface{} `json:"sys_notice"`
	Theme      interface{} `json:"theme"`
	TopPhoto   string      `json:"top_photo"`
	Vip        Vip         `json:"vip"`
}

type Nameplate struct {
	Condition  string `json:"condition"`
	Image      string `json:"image"`
	ImageSmall string `json:"image_small"`
	Level      string `json:"level"`
	Name       string `json:"name"`
	Nid        int64  `json:"nid"`
}

type Official struct {
	Desc  string `json:"desc"`
	Role  int64  `json:"role"`
	Title string `json:"title"`
	Type  int64  `json:"type"`
}

type Pendant struct {
	Expire       int64  `json:"expire"`
	Image        string `json:"image"`
	ImageEnhance string `json:"image_enhance"`
	Name         string `json:"name"`
	Pid          int64  `json:"pid"`
}

type Vip struct {
	AvatarSubscript int64 `json:"avatar_subscript"`
	Label           struct {
		LabelTheme string `json:"label_theme"`
		Path       string `json:"path"`
		Text       string `json:"text"`
	} `json:"label"`
	NicknameColor string `json:"nickname_color"`
	Status        int64  `json:"status"`
	ThemeType     int64  `json:"theme_type"`
	Type          int64  `json:"type"`
}
