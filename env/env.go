package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

func Load(cfg any) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("can't load .env file")
	}

	configType := reflect.TypeOf(cfg)
	configVal := reflect.ValueOf(cfg)
	if configVal.Kind() == reflect.Ptr {
		configVal = configVal.Elem()
	}
	configType = configVal.Type()
	if configType.Kind() != reflect.Struct {
		panic("Load: cfg argument must be a pointer to a struct")
	}

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		fieldValue := configVal.FieldByName(field.Name)

		val := os.Getenv(field.Tag.Get("env"))
		switch field.Type.Kind() {
		case reflect.String:
			fieldValue.SetString(val)
		case reflect.Int:
			if atoi, err := strconv.Atoi(val); err == nil {
				fieldValue.SetInt(int64(atoi))
			}
		case reflect.Int64:
			if f, err := strconv.ParseInt(val, 10, 64); err == nil {
				fieldValue.SetInt(f)
			}
		case reflect.Bool:
			fieldValue.SetBool(val == "true")
		case reflect.Float64:
			if f, err := strconv.ParseFloat(val, 64); err == nil {
				fieldValue.SetFloat(f)
			}
		}
	}

	return nil
}
