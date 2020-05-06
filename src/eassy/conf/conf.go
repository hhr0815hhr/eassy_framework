package conf

import (
	"github.com/jinzhu/configor"
)

type MysqlConfig struct {
	Master map[string]struct {
		Host string
		Port int
		User string
		Pass string
	}
	Slaver map[string]struct {
		Host string
		Port int
		User string
		Pass string
	}
}

//type ServerConfig struct {
//	Addr []string
//}

type EtcdConfig struct {
	Etcd map[string]string
}

var (
	Mysql   MysqlConfig
	EtcdCfg EtcdConfig
	//Server ServerConfig
)

func GetConf(isDev bool) {
	var configor1 *configor.Configor
	if isDev {
		configor1 = configor.New(&configor.Config{Environment: "development"})
	} else {
		configor1 = configor.New(&configor.Config{Environment: "production"})
	}
	//mysql
	err := configor1.Load(&Mysql, "../conf/mysql.yml")
	if err != nil {
		panic(err)
	}
	//etcd
	err = configor1.Load(&EtcdCfg, "../conf/etcd.yml")
	if err != nil {
		panic(err)
	}

	//server
	//err = configor1.Load(&Server,"../conf/server.yml")
	//if err != nil {
	//	panic(err)
	//}

}
