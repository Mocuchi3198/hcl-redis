package translator

import (
	"testing"
	"hcl-redis/types"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"fmt"
	"slices"
)

func TestTranslator(test *testing.T) {
	commands := []struct {
		name string
		commandHcl string
		expected CommandRedis
	}{
		{
			name: "create <string>",
			commandHcl: `create "especie" {
				type = "string"
				value = "perro"
			}`,
			expected: CommandRedis{"SET", []string{"especie", "perro"}},
		},
		{
			name: "create <hash>",
			commandHcl: `create "animal" {
				type = "hash"
				fields = {
					especie = "hamster sirio"
					nombre = "mandarino"
					edad = 1
				}
			}`,
			expected: CommandRedis{"HSET", []string{"animal",
				"especie", "hamster sirio",
				"nombre", "mandarino",
				"edad", "1"}},
		},
		{
			name: "create <list>",
			commandHcl: `create "tipos_mamiferos" {
				type = "list"
				values = ["gato", "perro", "hasmter", "chiguire"]
			}`,
			expected: CommandRedis{"RPUSH", []string{"tipos_mamiferos", "gato", "perro", "hasmter", "chiguire"}},
		},
		{
			name: "assign <hash>",
			commandHcl: `assign "animal" {
				fields = {
					color = "marron"
				}
			}`,
			expected: CommandRedis{"HSET", []string{"animal", "color", "marron"}},
		},
		{
			name: "add <list>",
			commandHcl: `add "tipos_mamiferos" {
				values = ["danta", "picure"]
			}`,
			expected: CommandRedis{"RPUSH", []string{"tipos_mamiferos", "danta", "picure"}},
		},
		{
			name: "get <string>",
			commandHcl: `get {
				name = "especie"
			}`,
			expected: CommandRedis{"GET", []string{"especie"}},
		},
		{
			name: "get <hash>",
			commandHcl: `get {
				name = "animal"
				field = "especie"
			}`,
			expected: CommandRedis{"HGET", []string{"animal", "especie"}},
		},
		{
			name: "delete <string>",
			commandHcl: `delete {
				name = "especie"
			}`,
			expected: CommandRedis{"DEL", []string{"especie"}},
		},
		{
			name: "delete <hash>",
			commandHcl: `delete {
				name = "animal"
				field = "color"
			}`,
			expected: CommandRedis{"HDEL", []string{"animal", "color"}},
		},
	}

	for _, e := range commands {
		var document types.Document
		err := hclsimple.Decode("input.hcl", []byte(e.commandHcl), nil, &document)
		if err != nil {
			test.Fatalf("%v\n", err)
		}
		err = Translator(&document, func(cmd CommandRedis) error {
			if e.expected.Command != cmd.Command {
				return fmt.Errorf("Error: '%s' dato esperado '%s', recibido '%s'", e.name, e.expected.Command, cmd.Command)
			}
			if len(e.expected.Args) != len(cmd.Args) {
				return fmt.Errorf("Error: '%s' dato esperado '%d' recibido '%d'", e.name, len(e.expected.Args), len(cmd.Args))
			}
			for i, element := range e.expected.Args {
				if e.expected.Command == "HSET" {
					if !slices.Contains(cmd.Args, element) {
						            return fmt.Errorf("Error: '%s' esperaba '%s' en HSET, no encontrado", e.name, element)
					}
				} else {
					if cmd.Args[i] != element {
						return fmt.Errorf("Error: '%s' dato esperado '%s', recibido '%s'", e.name, element, cmd.Args[i])
					}
				}
			}
			return nil
		})
		if err != nil {
			test.Fatal(err)
		}
	}
}
/*
func TestTranslateCreate() {}
func TestTranslateAssign() {}
func TestTranslateAdd() {}
func TestTranslateGet() {}
func TestTranslateDelete() {}
*/