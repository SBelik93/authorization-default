package models

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strings"
	"syscall"
)

type Environment struct {
	ConnectionDbString string `json:"connection_db_string"`
	Port               string `json:"port"`
	Contour            string `json:"contour"`
	envString
}

type envString struct {
	variables          map[string]interface{}
	variablesMarshaled []byte
}
type EnvironmentProd struct {
	ConnectionDbStringProd string `json:"connection_db_string_prod"`
	Port                   string `json:"port"`
	Contour                string `json:"contour"`
}

type EnvironmentDev struct {
	ConnectionDbStringDev string `json:"connection_db_string_dev"`
	Port                  string `json:"port"`
	Contour               string `json:"contour"`
}

type EnvironmentLocal struct {
	ConnectionDbStringLocal string `json:"connection_db_string_local"`
	Port                    string `json:"port"`
	Contour                 string `json:"contour"`
}

func NewEnvironment() (envs Environment, err error) {
	err = godotenv.Load() //Load .env file
	if err != nil {
		return
	}
	contour := os.Getenv("contour")
	if contour == "" {
		err = fmt.Errorf(`not found env "contour" with one of the values: local, prod, dev`)
		return
	}
	switch true {
	case contour == "dev":
		envsDev := EnvironmentDev{}
		ParseEnvString(&envsDev)
		envs = Environment{
			ConnectionDbString: envsDev.ConnectionDbStringDev,
			Port:               envsDev.Port,
			Contour:            envsDev.Contour,
		}
	case contour == "prod":
		envsProd := EnvironmentProd{}
		ParseEnvString(&envsProd)
		envs = Environment{
			ConnectionDbString: envsProd.ConnectionDbStringProd,
			Port:               envsProd.Port,
			Contour:            envsProd.Contour,
		}

	case contour == "local":
		envsLocal := EnvironmentLocal{}
		ParseEnvString(&envsLocal)
		envs = Environment{
			ConnectionDbString: envsLocal.ConnectionDbStringLocal,
			Port:               envsLocal.Port,
			Contour:            envsLocal.Contour,
		}
	}
	err = envs.Validate()
	return
}

func (e Environment) Validate() (err error) {
	refValue := reflect.ValueOf(e)
	numField := refValue.NumField()
	for i := 0; i < numField; i++ {
		if refValue.Field(i).String() == "" {
			tag := refValue.Type().Field(i).Tag.Get("json")
			contour := fmt.Sprintf("_%s", e.Contour)
			if len(strings.Split(tag, "_")) == 1 {
				contour = ""
			}
			fieldName := fmt.Sprintf("%v", refValue.Type().Field(i).Name)
			err = fmt.Errorf(`can not find env tag: %v%v, struct field name: %v`,
				tag,
				contour,
				fieldName,
			)
			return
		}
	}
	return
}

func ParseEnvString(model interface{}) *envString {
	e := envString{}
	var unVars = syscall.Environ()
	e.variables = make(map[string]interface{})
	for _, unVar := range unVars {
		data := strings.SplitN(unVar, "=", 2)
		if len(data) == 2 {
			e.variables[data[0]] = data[1]
		}
	}
	e.variablesMarshaled, _ = json.Marshal(&e.variables)
	return e.Env(model)
}

func (e *envString) Env(model interface{}) *envString {
	_ = json.Unmarshal(e.variablesMarshaled, model)
	return e
}
