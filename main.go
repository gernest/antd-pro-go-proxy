package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gernest/antd-pro-go-proxy/api"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "antd-pro-go-proxy"
	app.Usage = "a backend web server for antd-pro"
	app.Action = func(ctx *cli.Context) error {
		port := ctx.Int("port")
		h := api.Service()
		log.Printf("starting server at http://localhost:%d\n", port)
		return http.ListenAndServe(fmt.Sprintf(":%d", port), h)
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
