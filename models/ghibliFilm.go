package models

type GhibliFilm struct {
	ID                     string `json:"id,omitempty"`
	Title                  string `json:"title,omitempty"`
	OriginalTitle          string `json:"original_title,omitempty"`
	OriginalTitleRomanised string `json:"original_title_romanised,omitempty"`
	Description            string `json:"description,omitempty"`
	Director               string `json:"director,omitempty"`
	Producer               string `json:"producer,omitempty"`
	ReleaseDate            string `json:"release_date,omitempty"`
	RunningTime            string `json:"running_time,omitempty"`
	RtScore                string `json:"rt_score,omitempty"`
	People                 Urls   `json:"people,omitempty"`
	Species                Urls   `json:"species,omitempty"`
	Locations              Urls   `json:"locations,omitempty"`
	Vehicles               Urls   `json:"vehicles,omitempty"`
	Url                    string `json:"url,omitempty"`
}

type Urls []string
