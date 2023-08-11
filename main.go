package main

import (
	"fmt"
	"gcli/cmd"
	"gcli/internal/controller"
	"gcli/internal/history"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if err := initialize(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initialize() error {
	////////////////////////////////////////////////////////////////
	//// Viper Setup

	if err := initConfig(); err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("home dir not found %w", err)
	}

	viper.SetDefault("history.path", filepath.Join(home, ".config/gcli/history.json"))
	viper.SetDefault("current_conversation_id", "init")

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("config writing error %w", err)
	}

	////////////////////////////////////////////////////////////////
	//// Globals Setup

	if err := history.InitStorage(viper.GetString("history.path")); err != nil {
		return fmt.Errorf("history init error %w", err)
	}

	if err := controller.InitFlowController(); err != nil {
		return fmt.Errorf("controller init error %w", err)
	}
	return nil
}

func initConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("home dir not found %w", err)
	}

	pathToConfigDir := filepath.Join(home, ".config/gcli")
	if err = os.MkdirAll(pathToConfigDir, os.ModePerm); err != nil {
		return fmt.Errorf("can not create config dir %w", err)
	}

	pathToConfig := fmt.Sprintf("%s/config.yaml", pathToConfigDir)

	_, err = os.Stat(pathToConfig)
	if err != nil {
		f, err := os.Create(pathToConfig)
		_ = f.Close()
		if err != nil {
			return fmt.Errorf("config creation error %w", err)
		}
		fmt.Println("Config file created:", pathToConfig)
	}

	viper.SetConfigFile(pathToConfig)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("config parsing error %w", err)
	}
	return nil
}
