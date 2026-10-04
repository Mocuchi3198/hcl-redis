package translator

import (
	"fmt"
    "hcl-redis/types"
)

type CommandRedis struct {
	Command string
	Args []string
}

func Translator(document *types.Document, callback func(CommandRedis) error) error {

	for _, c := range document.Create {
        if err := translateCreate(c, callback); err != nil {
            return err
        }
    }

    for _, a := range document.Assign {
        if err := translateAssign(a, callback); err != nil {
            return err
        }
    }

    for _, a := range document.Add {
        if err := translateAdd(a, callback); err != nil {
            return err
        }
    }

    for _, g := range document.Get {
        if err := translateGet(g, callback); err != nil {
            return err
        }
    }

    for _, d := range document.Delete {
        if err := translateDelete(d, callback); err != nil {
            return err
        }
    }

    return nil
}

func translateCreate(create types.Create, callback func(CommandRedis) error) error {
    switch create.Type {
    case "string":
        if create.Value == nil {
            return fmt.Errorf("create '%s': type 'string' requiere 'value'", create.Name)
        }
        return callback(CommandRedis{"SET", []string{create.Name, *create.Value}})

    case "hash":
        if len(create.Fields) == 0 {
            return fmt.Errorf("create '%s': type 'hash' requiere 'fields'", create.Name)
        }
        args := []string{create.Name}
        for k, v := range create.Fields {
            args = append(args, k, v)
        }
        return callback(CommandRedis{"HSET", args})

    case "list":
        if len(create.Values) == 0 {
            return fmt.Errorf("create '%s': type 'list' requiere 'values'", create.Name)
        }
        args := append([]string{create.Name}, create.Values...)
        return callback(CommandRedis{"RPUSH", args})

    case "set":
        if len(create.Values) == 0 {
            return fmt.Errorf("create '%s': type 'set' requiere 'values'", create.Name)
        }
        args := append([]string{create.Name}, create.Values...)
        return callback(CommandRedis{"SADD", args})

    default:
        return fmt.Errorf("create '%s': tipo desconocido '%s'", create.Name, create.Type)
    }
}

func translateAssign(assign types.Assign, callback func(CommandRedis) error) error {
    args := []string{assign.Name}
    for k, v := range assign.Fields {
        args = append(args, k, v)
    }
    return callback(CommandRedis{"HSET", args})
}

func translateAdd(add types.Add, callback func(CommandRedis) error) error {
    args := append([]string{add.Name}, add.Values...)
    return callback(CommandRedis{"RPUSH", args})
}

func translateGet(get types.Get, emitir func(CommandRedis) error) error {
    if get.Field != nil {
        return emitir(CommandRedis{"HGET", []string{get.Name, *get.Field}})
    }
    return emitir(CommandRedis{"GET", []string{get.Name}})
}

func translateDelete(delete types.Delete, emitir func(CommandRedis) error) error {
    if delete.Field != nil {
        return emitir(CommandRedis{"HDEL", []string{delete.Name, *delete.Field}})
    }
    return emitir(CommandRedis{"DEL", []string{delete.Name}})
}