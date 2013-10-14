gorevel
=======

Go语言Revel框架学习，站点使用Revel框架、Qbs构建。

配置文件在 src/revelapp/conf 目录中，主配置app.conf，自定义配置my.conf (包括数据库配置、邮件发送配置)。

默认的数据库是mysql，数据库名gorevel，表结构不需要创建，由程序运行时自动创建。

###Requirements

- Go1.0+
- github.com/robfig/revel
- github.com/robfig/revel/revel
- github.com/coocood/qbs
- github.com/coocood/mysql
- code.google.com/p/go-uuid/uuid
- github.com/disintegration/imaging

###Install

    $ git clone git://github.com/goofcc/gorevel.git
    $ cd gorevel

Linux/Unix/OS X:

    $ ./install.sh
    $ ./run.sh

Windows:

    > install.bat
    > run.bat
    
打开浏览器访问 [http://localhost:9000](http://localhost:9000)

