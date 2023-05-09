dep:
	go mod tidy

kitex_gen_base:
	kitex --thrift-plugin validator -module github.com/kalandramo/appdemo ./idl/base.thrift

kitex_gen_user:
	kitex --thrift-plugin validator -module github.com/kalandramo/appdemo ./idl/user.thrift
