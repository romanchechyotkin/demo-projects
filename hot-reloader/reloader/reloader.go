package reloader

import (
	"log"
	"os"
	"time"

	"hot_reloader/config"

	"github.com/ilyakaznacheev/cleanenv"
)

type Reloadable interface {
	Reload(cfg any)
}

type Reloader struct {
	name       string
	configPath string
	config     any

	reloadables []Reloadable
	done        chan struct{}
}

func New(name, configPath string) *Reloader {
	r := &Reloader{
		name:        name,
		configPath:  configPath,
		reloadables: make([]Reloadable, 0),
		done:        make(chan struct{}),
	}

	r.readConfig()

	return r
}

func (r *Reloader) Start() {
	go r.run()
}

func (r *Reloader) Stop() {
	r.done <- struct{}{}
}

func (r *Reloader) Register(reloadable Reloadable) {
	r.reloadables = append(r.reloadables, reloadable)
}

func (r *Reloader) readConfig() {
	if _, err := os.Stat(r.configPath); err != nil {
		log.Println("failed to open config file", err)
		return
	}

	var cfg config.Config

	err := cleanenv.ReadConfig(r.configPath, &cfg)
	if err != nil {
		log.Println("failed to read config", err)
		return
	}

	r.config = &cfg
}

func (r *Reloader) run() {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-r.done:
			log.Printf("%s reloder stopped\n", r.name)
			return
		case <-ticker.C:
			log.Println("reloader started")

			r.readConfig()

			for _, reloadable := range r.reloadables {
				reloadable.Reload(r.config)
			}
		}
	}
}
