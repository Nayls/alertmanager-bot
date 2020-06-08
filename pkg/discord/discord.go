package discord

import (
	_alertmanager "alertmanager-bot/pkg/alertmanager"
	_utils "alertmanager-bot/pkg/utils"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// GetStatusColor ...
func GetStatusColor(status string) int {
	switch _status := strings.ToLower(status); _status {
	case "firing":
		return 15746887 // color Red
	case "resolved":
		return 4437377 // color Green
	}
	return 16777215
}

func DiscordEmbedTemplate(alertman *_alertmanager.AlertManOut, discord *discordgo.WebhookParams) {
	for status, alerts := range _utils.GetGroupedAlerts(alertman) {
		embedsSummary := ""
		embedsDescription := ""
		embedsOwners := ""
		embedsOwnersSlice := []string{""}
		embedsDescriptionAlert := ""
		embedsGrafanaDashboard := ""
		embedsGrafanaDashboardURL := ""
		embedsRunbook := ""

		// Summary - alert summary
		if alertman.CommonAnnotations.Summary != "" {
			embedsSummary = fmt.Sprintf("Summary: %s\n", alertman.CommonAnnotations.Summary)
		}

		for _, alert := range alerts {
			// Embeds.Description - alert description
			if alert.Annotations.Description != "" {
				embedsDescription = fmt.Sprintf("Description: %s\n", alert.Annotations.Description)
			}

			embedsOwnersSlice = append(embedsOwnersSlice, _utils.GetOwnerSeverityMention(alert.Labels["severity"])...)

			// Embeds.Description - alert info
			embedsDescriptionAlert += "```"
			if alert.Labels["severity"] != "" {
				embedsDescriptionAlert += fmt.Sprintf("Severity: %s\n", alert.Labels["severity"])
			}
			if alert.Labels["job"] != "" {
				embedsDescriptionAlert += fmt.Sprintf("Job: %s\n", alert.Labels["job"])
			}
			if alert.Labels["pod"] != "" {
				embedsDescriptionAlert += fmt.Sprintf("Pod: %s\n", alert.Labels["pod"])
			}
			if alert.Labels["namespace"] != "" {
				embedsDescriptionAlert += fmt.Sprintf("Namespace: %s\n", alert.Labels["namespace"])
			}
			if alert.Labels["Instance"] != "" {
				embedsDescriptionAlert += fmt.Sprintf("Instance: %s\n", alert.Labels["instance"])
			}
			if alert.Labels["environment"] != "" {
				embedsDescriptionAlert += fmt.Sprintf("Environment: %s\n", alert.Labels["environment"])
			}
			embedsDescriptionAlert += "```"

			if alert.GeneratorURL != "" {
				embedsRunbook = fmt.Sprintf("[Runbook](%s)", alert.GeneratorURL)
			}
			if alert.Labels["grafana_dashboard_image"] != "" {
				embedsGrafanaDashboard = fmt.Sprintf("[Grafana](%s)", alert.Labels["grafana_dashboard_image"])
				embedsGrafanaDashboardURL = alert.Labels["grafana_dashboard_image"]
			}
		}

		// Embeds.Description - unique owners
		embedsOwners = _utils.GetUniqueOwners(embedsOwnersSlice)
		if embedsOwners != "" {
			discord.Content = fmt.Sprintf("%s %s\n", alertman.CommonLabels.Alertname, embedsOwners)
		}

		discord.Embeds = append(discord.Embeds, &discordgo.MessageEmbed{
			Title:       alertman.CommonLabels.Alertname,                                                                                                                                       // Discord.Embeds.Title
			Color:       GetStatusColor(status),                                                                                                                                                // Discord.Embeds.Color
			Description: fmt.Sprintf("%s %s %s %s %s %s", embedsSummary, embedsDescription, embedsOwners, embedsDescriptionAlert, embedsRunbook, embedsGrafanaDashboard), // Discord.Embeds.Description
			Image: &discordgo.MessageEmbedImage{
				URL: _utils.GrafanaDashboardToExternalURL(embedsGrafanaDashboardURL, viper.GetString("grafana.external_url")),
			},
		})
	}
}
