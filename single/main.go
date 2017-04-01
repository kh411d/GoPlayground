package main

import (
    "net/http"
    "os"
    "runtime"

    "github.com/facebookgo/grace/gracehttp"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    config "github.com/spf13/viper"
    "github.com/urfave/cli"
)

func SetConfig(port, configPath string) {
    config.SetConfigName("config")
    config.AddConfigPath("$GOPATH/src/github.com/kh411d/MyGoPlayground")

    if len(port) != 0 {
        config.Set("port", port)
    }

}

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())

}

func main() {
    app := cli.NewApp()

    app.Name = "Internal Tool"
    app.Usage = "Automator service"

    app.Commands = []cli.Command{
        {
            Name:        "cli-tool",
            Usage:       "Run tool",
            Description: ``,
            Action: func(c *cli.Context) error {
                SetConfig(c.String("port"), c.String("config"))

                e := echo.New()
                e.GET("/hello", func(c echo.Context) error {
                    return c.String(http.StatusOK, "/hello")
                })

                std := standard.New(":" + config.GetString("port"))
                std.SetHandler(echo)
                gracehttp.Serve(std.Server)
            },
        },
    }

    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "config, c",
            Value: "config.yaml",
            Usage: "Setting up config",
        },
        cli.StringFlag{
            Name:  "port, p",
            Value: "3000",
            Usage: "Setting up port",
        },
    }

    app.Flags = append(app.Flags, []cli.Flag{}...)
    app.Run(os.Args)
}
