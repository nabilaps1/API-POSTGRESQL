package configs

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

/*
Ponteiro que referencia a struct abaixo.
Garante segurança ja que ninguem de fora deste pacote consegue acesso às configuracoes daqui.
*/

var cfg *config

// struct principal -> Leitura (junta as duas secundarias)
type config struct {
	API APIConfig
	DB  DBConfig
}

// structs secundarias -> servidor e banco de dados
type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000") // arquivo toml
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432") // porta padrao postgre
}

// Retorna um erro caso nao consiga carregar as configuracoes
func Load() error {
	viper.SetConfigName("config") // nome arquivo (sem extensao)
	viper.SetConfigType("toml")   // NECESSARIO, o no nome do arquivo nao tem a extensao
	viper.AddConfigPath(".")      // caminho para achar o arquivo de config (vamos colocar smp ao lado do binario)

	err := viper.ReadInConfig() // encontra e le o arquivo de configuracao

	if err != nil {
		// validando tipos de erros
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { // se erro != de FileNotFound, beleza (ja que setamos algumas configs padrao)
			fmt.Printf("File not found")
			return err
		} else { // achou config file, mas outro erro foi produzido
			fmt.Printf("Error type : %s\n", reflect.TypeOf(err))
		}
	}

	cfg = new(config) // cria um ponteiro da struct (mesma coisa que: cfg = &config)

	// atribuir os dados na struct
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil // se nao há erro, retorna nil
}

// acessar as informacoes
func GetDB() DBConfig { // retorna config banco de dados
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
