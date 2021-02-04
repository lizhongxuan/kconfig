package kconfig

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
)

var (
	kConfig []byte
)

func InitConfig(filePath string, isJson bool, envFilePath ...string) {
	// 如果有环境变量,则使用环境变量路径
	if len(envFilePath) > 0 {
		if envpath := getEnvDf(envFilePath[0]); envpath != "" {
			filePath = envpath
		}
	}
	fmt.Println("kconfig path:", filePath)
	// 解析yaml文件
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	if isJson { // json
		kConfig = file
		// check json
		t := make(map[string]interface{})
		if err = json.Unmarshal(kConfig, &t); err != nil {
			panic(err)
		}
	} else { // yaml
		kConfig,err = yaml.YAMLToJSON(file)
		if err != nil {
			panic(err)
		}
	}
}

func GetString(keyPath string, df ...string) string {
	if kConfig == nil {
		return ""
	}
	v := gjson.GetBytes(kConfig, keyPath).String()
	if v == "" && len(df) > 0 {
		return df[0]
	}
	return v
}

func GetInt64(keyPath string, df ...int64) int64 {
	if kConfig == nil {
		return 0
	}
	res := gjson.GetBytes(kConfig, keyPath)
	if !res.Exists() && len(df) > 0 {
		return df[0]
	}
	return res.Int()
}

func GetStringArray(keyPath string) []string {
	if kConfig == nil {
		return nil
	}
	arr := make([]string, 0)
	res := gjson.GetBytes(kConfig, keyPath)
	if !res.Exists() {
		return arr
	}
	t := res.Array()
	for _, v := range t {
		arr = append(arr, v.String())
	}
	return arr
}

func GetInt64Array(keyPath string) []int64 {
	if kConfig == nil {
		return nil
	}
	arr := make([]int64, 0)
	res := gjson.GetBytes(kConfig, keyPath)
	if !res.Exists() {
		return arr
	}
	t := res.Array()
	for _, v := range t {
		arr = append(arr, v.Int())
	}
	return arr
}

func GetStringMap(keyPath string) map[string]string {
	if kConfig == nil {
		return nil
	}
	m := make(map[string]string)
	res := gjson.GetBytes(kConfig, keyPath)
	if !res.Exists() {
		return m
	}
	t := res.Map()
	for k, v := range t {
		m[k] = v.String()
	}
	return m
}

func GetInt64Map(keyPath string) map[string]int64 {
	if kConfig == nil {
		return nil
	}
	m := make(map[string]int64)
	res := gjson.GetBytes(kConfig, keyPath)
	if !res.Exists() {
		return m
	}
	t := res.Map()
	for k, v := range t {
		m[k] = v.Int()
	}
	return m
}

func getEnvDf(key string, df ...string) string {
	value := os.Getenv(key)
	if value == "" && len(df) > 0 {
		return df[0]
	}
	return value
}
