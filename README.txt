# hcl-redis

A Redis dialect written in HCL.

Instead of writing Redis commands (`SET`, `HSET`, `HGET`), you write
HCL blocks with clear verbs (`create`, `assign`, `add`, `get`, `delete`).
hcl-redis translates your HCL into Redis commands and executes them.

## Why?

Because Redis commands are cryptic:

- `SET nombre "Ana"` — does it configure or save?
- `HSET persona nombre "Ana"` — what does the `H` mean?
- `HGET persona nombre` — why `H` before `GET`?

This dialect uses explicit verbs and readable blocks:

```hcl
create "nombre" {
  type  = "string"
  value = "Ana"
}

create "persona" {
  type = "hash"
  fields = {
    nombre = "Ana"
    edad   = "30"
  }
}

assign "persona" {
  fields = {
    edad = "31"
  }
}

get {
  name  = "persona"
  field = "nombre"
}

delete {
  name = "persona"
}
________________________

Installation

Requires Go 1.21+ and a running Redis server.

git clone https://github.com/youruser/hcl-redis.git
cd hcl-redis
go build -o hcl-redis .

_______________________

Usage

File mode

./hcl-redis apply comandos.hcl 127.0.0.1:6379

Executes all blocks in the file against Redis.

______________________

REPL mode

./hcl-redis repl 127.0.0.1:6379

Opens an interactive interpreter. Write HCL and it runs immediately.

hcl-redis REPL
> create "nombre" {
  type  = "string"
  value = "Ana"
}
OK
> get {
  name = "nombre"
}
"Ana"
> .exit

______________________

Syntax

- create

Creates a new key. The type determines the required fields.

Type  | Required fields
-------------------------
string| value
hash  | fields
list  | values
set   | values

Example:

create "nombre" {
  type  = "string"
  value = "Ana"
}

create "persona" {
  type = "hash"
  fields = {
    nombre = "Ana"
    edad   = "30"
  }
}

create "tareas" {
  type   = "list"
  values = ["buy bread", "study Go"]
}

create "colores" {
  type   = "set"
  values = ["red", "green", "blue"]
}

=======================

- assign

Modifies fields of an existing hash.

assign "persona" {
  fields = {
    edad = "31"
  }
}

- add

Appends values to an existing list or set.

add "tareas" {
  values = ["call the doctor"]
}

======================

- get

Reads a value. With field, reads a field from a hash.

get {
  name = "nombre"
}

get {
  name  = "persona"
  field = "edad"
}

=====================

- delete

Removes a whole key, or a field from a hash.

delete {
  name = "nombre"
}

delete {
  name  = "persona"
  field = "edad"
}

____________________________

Translation to Redis

hcl-redis                  |  Redis
-------------------------------------
create with type = "string"|  SET
create with type = "hash"  |  HSET
create with type = "list"  |  RPUSH
create with type = "set"   |  SADD
assign with fields         |  HSET
add with values            |  RPUSH
get with name              |  GET
get with name and field    |  HGET
delete with name           |  DEL
delete with name and field |  HDEL

_________________________________

Project structure

hcl-redis/
├── main.go               # Entry point (CLI)
├── types/                # Structs for HCL blocks
├── parser/               # HCL → Document
├── translator/           # Document → Redis commands
├── client/               # Redis commands → server
├── repl/                 # Interactive REPL
└── examples/
    └── comandos.hcl

_______________________________

Tests

go test ./...

_______________________________

Status

Work in progress. Currently supports:

    String, Hash, List, Set.

    File mode and REPL mode.

    Unit tests for parser, translator, client and block balancing.

Pending:

    More commands (exists, incr, lrange, smembers).

    Server mode (other programs speak HCL over TCP).

    Publishing on GitHub.