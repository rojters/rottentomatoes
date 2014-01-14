package rottentomatoes

import (
	"strconv"
	"time"
)

// MovieInfo returns movie information using id of movie.
func (c *client) MovieInfo(id string) (movie Movie, err error) {

	time.Sleep(1 * time.Second)

	endpoint := c.getEndpoint("MovieInfo", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movie, err = unmarshalMoviesInfo(data)
	return
}

// MovieCast returns the abridged cast members of a movie.
func (c *client) MovieCast(id string) (castList []Cast, err error) {

	time.Sleep(1 * time.Second)

	endpoint := c.getEndpoint("MovieCast", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	castList, err = unmarshalCastInfo(data)
	return
}

// MovieClips returns clips, trailers etc. for a movie.
func (c *client) MovieClips(id string) (clips []Clip, err error) {

	time.Sleep(1 * time.Second)

	endpoint := c.getEndpoint("MovieClips", id)
	urlParams := c.prepareUrl(nil)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	clips, err = unmarshalClips(data)
	return
}

// MovieReviews returns reviews for a movie.
func (c *client) MovieReviews(id string, review_type string, page_limit int, page int, country string) (reviews []Review, total int, err error) {

	time.Sleep(1 * time.Second)

	endpoint := c.getEndpoint("MovieReviews", id)
	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	q := map[string]string{
		"review_type": review_type,
		"page_limit":  page_limit_t,
		"page":        page_t,
		"country":     country,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	reviews, total, err = unmarshalReviews(data)
	return
}

// MovieSimilar returns similar movies.
func (c *client) MovieSimilar(id string, limit int) (movies []Movie_, err error) {

	time.Sleep(1 * time.Second)

	endpoint := c.getEndpoint("MovieSimilar", id)
	limit_t := strconv.Itoa(limit)

	q := map[string]string{
		"limit": limit_t,
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, err = unmarshalMovies(data)
	return
}

// MovieAlias returns movie using alternative id, such as IMDB id.
func (c *client) MovieAlias(id string) (movie Movie, err error) {

	time.Sleep(1 * time.Second)

	endpoint := c.getEndpoint("MovieAlias", id)

	q := map[string]string{
		"id":   id,
		"type": "imdb",
	}

	urlParams := c.prepareUrl(q)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movie, err = unmarshalMoviesInfo(data)
	return
}

// MoviesSearch returns a list movies that matches the query string.
func (c *client) MoviesSearch(q string, page_limit int, page int) (movies []Movie_, total int, err error) {

	time.Sleep(1 * time.Second)

	endpoint := c.getEndpoint("MoviesSearch", "")

	page_limit_t := strconv.Itoa(page_limit)
	page_t := strconv.Itoa(page)

	queryParams := map[string]string{
		"q":          q,
		"page_limit": page_limit_t,
		"page":       page_t,
	}

	urlParams := c.prepareUrl(queryParams)
	data, err := c.request(endpoint + urlParams)

	if err != nil {
		return
	}

	movies, total, err = unmarshalSearch(data)
	return
}