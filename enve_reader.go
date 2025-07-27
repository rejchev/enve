package enve

import (
	"bufio"
	"io"
	"strings"
)

var _ IEnveSource = (*ReaderSource)(nil)

type ReaderSource struct {
	ior io.Reader
	d byte
}

func NewReaderSource(ior io.Reader) *ReaderSource {
	return NewReaderSourceE(ior, '\n')
}

func NewReaderSourceE(ior io.Reader, delim byte) *ReaderSource {
	return &ReaderSource{ior: ior, d: delim}
}

func (r *ReaderSource) GetEnvs() (map[string]string, error) {
	m := make(map[string]string)

	if r.ior != nil {
		l := ""
		err := (error)(nil)
		reader := bufio.NewReader(r.ior)

		for {
			if l, err = reader.ReadString(r.d); err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}

			if l[len(l)-1] == '\n' {
				drop := 1
				if len(l) > 1 && l[len(l)-2] == '\r' {
					drop = 2
				}
				l = l[:len(l)-drop]
			}

			if kv := strings.Split(l, "="); len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
	}

	return m, nil
}
