# Warden: A Multi-LLM AI Framework

![1500x500 (21)](https://github.com/user-attachments/assets/1f5870be-ae28-407a-bea2-2f4bcf18e230)

## **Table of Contents**
- [Overview](#overview)
- [Core Features](#core-features)
- [Quick Start](#quick-start)
- [Architecture Overview](#architecture-overview)
- [Supported Platforms](#supported-platforms)
- [Extensibility](#extensibility)
- [Examples](#examples)
- [Future Roadmap](#future-roadmap)
- [Contributing](#contributing)
- [License](#license)

---

## **Overview**
**Warden** is a highly extensible and modular **Go-based framework** for building AI-powered agents capable of interacting with multiple large language models (LLMs) and platforms. It offers a seamless integration with modern AI models such as **DeepSeek**, **OpenAI**, **Claude**, and **Grok**, while providing cross-platform compatibility, including CLI, Twitter, and Discord.

With **Warden**, developers can define, extend, and share intelligent agents that are capable of responding contextually, managing memory, and collaborating with other agents. Its simple but powerful design makes it the perfect starting point for creating scalable AI-powered applications.

### **Key Highlights**
- **Go-Powered Performance**: A fast, memory-safe, and highly concurrent system.
- **Pluggable Multi-LLM Support**: Integrate with DeepSeek, OpenAI GPT-4, Claude, Grok, and more.
- **Cross-Platform Compatibility**: CLI, Twitter, and Discord bots.
- **Extensible Architecture**: Build, export, and extend agents with ease.
- **Stateful Conversations**: Persist memory for contextual interactions.
- **Secure and Modular Design**: Isolated environments for agent execution.

<img width="726" alt="image_25_1" src="https://github.com/user-attachments/assets/6139ac25-a9c3-4bb9-8875-7946f79ba69e" />


---

## **Core Features**

### **1. Agent-Based Architecture**
- **Pluggable Agents**: Define different types of agents with distinct AI capabilities.
- **Multi-Agent System**: Create, manage, and collaborate with multiple agents.
- **Modular Extensibility**: Dynamically add functionalities to agents.
- **Serialization and Sharing**: Export agents as `.agent` files and share across environments.

### **2. Memory and State Management**
- **Context-Aware Conversations**: Persist memory across interactions for contextual replies.
- **Shared State System**: Enable agents to collaborate by sharing knowledge.
- **Custom Knowledge Injection**: Inject external data into agents for enhanced intelligence.

### **3. Multi-LLM Integration**
- **Supported AI Providers**:
  - OpenAI (e.g., GPT-4)
  - GPT-4-Free
  - Claude (Anthropic)
  - DeepSeek
  - Grok (xAI)
- **Custom Model Selection**: Choose models dynamically for specific tasks.
- **Easy API Key Management**: Simplify LLM integrations with modular API management.

### **4. Cross-Platform Interaction**
- **CLI Chatbot**: Run AI conversations in the terminal.
- **Discord Bot**: Bring AI-powered interactions to your Discord server.
- **Twitter Integration**: Create, manage, and post tweets using AI-generated content.

### **5. Security and Reliability**
- **Memory Safety**: Built with Go's safe and efficient memory management.
- **Isolated Agent Execution**: Sandbox environments for safe agent operation.
- **API Security**: Proper handling of sensitive API keys and configurations.

### **6. Tooling and Extensibility**
- **Custom Tool System**: Define, register, and use external tools for agents.
- **Dynamic Function Calling**: Extend agent behavior with user-defined functions.
- **Dynamic Responses**: Connect to external APIs and databases for real-time intelligence.

---

## **Quick Start**

### **Step 1: Clone the Repository**
```bash
git clone https://github.com/warden-ai/Warden.git
cd Warden
```

### **Step 2: Install Dependencies**
```bash
go mod tidy
```

### **Step 3: Start a Supported API (e.g., GPT-4-Free)**
```bash
g4f api
```

### **Step 4: Run the Framework**

#### **4.1. Start the Full Application**
```bash
go run main.go
```

#### **4.2. Start the Discord Bot**
Set the `DISCORD_BOT_TOKEN` as an environment variable:
```bash
export DISCORD_BOT_TOKEN=your_discord_bot_token
go run main.go
```

#### **4.3. Run CLI Example**
```bash
go run main.go
```
Choose **CLI Chatbot** from the menu to interact with the AI.

---

## **Architecture Overview**
**Warden** follows a modular architecture with clean separation of concerns. Its components include:

1. **Agent Layer**: Handles interactions with various LLMs such as OpenAI, DeepSeek, Claude, and Grok.
2. **Platform Layer**: Manages integration with platforms like CLI, Discord, and Twitter.
3. **Memory and State Layer**: Provides shared memory for stateful agent interactions.
4. **Tooling Layer**: Enables custom tools and function calls for dynamic agent behavior.

---

## **Supported Platforms**
### **1. CLI**
- Fully interactive terminal-based chatbot.
- Option to select different LLMs for conversations.
  
### **2. Discord**
- Easy-to-configure bot for Discord servers.
- Allows users to chat with AI-powered agents directly in Discord channels.

### **3. Twitter**
- Create and post tweets using AI.
- Manage OAuth authentication for secure access to Twitter APIs.

---

## **Extensibility**
### **Adding New LLMs**
1. Create a new file in the `agent/` directory (e.g., `new_llm.go`).
2. Implement the `BaseAgent` interface with custom API logic.
3. Register the new LLM in the CLI or Discord bot logic.

### **Adding New Platforms**
1. Add a new directory (e.g., `new_platform/`).
2. Implement a handler for the platform's APIs.
3. Integrate the platform into the `main.go` or CLI.

---

## **Examples**
### **Example 1: Running GPT-4-Free Locally**
1. Start the GPT-4-Free API:
   ```bash
   g4f api
   ```
2. Run the Warden app:
   ```bash
   go run main.go
   ```

### **Example 2: Exporting and Importing Agents**
- Export an agent:
  ```go
  // Serialize the agent to a file
  agent.ExportToFile("agent_name.agent")
  ```
- Import an agent:
  ```go
  // Load the agent from a file
  loadedAgent := agent.ImportFromFile("agent_name.agent")
  ```

---

## **Future Roadmap**
1. **Web-Based Interface**:
   - Add a user-friendly web dashboard for managing agents.
2. **Enhanced Memory Features**:
   - Support for long-term memory and dynamic context injection.
3. **Additional Platforms**:
   - Slack, Telegram, and WhatsApp integrations.
4. **Advanced Collaboration**:
   - Allow multiple agents to collaborate on tasks in real time.

---

## **Contributing**
Contributions are welcome! Please fork the repository, create a feature branch, and submit a pull request. For major changes, open an issue first to discuss what you would like to add.

---

## **License**
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

Start building with **Warden**, the ultimate framework for AI-powered agents. 🚀