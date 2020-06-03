package syrupPlum

import (
	"github.com/Unknwon/goconfig"
	"github.com/pkg/errors"
)

type Option struct {
	ConfigPath string //使用配置文件初始化config
	SavePath   string //数据保存目录地址

}

func (option *Option) CheckHealthy() (bool, error) {
	if _, err := PathExists(option.SavePath); err != nil {
		return false, err
	}
	return true, nil
}

func InitOptionWithConfigFile(configPath string) (*Option, error) {
	cfg, err := goconfig.LoadConfigFile(configPath)
	if err != nil {
		return nil, err
	}
	savePath, err := cfg.GetValue("syrup_plum", "save_path")
	if err != nil {
		return nil, err
	}
	if savePath == "" {
		return nil, errors.New("save_path can't be empty")
	}
	return InitOption(savePath)
}

func InitOption(savePath string) (*Option, error) {
	//判断savePath 是否有效
	if _, err := PathExists(savePath); err != nil {
		return nil, err
	}
	return &Option{
		SavePath: savePath,
	}, nil
}
