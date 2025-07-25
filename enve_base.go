package enve

import (
	"os"
	"strings"
)

var _ IEnveSource = (*EnvironSource)(nil)

type EnvironSource byte

func (e *EnvironSource) GetEnvs() (map[string]string, error) {
	m := make(map[string]string)

	if envs := os.Environ(); len(envs) != 0 {
		for _, e := range envs {
			if kv := strings.Split(e, "="); len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
	}

	return m, nil
}
