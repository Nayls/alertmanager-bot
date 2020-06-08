// URLs provide a [uniform way to locate resources](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Here's how to parse URLs in Go.

package main

import (
	_utils "alertmanager-bot/pkg/utils"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	dashboard_url := "https://grafana.localhost:9090/render/d-solo/a87fb0d919ec0ea5f6543124e16c42a5/kubernetes-compute-resources-namespace-workloads?orgId=1&refresh=10s&from=1591256966344&to=1591260566344&var-datasource=default&var-type=deployment&var-cluster=&var-namespace=autoqa&panelId=2&width=1000&height=500&tz=Europe%2FMoscow"
	dashboard_app_url := "http://localhost:8080/render/d-solo/a87fb0d919ec0ea5f6543124e16c42a5/kubernetes-compute-resources-namespace-workloads?orgId=1&refresh=10s&from=1591256966344&to=1591260566344&var-datasource=default&var-type=deployment&var-cluster=&var-namespace=autoqa&panelId=2&width=1000&height=500&tz=Europe%2FMoscow"

	_utils.GrafanaDashboardToInternalURL(dashboard_url, viper.GetString("grafana.basic_username"), viper.GetString("grafana.basic_password"))
	_utils.GrafanaDashboardToExternalURL(dashboard_url, viper.GetString("grafana.external_url"))
	_utils.GrafanaDashboardExternalToInternalURL(dashboard_app_url, viper.GetString("grafana.internal_url"), viper.GetString("grafana.basic_username"), viper.GetString("grafana.basic_password"))
}
