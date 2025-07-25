package enve

import (
	"io"
	"bufio"
	"strings"
)

var _ IEnveSource = (*ReaderSource)(nil)

type ReaderSource struct {
	ior io.Reader
}

func NewReaderSource(ior io.Reader) *ReaderSource {
	return &ReaderSource { ior: ior }
}

func (r *ReaderSource) GetEnvs() (map[string]string, error) {
	m := make(map[string]string)

	if r.ior != nil {
		l := ""
		err := (error)(nil)
		reader := bufio.NewReader(r.ior)

		for {
			if l, err = reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}
			
			if kv := strings.Split(l, "="); len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
	}

	return m, nil
}
