package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/huijaaja42/reminder-bot/config"
	"github.com/huijaaja42/reminder-bot/model"
	"github.com/objectbox/objectbox-go/objectbox"
)

var s *discordgo.Session
var BotConfig *config.BotConfig
var box *model.ReminderBox

func initObjectBox() *objectbox.ObjectBox {
	objectBox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()
	if err != nil {
		log.Fatalf("Cannot open database: %v", err)
	}
	return objectBox
}

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

			var err error
			var user string
			content := ""

			if i.User != nil {
				user = i.User.ID
			} else {
				user = i.Member.User.ID
			}

			channel := i.ChannelID

			switch command {
			case "add":
				content, err = addCommand(optionMap, user, channel)
				if err != nil {
					content = err.Error()
				}
			case "remove":
				content = "error: invalid id"
				if opt, ok := optionMap["id"]; ok {
					idString := opt.StringValue()
					id, err := strconv.Atoi(idString)
					if err != nil {
						break
					}

					err = removeCommand(uint64(id), user)
					if err != nil {
						content = err.Error()
						break
					}

					content = "Reminder removed"
				}
			case "list":
				content, err = listCommand(user, channel)
				if err != nil {
					content = err.Error()
				}
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
		query := box.Query(model.Reminder_.Time.LessThan(0))
		query.SetInt64Params(model.Reminder_.Time, time.Now().Unix())

		reminders, err := query.Find()
		if err != nil {
			log.Fatalf("Database error: %v", err)
		}

		for _, r := range reminders {
			s.ChannelMessageSend(r.Channel, fmt.Sprintf("<@!%s>\n%s", r.User, r.Text))
			err := box.RemoveId(r.Id)
			if err != nil {
				log.Fatalf("Cannot remove reminder: %v", err)
			}
		}

		time.Sleep(time.Duration(BotConfig.Interval) * time.Second)
	}
}

func main() {
	ob := initObjectBox()
	defer ob.Close()
	box = model.BoxForReminder(ob)

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
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
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
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	//s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", []*discordgo.ApplicationCommand{}) // Remove all commands

	log.Println("Gracefully shutting down.")
}
