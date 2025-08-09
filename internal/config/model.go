package config

import (
	"strings"
	"time"

	"github.com/iamBelugaa/go-boilerplate/pkg/validation"
)

// Environment represents the runtime deployment environment.
type Environment string

// Supported environment constants.
var (
	EnvironmentStaging     Environment = "STAGING"
	EnvironmentProduction  Environment = "PRODUCTION"
	EnvironmentDevelopment Environment = "DEVELOPMENT"
)

// String returns a lowercase string representation of the Environment.
// Defaults to "development" if the value is unrecognized.
func (e Environment) String() string {
	switch e {
	case EnvironmentStaging:
		return "staging"
	case EnvironmentProduction:
		return "production"
	case EnvironmentDevelopment:
		return "development"
	default:
		return "development"
	}
}

// ToEnvironment parses a string and returns a matching Environment value.
// It accepts common aliases (e.g., "prod", "uat", "dev").
// Defaults to EnvironmentDevelopment if the input is not recognized.
func ToEnvironment(str string) Environment {
	switch strings.ToLower(str) {
	case "prod", "production":
		return EnvironmentProduction
	case "staging", "uat", "qa", "testing":
		return EnvironmentStaging
	case "dev", "develop", "development", "local":
		return EnvironmentDevelopment
	default:
		return EnvironmentDevelopment
	}
}

// Logging defines the log level and output destinations for the service.
type Logging struct {
	// Level defines the log verbosity: "debug", "info", "warn", or "error".
	Level string `json:"level" koanf:"level" validate:"required"`

	// OutputPaths defines destinations for logs: "stderr", "stdout", or file paths.
	OutputPaths []string `json:"outputPaths" koanf:"output_paths" validate:"required"`
}

// Validate checks that the Logging configuration is valid.
func (l *Logging) Validate() error {
	return validation.Check(l)
}

// Service contains high-level application metadata and environment details.
type Service struct {
	// Name uniquely identifies the service/application.
	Name string `json:"serviceName" koanf:"service_name" validate:"required"`

	// Version specifies the application release version.
	Version string `json:"serviceVersion" koanf:"service_version" validate:"required"`

	// Environment specifies the deployment environment.
	Environment Environment `json:"environment" koanf:"service_environment" validate:"required"`
}

// IsProduction returns true if the service is running in the Production environment.
func (s *Service) IsProduction() bool {
	return s.Environment == EnvironmentProduction
}

// Validate checks that the Service configuration is valid.
func (s *Service) Validate() error {
	return validation.Check(s)
}

// Server defines HTTP server settings for request handling and timeouts.
type Server struct {
	// Host is the IP or hostname where the server binds (e.g., "0.0.0.0").
	Host string `json:"host" koanf:"server_host" validate:"required"`

	// Port is the TCP port number the server listens on.
	Port uint `json:"port" koanf:"server_port" validate:"required"`

	// ReadTimeout is the maximum duration allowed for reading the request.
	ReadTimeout time.Duration `json:"readTimeout" koanf:"server_read_timeout" validate:"required"`

	// WriteTimeout is the maximum duration before timing out the response write.
	WriteTimeout time.Duration `json:"writeTimeout" koanf:"server_write_timeout" validate:"required"`

	// IdleTimeout is the maximum time to wait for the next request before closing.
	IdleTimeout time.Duration `json:"idleTimeout" koanf:"server_idle_timeout" validate:"required"`

	// ShutdownTimeout is the grace period before forcefully terminating the server.
	ShutdownTimeout time.Duration `json:"shutdownTimeout" koanf:"server_shutdown_timeout" validate:"required"`
}

// Validate checks that the Server configuration is valid.
func (s *Server) Validate() error {
	return validation.Check(s)
}

// Database contains all database connection pool and authentication settings.
type Database struct {
	// Host is the database server address.
	Host string `json:"host" koanf:"db_host" validate:"required"`

	// Port is the database server TCP port.
	Port int `json:"port" koanf:"db_port" validate:"required"`

	// User is the username for authentication.
	User string `json:"user" koanf:"db_user" validate:"required"`

	// Password is the authentication password (optional for some DBs).
	Password string `json:"password" koanf:"db_password"`

	// Name is the specific database/schema to connect to.
	Name string `json:"name" koanf:"db_name" validate:"required"`

	// SSLMode controls SSL behavior ("disable", "require", etc.).
	SSLMode string `json:"sslMode" koanf:"db_ssl_mode" validate:"required"`

	// MaxOpenConns is the maximum number of open connections.
	MaxOpenConns int `json:"maxOpenConns" koanf:"db_max_open_conns" validate:"required"`

	// MaxIdleConns is the maximum number of idle connections.
	MaxIdleConns int `json:"maxIdleConns" koanf:"db_max_idle_conns" validate:"required"`

	// ConnMaxLifetime is the maximum lifetime (in seconds) of a connection.
	ConnMaxLifetime int `json:"connMaxLifetime" koanf:"db_conn_max_lifetime" validate:"required"`

	// ConnMaxIdleTime is the maximum idle time (in seconds) for a connection.
	ConnMaxIdleTime int `json:"connMaxIdleTime" koanf:"db_conn_max_idle_time" validate:"required"`
}

// Validate checks that the Database configuration is valid.
func (db *Database) Validate() error {
	return validation.Check(db)
}

// HealthChecks configures periodic health verification for dependencies
// like databases or APIs or other services.
type HealthChecks struct {
	// Enabled turns health checks on or off.
	Enabled bool `json:"enabled" koanf:"enabled"`

	// Checks is a list of named health checks to perform (e.g., ["database", "api"]).
	Checks []string `json:"checks" koanf:"checks"`

	// Timeout is the maximum time to wait for each check before failing.
	Timeout time.Duration `json:"timeout" koanf:"timeout" validate:"min=1s"`

	// Interval is the frequency between running checks.
	Interval time.Duration `json:"interval" koanf:"interval" validate:"min=1s"`
}

// Validate checks that the HealthChecks configuration is valid.
func (hc *HealthChecks) Validate() error {
	return validation.Check(hc)
}

// Config is the top-level configuration struct aggregating all sub-configs.
type Config struct {
	// Server configures HTTP server behavior.
	Server *Server `json:"server" koanf:"server" validate:"required"`

	// Logging configures log level and destinations.
	Logging *Logging `json:"logging" koanf:"logging" validate:"required"`

	// Database configures database connection pooling and authentication.
	Database *Database `json:"database" koanf:"database" validate:"required"`

	// Service contains application name, version, and environment.
	Service *Service `json:"application" koanf:"application" validate:"required"`

	// HealthChecks defines periodic checks for service dependencies.
	HealthChecks *HealthChecks `json:"healthChecks" koanf:"health_checks" validate:"required"`
}
