package gag

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot/yatzie/shared/registry"
	"github.com/go-telegram-bot/yatzie/shared/utils"

	"github.com/tucnak/telebot"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type GagJson struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Src   string `json:"src"`
}
type GagsJson []GagJson

type GagPlugin struct {
}

func (m *GagPlugin) Run(bot *telebot.Bot, config util.Config, message telebot.Message) {

	if message.Text == config.CommandPrefix+"gag" {
		gags, err := getImages("http://api-9gag.herokuapp.com/")
		if err != nil {
			//bot.SendMessage(message.Chat, strings.Replace(gags[rand.Intn(len(gags))].Src, `\/`, "/", -1), nil)
			util.SendPhoto(strings.Replace(gags[rand.Intn(len(gags))].Src, `\/`, "/", -1), message, bot)

		}

	}

}

func getImages(url string) (GagsJson, error) {
	var data GagsJson
	r, err := http.Get(url)
	fmt.Println(url)

	if err != nil {
		return data, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	for _, i := range data {
		log.Println(
			url + i.Src)
	}
	return data, err
}

func init() {
	my := &GagPlugin{}
	plugin_registry.RegisterPlugin(my)
	plugin_registry.RegisterCommand("gag", "Display some random gag ")

}
