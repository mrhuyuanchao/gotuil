package gotuil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Configuration 配置信息
var Configuration *SettingInfo

func init() {
	appPath, _ := os.Getwd()
	fileName := filepath.Join(appPath, "conf")
	conf := &SettingInfo{}
	err := loadJSONConfig(filepath.Join(fileName, "config.json"), conf)
	if err != nil {
		err = loadYamlConfig(filepath.Join(fileName, "config.yaml"), conf)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	Configuration = conf
}
func loadJSONConfig(fileName string, v interface{}) error {
	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(v)
	return err
}

func loadYamlConfig(fileName string, v interface{}) error {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = yaml.Unmarshal(fileData, v)
	return err
}

// LoadConfig 加载配置文件
func LoadConfig(fileName string, v interface{}) error {
	ext := path.Ext(fileName)
	if ext == ".json" {
		return loadJSONConfig(fileName, v)
	}
	if ext == ".yaml" {
		return loadYamlConfig(fileName, v)
	}
	return fmt.Errorf(fmt.Sprintf("%s文件暂不支持", ext))
}
