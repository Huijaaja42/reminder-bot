package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/huijaaja42/reminder-bot/config"
)

type reminder struct {
	id      uuid.UUID
	channel string
	time    int64
	text    string
}

type user struct {
	sync.Mutex
	id    string
	queue []reminder
}

var users struct {
	sync.RWMutex
	userMap map[string]*user
}

func init() {
	users.userMap = make(map[string]*user)
}

func newUser(id string) *user {
	return &user{
		id:    id,
		queue: []reminder{},
	}
}

var s *discordgo.Session
var BotConfig *config.BotConfig

func init() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot parse config: %v", err)
	}
	BotConfig = &config.Bot

	s, err = discordgo.New("Bot " + BotConfig.Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "reminder",
			Description: "Reminder Bot commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "add",
					Description: "Add a reminder",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "text",
							Description: "Reminder text",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "time",
							Description: "Reminder time",
							Required:    true,
						},
					},
				},
				{
					Name:        "remove",
					Description: "Remove a reminder",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "id",
							Description: "Reminder id",
							Required:    true,
						},
					},
				},
				{
					Name:        "list",
					Description: "List all reminders",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"reminder": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			command := i.ApplicationCommandData().Options[0].Name
			options := i.ApplicationCommandData().Options[0].Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			content := ""

			if _, ok := users.userMap[i.Member.User.ID]; !ok {
				users.Lock()
				users.userMap[i.Member.User.ID] = newUser(i.Member.User.ID)
				users.Unlock()
			}

			u := users.userMap[i.Member.User.ID]

			switch command {
			case "add":
				var err error
				content, err = u.handleAddCommand(optionMap, i)
				if err != nil {
					content = err.Error()
				}
			case "remove":
				content = "error: invalid id"
				if opt, ok := optionMap["id"]; ok {
					err := u.removeReminder(opt.StringValue())
					if err != nil {
						content = err.Error()
						break
					}
					content = "Reminder removed"
				}
			case "list":
				content = u.listReminders()
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
		},
	}
)

func notifier() {
	for {
		users.RLock()
		for _, u := range users.userMap {
			u.Lock()
			for _, r := range u.queue {
				if r.time <= time.Now().Unix() {
					s.ChannelMessageSend(r.channel, fmt.Sprintf("<@!%s>\n%s", u.id, r.text))
					u.Unlock()
					err := u.removeReminder(r.id.String())
					u.Lock()
					if err != nil {
						log.Fatalf("Cannot remove reminder: %v", err)
					}
				}
			}
			u.Unlock()
		}
		users.RUnlock()
		time.Sleep(time.Duration(BotConfig.Interval) * time.Second)
	}
}

func main() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		log.Printf("Invite: https://discord.com/api/oauth2/authorize?client_id=%v&permissions=2048&scope=bot%%20applications.commands", s.State.User.ID)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, BotConfig.GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	go notifier()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Removing commands...")

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, BotConfig.GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	//s.ApplicationCommandBulkOverwrite(s.State.User.ID, BotConfig.GuildID, []*discordgo.ApplicationCommand{}) // Remove all commands

	log.Println("Gracefully shutting down.")
}
