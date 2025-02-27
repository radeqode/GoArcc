package config

import (
	"github.com/deqode/GoArcc/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
)

type FileInformation struct {
	Path string
	Name string
}

//Config : including all the configuration
type Config struct {
	Grpc               GrpcServerConfig        `mapstructure:"GRPC"`
	Graphql            GraphqlServerConfig     `mapstructure:"GRAPHQL"`
	Rest               RestServerConfig        `mapstructure:"REST"`
	HealthCheck        HealthCheckServerConfig `mapstructure:"HEALTH_CHECK"`
	Logger             LoggerConfig            `mapstructure:"LOGGER"`
	Postgres           PostgresConfig          `mapstructure:"POSTGRES"`
	Metrics            MetricsConfig           `mapstructure:"METRICS"`
	Jaeger             JaegerServerConfig      `mapstructure:"JAEGER"`
	Auth               AuthConfig              `mapstructure:"AUTH"`
	GithubVCSConfig    VCSSConfig              `mapstructure:"GITHUB_VCS_CONFIG"`
	CadenceConfig      CadenceConfig           `mapstructure:"CADENCE_CONFIG"`
	SupportedVcsConfig []string
}

type CadenceConfig struct {
	Domain  string `mapstructure:"DOMAIN"`
	Service string `mapstructure:"SERVICE"`
	Port    string `mapstructure:"PORT"`
	Host    string `mapstructure:"HOST"`
}

// GrpcServerConfig GrpcServerConfig: gRPC  server configuration
// Timeout is the request timeout:
// any client request take longer then the given timeout will automatically cancelled.
type GrpcServerConfig struct {
	Port           string `mapstructure:"PORT"`
	Host           string `mapstructure:"HOST"`
	RequestTimeout int    `mapstructure:"REQUEST_TIMEOUT"`
}

// GraphqlServerConfig GraphqlServerConfig: Graphql server configuration
type GraphqlServerConfig struct {
	Port           string `mapstructure:"PORT"`
	Host           string `mapstructure:"HOST"`
	RequestTimeout int    `mapstructure:"REQUEST_TIMEOUT"`
}

// RestServerConfig RestServerConfig: Rest Implementation config
type RestServerConfig struct {
	Port           string `mapstructure:"PORT"`
	Host           string `mapstructure:"HOST"`
	RequestTimeout int    `mapstructure:"REQUEST_TIMEOUT"`
}

// HealthCheckServerConfig HealthCheckServerConfig: Configuration about health check
type HealthCheckServerConfig struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

// LoggerConfig LoggerConfig: Zapier log level
type LoggerConfig struct {
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

// PostgresConfig PostgresConfig: detail config about the postgres database
type PostgresConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	DbName   string `mapstructure:"DB_NAME"`
	SslMode  string `mapstructure:"SSL_MODE"`
	Driver   string `mapstructure:"PG_DRIVER"`
}

//MetricsConfig : detail config about the Metrics
type MetricsConfig struct {
	URL         string `mapstructure:"URL"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
}

//JaegerServerConfig : detail config about the Jaeger
type JaegerServerConfig struct {
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
	LogSpans    string `mapstructure:"LOG_SPANS"`
}

// AuthConfig Authentication config: details provided by Auth0
type AuthConfig struct {
	Auth0ClientID     string `mapstructure:"AUTH0_CLIENT_ID"`
	Auth0Domain       string `mapstructure:"AUTH0_DOMAIN"`
	Auth0ClientSecret string `mapstructure:"AUTH0_CLIENT_SECRET"`
	Auth0CallBackURL  string `mapstructure:"AUTH0_CALL_BACK_URL"`
}

//VCSSConfig like github,gitlab,bitbucket config
type VCSSConfig struct {
	IType        int    `mapstructure:"ITYPE"`
	Provider     string `mapstructure:"PROVIDER"`
	URLTemplate  string `mapstructure:"URL_TEMPLATE"`
	ClientID     string `mapstructure:"CLIENT_ID"`
	RedirectURI  string `mapstructure:"REDIRECT_URI"`
	State        string `mapstructure:"STATE"`
	Scope        string `mapstructure:"SCOPE"`
	ResponseType string `mapstructure:"RESPONSE_TYPE"`
	ClientSecret string `mapstructure:"CLIENT_SECRET"`
	Name         string `mapstructure:"NAME"`
}

//GetVcsConfig : will give the particular vcs config
func GetVcsConfig(name string, vcsConfig []VCSSConfig) *VCSSConfig {
	for _, v := range vcsConfig {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

// LoadConfig config file from given path
func LoadConfig(filename, path string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(filename)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

// ParseConfig file from the given viper
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		logger.Log.Fatal("unable to decode into struct", zap.Error(err))
		return nil, err
	}
	return &c, nil
}

// GetConfigName get the path from local or docker
func GetConfigName() string {
	fileName := os.Getenv("CONFIG_NAME")
	if fileName != "" {
		return fileName
	}
	return "config"
}

func GetConfigDirectory() string {
	filePath := os.Getenv("CONFIG_DIRECTORY")
	if filePath != "" {
		return filePath
	}

	return "./../../"
}

// GetConfig : will get the config
func GetConfig() *Config {
	configFileName := GetConfigName()
	configFileDirectory := GetConfigDirectory()
	logger.Log.Info("Config Details", zap.String("configFileDirectory", configFileDirectory), zap.String("configFileName", configFileName))

	cfgFile, configFileLoadError := LoadConfig(configFileName, configFileDirectory)
	if configFileLoadError != nil {
		logger.Log.Fatal("unable to get config", zap.Error(configFileLoadError))
		panic(configFileLoadError)
	}

	cfg, parseError := ParseConfig(cfgFile)
	if parseError != nil {
		logger.Log.Fatal("unable to get config", zap.Error(parseError))
		panic(parseError)
	}

	cfg.SupportedVcsConfig = supportedVcsConfig()
	return cfg
}

// SupportedVcsConfig add supported type from here.
func supportedVcsConfig() []string {
	return []string{"github"}
}
