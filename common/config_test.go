package common

import (
	"os"
	"reflect"
	"testing"
)

func Test_loadConfig(t *testing.T) {

	tests := []struct {
		name    string
		wantC   *Config
		wantErr bool
	}{
		// {"hoge", &Config{}, false},
		{"gae", &Config{App: App{
			Name: "sample-prd",
			Env:  "prd",
		}, GAE: GAE{
			DeploymentID: "427554653150974968",
			MemoryMB:     256,
		}}, false},
		// {"k", &Config{}, false},
		// {"gae & k", &Config{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			os.Setenv("GAE_DEPLOYMENT_ID", "427554653150974968")
			os.Setenv("GAE_MEMORY_MB", "256")
			os.Setenv("APP_NAME", "sample-prd")
			os.Setenv("APP_ENV", "prd")

			gotC, err := LoadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("loadConfig() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
