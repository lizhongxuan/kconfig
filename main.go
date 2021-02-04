package main

import (
	"encoding/json"
	"fmt"
	log "github.com/golang/glog"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	KConfig *config
)

type config struct {
	Pool map[string]interface{}
}

func main() {
	configPath := "test.yaml"
	// 解析yaml文件
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("read file error: %v", err)
		return
	}
	KConfig = &config{}
	if err = yaml.Unmarshal(yamlFile, &KConfig.Pool); err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return
	}
	fmt.Println(KConfig.Pool)
	data,err := json.Marshal(KConfig.Pool)
	if err != nil {
		log.Fatalf("Marshal: %v", err)
		return
	}

	res := gjson.GetBytes(data,"arr.2.a.zz2")
	fmt.Println(res.String())
}
