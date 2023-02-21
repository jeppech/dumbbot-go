package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var cooldownLock = make(chan bool, 1)
var voiceLock = make(chan bool, 1)

func on_message(trigger string, callback MessageAuthorCallback) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		msg := strings.ToLower(m.Content)

		// check if the message is it matches the given trigger
		if !strings.HasPrefix(msg, trigger) {
			return
		}

		// Delete the message
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)

		if err != nil {
			fmt.Println("Error deleting message:", err)
			return
		}

		if len(cooldownLock) == 1 {
			return
		}

		parts := strings.SplitN(msg, " ", 2)
		triggers := strings.Split(parts[1], ",")
		sounds := make([]*Sound, 0)

		for _, trigger := range triggers {
			sound := find_sound(strings.TrimSpace(trigger))
			if sound != nil {
				sounds = append(sounds, sound)
			}
		}

		if len(sounds) == 0 {
			return
		}

		cooldownLock <- true

		go func() {
			time.Sleep(MSG_COOLDOWN)
			<-cooldownLock
		}()

		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			return
		}

		// Find the guild for that channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			// Could not find guild.
			return
		}

		if len(voiceLock) == 1 {
			return
		}

		voiceLock <- true

		callback(s, g, m, sounds)
		<-voiceLock
	}
}
