package main

import "log"
import "os"
import "bufio"
import "strings"

import "github.com/go-telegram-bot-api/telegram-bot-api"

var englishNouns []string

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines[] string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func main() {
    var err error

    englishNouns, err = readLines("english_nouns.txt")
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

    for u := range updates {
        go handleUpdate(bot, u)
    }
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    log.Printf("[%s] %s",
            update.Message.From.UserName, 
            update.Message.Text)

    if (strings.Index(update.Message.Text, "/start") == 0) {
        handleStartCommand(bot, update)
    } else {
        handleTurn(bot, update)
    }
}

func handleTurn(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    reply := gameTurn(update.Message.Text)
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
    bot.Send(msg)
}

func gameTurn(playerTurn string) string {
    words := strings.Fields(playerTurn)
    switch {
    case len(words) == 0:
        return "I don't see any words here."
    case len(words) > 1:
        return "There are two many words! Which do you choose?"
    }

    return "Uhm, I guess I should give you some word."
}

func handleStartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    reply := "Hi. This is a simple word chain game bot.\n\n" +
        "Rules are simple: you must reply with a word that starts with a letter the previous word ended with.\n\n" +
        "Just type a word to start playing!"
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
    bot.Send(msg)
}
