/*
 * @Author: ww
 * @Date: 2022-07-07 06:20:51
 * @Description:
 * @FilePath: /danmuplay/cfg/cfg.go
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