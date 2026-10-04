package client

import (
    "context"
    "fmt"
    "time"
    "github.com/redis/go-redis/v9"
    "hcl-redis/translator"
)

type Client struct {
    rdb *redis.Client
    ctx context.Context
}

func NewClient(addr string) *Client {
    return &Client{
        rdb: redis.NewClient(&redis.Options{
            Addr:         addr,
            DialTimeout:  5 * time.Second,
            ReadTimeout:  3 * time.Second,
            WriteTimeout: 3 * time.Second,
        }),
        ctx: context.Background(),
    }
}

func (cli *Client) Close() error {
    return cli.rdb.Close()
}

func (cli *Client) Execute(cmd translator.CommandRedis) error {
	fmt.Printf("DEBUG: ejecutando %s %v\n", cmd.Command, cmd.Args)
    var err error

    switch cmd.Command {
    case "SET":
        if len(cmd.Args) != 2 {
            err = fmt.Errorf("(error) ERR syntax error")
        } else {
            err = cli.rdb.Set(cli.ctx, cmd.Args[0], cmd.Args[1], 0).Err()
        }

    case "GET":
        val, e := cli.rdb.Get(cli.ctx, cmd.Args[0]).Result()
        if e == redis.Nil {
            fmt.Println("(nil)")
            return nil
        }
        err = e
        if err == nil {
            fmt.Println(val)
        }

    case "HSET":
        //err = cli.rdb.HSet(cli.ctx, cmd.Args[0], cmd.Args[1:]).Err()
        args := make([]interface{}, len(cmd.Args)-1)
	    for i, v := range cmd.Args[1:] {
	        args[i] = v
	    }
	    err = cli.rdb.HSet(cli.ctx, cmd.Args[0], args...).Err()

    case "HGET":
        val, e := cli.rdb.HGet(cli.ctx, cmd.Args[0], cmd.Args[1]).Result()
        if e == redis.Nil {
            fmt.Println("(nil)")
            return nil
        }
        err = e
        if err == nil {
            fmt.Println(val)
        }

    case "RPUSH":
    	args := make([]interface{}, len(cmd.Args)-1)
	    for i, v := range cmd.Args[1:] {
	        args[i] = v
	    }
        err = cli.rdb.RPush(cli.ctx, cmd.Args[0], args...).Err()


    case "SADD":
    	args := make([]interface{}, len(cmd.Args)-1)
	    for i, v := range cmd.Args[1:] {
	        args[i] = v
	    }
        err = cli.rdb.SAdd(cli.ctx, cmd.Args[0], args...).Err()

    case "DEL":
        err = cli.rdb.Del(cli.ctx, cmd.Args[0]).Err()

    case "HDEL":
        err = cli.rdb.HDel(cli.ctx, cmd.Args[0], cmd.Args[1]).Err()
    case "PING":
        err = cli.rdb.Ping(cli.ctx).Err()
    default:
        return fmt.Errorf("comando no soportado: %s", cmd.Command)
    }

    if err != nil {
        return fmt.Errorf("error ejecutando %s: %w", cmd.Command, err)
    }
    return nil
}