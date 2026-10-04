package main

import (
	//"log"
	"os"
	"fmt"
    "hcl-redis/translator"
    "hcl-redis/client"
    "hcl-redis/parser"
    "hcl-redis/repl"
)

func main() {
	fmt.Println("Hola hcl-redis")

	if len(os.Args) < 2 {
	    fmt.Println("Uso: hcl-redis apply <archivo.hcl>")
	    os.Exit(1)
  	}

  	command := os.Args[1]
    addr := "127.0.0.1:6379"

    if command == "apply" && len(os.Args) >= 4 {
        addr = os.Args[3]
    } else if command == "repl" && len(os.Args) >= 3 {
        addr = os.Args[2]
    }

    cli := client.NewClient(addr)
    defer cli.Close()

    switch command {
    case "apply":
        if len(os.Args) < 3 {
            fmt.Println("Uso: hcl-redis apply <archivo.hcl> [addr]")
            os.Exit(1)
        }
        file := os.Args[2]
        doc, err := parser.Parse(file)
        if err != nil {
            fmt.Println("Error al parsear:", err)
            os.Exit(1)
        }
        err = translator.Translator(doc, func(cmd translator.CommandRedis) error {
            return cli.Execute(cmd)
        })
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }
        fmt.Println("OK")

    case "repl":
        repl.Run(cli)

    default:
        fmt.Println("Comando desconocido:", command)
        os.Exit(1)
    }
}

/*
create "nombre" {
  type = "string"
  value = "Ana"
}

create "persona" {
  type = "hash"
  fields = {
    nombre = "Ana"
    edad = 30
  }
}

create "tareas" {
  type = "list"
  values = ["comprar pan", "estudiar"]
}

assign "persona" {
	fields = {
  	edad = 31
	}
}

add "tareas" {
  values = ["llamar al médico"]
}

get "nombre"

get "persona" {
  field = "nombre"
}

delete "persona" {
  field = "edad"
}

delete "persona"*/
