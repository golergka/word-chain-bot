package main

import "fmt"
import "strings"

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Game struct {
	score	int
	bot		*tgbotapi.BotAPI
	chatId	int
	dict	Dictionary
	last	rune
}

func MakeGame(bot *tgbotapi.BotAPI, chatId int, dict Dictionary) Game {
	return Game{score:0, bot:bot, chatId:chatId, dict:dict}
}

func (game *Game) Start() {
    game.bot.Send(tgbotapi.NewMessage(game.chatId, 
        "Starting a new game!"))
    firstRune := game.dict.RandomFirst()
    firstWord := game.dict.RandomWord(firstRune)
	game.last = []rune(firstWord)[len(firstWord) - 1]
    game.bot.Send(tgbotapi.NewMessage(game.chatId, 
        firstWord))
}

func (game *Game) sendReply(reply string) {
	game.bot.Send(tgbotapi.NewMessage(game.chatId, reply))
}

func (game *Game) Turn(playerTurn string) {

    words := strings.Fields(playerTurn)

	switch
	{
    case len(words) == 0:
        game.sendReply("I don't see any words here.")
    case len(words) > 1:
        game.sendReply("There are two many words! Which do you choose?")
		break
	case !game.dict.Contains(playerTurn):
		game.sendReply(fmt.Sprintf("I don't know this word, %v", playerTurn))
		break
	case []rune(playerTurn)[0] != game.last:
		game.sendReply(fmt.Sprintf("You have to send the word starting with %v", game.last))
		break
	default: 
		game.score++;
		lastRune := []rune(playerTurn)[len(playerTurn) - 1]
		reply := game.dict.RandomWord(lastRune)
		game.sendReply(reply)
		game.last = []rune(reply)[len(reply) - 1]
	}
}
