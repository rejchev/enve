package enve

import (
	"testing"
)

var _ IEnveSource = (*exampleSource)(nil)

type exampleSource map[string]string 

func (e *exampleSource) GetEnvs() (map[string]string, error) {
	return *e, nil
}

var envs = exampleSource {
	"APP_NAME": "enve",
	"APP_VERSION": "1",
	"APP_DEV": "true",
}

type EnvBasedConfig struct {
	AppName string `json:"APP_NAME"`
	AppVersion int `json:"APP_VERSION,string"`
	AppDev bool `json:"APP_DEV,string"`
}

func Test_That_ExampleSource_Success(t *testing.T) {
	buf := EnvBasedConfig{}

	expected := EnvBasedConfig {
		AppName: "enve",
		AppVersion: 1,
		AppDev: true,
	}
	
	if err := Parse(&buf, &envs); err != nil {
		t.Errorf("Err: %s", err)
	}

	if buf != expected {
		t.Errorf("Expected %v, got %v", expected, buf)
	}
}

type EnvEvironBasedConfig struct {
	VsCodePID int `json:"VSCODE_PID,string"`
}

func Test_That_EnveEnvironSource_Success(t *testing.T) {
	buf := EnvEvironBasedConfig{}

	expected := EnvEvironBasedConfig {
		VsCodePID: 9172,
	}

	switch m, err := new(EnvironSource).GetEnvs(); {
	case err != nil:
		t.Errorf("err on build environ map: %v", err)

	default:
		
		m["VSCODE_PID"] = "9172"

		ex := exampleSource(m)

		if err := Parse(&buf, &ex); err != nil {
			t.Errorf("Err: %s", err)
		}

		if buf != expected {
			t.Errorf("Expected %v, got %v", expected, buf)
		}
	}
}