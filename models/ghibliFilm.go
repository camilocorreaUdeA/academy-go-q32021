package models

type GhibliFilm struct {
	ID                     string `json:"id"`
	Title                  string `json:"title"`
	OriginalTitle          string `json:"original_title"`
	OriginalTitleRomanised string `json:"original_title_romanised"`
	Description            string `json:"description"`
	Director               string `json:"director"`
	Producer               string `json:"producer"`
	ReleaseDate            string `json:"release_date"`
	RunningTime            string `json:"running_time"`
	RtScore                string `json:"rt_score"`
	People                 Urls   `json:"people"`
	Species                Urls   `json:"species"`
	Locations              Urls   `json:"locations"`
	Vehicles               Urls   `json:"vehicles"`
	Url                    string `json:"url"`
}

type Urls []string
