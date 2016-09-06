package config

var (
	Version  = "UNSET"
	Revision = "UNSET"
)

type Config struct {
	Version     string
	Revision    string
	StatePath   string
	HomePath    string
	FlagVerbose bool
	FlagJSONLog bool
	FlagAPIPort int
}

func (c *Config) setDefaults() {
	c.Version = Version
	c.Revision = Revision
	c.StatePath = "./state.yaml"
	c.FlagVerbose = false
	c.FlagJSONLog = false
	c.FlagAPIPort = 1015
}

func Configuration() (c Config) {
	c.setDefaults()
	return c
}
