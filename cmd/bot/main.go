package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ADG08/ADGMusic/internal/infrastructure/discord"
	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Erreur lors de la création de la session Discord: %v", err)
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuilds | discordgo.IntentGuildVoiceStates

	dg.LogLevel = discordgo.LogDebug

	commandHandler := discord.NewCommandHandler(dg)
	commandHandler.RegisterCommands()

	err = dg.Open()
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la connexion: %v", err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	dg.Close()
}
