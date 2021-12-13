package spotify

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type SpotifyClient interface {
	GetSongGenres(SpotifyID) (GenreList, error)
	GetArtistGenres(SpotifyID) (GenreList, error)
	GetPlaylistGenres(SpotifyID) (GenreList, error)
}

type spotifyClient struct {
	httpClient http.Client
}

var _ SpotifyClient = &spotifyClient{}

func NewSpotifyClient(log *logrus.Entry) {
	log.Infof("Initializing new SpotifyClient")

	token := 
}

func (client *spotifyClient) GetSongGenres(id SpotifyID) (GenreList, error) {

}

func (client *spotifyClient) GetArtistGenres(id SpotifyID) (GenreList, error) {

}

func (client *spotifyClient) GetPlaylistGenres(id SpotifyID) (GenreList, error) {

}
