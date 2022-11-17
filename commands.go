package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func (u *user) addReminder(channel string, t int64, text string) {
	u.Lock()
	defer u.Unlock()

	u.queue = append(u.queue, reminder{
		id:      uuid.New(),
		channel: channel,
		time:    t,
		text:    text,
	})
}

func (u *user) listReminders() string {
	u.Lock()
	defer u.Unlock()

	list := "**Your scheduled reminders:**\n"
	for _, e := range u.queue {
		list += fmt.Sprintf("`%v` at <t:%v:f> %s\n", e.id, e.time, e.text)
	}

	return list
}

func (u *user) removeReminder(id string) error {
	u.Lock()
	defer u.Unlock()

	for i, e := range u.queue {
		if id == e.id.String() {
			u.queue = append(u.queue[:i], u.queue[i+1:]...)
			return nil
		}
	}

	return errors.New("error: invalid id")
}

func (u *user) handleAddCommand(optionMap map[string]*discordgo.ApplicationCommandInteractionDataOption, i *discordgo.InteractionCreate) (string, error) {
	var text string
	var timeString string

	if opt, ok := optionMap["text"]; ok {
		text = opt.StringValue()
	}

	if opt, ok := optionMap["time"]; ok {
		timeString = opt.StringValue()
	}

	if text == "" || timeString == "" {
		return "", errors.New("error: invalid input")
	}

	if strings.HasPrefix(timeString, "in") {
		r := regexp.MustCompile(`in\s*(?P<value>\d+)\s*(?P<unit>s|min|h|d|w|mon|y)`)

		match := r.FindStringSubmatch(timeString)
		if match == nil {
			return "", errors.New("error: invalid time format")
		}

		var unit string
		var valueString string

		for i, v := range r.SubexpNames() {
			if v == "unit" {
				unit = match[i]
			}
			if v == "value" {
				valueString = match[i]
			}
		}

		value, err := strconv.Atoi(valueString)
		if err != nil {
			return "", errors.New("error: invalid time format")
		}

		t := time.Now()

		switch unit {
		case "s":
			t = t.Add(time.Second * time.Duration(value))
			u.addReminder(i.ChannelID, t.Unix(), text)
			return fmt.Sprintf("Will remind you at <t:%v:f> <t:%v:R>", t.Unix(), t.Unix()), nil
		case "min":
			t = t.Add(time.Minute * time.Duration(value))
			u.addReminder(i.ChannelID, t.Unix(), text)
			return fmt.Sprintf("Will remind you at <t:%v:f> <t:%v:R>", t.Unix(), t.Unix()), nil
		case "h":
			t = t.Add(time.Hour * time.Duration(value))
			u.addReminder(i.ChannelID, t.Unix(), text)
			return fmt.Sprintf("Will remind you at <t:%v:f> <t:%v:R>", t.Unix(), t.Unix()), nil
		case "d":
			t = t.AddDate(0, 0, value)
			u.addReminder(i.ChannelID, t.Unix(), text)
			return fmt.Sprintf("Will remind you at <t:%v:f> <t:%v:R>", t.Unix(), t.Unix()), nil
		case "w":
			t = t.AddDate(0, 0, value*7)
			u.addReminder(i.ChannelID, t.Unix(), text)
			return fmt.Sprintf("Will remind you at <t:%v:f> <t:%v:R>", t.Unix(), t.Unix()), nil
		case "mon":
			t = t.AddDate(0, 1, value)
			u.addReminder(i.ChannelID, t.Unix(), text)
			return fmt.Sprintf("Will remind you at <t:%v:f> <t:%v:R>", t.Unix(), t.Unix()), nil
		case "y":
			t = t.AddDate(1, 0, value)
			u.addReminder(i.ChannelID, t.Unix(), text)
			return fmt.Sprintf("Will remind you at <t:%v:f> <t:%v:R>", t.Unix(), t.Unix()), nil
		default:
			return "", errors.New("error: invalid time format")
		}
	}

	t, err := dateparse.ParseStrict(timeString)
	if err != nil {
		return "", errors.New("error: invalid time format")
	}
	u.addReminder(i.ChannelID, t.Unix(), text)
	return fmt.Sprintf("Will remind you at <t:%v:f> <t:%v:R>", t.Unix(), t.Unix()), nil
}
