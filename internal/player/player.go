package player

import (
	"context"
	"fmt"
	"gotunes/internal/queue"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Player gestisce la ripoduzione per il server
type Player struct {
	mu          sync.Mutex
	VoiceConn   *discordgo.VoiceConnection
	Queue       *queue.Queue
	Playing     bool
	Skip        chan bool
	Stop        chan bool
	CurrentSong *queue.Song
}

// .ctor per creare un nuovo player
func New() *Player {
	return &Player{
		Queue: &queue.Queue{},
		Skip:  make(chan bool, 1),
		Stop:  make(chan bool, 1),
	}
}

func GetStreamURL(query string) (string, string, error) {

	// Se non e' URL cerco su  YouTube
	searchQuery := query
	if len(query) < 4 || query[:4] != "http" {
		searchQuery = "ytsearch1:" + query
	}

	// Creo timeout per evitare blocchi
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//TODO: capire se usare un tool custom
	cmd := exec.CommandContext(
		ctx,
		"yt-dlp",
		"--get-url",
		"--get-title",
		"-f", "bestaudio",
		"--no-playlist",
		searchQuery,
	)

	output, err := cmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("yt-dlp error: %w", err)
	}

	// Splitto l'output
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) < 2 {
		return "", "", fmt.Errorf("risposta yt-dlp inaspettata")
	}

	title := lines[0]
	url := lines[1]

	return title, url, nil // Torno titolo - url
}

// Start avvia il loop di riproduzione
func (p *Player) Start(s *discordgo.Session, guildID string) {

}
