<p align="center">
  <h1 align="center">ENVE</h1>
  <p align="center">An extended, too simple, zero-dependencies library to parse environment variables into structs</p>
</p>

<div align="center">

  [![Test & Build](https://github.com/rejchev/enve/actions/workflows/build.yml/badge.svg)](https://github.com/rejchev/enve/actions/workflows/build.yml)
  [![Go Reference](https://pkg.go.dev/badge/github.com/rejchev/enve.svg)](https://pkg.go.dev/github.com/rejchev/enve)
  [![Go Report Card](https://goreportcard.com/badge/github.com/rejchev/enve)](https://goreportcard.com/report/github.com/rejchev/enve)
  
</div>

### About
An extended, too simple, zero-dependencies library to parse environment variables into structs.\
Library provides interfaces for implementing and using environment variable sources, allowing you to use multiple environment variable sources at once.\
`Enve` based on `encoding/json` lib

### Installation
```bash
go get github.com/rejchev/enve
```

### Usage
```go
type config struct {
    AppName string  `json:"APP_NAME"`
    IsDebug bool    `json:"IS_DEBUG,string"`    
}

c := config{}

// Using os.Environ() as envs source
err := enve.Parse(&c, new(enve.EnvironSource))

// Using io.Reader as envs source
envs := "APP_NAME=testname\n" +
        "IS_DEBUG=true"

reader := strings.Reader(envs)

err := enve.Parse(&c, enve.NewReaderSource(reader))

// Using combines of sources
err := enve.Parse(&c, enve.NewReaderSource(reader), new(enve.EnvironSource))

// Using custom source
var _ enve.IEnveSource = (*CustomEnveSource)(nil)

type CustomEnveSource byte

func (e *EnvironSource) GetEnvs() (map[string]string, error) {
    m := make(map[string]string)

    // some logic
    
    return m, nil
}

err := enve.Parse(&c, new(CustomEnveSource))
```