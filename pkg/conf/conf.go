package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"sync"
)

// Configuration provides configuration for application.
type Configuration struct {
	mu        sync.RWMutex
	onChanges []func(*Configuration)
	loaded    bool

	client *viper.Viper
}

func (c *Configuration) SetConfigType(in string) {
	c.client.SetConfigType(in)
}

func (c *Configuration) AddConfigPath(in string) {
	c.client.AddConfigPath(in)
}

func (c *Configuration) AddOnChange(fn func(*Configuration)) {
	c.onChanges = append(c.onChanges, fn)
}

func (c *Configuration) emitOnChange() {
	for _, fn := range c.onChanges {
		fn(c)
	}
}

func NewConfiguration() *Configuration {
	return &Configuration{
		mu:        sync.RWMutex{},
		onChanges: make([]func(*Configuration), 0),
		loaded:    false,
		client:    viper.New(),
	}
}

var cfg *Configuration

func init() {
	cfg = NewConfiguration()
}

func Load(key string, s interface{}) error {
	if err := cfg.client.Sub(key).Unmarshal(s); err != nil {
		return err
	}

	cfg.AddOnChange(func(configuration *Configuration) {
		_ = cfg.client.Sub(key).Unmarshal(s)
	})

	return nil
}

func Init(path, ext string) {
	cfg.SetConfigType(ext)
	cfg.AddConfigPath(path)

	err := cfg.client.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg.client.WatchConfig()
	cfg.client.OnConfigChange(func(in fsnotify.Event) {
		cfg.emitOnChange()
	})
}
