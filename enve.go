package enve

import (
	"encoding/json"
	"fmt"
	"maps"
)

type IEnveSource interface {
	GetEnvs() (map[string]string, error)
}

func GetEnvs(i []IEnveSource) ([]byte, error) {
	m := make(map[string]string)

	for _, env := range i {
		if env == nil {
			continue;
		}
		
		switch n, err := env.GetEnvs(); {
		case err != nil:
			continue

		default:
			maps.Copy(m, n)
		}
	}

	return json.Marshal(m)
}

func Parse(v any, s ...IEnveSource) error {
	if s == nil {
		return fmt.Errorf("source must be present")
	}

	switch b, err := GetEnvs(s); {
	case err != nil:
		return err
	default:
		return json.Unmarshal(b, v)
	}
}
