package redis

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/guru-invest/guru.feeder.marketdata/src/health"
	"github.com/guru-invest/guru.framework/src/helpers/errors"
)

var pool *DbnoPool

type DbnoPool struct {
	*redis.Pool
	dbno int
}

func (p *DbnoPool) Get() redis.Conn {
	conn := p.Pool.Get()
	return conn
}

func InitPool(poolSize int, server string, port string, database int) {
	pool = &DbnoPool{
		&redis.Pool{
			MaxIdle:   poolSize,
			MaxActive: poolSize,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", server+":"+port)
				if err != nil {
					log.Printf("ERROR: fail init redis pool: %s", err.Error())
				}
				return conn, err
			},
			Wait: true,
		}, database,
	}
}

func Set(key string, value string) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return errors.Throw(err, "error on save information on redis")
	}
	return nil
}

func Get(key string) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	ret, err := conn.Do("GET", key)
	if err != nil {
		return nil, errors.Throw(err, "error on get information on redis")
	}
	return ret, nil
}

func GetAll(keys []interface{}) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	values, err := redis.Strings(conn.Do("MGET", keys...))
	if err != nil {
		health.HEALTH.ConnectionStatus = 503
		return nil, errors.Throw(err, "error on get information on redis")
	}
	return values, nil
}
