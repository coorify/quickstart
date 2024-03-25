package main

import (
	"embed"
	"io/fs"
	"os"
	"os/signal"
	"syscall"

	"github.com/coorify/backend"
	"github.com/coorify/quickstart/option"
	"github.com/coorify/quickstart/plugin"
	"github.com/coorify/quickstart/router"
	"github.com/jinzhu/configor"
	_ "github.com/joho/godotenv/autoload"
)

//go:embed web/dist/*
var frontend embed.FS

func loadOpt(opt interface{}) error {
	loader := configor.New(&configor.Config{
		ENVPrefix: "BE",
	})

	files := os.Getenv("BE_CONFIG_FILE")
	if files == "" {
		files = "config.yml"
	}

	return loader.Load(opt, files)
}

func main() {
	opt := &option.Option{}
	if err := loadOpt(opt); err != nil {
		panic(err)
	}

	app := backend.NewServer(opt)

	fe, _ := fs.Sub(frontend, "web/dist")
	app.Frontend(fe)

	plugin.Setup(app)
	router.Setup(app)

	// as desktop
	// if err := desktop.Run(app); err != nil {
	// 	panic(err)
	// }

	if err := app.Start(); err != nil {
		panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT)
	if err := app.Stop(syscall.SIGQUIT != <-sigchan); err != nil {
		panic(err)
	}
}
