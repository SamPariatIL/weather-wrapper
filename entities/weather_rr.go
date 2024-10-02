package entities

type Coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type CurrentWeather struct {
	Coord   `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Haze        string `json:"haze"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp        float32 `json:"temp"`
		FeelsLike   float32 `json:"feels_like"`
		TempMin     float32 `json:"temp_min"`
		TempMax     float32 `json:"temp_max"`
		Pressure    int     `json:"pressure"`
		Humidity    int     `json:"humidity"`
		SeaLevel    int     `json:"sea_level"`
		GroundLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float32 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		SunRise int    `json:"sunrise"`
		SunSet  int    `json:"sunset"`
	} `json:"sys"`
	TimeZone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	COD      int    `json:"cod"`
}

type Forecast struct {
	COD     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp        float32 `json:"temp"`
			FeelsLike   float32 `json:"feels_like"`
			TempMin     float32 `json:"temp_min"`
			TempMax     float32 `json:"temp_max"`
			Pressure    int     `json:"pressure"`
			SeaLevel    int     `json:"sea_level"`
			GroundLevel int     `json:"grnd_level"`
			Humidity    int     `json:"humidity"`
			TempKf      float32 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility int     `json:"visibility"`
		Pop        float32 `json:"pop"`
		Sys        struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
	City struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Coord
		Country    string `json:"country"`
		Population int    `json:"population"`
		TimeZone   int    `json:"timezone"`
		SunRise    int    `json:"sunrise"`
		SunSet     int    `json:"sunset"`
	} `json:"city"`
}
