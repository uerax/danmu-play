/*
 * @Author: ww
 * @Date: 2022-07-07 06:20:51
 * @Description:
 * @FilePath: \danmu-play\cfg\cfg.go
 */
package cfg

import "github.com/uerax/goconf"

var Config *goconf.CfgFile

func Init(path string) error {
	cf := goconf.NewCfgFile()
	err := cf.ReadAll(path)
	if err != nil {
		return err
	}
	Config = cf
	return nil
}

func GetValue(param ...string) (interface{}, error) {
	return Config.GetValue(param...)
}

func GetStringWithDefault(def string, param ...string) string {
	i, err := Config.GetValue(param...)
	if err != nil {
		return def
	}
	return i.(string)
}

func GetIntWithDefault(def int, param ...string) int {
	i, err := Config.GetValue(param...)
	if err != nil {
		return def
	}
	return i.(int)
}

func GetInt64WithDefault(def int64, param ...string) int64 {
	i, err := Config.GetValue(param...)
	if err != nil {
		return def
	}
	return int64(i.(int))
}
