package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/asphaltbuffet/go-dungeon/pkg/dungeon"
)

func Test_LoadLanguage(t *testing.T) {
	filename := "../../languages/en_default.json"

	got, err := LoadJSON(filename)

	assert.NoError(t, err)
	assert.IsType(t, &dungeon.Dungeon{}, got)
}
