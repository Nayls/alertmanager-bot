package main

import (
	_alertmanager "alertmanager-bot/pkg/alertmanager"
	_discord "alertmanager-bot/pkg/discord"
	_utils "alertmanager-bot/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/resty.v1"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Config create private data struct to hold config options.
type Config struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
	Debug    string `yaml:"debug"`
	Discord  struct {
		Username  string `yaml:"username"`
		Webhook   string `yaml:"webhook"`
		AvatarURL string `yaml:"avatar_url"`
	} `yaml:"discord"`
	Grafana struct {
		InternalURL   string `yaml:"internal_url"`
		ExternalURL   string `yaml:"external_url"`
		BasicUsername string `yaml:"basic_username"`
		BasicPassword string `yaml:"basic_password"`
	} `yaml:"grafana"`
}

// Create a new config instance.
var (
	conf *Config
)

// Read the config file from the current directory and marshal
// into the conf config struct.
func getConf() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("GetParseURL parse error: ", err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		log.Fatal("GetParseURL parse error: ", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("APP")

	viper.SetDefault("hostname", "")
	viper.SetDefault("port", "8080")
	viper.SetDefault("debug", "false")
	viper.SetDefault("discord.username", "Discord AlertBot")
	viper.SetDefault("discord.webhook", "")
	viper.SetDefault("discord.avatar_url", "")
	viper.SetDefault("grafana.internal_url", "http://localhost:8080/render/d-solo/")
	viper.SetDefault("grafana.external_url", "")
	viper.SetDefault("grafana.basic_username", "")
	viper.SetDefault("grafana.basic_password", "")

	conf.Hostname = viper.GetString("hostname")
	conf.Port = viper.GetString("port")
	conf.Debug = viper.GetString("debug")
	conf.Discord = struct {
		Username  string `yaml:"username"`
		Webhook   string `yaml:"webhook"`
		AvatarURL string `yaml:"avatar_url"`
	}{
		Username:  viper.GetString("discord.username"),
		Webhook:   viper.GetString("discord.webhook"),
		AvatarURL: viper.GetString("discord.avatar_url"),
	}
	conf.Grafana = struct {
		InternalURL   string `yaml:"internal_url"`
		ExternalURL   string `yaml:"external_url"`
		BasicUsername string `yaml:"basic_username"`
		BasicPassword string `yaml:"basic_password"`
	}{
		InternalURL:   viper.GetString("grafana.internal_url"),
		ExternalURL:   viper.GetString("grafana.external_url"),
		BasicUsername: viper.GetString("grafana.basic_username"),
		BasicPassword: viper.GetString("grafana.basic_password"),
	}

	return conf
}

// Initialization routine.
func init() {
	conf = getConf()
	log.Printf("Alermanager-Bot is Started")
	if conf.Discord.Webhook == "" {
		log.Printf("discord.webhook not found, default value is used \"%s\"", conf.Discord.Webhook)
	}
	if conf.Grafana.InternalURL == "" {
		log.Printf("grafana.internal_url not found, default value is used \"%s\"", conf.Grafana.InternalURL)
	}
	if conf.Grafana.ExternalURL == "" {
		log.Printf("grafana.external_url not found, default value is used \"%s\"", conf.Grafana.ExternalURL)
	}
}

func main() {
	if conf.Debug == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	grafana := r.Group("/render")
	{
		grafana.GET("/d-solo/*anypath", func(c *gin.Context) {
			dashboard_url := _utils.GrafanaDashboardExternalToInternalURL(c.Request.URL.String(), viper.GetString("grafana.internal_url"), viper.GetString("grafana.basic_username"), viper.GetString("grafana.basic_password"))

			res, _ := resty.R().
				SetHeader("Content-Type", "image/png").
				SetHeader("Accept", "image/png").
				Get(dashboard_url)

			if res.StatusCode() == http.StatusOK {
				c.Writer.Write(res.Body())
			}
		})
	}

	bot := r.Group("/discord")
	{
		bot.POST("/embeds", func(c *gin.Context) {
			b, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Panicln(err)
				return
			}
			if conf.Debug == "true" {
				log.Printf("[GIN-debug] /discord/embeds Request.Body: %v\n", string(b))
			}

			alertman := _alertmanager.AlertManOut{}
			alertman.Alerts = make([]*_alertmanager.AlertManAlert, 0)
			err = json.Unmarshal(b, &alertman)
			if err != nil {
				log.Panicln(err)
				return
			}

			// ALERT TEMPLATE
			discord := discordgo.WebhookParams{}
			discord.Embeds = make([]*discordgo.MessageEmbed, 0)
			discord.Username = conf.Discord.Username   // Discord.Username
			discord.AvatarURL = conf.Discord.AvatarURL // Discord.AvatarURL

			// Заполнение информации в Discord.Embeds
			_discord.DiscordEmbedTemplate(&alertman, &discord)

			// Отправка в Discord
			_, _ = resty.R().
				SetHeader("Content-Type", "application/json").
				SetBody(discord).
				Post(conf.Discord.Webhook)

			if gin.IsDebugging() {
				jsonAlertman, err := json.MarshalIndent(alertman, "", "  ")
				if err != nil {
					log.Panicln(err)
					return
				} else {
					log.Printf("[GIN-debug] jsonAlertman error:\n\"alertmanager\":%v\n", string(jsonAlertman))
				}

				jsonDiscord, err := json.MarshalIndent(alertman, "", "  ")
				if err != nil {
					log.Panicln(err)
					return
				} else {
					log.Printf("[GIN-debug] jsonDiscord error:\n\"discord\":%v\n", string(jsonDiscord))
				}
			}

			c.JSON(http.StatusOK, map[string]interface{}{
				"alertmanager": alertman,
				"discord":      discord,
			})
		})
	}

	log.Fatal(r.Run(conf.Hostname + ":" + conf.Port))
}
