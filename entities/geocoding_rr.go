package entities

type Geocode struct {
	Name       string `json:"name"`
	LocalNames *struct {
		UR *string `json:"ur"`
		LV *string `json:"lv"`
		EO *string `json:"eo"`
		MY *string `json:"my"`
		AR *string `json:"ar"`
		ZH *string `json:"zh"`
		NE *string `json:"ne"`
		RU *string `json:"ru"`
		OC *string `json:"oc"`
		UK *string `json:"uk"`
		FA *string `json:"fa"`
		KU *string `json:"ku"`
		DE *string `json:"de"`
		TH *string `json:"th"`
		EL *string `json:"el"`
		PT *string `json:"pt"`
		MS *string `json:"ms"`
		FR *string `json:"fr"`
		ML *string `json:"ml"`
		HE *string `json:"he"`
		KN *string `json:"kn"`
		HI *string `json:"hi"`
		TA *string `json:"ta"`
		PA *string `json:"pa"`
		TE *string `json:"te"`
		BN *string `json:"bn"`
		ES *string `json:"es"`
		EN *string `json:"en"`
		KO *string `json:"ko"`
		CS *string `json:"cs"`
		JA *string `json:"ja"`
	} `json:"local_names"`
	Coord
	Country string  `json:"country"`
	State   *string `json:"state"`
}
