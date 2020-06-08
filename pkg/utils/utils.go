package utils

import (
	_alertmanager "alertmanager-bot/pkg/alertmanager"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// GetOwnerSeverityMention ...
func GetOwnerSeverityMention(severity string) []string {
	owners := viper.GetStringSlice("severity." + strings.ToLower(severity))

	if owners != nil {
		return owners
	}

	return nil
}

// GetUniqueStringSlice return []string with unique value
func GetUniqueStringSlice(payload []string) []string {
	check := make(map[string]int)
	d := append(payload)
	res := make([]string, 0)

	for _, val := range d {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res, letter)
	}

	return res
}

// GetUniqueOwners return string with unique value
func GetUniqueOwners(payload []string) string {
	owners := GetUniqueStringSlice(payload)
	res := ""

	for index, owner := range owners {
		if index != len(owner)-1 {
			res += fmt.Sprintf("%v ", owner) // Discord.Embed.Description
		} else {
			res += fmt.Sprintf("%v\n", owner) // Discord.Embed.Description
		}
	}

	return res
}

// GetGroupedAlerts ...
func GetGroupedAlerts(payload *_alertmanager.AlertManOut) map[string][]_alertmanager.AlertManAlert {
	res := make(map[string][]_alertmanager.AlertManAlert, 0)

	for _, alert := range payload.Alerts {
		res[alert.Status] = append(res[alert.Status], *alert)
	}

	return res
}
