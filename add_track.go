package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Sound struct {
	Name        string
	Description string
	Data        [][]byte
}

var soundList []*Sound

// adds a track to the soundMap
func add_track(trigger string, filename string, description string) {
	sound := &Sound{
		Name:        trigger,
		Description: description,
	}

	file, err := os.Open("./sounds/" + filename)
	if err != nil {
		log.Fatalln("Error opening dca file :", err)
	}

	var opuslen int16

	buf := make([][]byte, 0)
	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				log.Fatalln(err)
			}
			break
		}

		if err != nil {
			log.Fatalln("Error reading from dca file :", err)
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			log.Fatalln("Error reading from dca file :", err)
		}

		// Append encoded pcm data to the buffer.
		buf = append(buf, InBuf)
	}

	sound.Data = buf
	soundList = append(soundList, sound)
	fmt.Println(sound.Name, "-", sound.Description)
}

func find_user_voice_channel(vs []*discordgo.VoiceState, userID string, channelID string) *discordgo.VoiceState {
	for _, v := range vs {
		// if v.ChannelID == channelID {
		// 	return v
		// }
		if v.UserID == userID {
			return v
		}
	}
	return nil
}

func find_sound(name string) *Sound {
	for _, s := range soundList {
		if s.Name == name {
			return s
		}
	}
	return nil
}
