package handlers

import "github.com/camilocorreaUdeA/academy-go-q32021/models"

func filmObjectToRecord(film models.GhibliFilm) []string {
	return []string{
		film.ID,
		film.Title,
		film.OriginalTitle,
		film.OriginalTitleRomanised,
		film.Description,
		film.Director,
		film.Producer,
		film.ReleaseDate,
		film.RunningTime,
		film.RtScore,
		film.People[0],
		film.Species[0],
		film.Locations[0],
		film.Vehicles[0],
		film.Url,
	}
}
