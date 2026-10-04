package types

import (
	"testing"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func TestDocument(test *testing.T) {
	commands := []string {
		`create "especie" {
			type = "string"
			value = "perro"
		}`,
		`create "animal" {
			type = "hash"
			fields = {
				especie = "hamster sirio"
				nombre = "mandarino"
				edad = 1
			}
		}`,
		`create "tipos_mamiferos" {
			type = "list"
			values = ["gato", "perro", "hasmter", "chiguire"]
		}`,
		`assign "animal" {
			fields = {
				color = "marron"
			}
		}`,
		`add "tipos_mamiferos" {
			values = ["danta", "picure"]
		}`,
		`get {
			name = "especie"
		}`,

		`get {
			name = "animal"
			field = "especie"
		}`,
		`delete  {
			name = "especie"
		}`,
		`delete {
			name = "animal"
			field = "color"
		}`,
	}

	for i, e := range commands {
		var document Document
		err := hclsimple.Decode("input.hcl", []byte(e), nil, &document)
		if err != nil {
			test.Fatalf("%v | %v\n", err, i)
		}
	}
}