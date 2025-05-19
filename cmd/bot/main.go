package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ADG08/ADGMusic/internal/infrastructure/adapters/discord"
	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening Discord connection: %v", err)
	}
	defer dg.Close()

	voiceManager := discord.NewVoiceManager(dg)
	commandHandler := discord.NewCommandHandler(voiceManager)
	dg.AddHandler(commandHandler.HandleCommand)

	// Initialisation de ton service (exemple)
	// activeChannelService := services.NewActiveChannelService(repo)

	// Lancement du timer dans une goroutine
	go func() {
		for {
			// Timer aléatoire entre 1 et 5 minutes
			wait := time.Duration(rand.Intn(5)+1) * time.Minute
			time.Sleep(wait)

			// Récupère un channel actif aléatoire
			// channel, err := activeChannelService.GetRandomActiveChannel(context.Background())
			// if err == nil && channel != nil {
			//     // Fais rejoindre le bot à ce channel
			//     // ex: discordAdapter.JoinVoice(channel.GuildID, channel.ChannelID)
			// }
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}
