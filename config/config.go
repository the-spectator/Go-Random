package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	appName           string
	appPort           int
	safeWordFilePath  string
	swearWordFilePath string
)

/*Load configuration variables */
func Load() {
	viper.SetDefault("APP_NAME", "app")
	viper.SetDefault("APP_PORT", "8002")

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

/*AppName Gives App Name */
func AppName() string {
	if appName == "" {
		appName = ReadEnvString("APP_NAME")
	}
	return appName
}

/*AppPort Gives App PORT */
func AppPort() int {
	port := os.Getenv("PORT")
	logger.Printf("ENV (%v)", port)
	if port != "" {
		appPort, err := strconv.Atoi(port)
		logger.Printf("ERR (%v)", err)
		if err == nil {
			logger.Printf("AToi (%v)", port)
			return appPort
		}
	}

	logger.Printf("ReadEnvInt %v", port)
	if appPort == 0 {
		appPort = ReadEnvInt("APP_PORT")
	}
	return appPort
}

/*ReadEnvInt reads environment variable as Int*/
func ReadEnvInt(key string) int {
	checkIfSet(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid integer", key))
	}
	return v
}

/*SafeWordFilePath Gives SafeWordFilePath */
func SafeWordFilePath() string {
	if safeWordFilePath == "" {
		safeWordFilePath = ReadEnvString("SAFE_WORD_FILE_PATH")
	}
	return safeWordFilePath
}

/*SwearWordFilePath Gives SafeWordFilePath */
func SwearWordFilePath() string {
	if swearWordFilePath == "" {
		swearWordFilePath = ReadEnvString("SWEAR_WORD_FILE_PATH")
	}
	return swearWordFilePath
}

/*ReadEnvString reads environment variable as String*/
func ReadEnvString(key string) string {
	checkIfSet(key)
	return viper.GetString(key)
}

/*ReadEnvBool reads environment variable as Boolean*/
func ReadEnvBool(key string) bool {
	checkIfSet(key)
	return viper.GetBool(key)
}

func checkIfSet(key string) {
	if !viper.IsSet(key) {
		err := errors.New(fmt.Sprintf("Key %s is not set", key))
		panic(err)
	}
}
