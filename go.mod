module goProjects

go 1.17

require github.com/spf13/cobra v1.3.0

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/wangmingyang1994/golearn v1.3.1
)

replace github.com/wangmingyang1994/golearn => ./staging/src/github.com/wangmingyang1994/golearn
