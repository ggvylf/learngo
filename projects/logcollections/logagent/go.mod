module github.com/ggvylf/learngo/projects/logcollections/logagent

go 1.15

require (
	github.com/Shopify/sarama v1.27.2 //ct
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/hpcloud/tail v1.0.0 //ct
	// github.com/ggvylf/learngo/projects/logcollections/logagent/conf v0.0.0
	// github.com/ggvylf/learngo/projects/logcollections/logagent/kafka v0.0.0
	// github.com/ggvylf/learngo/projects/logcollections/logagent/taillog v0.0.0
	// github.com/ggvylf/learngo/projects/logcollections/logagent/etcd v0.0.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	golang.org/x/sys v0.0.0-20201101102859-da207088b7d1 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect

)

// replace (
// 	github.com/ggvylf/learngo/projects/logcollections/logagent/conf => ./conf
// 	github.com/ggvylf/learngo/projects/logcollections/logagent/kafka => ./kafka
// 	github.com/ggvylf/learngo/projects/logcollections/logagent/taillog => ./taillog
// 	github.com/ggvylf/learngo/projects/logcollections/logagent/etcd => ./etcd
// )
