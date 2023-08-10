package loader

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/asphaltbuffet/go-dungeon/pkg/dungeon"
)

func LoadJSON(path string) (*dungeon.Dungeon, error) {
	// read file into []byte
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("opening language file: %w", err)
	}

	var d dungeon.Dungeon

	// unmarshal []byte into dungeon.Dungeon
	err = json.Unmarshal(b, &d)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling dungeon: %w", err)
	}

	return &d, nil
}
