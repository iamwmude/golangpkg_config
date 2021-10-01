package config

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"gotest.tools/assert"
)

const TEST_CONFIG_NAME = "test_config.ini"

func init() {
	copyTestConfigFile()
	Init()
}

func copyTestConfigFile() {
	input, err := ioutil.ReadFile(TEST_CONFIG_NAME)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(getConfigFolderPath()+TEST_CONFIG_NAME, input, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteTestConfigFile() {
	if err := os.Remove(getConfigFolderPath() + TEST_CONFIG_NAME); err != nil {
		log.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	defer deleteTestConfigFile()

	type args struct {
		section      string
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				section:      "test",
				key:          "val_string",
				defaultValue: "none",
			},
			want: "test",
		},
		{
			name: "simple default value",
			args: args{
				section:      "test",
				key:          "val_string_1",
				defaultValue: "none",
			},
			want: "none",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.section, tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestGetInt(t *testing.T) {
	defer deleteTestConfigFile()

	type args struct {
		section      string
		key          string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple",
			args: args{
				section:      "test",
				key:          "val_int",
				defaultValue: -5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.section, tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	defer deleteTestConfigFile()

	type args struct {
		section      string
		key          string
		defaultValue bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple",
			args: args{
				section:      "test",
				key:          "val_bool",
				defaultValue: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBool(tt.args.section, tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFloat64(t *testing.T) {
	defer deleteTestConfigFile()

	type args struct {
		section      string
		key          string
		defaultValue float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "simple",
			args: args{
				section:      "test",
				key:          "val_float",
				defaultValue: 9.2,
			},
			want: 5.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFloat64(tt.args.section, tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
	defer deleteTestConfigFile()

	_, err := loadConfig().GetSection("test")
	assert.Equal(t, err, nil)

	_, err = loadConfig().GetSection("test1")
	assert.Error(t, err, err.Error())
}

func Test_getConfigFolderPath(t *testing.T) {
	defer deleteTestConfigFile()

	_, err := os.Stat(getConfigFolderPath() + TEST_CONFIG_NAME)
	assert.Equal(t, os.IsNotExist(err), false)
	_, err = os.Stat(getConfigFolderPath() + "abcdefg")
	assert.Equal(t, os.IsNotExist(err), true)
}
