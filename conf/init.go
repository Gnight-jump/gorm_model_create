/**
初始化结构体对象
*/
package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Tconf Conf

func Init() error {
	ReConf, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		fmt.Print(err)
		return err
	}
	yaml.Unmarshal(ReConf, &Tconf)

	fmt.Println("数据库配置：", Tconf.Dbase)
	fmt.Println("目标表单配置：", Tconf.TargetModel)
	return nil
}
