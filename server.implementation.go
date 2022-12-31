package main

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ServerImplementation struct {
	Lock sync.Mutex
	DB   *gorm.DB
}

func (s *ServerImplementation) UploadMovie(ctx echo.Context) error {
	var newMovie Movie
	err := ctx.Bind(&newMovie)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	s.Lock.Lock()
	defer s.Lock.Unlock()
	tx := s.DB.Create(&newMovie)
	if tx.Error != nil {
		return ctx.JSON(http.StatusBadRequest, tx.Error)
	}
	return ctx.JSON(http.StatusOK, newMovie)
}

// Get Movies by cast member
// (GET /movies/castmember/{castmember})
func (s *ServerImplementation) GetMovieByCastMember(ctx echo.Context, castmember string) error {
	return nil

}

// Get Movies by genre
// (GET /movies/genre/{genre})
func (s *ServerImplementation) GetMovieBygenre(ctx echo.Context, genre string) error {
	return nil

}

// Get Movies by name
// (GET /movies/name/{name})
func (s *ServerImplementation) GetMovieByName(ctx echo.Context, name string) error {
	return nil

}

// Get Movies by year
// (GET /movies/year/{year})
func (s *ServerImplementation) GetMovieByYear(ctx echo.Context, year int64) error {
	var movies []Movie
	tx := s.DB.Where("year = ?", year).Find(&movies)
	if tx.Error != nil {
		return ctx.JSON(http.StatusBadRequest, tx.Error.Error())
	}
	return ctx.JSON(http.StatusOK, movies)
}
