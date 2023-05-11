package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

//iniファイル読み込み

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

// グローバル変数として宣言
var Config ConfigList

func init() {

	cfg, _ := ini.Load("config.ini")

	Config = ConfigList{
		//MustInt,MustStringはエラーなく必ず値を返す関数。値が存在しない場合に設定値を返す。
		Port:   cfg.Section("web").Key("post").MustInt(8080),
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
		//.String()は値が存在しない場合、Stringの初期値の空文字を返す
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("DbName = %v\n", Config.DbName)
	fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)
}
