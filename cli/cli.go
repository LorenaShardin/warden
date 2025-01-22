package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-agents/agent"
	"go-agents/discord"
)

// DisplayBanner shows the ASCII art for Warden
func displayBanner() {
	asciiArt := `
                 ~                                 ~
                ~~~                               ~~~
             ~~~~~~~~       <->     v1       ~~~~~~~~
           ~~~~~~~~~~~                   ~~~~~~~~~~~~
           ~~~~~~~~~~~    {-------}     ~~~~~~~~~~~~
             ~~~~~~~~     ~~~~~~~~~     ~~~~~~~~
                 ~~~    ~~~~~~~~~~~~~    ~~~
                 ~~   ~~~~~~~~~~~~~~~~~   ~~
                     ~~~~~~~~~~~~~~~~~~~
                   ~~~~~~~~~~~~~~~~~~~~~~~
                 ~~~~~~~~~~~~~~~~~~~~~~~~~~~
              ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           <~~~~~~~~~~~~~~~~~~   ~~~~~~~~~~~~~~~~~>
               / /                ||                \ \
              / /                (||)                \ \
             / /                  ||                  \ \
             \/                                          \/

	██╗    ██╗ █████╗ ██████╗ ██████╗ ███████╗███╗   ██╗
	██║    ██║██╔══██╗██╔══██╗██╔══██╗██╔════╝████╗  ██║
	██║ █╗ ██║███████║██████╔╝██║  ██║█████╗  ██╔██╗ ██║
	██║███╗██║██╔══██║██╔═══╝ ██║  ██║██╔══╝  ██║╚██╗██║
	╚███╔███╔╝██║  ██║██║     ██████╔╝███████╗██║ ╚████║
	 ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝     ╚═════╝ ╚══════╝╚═╝  ╚═══╝
	`
	fmt.Println(asciiArt)
}

// StartCLI starts the chatbot CLI loop
func StartCLI() {
	displayBanner()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to WARDEN: A Multi-LLM AI Framework!")
	fmt.Println("Choose an option:")
	fmt.Println("1. Run AI Chatbot in CLI")
	fmt.Println("2. Run Discord Bot")
	fmt.Println("3. Exit")

	fmt.Print("Enter choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		startCLIChat()
	case "2":
		discord.StartDiscordBot()
	case "3":
		fmt.Println("Exiting... Goodbye!")
		return
	default:
		fmt.Println("Invalid choice")
	}
}

func startCLIChat() {
	fmt.Println("Choose an LLM:")
	fmt.Println("1. GPT-4 Free")
	fmt.Println("2. OpenAI (GPT-4)")
	fmt.Println("3. DeepSeek")
	fmt.Println("4. Claude (Anthropic)")
	fmt.Println("5. Grok (xAI)")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var aiAgent interface {
		SendMessage(string) (string, error)
	}

	switch choice {
	case "1":
		aiAgent = agent.NewGPT4FreeAgent("You are an AI assistant.")
	case "2":
		fmt.Print("Enter OpenAI API Key: ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		aiAgent = agent.NewOpenAIAgent(apiKey, "You are an AI assistant.")
	case "3":
		fmt.Print("Enter DeepSeek API Key: ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		aiAgent = agent.NewDeepSeekAgent(apiKey, "You are an AI assistant.")
	case "4":
		fmt.Print("Enter Claude API Key: ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		aiAgent = agent.NewClaudeAgent(apiKey, "You are an AI assistant.")
	case "5":
		fmt.Print("Enter Grok API Key: ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		aiAgent = agent.NewGrokAgent(apiKey, "You are an AI assistant.")
	default:
		fmt.Println("Invalid choice")
		return
	}

	for {
		fmt.Print("USER: ")
		userMessage, _ := reader.ReadString('\n')
		userMessage = strings.TrimSpace(userMessage)

		if userMessage == "exit" {
			break
		}

		reply, err := aiAgent.SendMessage(userMessage)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("AI:", reply)
		}
	}
}
