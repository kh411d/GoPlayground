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
            Name:        "tool",
            Usage:       "Run tool",
            Description: ``,
            Action: func(c *cli.Context) {
                SetConfig(c.String("port"), c.String("config"))

                e := echo.New()
                e.POST("/hello", func(c echo.Context) error {
                    type Bodyraw struct {
                        Payload string `json:"payload"`
                    }
                    opt := new(Bodyraw)
                    if err := c.Bind(&opt); err != nil {
                        return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON input(parse error).")
                    }

                    return c.String(http.StatusOK, opt.Payload)
                })

                std := standard.New(":" + config.GetString("port"))
                std.SetHandler(e)
                gracehttp.Serve(std.Server)
            },
            Flags: []cli.Flag{
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
            },
        },
    }

    app.Run(os.Args)
}
