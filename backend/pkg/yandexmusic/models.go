package yandexmusic

type Account struct {
	UID   int64  `json:"uid"`
	Login string `json:"login"`
}
type GetAccountResult struct {
	Account        Account `json:"account"`
	Subeditor      bool    `json:"subeditor"`
	SubeditorLevel int     `json:"subeditorLevel"`
	PretrialActive bool    `json:"pretrialActive"`
	Masterhub      struct {
		ActiveSubscriptions    []string `json:"activeSubscriptions"`
		AvailableSubscriptions []string `json:"availableSubscriptions"`
	} `json:"masterhub"`
	Plus struct {
		HasPlus             bool `json:"hasPlus"`
		IsTutorialCompleted bool `json:"isTutorialCompleted"`
	} `json:"plus"`
	HasOptions   []string `json:"hasOptions"`
	DefaultEmail string   `json:"defaultEmail"`
	BarBelow     struct {
		AlertID   string `json:"alertId"`
		Text      string `json:"text"`
		BgColor   string `json:"bgColor"`
		TextColor string `json:"textColor"`
		AlertType string `json:"alertType"`
		Button    struct {
			Text      string `json:"text"`
			BgColor   string `json:"bgColor"`
			TextColor string `json:"textColor"`
			URI       string `json:"uri"`
		} `json:"button"`
		CloseButton       bool `json:"closeButton"`
		CloseButtonStyles struct {
			White struct {
				BgColor     string `json:"bgColor"`
				TextColor   string `json:"textColor"`
				BorderColor string `json:"borderColor"`
			} `json:"white"`
			Black struct {
				BgColor     string `json:"bgColor"`
				TextColor   string `json:"textColor"`
				BorderColor string `json:"borderColor"`
			} `json:"black"`
		} `json:"closeButtonStyles"`
	} `json:"bar-below"`
	Userhash string `json:"userhash"`
}
