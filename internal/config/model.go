package config

// Environment defines the type for representing different runtime environments.
type Environment string

// Supported environment constants.
var (
	EnvironmentStaging     Environment = "STAGING"
	EnvironmentProduction  Environment = "PRODUCTION"
	EnvironmentDevelopment Environment = "DEVELOPMENT"
)

// String returns string representation of the Environment.
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
