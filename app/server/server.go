package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/willdavsmith/spotisort/app/clients/spotify"
)

type Server struct {
	httpClient    http.Client
	spotifyClient spotify.SpotifyClient
}

func (s *Server) RunWebServer() {
	r := mux.NewRouter()
	apiSubrouter := r.PathPrefix("/api/").Subrouter()
	apiSubrouter.HandleFunc("/classify", s.classifySong)
	apiSubrouter.HandleFunc("/organize", s.organizeLikedSongs)

	srv := http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.spotifyClient = s.spotifyClient

	fmt.Println("Spotisort running on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}

func (s *Server) classifySong(w http.ResponseWriter, r *http.Request) {
	// Get song ID from request body
	// Get playlist genres list
	// Classify song by playlist list
}

func (s *Server) organizeLikedSongs(w http.ResponseWriter, r *http.Request) {
	// Get liked playlist
	// for each song in liked playlist, classifySong
}

func (s *Server) generateSongGenres(w http.ResponseWriter, r *http.Request) {
	requestBody, httpStatusCode, err := getJSONRequestBody(w, r)
	if err != nil {
		respondWithError(w, err.Error(), httpStatusCode)
		return
	}

	if req, ok := requestBody.(SpotifyAPIRequestBody); !ok {
		respondWithError(w, "Cannot parse request body", http.StatusBadRequest)
		return
	}
}

func (s *Server) generatePlaylistGenres(w http.ResponseWriter, r *http.Request) {
	//TODO Spotify limits you at 100 tracks. Need to set offset.
	id, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", id)
	bearer := fmt.Sprintf("Bearer %s", accessToken)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	resp, err := api.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data Tracks

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln(err)
	}

	keywords := make(map[string]int, 0)
	total := 0

	for i := 0; i < len(data.Items); i++ {
		url := data.Items[i].Track.ItemArtists[0].Href

		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", bearer)

		resp, err := api.client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var artist Artist

		if err := json.Unmarshal(body, &artist); err != nil {
			log.Fatalln(err)
		}

		for j := 0; j < len(artist.Genres); j++ {
			fields := strings.Fields(artist.Genres[j])
			for k := 0; k < len(fields); k++ {
				keywords[fields[k]] += 1
				total += 1
			}
		}
	}

	fmt.Println(keywords)

	w.WriteHeader(http.StatusAccepted)
}

func respondWithError(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func getJSONRequestBody(w http.ResponseWriter, r *http.Request) (body interface{}, httpStatusCode int, err error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		err := fmt.Errorf("Content Type is not application/json")
		return nil, http.StatusUnsupportedMediaType, err
	}

	var req interface{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return nil, http.StatusBadRequest, err
	}

	return &req, http.StatusOK, nil
}
