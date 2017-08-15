package define

import (
    "flag"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-ini/ini"
    "log"
)

var (
    Connection string
    

    Environment string
    Config string
    

    Db *sql.DB
)

func init () {
    flag.StringVar(&Environment, "env","dev","server run environment")
    flag.StringVar(&Config, "conf","","config path")
    
    flag.Parse()

    conf, err := ini.Load(Config)
    if err != nil {
        log.Panic("init config error")
    }

    Connection = conf.Section(Environment).Key("Connection").String()
    

    Db, _ = sql.Open("mysql", Connection)
}