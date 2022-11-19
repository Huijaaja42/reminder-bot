package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/bwmarrin/discordgo"
	"github.com/huijaaja42/reminder-bot/model"
)

func addCommand(optionMap map[string]*discordgo.ApplicationCommandInteractionDataOption, user string, channel string) (string, error) {
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

	t := time.Now()

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

		switch unit {
		case "s":
			t = t.Add(time.Second * time.Duration(value))
		case "min":
			t = t.Add(time.Minute * time.Duration(value))
		case "h":
			t = t.Add(time.Hour * time.Duration(value))
		case "d":
			t = t.AddDate(0, 0, value)
		case "w":
			t = t.AddDate(0, 0, value*7)
		case "mon":
			t = t.AddDate(0, 1, value)
		case "y":
			t = t.AddDate(1, 0, value)
		default:
			return "", errors.New("error: invalid time format")
		}
	} else {
		var err error
		t, err = dateparse.ParseStrict(timeString)
		if err != nil {
			return "", errors.New("error: invalid time format")
		}
	}

	id, err := box.Put(&model.Reminder{
		User:    user,
		Channel: channel,
		Time:    t.Unix(),
		Text:    text,
	})
	if err != nil {
		log.Printf("Database error: %v", err)
		return "", errors.New("error: database error")
	}

	return fmt.Sprintf("Will remind you at <t:%v:f> id: `%v`", t.Unix(), id), nil
}

func removeCommand(id uint64, user string) error {
	r, err := box.Get(id)
	if err != nil || r == nil || r.User != user {
		return errors.New("error: invalid id")
	}

	err = box.RemoveId(r.Id)
	if err != nil {
		return errors.New("error: invalid id")
	}

	return nil
}

func listCommand(user string, channel string) (string, error) {
	query := box.Query(model.Reminder_.User.Equals("", false))
	query.SetStringParams(model.Reminder_.User, user)

	reminders, err := query.Find()
	if err != nil {
		return "", errors.New("error: database error")
	}

	list := "**Your scheduled reminders:**\n"

	for _, r := range reminders {
		if r.Channel == channel {
			list += fmt.Sprintf("id: `%v` at <t:%v:f> in this channel:\n> %s\n", r.Id, r.Time, r.Text)
			continue
		}

		st, _ := s.Channel(r.Channel)
		if st.Type == discordgo.ChannelTypeDM {
			list += fmt.Sprintf("id: `%v` at <t:%v:f> in DM:\n> %s\n", r.Id, r.Time, r.Text)
			continue
		}

		list += fmt.Sprintf("id: `%v` at <t:%v:f> in <#%v>:\n> %s\n", r.Id, r.Time, r.Channel, r.Text)
	}

	return list, nil
}
