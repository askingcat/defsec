package database

import (
	defsecTypes "github.com/aquasecurity/defsec/pkg/types"
)

type Database struct {
	MSSQLServers      []MSSQLServer
	MariaDBServers    []MariaDBServer
	MySQLServers      []MySQLServer
	PostgreSQLServers []PostgreSQLServer
}

type MariaDBServer struct {
	defsecTypes.Metadata
	Server
}

type MySQLServer struct {
	defsecTypes.Metadata
	Server
}

type PostgreSQLServer struct {
	defsecTypes.Metadata
	Server
	Config PostgresSQLConfig
}

type PostgresSQLConfig struct {
	defsecTypes.Metadata
	LogCheckpoints       defsecTypes.BoolValue
	ConnectionThrottling defsecTypes.BoolValue
	LogConnections       defsecTypes.BoolValue
}

type Server struct {
	defsecTypes.Metadata
	EnableSSLEnforcement      defsecTypes.BoolValue
	MinimumTLSVersion         defsecTypes.StringValue
	EnablePublicNetworkAccess defsecTypes.BoolValue
	FirewallRules             []FirewallRule
}

type MSSQLServer struct {
	defsecTypes.Metadata
	Server
	ExtendedAuditingPolicies []ExtendedAuditingPolicy
	SecurityAlertPolicies    []SecurityAlertPolicy
}

type SecurityAlertPolicy struct {
	defsecTypes.Metadata
	EmailAddresses     []defsecTypes.StringValue
	DisabledAlerts     []defsecTypes.StringValue
	EmailAccountAdmins defsecTypes.BoolValue
}

type ExtendedAuditingPolicy struct {
	defsecTypes.Metadata
	RetentionInDays defsecTypes.IntValue
}

type FirewallRule struct {
	defsecTypes.Metadata
	StartIP defsecTypes.StringValue
	EndIP   defsecTypes.StringValue
}
