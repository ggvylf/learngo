module github.com/ggvylf/learngo/projects/logcollections/logagent

go 1.15

require (
	github.com/ggvylf/learngo/projects/logcollections/logagent/conf v0.0.0
	github.com/ggvylf/learngo/projects/logcollections/logagent/kafka v0.0.0
	github.com/ggvylf/learngo/projects/logcollections/logagent/taillog v0.0.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.62.0
)

replace (
	github.com/ggvylf/learngo/projects/logcollections/logagent/conf => ./conf
	github.com/ggvylf/learngo/projects/logcollections/logagent/kafka => ./kafka
	github.com/ggvylf/learngo/projects/logcollections/logagent/taillog => ./taillog
)
