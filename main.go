package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type MessageAuthorCallback func(s *discordgo.Session, g *discordgo.Guild, m *discordgo.MessageCreate, sounds []*Sound)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

const MSG_COOLDOWN = 2 * time.Second

var token string

func main() {
	if token == "" {
		fmt.Println("No token provided. Please run: bgame -t <bot token>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	add_track("mlg", "horn.dca", "BWEEW BW BW BWEEEEEWWWW")
	add_track("sepan", "sepan.dca", "SEPAN!")
	add_track("youlie", "youlie.dca", "YOU LIE!")
	add_track("waffor", "waffornoe.dca", "Wa'for noe?")
	add_track("waffor2", "waffornoe2.dca", "Wa'for noe?")
	add_track("hund", "hund.dca", "Det må simpelthen være en langhåret hund")
	add_track("afrifugl", "afrikansk.dca", "Det må være en afrikans fugl?")
	add_track("and", "and.dca", "And! jeg sagde and!")
	add_track("viskavinde", "viskavinde.dca", "Vi ska v-v-vii vi ska prøve den 50 kroner")
	add_track("wasadusaa", "wasadusaa.dca", "Wa' saa' du' såå?")
	add_track("klarlars", "klarlars.dca", "Er du klar til et spørgsmål mere lars?")
	add_track("haehaehae", "haehaehae.dca", "haehaehae")
	add_track("kanin", "kanin.dca", "En kanin? haehae uhaehue ahhhaa")
	add_track("jegspiste", "jegspiste.dca", "åårh hvordan sagde, mmhjeg spist nogen gange en")
	add_track("kanarie", "kanarie.dca", "Kanariefugl?")
	add_track("barber", "barberfaar.dca", "En maskinøøh, der barber hårene på en får")
	add_track("stutteri", "jegkenderikk.dca", "hahaehae, jeg kender ikke så maaajg")
	add_track("eriklar", "jaer.dca", "Er i klar deromme? jaaeer")
	add_track("coke", "coke.dca", "JEG SKAL DA HA NOGET COKE")
	add_track("ikksejt", "ikksejt.dca", "Det er slet ikk sejt!")
	add_track("detgoerdetikk", "detgoerdetikk.dca", "Det gør det ik-ke!")
	add_track("deeznutz", "deeznutz.dca", "Deez Nutz!")
	add_track("fuckbloody", "fuckbloody.dca", "No you fuck bloody basterd")
	add_track("adresse", "adresse.dca", "Mandeøvænget")
	add_track("fugleinfluenza", "fugleinfluenza.dca", "Jajaeeee tænker på eeh fujhler influenza")
	add_track("hvemervaek", "hvemervaek.dca", "Hvem er væk?")
	add_track("nr31", "nr31.dca", "Nummer 31?")
	add_track("solsort", "solsort.dca", "Det er lisom sol SORT")
	add_track("wasaydu", "wasaydu.dca", "Wha say dy?")
	add_track("iorden", "iorden.dca", "Nåaår okay, det i orden")
	add_track("fugleinfluenza2", "fugleinfluenza2.dca", "Men hva me fugler influensja")
	add_track("hvisdukoere", "hvis-du-koere.dca", "Hvis du køre, så stikker af..")
	add_track("buresinde", "du-skal-bures-inde.dca", "Du skal fande bures inde mand")
	add_track("nyekoner", "nyekoner.dca", "jeg havde andre interesser hahehaeehaee")
	add_track("havremaelk", "havremaelk.dca", "du burde skydes")
	add_track("wc", "wc.dca", "det svineri vil jeg ikk ha")
	add_track("gayyy", "gayyy.dca", "hah gayyyyyy")

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)
	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	create_bgame_command(dg)

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(guildCreate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("BGame.gl bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, "no you bloody!")
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			// _, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready!")
			return
		}
	}
}

func playSounds(s *discordgo.Session, guildID, channelID string, sounds []*Sound) {
	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)

	if err != nil {
		fmt.Println("Error joining voice channel:", err)
		return
	}

	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	vc.Speaking(true)

	for _, sound := range sounds {
		// load the sound if it is not already loaded

		// Send the buffer data.
		for _, buff := range sound.Data {
			vc.OpusSend <- buff
		}

		time.Sleep(150 * time.Millisecond)
	}

	// Stop speaking
	vc.Speaking(false)
}
