package encodor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Message is a Telegram object that can be found in an update.
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// Update is a Telegram object that the handler receives every time an user
// interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// A Telegram Chat indicates the conversation to which the message belongs.
type Chat struct {
	Id int `json:"id"`
}

func beghiloszMapping(letter rune) rune {
	var char_map = map[rune]rune{
		'B': '8',
		'E': '3',
		'G': '6',
		'H': '4',
		'I': '1',
		'L': '7',
		'O': '0',
		'S': '5',
		'Z': '2',
	}
	new_letter, exists := char_map[letter]
	if !exists {
		return letter
	} else {
		return new_letter
	}
}

func Beghilosz(input string) string {
	words := strings.Fields(input)
	encoded_words := make([]string, len(words))
	for i, word := range words {
		word = strings.ToUpper(word)
		if !(strings.HasPrefix(word, "#") || strings.HasPrefix(word, "@")) {
			word = strings.Map(beghiloszMapping, word)
			word = reverse(word)
		}
		encoded_words[i] = word
	}
	encoded_words = reverseSlice(encoded_words)
	output := strings.Join(encoded_words, " ")
	return output
}

func reverse(input string) string {
	var result string
	for _, letter := range input {
		result = string(letter) + result
	}
	return result
}

func reverseSlice(input []string) []string {
	for i := 0; i < len(input)/2; i++ {
		input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
	}
	return input
}

// parseTelegramRequest handles incoming update from the Telegram web hook
func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

// HandleTelegramWebHook sends a message back to the chat in encoded form
func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	// Parse incoming request
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	//BEGHILOSZ encode incomming message
	encoded := Beghilosz(update.Message.Text)

	// Send the punchline back to Telegram
	var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, encoded)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, reponse body is %s",
			errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("Encoded text %s successfuly distributed to chat id %d",
			encoded, update.Message.Chat.Id)
	}
}

// sendTextToTelegramChat sends a text message to the Telegram chat identified
// by its chat Id
func sendTextToTelegramChat(chatId int, text string) (string, error) {

	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}
