

# 测试

``` shell
# server
go run *.go -server=true -P=compact -buffered=true -framed=true -addr=localhost:9090 -secure=false
# client
go run *.go -server=false -P=compact -buffered=true -framed=true -addr=localhost:9090 -secure=false
```

# Thrift reference

* [IDL](https://thrift.apache.org/docs/idl.html)
* [Thrift types](https://thrift.apache.org/docs/types)
* [go get started](https://github.com/apache/thrift/tree/master/tutorial/go)

