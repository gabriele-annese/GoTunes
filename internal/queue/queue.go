package queue

import (
	"sync"
)

// Rappresenta la canzone nella coda
type Song struct {
	Title     string
	URL       string // URL youtube, spotify
	StreamURL string // URL audio diretto
	RequestBy string
}

// Rappresenta la coda vera e propria
type Queue struct {
	mu    sync.Mutex
	songs []Song
}

// Add aggiunge una canzone alla coda
func (q *Queue) Add(s Song) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// aggiungo la song all'array
	q.songs = append(q.songs, s)
}

// Next recupera la prima canzone dalla coda, la riuomve dalla coda e la ritorna
func (q *Queue) Next() (Song, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.songs) == 0 {
		return Song{}, false
	}

	s := q.songs[0]
	q.songs = q.songs[1:]

	return s, true
}

// Peek recupera la prima canzone dalla coda senza rimoverla
func (q *Queue) Peek() (Song, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.songs) == 0 {
		return Song{}, false
	}

	// Recupero la prima canzone dalla lista
	song := q.songs[0]
	return song, true
}

// List ritorna la lista di tutte le canzoni in coda
func (q *Queue) List() []Song {
	q.mu.Lock()
	defer q.mu.Unlock()

	song_list := []Song{}
	queue_list := append(song_list, q.songs...)
	return queue_list
}

// Clear pulisce tutta la queue
func (q *Queue) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()

	// pulisco la queue
	q.songs = []Song{}
}

// Len restituisce il numero di canzoni nella queue
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	len_songs := len(q.songs)
	return len_songs
}
