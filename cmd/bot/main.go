package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Carico il file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errore durante il caricamento del file .env")
	}

	// Recuper il token
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN non valorizzato")
	}

	// Creo la sessione Discord
	dis_session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Errore durante la creazione della sessione discrod: ", err)
	}

	// Definisco le proprieta Intents
	dis_session.Identify.Intents = discordgo.IntentGuilds | discordgo.IntentGuildMessages | discordgo.IntentGuildVoiceStates

	dis_session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf(
			"Il bot si e' connesso come %v#%v",
			s.State.User.Username,
			s.State.User.Discriminator)
		s.UpdateGameStatus(0, "Cazzo vuoi colgione")
	})

	// Apri la connessione
	err = dis_session.Open()
	if err != nil {
		log.Fatal("Errore durante l'apertura della connessione: ", err)
	}
	defer dis_session.Close()

	log.Println("DuxMusicBot e' in esecuzione.\nTROIE il duce e' qui")

	// CTRL + C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Alla prossima sudditi coglioni.")
}
