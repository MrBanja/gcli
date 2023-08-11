package history

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/MrBanja/openaiAPI"
	"github.com/spf13/viper"
)

var storage *viper.Viper

func InitStorage(pathToDB string) error {
	if storage != nil {
		return fmt.Errorf("storage already initialized")
	}
	v := viper.New()
	_, err := os.Stat(pathToDB)
	if err != nil {
		f, err := os.Create(pathToDB)
		if err != nil {
			return err
		}
		_, err = f.Write([]byte("{}"))
		err = f.Close()
		if err != nil {
			return err
		}
	}

	v.SetConfigFile(pathToDB)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	storage = v
	return nil
}

func Get(conversationID string) ([]openaiAPI.Message, error) {
	var messages []openaiAPI.Message
	if err := storage.UnmarshalKey(conversationID, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func Set(conversationID string, messages []openaiAPI.Message) error {
	storage.Set(conversationID, messages)
	return storage.WriteConfig()
}

func Keys() []string {
	return storage.AllKeys()
}

func Delete(conversationID string) error {
	cfg := storage.AllSettings()
	vals := cfg

	parts := strings.Split(conversationID, ".")
	for i, k := range parts {
		v, ok := vals[k]
		if !ok {
			// Doesn't exist no action needed
			break
		}

		switch len(parts) {
		case i + 1:
			// Last part so delete.
			delete(vals, k)
		default:
			m, ok := v.(map[string]interface{})
			if !ok {
				return fmt.Errorf("unsupported type: %T for %q", v, strings.Join(parts[0:i], "."))
			}
			vals = m
		}
	}

	b, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	if err = storage.ReadConfig(bytes.NewReader(b)); err != nil {
		return err
	}

	return storage.WriteConfig()
}
