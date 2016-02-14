package main

import "strings"

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Game struct {
	score	int
	bot		*tgbotapi.BotAPI
	chatId	int
	dict	Dictionary
}

func MakeGame(bot *tgbotapi.BotAPI, chatId int, dict Dictionary) Game {
	return Game{score:0, bot:bot, chatId:chatId, dict:dict}
}

func (game *Game) Start() {
    game.bot.Send(tgbotapi.NewMessage(game.chatId, 
        "Starting a new game!"))
    firstRune := game.dict.RandomFirst()
    firstWord := game.dict.RandomWord(firstRune)
    game.bot.Send(tgbotapi.NewMessage(game.chatId, 
        firstWord))
}

func (game *Game) Turn(update tgbotapi.Update) {
    reply := game.answer(update.Message.Text)
    game.score++;
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
    game.bot.Send(msg)
}

func (game Game) answer(playerTurn string) string {
    words := strings.Fields(playerTurn)
    switch {
    case len(words) == 0:
        return "I don't see any words here."
    case len(words) > 1:
        return "There are two many words! Which do you choose?"
    }

    return game.dict.RandomWord([]rune(playerTurn)[len(playerTurn) - 1])
}

