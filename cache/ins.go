package cache

import (
	"encoding/json"
	"errors"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/constants"
	"github.com/coocood/freecache"
	"runtime/debug"
	"strconv"
	"sync"
)

var ins *freecache.Cache
var once sync.Once

func GetIns() *freecache.Cache {
	once.Do(func() {
		ins = init_cache()
	})
	return ins
}

func init_cache() *freecache.Cache {
	config := (&configs.CacheConfig{}).Load()
	value, err := strconv.Atoi(config.Size)
	if err == nil && value != 0 {
		c := freecache.NewCache(value)
		//根据cache的大小进行设置
		debug.SetGCPercent(20)
		return c
	}
	return nil
}

func Set(key string, value interface{}, exp int) error {
	if GetIns() == nil {
		return errors.New(constants.ErrCacheInit)
	}

	k := []byte(key)
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = GetIns().Set(k, v, exp)
	if err != nil {
		return err
	}
	return nil
}
func Get(key string) ([]byte, error) {
	if GetIns() == nil {
		return nil, errors.New(constants.ErrCacheInit)
	}
	k := []byte(key)
	return GetIns().Get(k)
}