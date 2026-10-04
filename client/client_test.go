package client

import (
	"testing"
	"fmt"
	"hcl-redis/translator"
)

func TestNewClient(test *testing.T) {
	addr := "127.0.0.1:6379"
	result := NewClient(addr)
	if result == nil {
		error := fmt.Sprintf("Error: se esperaba un puntero a Client, pero se obtuvo nil: %v", result)
		test.Fatal(error)
	}
}

func TestClose(test *testing.T) {
	addr := "127.0.0.1:6379"
	client := NewClient(addr)
	result := client.Close()
	if result != nil {
		error := fmt.Sprintf("Error al cerrar conexion del *Client %v", result)
		test.Fatal(error)
	}

	result = client.Execute(translator.CommandRedis{
        Command: "PING",
    })
    if result == nil {
        test.Fatal("esperaba error tras cerrar")
    }
}

func TestExecute(test *testing.T) {
	
	tableTest := []struct {
		context string
		command translator.CommandRedis
		isFail bool
	} {
		{
			context: "test DEL <SET nombre>",
			command: translator.CommandRedis {
				Command: "DEL",
				Args: []string{"nombre"},
			},
			isFail: false,
		},
		{
			context: "test HDEL <HSET persona <nombre>>",
			command: translator.CommandRedis {
				Command: "HDEL",
				Args: []string{"persona", "nombre"},
			},
			isFail: false,
		},
		{
			context: "test DEL <HSET persona>",
			command: translator.CommandRedis {
				Command: "DEL",
				Args: []string{"persona"},
			},
			isFail: false,
		},
		{
			context: "test SET <nombre>",
			command: translator.CommandRedis {
				Command: "SET",
				Args: []string{"nombre", "Ana"},
			},
			isFail: false,
		},
		{
			context: "test HSET",
			command: translator.CommandRedis {
				Command: "HSET",
				Args: []string{"persona", "nombre", "Mandarino", "edad", "30"},
			},
			isFail: false,
		},
		{
			context: "test GET",
			command: translator.CommandRedis {
				Command: "GET",
				Args: []string{"nombre"},
			},
			isFail: false,
		},
		{
			context: "test HSET",
			command: translator.CommandRedis {
				Command: "HSET",
				Args: []string{"persona", "estado", "soltero"},
			},
			isFail: false,
		},
		{
			context: "test HGET <nombre>",
			command: translator.CommandRedis {
				Command: "HGET",
				Args: []string{"persona", "nombre"},
			},
			isFail: false,
		},
		{
			context: "test HGET <estado>",
			command: translator.CommandRedis {
				Command: "HGET",
				Args: []string{"persona", "estado"},
			},
			isFail: false,
		},
		{
			context: "test SET-ERROR <SET <Pais nombre URSS>>",
			command: translator.CommandRedis {
				Command: "SET",
				Args: []string{"Pais", "nombre", "URSS"},
			},
			isFail: true,
		},
	}

	for _, table := range tableTest {
		test.Run(table.context, func(t *testing.T) {
			addr := "127.0.0.1:6379"
			client := NewClient(addr)
			result := client.Execute(table.command)
			if table.isFail {
	            if result == nil {
	                t.Fatalf("esperaba error en: %s", table.context)
	            }
	        } else {
	            if result != nil {
	                t.Fatalf("no esperaba error en %s: %v", table.context, result)
	            }
	        }
		})
	}
}