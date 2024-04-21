package cmd

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gopkg.in/telebot.v3"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the application",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	_ = godotenv.Load()
	pref := telebot.Settings{
		Token:  os.Getenv("BOT_API"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(telebot.OnText, func(tCtx telebot.Context) error {
		return tCtx.Reply("Hello World")
	})

	b.Start()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
