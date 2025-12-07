package configs

type (
	Config struct {
		Service       Service       `mapstructure:"service"`
		Database      Database      `mapstructure:"database"`
		SpotifyConfig SpotifyConfig `mapstructure:"spotifyConfig"`
	}

	Service struct {
		Port       string `mapstructure:"port"`
		SecrestJWT string `mapstructure:"secretJWT"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}

	SpotifyConfig struct {
		ClientID     string `mapstructure:"cliendID"`
		ClientSecret string `mapstructure:"clientSecret"`
	}
)
