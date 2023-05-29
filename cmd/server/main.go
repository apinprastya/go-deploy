package main

import (
	"godeploy/pkg/config"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.Info("Go-Deploy")

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	logrus.Infof("Working dir: %s", dir)

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := &config.Config{}
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}
	conf.Apply(dir)

	/*rootPath := os.Getenv("ROOT_FOLDER")
	address := os.Getenv("SERVER_ADDRESS")
	port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	secret := os.Getenv("SECRET")
	srv := service.NewService(rootPath)
	api := api.NewApi(secret, srv)

	if err := api.Run(address, port); err != nil {
		fmt.Errorf("can not run server: %s", err.Error())
	}*/
}
