package config

import (
	"github.com/pesakit/pesakit/env"
	"github.com/techcraftlabs/mpesa"
)

const (
	homeMnoEnvVar          = "PESAKIT_MNO"
	homeMnoDefaultValue    = ""
	homeDirEnvVar          = "PESAKIT_HOME"
	homeDirDefaultValue    = ""
	configFileEnvVar       = "PESAKIT_CONFIG"
	configFileDefaultValue = ""
	debugModeEnvVar        = "PESAKIT_DEBUG_MODE"
	debugDefaultValue      = false
)

type (
	App struct {
		Home   string
		Config string
		Debug  bool
	}

	TigoPesa struct {
		App
		TigoPesaKey    string
		TigoPesaSecret string
	}

	Airtel struct {
		App
		AirtelKey    string
		AirtelSecret string
	}

	Mno struct {
		Value string
	}
)

func DefaultAppConf() *App {
	return &App{
		Home:   env.String(homeDirEnvVar, homeDirDefaultValue),
		Config: env.String(configFileEnvVar, configFileDefaultValue),
		Debug:  env.Bool(debugModeEnvVar, debugDefaultValue),
	}
}

func DefaultMnoConf() *Mno {
	return &Mno{
		Value: env.String(homeMnoEnvVar, homeMnoDefaultValue),
	}
}

type Info struct {
	App      *App
	Mpesa    *Mpesa
	TigoPesa *TigoPesa
	Airtel   *Airtel
}

func (m *Mpesa) ToConfig() *mpesa.Config {
	return &mpesa.Config{}
}
