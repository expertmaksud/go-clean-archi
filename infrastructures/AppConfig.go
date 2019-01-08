package infrastructures

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //use for postgres
	"github.com/spf13/viper"
)

//Config DB config
type Config struct {
	ConnectionString string
}

func init() {
	viper.SetConfigName("app")    // no need to include file extension
	viper.AddConfigPath("config") // set the path of your config file

}

//NewConfig set configuration
func NewConfig() *Config {
	return &Config{
		ConnectionString: getConnectionString(),
	}
}

func getConnectionString() string {
	var connString string
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {

		env := viper.GetString("enabled_env")

		host := viper.GetString(env + ".db.host")
		port := viper.GetInt(env + ".db.port")
		user := viper.GetString(env + ".db.user")
		dbname := viper.GetString(env + ".db.dbname")
		password := viper.GetString(env + ".db.password")
		sslmode := viper.GetString(env + ".db.sslmode")

		connString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

		fmt.Printf("\nDevelopment Config found:\n ConnString = %s\n", connString)
	}

	return connString
}

//ConnectDatabase connect with DB
func ConnectDatabase(config *Config) (*gorm.DB, error) {

	return gorm.Open("postgres", config.ConnectionString)

}
