package spotify

type SpotifyAPIRequestBody struct {
	ID SpotifyID `json:"id"`
}

type SpotifyID string

type GenreList []string
