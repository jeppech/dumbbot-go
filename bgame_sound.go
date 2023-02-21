package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func create_bgame_command(s *discordgo.Session) {
	var commands = []*discordgo.ApplicationCommand{
		{
			Name:        "bgame",
			Description: "No fuck you bloody",
			Type:        discordgo.ChatApplicationCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "say",
					Description:  "Det må du selv finde ud af",
					Required:     true,
					Autocomplete: true,
				},
			},
		},
	}

	var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"bgame": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			options := i.ApplicationCommandData().Options
			option_map := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

			for _, opt := range options {
				option_map[opt.Name] = opt
			}

			if i.Type == discordgo.InteractionApplicationCommandAutocomplete {
				choices := make([]*discordgo.ApplicationCommandOptionChoice, 0)
				if option, ok := option_map["say"]; ok {
					msg := strings.ToLower(option.StringValue())

					if len(msg) > 0 {
						parts := strings.Split(msg, ",")
						last := parts[len(parts)-1]
						if len(parts) > 1 {
							if len(last) > 0 {
								rest := strings.Join(parts[:len(parts)-1], ",")
								for _, sound_name := range soundList {
									if strings.HasPrefix(sound_name.Name, last) {
										choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
											Name:  fmt.Sprintf("%s,%s", rest, sound_name.Name),
											Value: fmt.Sprintf("%s,%s", rest, sound_name.Name),
										})
									}
								}
							}
						} else {
							for _, sound_name := range soundList {
								if strings.HasPrefix(sound_name.Name, last) {
									choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
										Name:  sound_name.Name,
										Value: sound_name.Name,
									})
								}
							}
						}
					} else {
						for i, sound_name := range soundList {
							if i == 24 {
								break
							}
							choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
								Name:  sound_name.Name,
								Value: sound_name.Name,
							})
						}
					}
				}

				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionApplicationCommandAutocompleteResult,
					Data: &discordgo.InteractionResponseData{
						Choices: choices,
					},
				})

				if err != nil {
					fmt.Println("Error responding to interaction:", err)
				}
				return
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsEphemeral,
					Content: "Tjek!",
				},
			})

			if err != nil {
				fmt.Println("Error responding to interaction:", err)
			}

			if option, ok := option_map["say"]; ok {
				msg := strings.ToLower(option.StringValue())
				sound_names := strings.Split(msg, ",")
				sounds := make([]*Sound, 0)

				for _, name := range sound_names {
					sound := find_sound(strings.TrimSpace(name))
					if sound != nil {
						sounds = append(sounds, sound)
					}
				}

				if len(sounds) == 0 {
					return
				}

				g, err := s.State.Guild(i.GuildID)

				if err != nil {
					fmt.Println("Error getting guild:", err)
					return
				}

				var default_chan_id string
				for _, gc := range g.Channels {
					if strings.ToLower(gc.Name) == "generel" {
						default_chan_id = gc.ID
					}
				}

				vs := find_user_voice_channel(g.VoiceStates, i.Member.User.ID, default_chan_id)

				if vs == nil {
					s.InteractionResponseDelete(i.Interaction)
					return
				}

				vc, err := s.ChannelVoiceJoin(g.ID, vs.ChannelID, false, true)

				if err != nil {
					fmt.Println("Error joining voice channel:", err)
					return
				}

				fmt.Printf("%s (%s) is playing %d sounds \n", i.Member.User.Username, i.Member.Nick, len(sounds))
				playSounds(s, g.ID, vs.ChannelID, sounds)

				vc.Disconnect()
			}

			s.InteractionResponseDelete(i.Interaction)
		},
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)

		if err != nil {
			fmt.Printf("Error creating command '%v': %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}
