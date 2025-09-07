package main

import (
	"fmt"
)
type Movie struct {
	Title  string
	Year   int
	Rating float64
	Genres []string
}

func main() {
	movies := []Movie{
		{
			Title:  "Майнкрафт КИНО",
			Year:   2025,
			Rating: 9.1,
			Genres: []string{"приключения", "фэнтези", "комедия"},
		},
		{
			Title:  "ФНАФ кино",
			Year:   2023,
			Rating: 9.2,
			Genres: []string{"ужасы"},
		},
		{
			Title:  "Минионы",
			Year:   2024,
			Rating: 8.6,
			Genres: []string{"фантастика", "драма", "приключения"},
		},
		{
			Title:  "Кот в сапогах: последнее желание",
			Year:   2023,
			Rating: 8.7,
			Genres: []string{"приключения", "драма"},
		},
		{
			Title:  "Форсаж 42",
			Year:   2025,
			Rating: 1.7,
			Genres: []string{"фантастика", "боевик", "семья"},
		},
	}
	fmt.Println("Все фильмы в базе:")
	printMovies(movies)

	highestRated := findHighestRated(movies)
	fmt.Printf("\nФильм с самым высоким рейтингом:\n")
	fmt.Printf("«%s» (%d) - %.1f/10\n", highestRated.Title, highestRated.Year, highestRated.Rating)

	fmt.Println("Поиск фильмов по жанру:")
	searchGenres := []string{"драма", "фантастика", "комедия"}
	for _, genre := range searchGenres {
		fmt.Printf("\nФильмы в жанре '%s':\n", genre)
		found := findMoviesByGenre(movies, genre)
		if len(found) == 0 {
			fmt.Printf("  Не найдено фильмов в жанре '%s'\n", genre)
		} else {
			for _, movie := range found {
				fmt.Printf("  - «%s» (%d) - %.1f/10\n", movie.Title, movie.Year, movie.Rating)
			}
		}
	}


	fmt.Println("Добавляем жанры фильму «Начало»:")
	for i := range movies {
		if movies[i].Title == "Начало" {
			movies[i].Genres = append(movies[i].Genres, "детектив")
			fmt.Printf("Добавлен жанр 'детектив' к фильму «%s»\n", movies[i].Title)
			fmt.Printf("Теперь жанры: %v\n", movies[i].Genres)
			break
		}
	}

	fmt.Println("Добавляем новый фильм:")
	
	newMovie := Movie{
		Title:  "Три богатыря и четвертый богатырь",
		Year:   2025,
		Rating: 9.5,
		Genres: []string{"криминал", "комедия", "боевик", "богатыри"},
	}
	movies = append(movies, newMovie)
	fmt.Printf("Добавлен фильм: «%s» (%d) - %.1f/10\n", newMovie.Title, newMovie.Year, newMovie.Rating)
	fmt.Printf("Жанры: %v\n", newMovie.Genres)

	fmt.Println("Фильмы в жанре 'комедия' после добавления:")
	comedies := findMoviesByGenre(movies, "комедия")
	for _, movie := range comedies {
		fmt.Printf("  - «%s» (%d) - %.1f/10\n", movie.Title, movie.Year, movie.Rating)
	}
}

func printMovies(movies []Movie) {
	for i, movie := range movies {
		fmt.Printf("%d. «%s» (%d) - %.1f/10\n", i+1, movie.Title, movie.Year, movie.Rating)
		fmt.Printf("   Жанры: %v\n", movie.Genres)
	}
}

func findHighestRated(movies []Movie) Movie {
	highest := movies[0]
	for _, movie := range movies {
		if movie.Rating > highest.Rating {
			highest = movie
		}
	}
	return highest
}

func findMoviesByGenre(movies []Movie, genre string) []Movie {
	var result []Movie
	for _, movie := range movies {
		for _, movieGenre := range movie.Genres {
			if movieGenre == genre {
				result = append(result, movie)
				break
			}
		}
	}
	return result
}