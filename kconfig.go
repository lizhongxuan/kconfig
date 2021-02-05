package kconfig

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"sync"
)

var (
	kConfig config
)

type config struct {
	data []byte
	pool sync.Map
}

func InitConfig(filePath string, isJson bool, envFilePath ...string) {
	// 如果有环境变量,则使用环境变量路径
	if len(envFilePath) > 0 {
		if envpath := getEnvDf(envFilePath[0]); envpath != "" {
			filePath = envpath
		}
	}
	//fmt.Println("kconfig path:", filePath)
	// 解析yaml文件
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	if isJson { // json
		kConfig.data = file
		// check json
		t := make(map[string]interface{})
		if err = json.Unmarshal(kConfig.data, &t); err != nil {
			panic(err)
		}
	} else { // yaml
		kConfig.data, err = yaml.YAMLToJSON(file)
		if err != nil {
			panic(err)
		}
	}
}

func GetString(keyPath string, df ...string) string {
	if kConfig.data == nil {
		return ""
	}
	if v, ok := kConfig.pool.Load(keyPath); ok {
		return v.(string)
	}
	v := gjson.GetBytes(kConfig.data, keyPath).String()
	if v == "" && len(df) > 0 {
		return df[0]
	}
	kConfig.pool.Store(keyPath, v)
	return v
}

func GetInt64(keyPath string, df ...int64) int64 {
	if kConfig.data == nil {
		return 0
	}
	if v, ok := kConfig.pool.Load(keyPath); ok {
		return v.(int64)
	}
	res := gjson.GetBytes(kConfig.data, keyPath)
	if !res.Exists() && len(df) > 0 {
		return df[0]
	}
	kConfig.pool.Store(keyPath, res.Int())
	return res.Int()
}

func GetStringArray(keyPath string) []string {
	if kConfig.data == nil {
		return nil
	}
	if v, ok := kConfig.pool.Load(keyPath); ok {
		return v.([]string)
	}
	arr := make([]string, 0)
	res := gjson.GetBytes(kConfig.data, keyPath)
	if !res.Exists() {
		return arr
	}
	t := res.Array()
	for _, v := range t {
		arr = append(arr, v.String())
	}
	kConfig.pool.Store(keyPath, arr)
	return arr
}

func GetInt64Array(keyPath string) []int64 {
	if kConfig.data == nil {
		return nil
	}
	if v, ok := kConfig.pool.Load(keyPath); ok {
		return v.([]int64)
	}
	arr := make([]int64, 0)
	res := gjson.GetBytes(kConfig.data, keyPath)
	if !res.Exists() {
		return arr
	}
	t := res.Array()
	for _, v := range t {
		arr = append(arr, v.Int())
	}
	kConfig.pool.Store(keyPath, arr)
	return arr
}

func GetStringMap(keyPath string) map[string]string {
	if kConfig.data == nil {
		return nil
	}
	if v, ok := kConfig.pool.Load(keyPath); ok {
		return v.(map[string]string)
	}
	m := make(map[string]string)
	res := gjson.GetBytes(kConfig.data, keyPath)
	if !res.Exists() {
		return m
	}
	t := res.Map()
	for k, v := range t {
		m[k] = v.String()
	}
	kConfig.pool.Store(keyPath, m)
	return m
}

func GetInt64Map(keyPath string) map[string]int64 {
	if kConfig.data == nil {
		return nil
	}
	if v, ok := kConfig.pool.Load(keyPath); ok {
		return v.(map[string]int64)
	}
	m := make(map[string]int64)
	res := gjson.GetBytes(kConfig.data, keyPath)
	if !res.Exists() {
		return m
	}
	t := res.Map()
	for k, v := range t {
		m[k] = v.Int()
	}
	kConfig.pool.Store(keyPath, m)
	return m
}

func getEnvDf(key string, df ...string) string {
	value := os.Getenv(key)
	if value == "" && len(df) > 0 {
		return df[0]
	}
	return value
}
