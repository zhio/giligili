package serializer

import (
	"giligili/model"
)

type Movie struct {
	ID 		uint 	`json:"id"`
	Movie_id int  `json:"movie_id"`
	Name string		`json:"name"`
	En_name string	`json:"en_name"`
	Category string `json:"category"'`
	Cover_url string 	`json:"cover_url"`
	Area string `json:"area"`
	Release_date int `json:"release_date"`
	Info string `json:"info"`
	Curation int `json:"duration"`
}
func BuildMovie(item model.Movie) Movie{
	return  Movie {
		ID:           item.ID,
		Movie_id:     item.Movie_id,
		Name:         item.Name,
		En_name:      item.En_name,
		Category:     item.Category,
		Cover_url:    "http:"+item.Cover_url,
		Area:         item.Area,
		Release_date: item.Release_date,
		Info:         item.Info,
		Curation:     item.Duration,
	}

}

func BuildMovies(items []model.Movie) (movies []Movie){
	for _,item := range items{
		movie := BuildMovie(item)
		movies = append(movies, movie)
	}
	return movies
}

