package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"gopkg.in/ini.v1"
)

var configMutex sync.RWMutex
var config *ini.File

func Get(section, key, defaultValue string) string {
	configMutex.RLock()
	defer configMutex.RUnlock()

	var value = defaultValue
	if v := config.Section(section).Key(key).String(); v != "" {
		value = v
	}

	return value
}

func GetInt(section, key string, defaultValue int) int {
	configMutex.RLock()
	defer configMutex.RUnlock()

	var value = defaultValue

	if v, err := config.Section(section).Key(key).Int(); nil == err {
		value = v
	}

	return value
}

func GetBool(section, key string, defaultValue bool) bool {
	configMutex.RLock()
	defer configMutex.RUnlock()

	var value = defaultValue
	if v, err := config.Section(section).Key(key).Bool(); nil == err {
		value = v
	}

	return value
}

func GetFloat64(section, key string, defaultValue float64) float64 {
	configMutex.RLock()
	defer configMutex.RUnlock()

	var value = defaultValue

	if v, err := config.Section(section).Key(key).Float64(); nil == err {
		value = v
	}

	return value
}

func Init() {
	configMutex.Lock()
	defer configMutex.Unlock()
	config = loadConfig()
}

// TODO: for config file watcher
func Update() {}

func loadConfig() *ini.File {
	configPathList := []interface{}{
		// Please add config file path here if the file is not located at config/
	}

	files, err := ioutil.ReadDir(getConfigFolderPath())
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".ini") {
			configPathList = append(configPathList, getConfigFolderPath()+file.Name())
		}
	}

	if len(configPathList) == 0 {
		return &ini.File{}
	}

	file, err := ini.Load(configPathList[0], configPathList...)
	if err != nil {
		log.Fatalf("load config failed. reason:%s", err.Error())
	}

	return file
}

func getConfigFolderPath() string {
	_, b, _, _ := runtime.Caller(0)
	return strings.TrimSuffix(filepath.Dir(b), "/pkg/config") + "/config/"
}
