package main

import (
	"context"
	"fmt"
	_ "github.com/mritd/logrus"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var botToken string

var rootCmd = &cobra.Command{
	Use: "tgid",
	RunE: func(cmd *cobra.Command, args []string) error {

		if botToken == "" {
			botToken = os.Getenv("BOT_TOKEN")
		}

		botCli, err := telego.NewBot(botToken, telego.WithDefaultLogger(false, true), telego.WithHealthCheck())
		if err != nil {
			return fmt.Errorf("failed to create telegram bot cli: %w", err)
		}

		botUpdates, err := botCli.UpdatesViaLongPolling(nil)
		if err != nil {
			return fmt.Errorf("failed to create telegram updates: %w", err)
		}

		botHandler, err := telegohandler.NewBotHandler(botCli, botUpdates)
		if err != nil {
			return fmt.Errorf("failed to create telegram handler: %w", err)
		}

		handlerCmd(botHandler)

		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		go botHandler.Start()
		logrus.Info("Telegram ID Bot Started...")

		<-ctx.Done()
		logrus.Info("Telegram ID Bot Shutting down...")
		botHandler.Stop()
		return nil
	},
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&botToken, "token", "", "Telegram Bot Token")
}
