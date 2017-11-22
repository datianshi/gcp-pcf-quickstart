/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ert

import (
	"encoding/json"
	"fmt"
	"omg-cli/config"
	"omg-cli/omg/tiles"
	"omg-cli/ops_manager"
)

type Properties struct {
	// Domains
	AppsDomain Value `json:".cloud_controller.apps_domain"`
	SysDomain  Value `json:".cloud_controller.system_domain"`
	// Networking
	NetworkingPointOfEntry    Value            `json:".properties.networking_point_of_entry"`
	TcpRouting                Value            `json:".properties.tcp_routing"`
	TcpRoutingReservablePorts Value            `json:".properties.tcp_routing.enable.reservable_ports"`
	GoRouterSSLCiphers        Value            `json:".properties.gorouter_ssl_ciphers"`
	HAProxySSLCiphers         Value            `json:".properties.haproxy_ssl_ciphers"`
	SkipSSLVerification       BooleanValue     `json:".ha_proxy.skip_cert_verify"`
	HAProxyForwardTLS         Value            `json:".properties.haproxy_forward_tls"`
	IngressCertificates       CertificateValue `json:".properties.networking_poe_ssl_cert"`
	// Application Containers
	ContainerDNSServers Value `json:".diego_cell.dns_servers"`
	// Application Security Groups
	SecurityAcknowledgement Value `json:".properties.security_acknowledgement"`
	// UAA
	ServiceProviderCredentials CertificateValue `json:".uaa.service_provider_key_credentials"`

	UaaDbChoice   Value        `json:".properties.uaa_database"`
	UaaDbIp       Value        `json:".properties.uaa_database.external.host"`
	UaaDbPort     IntegerValue `json:".properties.uaa_database.external.port"`
	UaaDbUsername Value        `json:".properties.uaa_database.external.uaa_username"`
	UaaDbPassword SecretValue  `json:".properties.uaa_database.external.uaa_password"`

	// Databases
	ErtDbChoice Value        `json:".properties.system_database"`
	ErtDbIp     Value        `json:".properties.system_database.external.host"`
	ErtDbPort   IntegerValue `json:".properties.system_database.external.port"`

	ErtDbAppUsageUsername            Value       `json:".properties.system_database.external.app_usage_service_username"`
	ErtDbAppUsagePassword            SecretValue `json:".properties.system_database.external.app_usage_service_password"`
	ErtDbAutoscaleUsername           Value       `json:".properties.system_database.external.autoscale_username"`
	ErtDbAutoscalePassword           SecretValue `json:".properties.system_database.external.autoscale_password"`
	ErtDbCloudControllerUsername     Value       `json:".properties.system_database.external.ccdb_username"`
	ErtDbCloudControllerPassword     SecretValue `json:".properties.system_database.external.ccdb_password"`
	ErtDbDiegoUsername               Value       `json:".properties.system_database.external.diego_username"`
	ErtDbDiegoPassword               SecretValue `json:".properties.system_database.external.diego_password"`
	ErtDbLocketUsername              Value       `json:".properties.system_database.external.locket_username"`
	ErtDbLocketPassword              SecretValue `json:".properties.system_database.external.locket_password"`
	ErtDbNetworkPolicyServerUsername Value       `json:".properties.system_database.external.networkpolicyserver_username"`
	ErtDbNetworkPolicyServerPassword SecretValue `json:".properties.system_database.external.networkpolicyserver_password"`
	ErtDbNfsUsername                 Value       `json:".properties.system_database.external.nfsvolume_username"`
	ErtDbNfsPassword                 SecretValue `json:".properties.system_database.external.nfsvolume_password"`
	ErtDbNotificationsUsername       Value       `json:".properties.system_database.external.notifications_username"`
	ErtDbNotificationsPassword       SecretValue `json:".properties.system_database.external.notifications_password"`
	ErtDbAccountUsername             Value       `json:".properties.system_database.external.account_username"`
	ErtDbAccountPassword             SecretValue `json:".properties.system_database.external.account_password"`
	ErtDbRoutingUsername             Value       `json:".properties.system_database.external.routing_username"`
	ErtDbRoutingPassword             SecretValue `json:".properties.system_database.external.routing_password"`
	ErtDbSilkUsername                Value       `json:".properties.system_database.external.silk_username"`
	ErtDbSilkPassword                SecretValue `json:".properties.system_database.external.silk_password"`

	// MySQL
	MySqlMonitorRecipientEmail Value `json:".mysql_monitor.recipient_email"`
}

type Value struct {
	Value string `json:"value"`
}

type IntegerValue struct {
	Value int `json:"value"`
}

type BooleanValue struct {
	Value bool `json:"value"`
}

type Secret struct {
	Value string `json:"secret"`
}

type SecretValue struct {
	Sec Secret `json:"value"`
}

type Certificate struct {
	PublicKey  string `json:"cert_pem"`
	PrivateKey string `json:"private_key_pem"`
}

type CertificateValue struct {
	Value Certificate `json:"value"`
}

type Resources struct {
	TcpRouter                    Resource `json:"tcp_router"`
	Router                       Resource `json:"router"`
	DiegoBrain                   Resource `json:"diego_brain"`
	ConsulServer                 Resource `json:"consul_server"`
	Nats                         Resource `json:"nats"`
	NfsServer                    Resource `json:"nfs_server"`
	MysqlProxy                   Resource `json:"mysql_proxy"`
	Mysql                        Resource `json:"mysql"`
	BackupPrepare                Resource `json:"backup-prepare"`
	DiegoDatabase                Resource `json:"diego_database"`
	Uaa                          Resource `json:"uaa"`
	CloudController              Resource `json:"cloud_controller"`
	HaProxy                      Resource `json:"ha_proxy"`
	MysqlMonitor                 Resource `json:"mysql_monitor"`
	ClockGlobal                  Resource `json:"clock_global"`
	CloudControllerWorker        Resource `json:"cloud_controller_worker"`
	DiegoCell                    Resource `json:"diego_cell"`
	LoggregatorTrafficcontroller Resource `json:"loggregator_trafficcontroller"`
	SyslogAdapter                Resource `json:"syslog_adapter"`
	SyslogScheduler              Resource `json:"syslog_scheduler"`
	Doppler                      Resource `json:"doppler"`
	SmokeTests                   Resource `json:"smoke-tests"`
	PushAppsManager              Resource `json:"push-apps-manager"`
	Notifications                Resource `json:"notifications"`
	NotificationsUi              Resource `json:"notifications-ui"`
	PushPivotalAccount           Resource `json:"push-pivotal-account"`
	PushUsageService             Resource `json:"push-usage-service"`
	Autoscaling                  Resource `json:"autoscaling"`
	AutoscalingRegisterBroker    Resource `json:"autoscaling-register-broker"`
	Nfsbrokerpush                Resource `json:"nfsbrokerpush"`
	Bootstrap                    Resource `json:"bootstrap"`
	MysqlRejoinUnsafe            Resource `json:"mysql-rejoin-unsafe"`
}

type Resource struct {
	RouterNames       []string `json:"elb_names,omitempty"`
	Instances         int      `json:"instances,omitempty"`
	InternetConnected bool     `json:"internet_connected"`
}

func (*Tile) Configure(cfg *config.Config, om *ops_manager.Sdk) error {
	if err := om.StageProduct(tile.Product); err != nil {
		return err
	}

	network := tiles.NetworkConfig(cfg.ErtSubnetName, cfg)

	networkBytes, err := json.Marshal(&network)
	if err != nil {
		return err
	}

	properties := Properties{
		AppsDomain:                 Value{cfg.AppsDomain},
		SysDomain:                  Value{cfg.SysDomain},
		NetworkingPointOfEntry:     Value{"external_non_ssl"},
		ContainerDNSServers:        Value{"8.8.8.8,8.8.4.4"},
		SkipSSLVerification:        BooleanValue{true},
		HAProxyForwardTLS:          Value{"disable"},
		IngressCertificates:        CertificateValue{Certificate{cfg.SslCertificate, cfg.SslPrivateKey}},
		TcpRouting:                 Value{"enable"},
		TcpRoutingReservablePorts:  Value{cfg.TcpPortRange},
		GoRouterSSLCiphers:         Value{"ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384"},
		HAProxySSLCiphers:          Value{"DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384"},
		SecurityAcknowledgement:    Value{"X"},
		ServiceProviderCredentials: CertificateValue{Certificate{cfg.SslCertificate, cfg.SslPrivateKey}},

		UaaDbChoice:   Value{"external"},
		UaaDbIp:       Value{cfg.ExternalSqlIp},
		UaaDbPort:     IntegerValue{cfg.ExternalSqlPort},
		UaaDbUsername: Value{cfg.ERTSqlUsername},
		UaaDbPassword: SecretValue{Secret{cfg.ERTSqlPassword}},

		ErtDbChoice:                      Value{"external"},
		ErtDbIp:                          Value{cfg.ExternalSqlIp},
		ErtDbPort:                        IntegerValue{cfg.ExternalSqlPort},
		ErtDbAppUsageUsername:            Value{cfg.ERTSqlUsername},
		ErtDbAppUsagePassword:            SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbAutoscaleUsername:           Value{cfg.ERTSqlUsername},
		ErtDbAutoscalePassword:           SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbCloudControllerUsername:     Value{cfg.ERTSqlUsername},
		ErtDbCloudControllerPassword:     SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbDiegoUsername:               Value{cfg.ERTSqlUsername},
		ErtDbDiegoPassword:               SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbLocketUsername:              Value{cfg.ERTSqlUsername},
		ErtDbLocketPassword:              SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbNetworkPolicyServerUsername: Value{cfg.ERTSqlUsername},
		ErtDbNetworkPolicyServerPassword: SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbNfsUsername:                 Value{cfg.ERTSqlUsername},
		ErtDbNfsPassword:                 SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbNotificationsUsername:       Value{cfg.ERTSqlUsername},
		ErtDbNotificationsPassword:       SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbAccountUsername:             Value{cfg.ERTSqlUsername},
		ErtDbAccountPassword:             SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbRoutingUsername:             Value{cfg.ERTSqlUsername},
		ErtDbRoutingPassword:             SecretValue{Secret{cfg.ERTSqlPassword}},
		ErtDbSilkUsername:                Value{cfg.ERTSqlUsername},
		ErtDbSilkPassword:                SecretValue{Secret{cfg.ERTSqlPassword}},

		MySqlMonitorRecipientEmail: Value{"admin@example.org"},
	}

	propertiesBytes, err := json.Marshal(&properties)
	if err != nil {
		return err
	}

	resoruces := Resources{
		TcpRouter: Resource{
			RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.TcpTargetPoolName)},
			InternetConnected: false,
			Instances:         3,
		},
		Router: Resource{
			RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.WssTargetPoolName), fmt.Sprintf("http:%s", cfg.HttpBackendServiceName)},
			InternetConnected: false,
		},
		DiegoBrain: Resource{
			RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.SshTargetPoolName)},
			InternetConnected: false,
		},
	}
	resorucesBytes, err := json.Marshal(&resoruces)
	if err != nil {
		return err
	}

	return om.ConfigureProduct(tile.Product.Name, string(networkBytes), string(propertiesBytes), string(resorucesBytes))
}
