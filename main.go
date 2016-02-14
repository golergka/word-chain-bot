package main

import "log"
import "os"
import "bufio"
import "strings"

import "github.com/go-telegram-bot-api/telegram-bot-api"

func makeDictionaryFromPath(path string) (*Dictionary, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    return MakeDictionary(bufio.NewScanner(file))
}

func main() {
    englishNounDict, err := makeDictionaryFromPath("english_nouns.txt")
    if err != nil {
        log.Panic(err)
    }

    token := os.Getenv("BOT_TOKEN")
    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Panic(err)
    }
    log.Printf("Authorized on account %s", bot.Self.UserName)

    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60

    updates, err := bot.GetUpdatesChan(updateConfig)
    if err != nil {
        log.Panic(err)
    }

    for update := range updates {
        go func() {
            log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

            if strings.Index(update.Message.Text, "/start") == 0 {
                handleStartCommand(bot, update)
            } else if game, ok := games[update.Message.Chat.ID]; ok {
                game.Turn(update)
            } else {
                newGame := MakeGame(bot, update.Message.Chat.ID, *englishNounDict)
                newGame.Start()
                games[update.Message.Chat.ID] = newGame
            }
        }()
    }
}

var games map[int]Game = make(map[int]Game)

func handleStartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    reply := "Hi. This is a simple word chain game bot.\n\n" +
        "Rules are simple: you must reply with a word that starts with a letter the previous word ended with.\n\n" +
        "Just type a word to start playing!"
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
    bot.Send(msg)
}
