# go-common

## 1. logging
日志模块，可以输出到控制台，也可以输出到文件。

### v1
支持根据文件大小来rotate。目前依赖了logrus和lumberjack。

### v2
封装了log标准库,提供了文件log和控制台log。目前不支持rotate，rotate可以利用linux的logrotate。

### v3
