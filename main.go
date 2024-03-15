package main

import (
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

// //go:embed web/dist
// var frontend embed.FS

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
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT)

	opt := &option.Option{}
	if err := loadOpt(opt); err != nil {
		panic(err)
	}

	app := backend.NewServer(opt)

	// fe, _ := fs.Sub(frontend, "web/dist")
	// app.Frontend(fe)

	plugin.Setup(app)
	router.Setup(app)

	if err := app.Start(); err != nil {
		panic(err)
	}

	s := <-sigchan

	if err := app.Stop(s != syscall.SIGQUIT); err != nil {
		panic(err)
	}
}
