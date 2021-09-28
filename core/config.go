package core

import (
	"cloud-task/global"
	"fmt"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
)

const defaultConfigFile = "config-local.yaml"

func Init() {
	vDefault := viper.New()
	v := viper.New()
	vDefault.SetConfigFile(getRealPath(defaultConfigFile))
	errD := vDefault.ReadInConfig()
	if errD != nil {
		fmt.Printf("read default config file %s fail: %s\n", defaultConfigFile, errD)
	}
	configs := vDefault.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, vl := range configs {
		v.SetDefault(k, vl)
	}
	//读取环境变量
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	global.GVA_CONFIG.System.Version = v.GetString("system.version")
}
func getRealPath(configFile string) string {
	fs := afero.NewOsFs()
	if _, err := fs.Open(configFile); err != nil {
		//未读取到文件,获取config.go文件的路径
		_, file, _, _ := runtime.Caller(0)
		dir := path.Dir(file)
		parentDir := path.Dir(dir)
		configFile = path.Join(parentDir, configFile)
	}
	return configFile
}
