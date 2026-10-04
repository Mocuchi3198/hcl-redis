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
  values = ["comprar pan", "estudiar Go"]
}

get {
  name = "persona"
  field = "nombre"
}

assign "persona" {
  fields = {
    estado = "soltero"
  }
}

get {
  name = "persona"
  field = "estado"
}

create "Animal" {
  type = "hash"
  fields = {
    nombre = "Tigro"
    edad   = "1"
  }
}