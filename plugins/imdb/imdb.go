package imdb

import (
	"encoding/json"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"io"
	"log"
	"strings"
)

type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Duration string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Actors   string `json:"Actors"`
	Id       string `json:"imdbID"`
	Image    string `json:"Poster"`
}

type IMDBPlugin struct {
	//whatever
}

func (m *IMDBPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {
	if strings.Contains(message.Text, config.CommandPrefix+"imdb") {
		imdbsearch := message.Text
		log.Println("Searching " + imdbsearch)
		imdbsearch = strings.Replace(imdbsearch, config.CommandPrefix+"imdb ", "", -1)
		imdbsearch = strings.Replace(imdbsearch, " ", "%20", -1)

		util.DecodeJson("http://www.imdbapi.com/?t="+imdbsearch, func(body io.ReadCloser) bool {
			var imdb Movie
			err := json.NewDecoder(body).Decode(&imdb)

			bot.SendMessage(message.Chat, imdb.Title+" - "+imdb.Year+"\n"+imdb.Genre+" - "+imdb.Duration+"\n"+imdb.Image+"\n"+"http://imdb.com/title/"+imdb.Id, nil)

			if err != nil {
				return false
			} else {
				return true
			}
		})

	}

}

func init() {
	plugin_registry.RegisterPlugin(&IMDBPlugin{})
	plugin_registry.RegisterCommand("imdb", "Search a movie on imdb")
}
