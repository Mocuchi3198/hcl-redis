package repl

import (
    "fmt"
    "io"
    "strings"

    "github.com/chzyer/readline"
    "hcl-redis/client"
    "hcl-redis/parser"
    "hcl-redis/translator"
)

var completer = readline.NewPrefixCompleter(
    readline.PcItem("create"),
    readline.PcItem("assign"),
    readline.PcItem("add"),
    readline.PcItem("get"),
    readline.PcItem("delete"),
    readline.PcItem(".help"),
    readline.PcItem(".clear"),
    readline.PcItem(".exit"),
)

func Run(cli *client.Client) {
    rl, err := readline.NewEx(&readline.Config{
        Prompt:       "> ",
        AutoComplete: completer,
        HistoryFile:  "/tmp/hcl-redis-history",
    })
    if err != nil {
        panic(err)
    }
    defer rl.Close()

    fmt.Println("hcl-redis REPL")
    fmt.Println("Escribe HCL. Usa .help para ayuda, .exit para salir.")
    fmt.Println()

    var buffer strings.Builder

    for {
        // Prompt según si estamos dentro de un bloque
        if buffer.Len() == 0 {
            rl.SetPrompt("> ")
        } else {
            rl.SetPrompt("  ")
        }

        line, err := rl.Readline()
        if err == readline.ErrInterrupt {
            if buffer.Len() == 0 {
                fmt.Println()
                break
            }
            buffer.Reset()
            continue
        } else if err == io.EOF {
            fmt.Println()
            break
        }

        line = strings.TrimRight(line, "\r\n")
        trim := strings.TrimSpace(line)

        // Comandos especiales (solo fuera de bloques)
        if buffer.Len() == 0 {
            if trim == "" {
                continue
            }
            if trim == ".exit" {
                break
            }
            if trim == ".help" {
                ShowHelp()
                continue
            }
            if trim == ".clear" {
                buffer.Reset()
                fmt.Println("Buffer limpiado.")
                continue
            }
        }

        buffer.WriteString(line)
        buffer.WriteString("\n")

        // Si no está balanceado, seguir acumulando
        if !Balanced(buffer.String()) {
            continue
        }

        // Parsear y ejecutar
        contend := buffer.String()
        buffer.Reset()

        doc, err := parser.ParseString(contend)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        err = translator.Translator(doc, func(cmd translator.CommandRedis) error {
            return cli.Execute(cmd)
        })
        if err != nil {
            fmt.Println("Error:", err)
        }
    }
}

func Balanced(text string) bool {
    count := 0
    for _, c := range text {
        if c == '{' {
            count++
        } else if c == '}' {
            count--
        }
    }
    return count == 0
}

func ShowHelp() {
    fmt.Print(`
Bloques disponibles:

  create "<nombre>" {
    type   = "string" | "hash" | "list" | "set"
    value  = "..."           # para string
    fields = { ... }         # para hash
    values = [ ... ]         # para list o set
  }

  assign "<nombre>" {
    fields = { campo = "valor" }
  }

  add "<nombre>" {
    values = ["a", "b", "c"]
  }

  get {
    name  = "<nombre>"
    field = "<campo>"        # opcional, solo para hash
  }

  delete {
    name  = "<nombre>"
    field = "<campo>"        # opcional, solo para hash
  }

Comandos del REPL:

  .help    Muestra esta ayuda
  .clear   Limpia el buffer actual
  .exit    Sale del REPL

Atajos de teclado:

  ← →      Mover cursor
  ↑ ↓      Historial
  Ctrl+R   Buscar en historial
  Tab      Autocompletar
`)
}