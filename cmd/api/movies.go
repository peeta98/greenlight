package main

import (
	"fmt"
	"github.com/peeta98/greenlight/internal/data"
	"net/http"
	"time"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of the Movie struct, containing the ID	we extracted from
	// the URL and some dummy data. Also notice that we deliberately haven't set a
	// value for the Year field.
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	// Create an envelope{"movies": movie} instance and pass it to writeJSON(), instead
	// of passing the plain movie struct.
	err = app.writeJSON(w, http.StatusOK, envelope{"movies": movie}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
