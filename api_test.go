package rottentomatoes

import (
	"reflect"
	"testing"
)

func TestMovieInfo(t *testing.T) {
	c, _ := NewClient()

	m, err := c.MovieInfo("14281")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := m.Title
	want := "The Big Lebowski"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieInfo = %+v,\nwant %+v", got, want)
	}
}

func TestMovieCast(t *testing.T) {
	c, _ := NewClient()

	cast, err := c.MovieCast("14281")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := cast[0].Name
	want := "Jeff Bridges"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieCast = %+v,\nwant %+v", got, want)
	}
}

func TestMovieClips(t *testing.T) {
	c, _ := NewClient()

	clips, err := c.MovieClips("14281")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := clips[0].Title
	want := "Big Lebowski Trailer"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieClips = %+v,\nwant %+v", got, want)
	}
}

func TestMovieReviews(t *testing.T) {
	c, _ := NewClient()

	reviews, _, err := c.MovieReviews("14281", "top_critic", 5, 1, "us")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := reviews[0].Quote
	want := "Adds up to considerably less than the sum of its often scintillating parts."

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieReviews = %+v,\nwant %+v", got, want)
	}
}

func TestMovieSimilar(t *testing.T) {
	c, _ := NewClient()

	similar, err := c.MovieSimilar("14281", 5)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := similar[0].Title
	want := "Fargo"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieSimilar = %+v,\nwant %+v", got, want)
	}
}

func TestMovieAlias(t *testing.T) {
	c, _ := NewClient()

	m, err := c.MovieAlias("0118715")

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := m.Title
	want := "The Big Lebowski"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MovieAlias = %+v,\nwant %+v", got, want)
	}
}

func TestMoviesSearch(t *testing.T) {
	c, _ := NewClient()

	result, _, err := c.MoviesSearch("The Big Lebowski", 1, 1)

	if err != nil {
		t.Errorf("error not nil: %+v", err)
	}

	got := result[0].Id
	want := "14281"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MoviesSearch = %+v,\nwant %+v", got, want)
	}
}