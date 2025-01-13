package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/sirupsen/logrus"
)

func handlerCmd(botHandler *telegohandler.BotHandler) {
	botHandler.Handle(func(bot *telego.Bot, update telego.Update) {
		logrus.Infof("User [%d] Request ID...", update.Message.From.ID)
		_, err := bot.SendMessage(&telego.SendMessageParams{
			ChatID:    update.Message.Chat.ChatID(),
			Text:      fmt.Sprintf("Your own ID is: *`%d`*", update.Message.From.ID),
			ParseMode: telego.ModeMarkdownV2,
			ReplyParameters: &telego.ReplyParameters{
				MessageID: update.Message.MessageID,
			},
		})
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
	}, telegohandler.CommandEqual("getid"))

	botHandler.Handle(func(bot *telego.Bot, update telego.Update) {
		logrus.Infof("User [%d] Request Group ID...", update.Message.From.ID)
		if update.Message.Chat.ID == update.Message.From.ID {
			_, _ = bot.SendMessage(&telego.SendMessageParams{
				ChatID:    update.Message.Chat.ChatID(),
				Text:      `⛔ *You're not in the group, you can't get the group ID\.*`,
				ParseMode: telego.ModeMarkdownV2,
				ReplyParameters: &telego.ReplyParameters{
					MessageID: update.Message.MessageID,
				},
			})
			return
		}
		var message = fmt.Sprintf("Your group ID is: *`%d`*", update.Message.Chat.ID)
		if update.Message.MessageThreadID != 0 {
			message = fmt.Sprintf("Your group ID is: *`%d`*\nYour group thread ID is: *`%d`*", update.Message.Chat.ID, update.Message.MessageThreadID)
		}
		_, err := bot.SendMessage(&telego.SendMessageParams{
			ChatID:    update.Message.Chat.ChatID(),
			Text:      message,
			ParseMode: telego.ModeMarkdownV2,
			ReplyParameters: &telego.ReplyParameters{
				MessageID: update.Message.MessageID,
			},
		})
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}

	}, telegohandler.CommandEqual("getgroupid"))
}
