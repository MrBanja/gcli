package history

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MrBanja/openaiAPI"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type History struct {
	pathToDB string
	storage  *viper.Viper
}

func initStorage(pathToDB string) (*viper.Viper, error) {
	v := viper.New()
	_, err := os.Stat(pathToDB)
	if err != nil {
		f, err := os.Create(pathToDB)
		if err != nil {
			return nil, err
		}
		_, err = f.Write([]byte("{}"))
		err = f.Close()
		if err != nil {
			return nil, err
		}
	}

	v.SetConfigFile(pathToDB)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func New(pathToDB string) (*History, error) {
	storage, err := initStorage(pathToDB)
	if err != nil {
		return nil, err
	}
	return &History{
		pathToDB: pathToDB,
		storage:  storage,
	}, nil
}

func (h *History) Get(conversationID string) ([]openaiAPI.Message, error) {
	var messages []openaiAPI.Message
	if err := h.storage.UnmarshalKey(conversationID, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func (h *History) Set(conversationID string, messages []openaiAPI.Message) error {
	h.storage.Set(conversationID, messages)
	return h.storage.WriteConfig()
}

func (h *History) Keys() []string {
	return h.storage.AllKeys()
}

func (h *History) Delete(conversationID string) error {
	cfg := h.storage.AllSettings()
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

	if err = h.storage.ReadConfig(bytes.NewReader(b)); err != nil {
		return err
	}

	return h.storage.WriteConfig()
}
