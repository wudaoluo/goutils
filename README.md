# goutils

 一些经常用到的库


 #服务发现

 ### consul启动命令
    ./consul agent -dev -ui -node=consul-dev -bind=10.211.55.4 -client=0.0.0.0

 ### 增加consul服务发现
    a:=servDiscover.NewConsul("carlo","10.211.55.4","10.211.55.4:8500",9090,[]string{"tomcat"})
    err:=a.ServDiscover()


 ### 参考:
    https://github.com/johng-cn/gf
    https://github.com/henrylee2cn/goutil



