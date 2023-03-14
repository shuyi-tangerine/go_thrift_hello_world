module github.com/shuyi-tangerine/go_thrift_hello_world

go 1.19

replace git.apache.org/thrift.git v0.18.1 => github.com/apache/thrift v0.18.1
replace github.com/apache/thrift v0.18.1 =>  git.apache.org/thrift.git v0.18.1

require (
	git.apache.org/thrift.git v0.18.1 // indirect
	github.com/apache/thrift v0.18.1 // indirect
)
