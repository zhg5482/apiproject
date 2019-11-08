# apiproject
   
## beego 安装
    cd $GOLANG/src
    go get github.com/astaxie/beego
    
## bee 安装
    cd $GOLANG/src
    go get github.com/beego/bee
    参考文档 ： https://www.cnblogs.com/zuxingyu/p/6016816.html
    https://www.cnblogs.com/xiaochuizi/p/9257521.html
    
## 创建项目
    bee api apiproject
    
## 启动项目
    启动 ./main.go
    自动编译启动 bee run (修改后自动生效)
    //bee run -gendoc=true -downdoc=true
    参考文档 ： https://blog.csdn.net/loongshawn/article/details/54114606
    
## models 配置
    models/models.go
    参考文档:https://blog.csdn.net/boss2967/article/details/82792076
    
## redis 连接池
    lib/refisCache.go
    参考 ：https://www.cnblogs.com/mafeng/p/6322957.html
# 项目相关配置 app.conf
# nginx 配置 golang.conf

# swagger 文档
    http://aip.golang.com/swagger/index.html

    