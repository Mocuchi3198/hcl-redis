package parser

import (
	"testing"
)

func TestParse(test *testing.T) {
	_, err := Parse("../comandos.hcl")
	if err != nil {
		test.Fatal(err)
	}
}

func TestParseCreateString(test *testing.T) {
	content := `create "nombreAnimal" {
		type = "string"
		value = "Azulejo"
	}`

	result, err := ParseString(content)

	if err != nil {
		test.Fatal(err)
	}

	animal := result.Create[0]

	if animal.Name != "nombreAnimal" {
		test.Fatalf("Error: se esperaba 'nombreAnimal' en 'animal.Name' pero se obtuvo %s\n", animal.Name)
	}

	if animal.Type != "string" {
		test.Fatalf("Error: se esperaba 'string' en 'animal.Type' pero se obtuvo %s\n", animal.Type)
	}

	if *animal.Value != "Azulejo" {
		test.Fatalf("Error: se esperaba 'Azulejo' en 'animal.Value' pero se obtuvo %s\n", *animal.Value)
	}
}

func TestParseCreateHash(test *testing.T) {
	content := `create "Animal" {
  	type = "hash"
	  fields = {
	    nombre = "Tigro"
	    edad   = "1"
  		}
	}`
	result, err := ParseString(content)
	if err != nil {
		test.Fatal(err)
	}

	if len(result.Create) != 1 {
		test.Fatalf("Error: se esperaba '1' create en 'len(result.Create)' pero se obtuvo %d\n", len(result.Create))
	}

	animal := result.Create[0]
	
	if animal.Name != "Animal" {
		test.Fatalf("Error: se esperaba 'Animal' en 'animal.Name' pero se obtuvo %s\n", animal.Name)
	}
	
	if animal.Type != "hash" {
		test.Fatalf("Error: se esperaba 'hash' en 'animal.Type' peros se obtuvo %s\n", animal.Type)
	}

	if animal.Fields["nombre"] != "Tigro" {
		test.Fatalf("Error: se esperaba 'Tigro' en 'animal.Fields[\"nombre\"]'pero se obtuvo %s", animal.Fields["nombre"])
	}

	if animal.Fields["edad"] != "1" {
		test.Fatalf("Error: se esperaba '1' en 'animal.Fields[\"edad\"]' pero se obtuvo %s", animal.Fields["edad"])
	}
}

func TestParseCreateList(test *testing.T) {
	content := `create "Animales" {
		type = "list"
		values = ["Conejo", "Tapir", "Cachicamo", "Sangre de Toro"]
	}
	`
	result, err := ParseString(content)

	if err != nil {
		test.Fatal(err)
	}

	if len(result.Create) != 1 {
		test.Fatalf("Error: se esperaba '1' en 'len(result.Create)' pero se obtuvo %d\n", len(result.Create))
	}

	animalsList := result.Create[0]

	if animalsList.Name != "Animales" {
		test.Fatalf("Error: se esperaba 'Animales' en 'animalsList.Name' pero se obtuvo %s\n", animalsList.Name)
	}

	if animalsList.Type != "list" {
		test.Fatalf("Error: se esperaba 'list' en 'animalsList.Type' pero se obtuvo %s\n", animalsList.Type)
	}

	if len(animalsList.Values) != 4 {
		test.Fatalf("Error: se esperaba '4' en 'len(animalsList.Values)' pero se obtuvo %d\n", len(animalsList.Values))
	}

	listExpectedAnimals := []string {
		"Conejo",
		"Tapir",
		"Cachicamo",
		"Sangre de Toro",
	}

	for i, e := range animalsList.Values {
		if listExpectedAnimals[i] != e {
			test.Fatalf("Error: se esperaba '%s' en 'listExpectedAnimals[%d] != e' pero se obtuvo %s\n", listExpectedAnimals[i], i, e)
		}
	}
}

func TestParseCreateSet(test *testing.T) {
    content := `create "Colores" {
        type = "set"
        values = ["rojo", "verde", "azul"]
    }`

    result, err := ParseString(content)
    if err != nil {
        test.Fatal(err)
    }

    if len(result.Create) != 1 {
        test.Fatalf("Error: se esperaba '1' en 'len(result.Create)' pero se obtuvo %d", len(result.Create))
    }

    colores := result.Create[0]

    if colores.Name != "Colores" {
        test.Errorf("Error: se esperaba 'Colores' en 'colores.Name' pero se obtuvo '%s'", colores.Name)
    }
    if colores.Type != "set" {
        test.Errorf("Error: se esperaba 'set' en 'colores.Type' pero se obtuvo '%s'", colores.Type)
    }
    if len(colores.Values) != 3 {
        test.Errorf("Error: se esperaba 3 valores en 'colores.Values' pero se obtuvo %d", len(colores.Values))
    }
}

func TestParseAssign(test *testing.T) {
	content := `assign "Animal" {
		fields = {
			nombre = "Fantasma"
			edad = "4"
		}
	}`

	result, err := ParseString(content)
	if err != nil {
		test.Fatal(err)
	}

	animal := result.Assign[0]

	if animal.Name != "Animal" {
		test.Fatalf("Error: se esperaba 'Animal' en 'animal.Name' y se obtuvo %s", animal.Name)
	}

	if animal.Fields["nombre"] != "Fantasma" {
		test.Fatalf("Error: se esperaba 'Fantasma' en 'animal.Fields[\"nombre\"]' pero se obtuvo %s", animal.Fields["nombre"])
	}

	if animal.Fields["edad"] != "4" {
		test.Fatalf("Error: se esperaba '4' en 'animal.Fields[\"edad\"]' pero se obtuvo %s", animal.Fields["edad"])
	}
}

func TestParseAdd(test *testing.T) {
	content := `add "animales" {
		values = ["Pasalido", "Mayo", "Antenas Largas", "Jade"]
	}`

	result, err := ParseString(content)

	if err != nil {
		test.Fatal(err)
	}

	animalsList := result.Add[0]

	if animalsList.Name != "animales" {
		test.Fatalf("Error: se esperaba 'animales' en 'animalsList.Name' pero se obtuvo %s", animalsList.Name)
	}

	if len(animalsList.Values) != 4 {
		test.Fatalf("Error: se esperaba '4' en 'len(animalsList.Values)' pero se obtuvo %d", len(animalsList.Values))
	}

	listExpectedAnimals := []string {
		"Pasalido",
		"Mayo",
		"Antenas Largas",
		"Jade",
	}

	for i, e := range animalsList.Values {
		if listExpectedAnimals[i] != e {
			test.Fatalf("Error: se esperaba '%s' en 'listExpectedAnimals[%d] != e' pero se obtuvo '%s'", listExpectedAnimals[i], i, e)
		}
	}
}

func TestParseGet(test *testing.T) {
	content := `get {
		name = "Animal"
	}
	
	get {
		name = "Animal"
		field = "edad"
	}`

	result, err := ParseString(content)
	if err != nil {
		test.Fatal(err)
	}

	if len(result.Get) != 2 {
		test.Fatalf("Error: se esperaba '2' en 'len(result.Get)' pero se obtuvo %d", len(result.Get))
	}

	animalSimple := result.Get[0]
	animalComplex := result.Get[1]

	if animalSimple.Name != "Animal" {
		test.Fatalf("Error: se esperaba 'Animal' en 'animalSimple.Name' pero se obtuvo %s", animalSimple.Name)
	}

	if animalSimple.Field != nil {
		test.Fatalf("Error: se esperaba 'nil' en 'animalSimple.Field' pero se obtuvo '%s'", *animalSimple.Field)
	}

	if animalComplex.Name != "Animal" {
		test.Fatalf("Error: se esperaba 'Animal' en 'animalComplex.Name' pero se obtuvo %s", animalComplex.Name)
	}

	if *animalComplex.Field != "edad" {
		test.Fatalf("Error: se esperaba 'edad' en 'animalComplex.Field' pero se obtuvo %s", *animalComplex.Field)
	}
}

func TestParseDelete(test *testing.T) {
	content := `delete {
		name = "nombreAnimal"
	}

	delete {
		name = "Animal"
		field = "edad"
	}`

	result, err := ParseString(content)
	if err != nil {
		test.Fatal(err)
	}

	if len(result.Delete) != 2 {
		test.Fatalf("Error: se esperaba '2' en 'len(result.Delete)' pero se obtuvo %d", len(result.Delete))
	}
	
	animalSimple := result.Delete[0]
	animalComplex := result.Delete[1]

	if animalSimple.Name != "nombreAnimal" {
		test.Fatalf("Error: se esperaba 'nombreAnimal' en 'animalSimple.Name' pero se obtuvo %s", animalSimple.Name)
	}
	
	if animalSimple.Field != nil {
		test.Fatalf("Error: se esperaba 'nil' en 'animalSimple.Field' pero se obtuvo %s", *animalSimple.Field)
	}

	if animalComplex.Name != "Animal" {
		test.Fatalf("Error: se esperaba 'Animal' en 'animalComplex.Name' pero se obtuvo %s", animalComplex.Name)
	}

	if *animalComplex.Field != "edad" {
		test.Fatalf("Error: se esperaba 'edad' en 'animalComplex.Field' pero se obtuvo %s", *animalComplex.Field)
	}
}

func TestParseStringError(test *testing.T) {
	content := `create "nombreAnimal" {
		type = "string"
		value = [1, 2, 3]
	}`

	_, err := ParseString(content)
	if err == nil {
		test.Fatal("Error: se esperaba un error pero se obtuvo nil")
	}
}