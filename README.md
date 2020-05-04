查询IP位置
---
基于命令行程序 [geoiplookup](https://linux.die.net/man/1/geoiplookup) 实现

IP数据库从[maxmind](https://www.maxmind.com/en/home) 下载免费版，或者通过[国内网盘](https://590m.com/dir/26620660-38698816-0b9da5) 下载。

将环境变量`GEO_FILE`设置为IP数据库的文件路径，如：
```
GEO_FILE=/usr/share/GeoIP/GeoLite2-City.mmdb
```

设置环境变量，Web服务器监听地址
```.env
MY_LISTEN=127.0.0.1:7000
```

用法
---
获取本机IP：
```
GET http://127.0.0.1:7000/ip
```

查询IP：
```
GET http://127.0.0.1:7000/ip/8.8.8.8
```

Demo
---
（推荐使用）

https://ip.liuboping.win:1443/my_ip

（不推荐）

https://ip.liuboping.win/my_ip