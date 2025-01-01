package discord

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"go-agents/agent"
)

// StartDiscordBot initializes the Discord bot
func StartDiscordBot() {
	token := os.Getenv("DISCORD_BOT_TOKEN") // Set in environment variables
	if token == "" {
		fmt.Println("Error: DISCORD_BOT_TOKEN is missing!")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Set message handler
	dg.AddMessageCreateHandler(handleMessage)

	// Open connection to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord connection:", err)
		return
	}

	fmt.Println("Bot is now running! Press CTRL+C to exit.")
	select {}
}

// Handles messages from users
func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	message := strings.ToLower(m.Content)
	aiAgent := agent.NewGPT4FreeAgent("You are a Discord AI assistant.")

	reply, err := aiAgent.SendMessage(message)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error processing request.")
	} else {
		s.ChannelMessageSend(m.ChannelID, reply)
	}
}
