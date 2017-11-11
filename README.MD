# 服务计算作业 Cloudgo
学号：15331366
# 框架
gin

参考[7 Frameworks To Build A REST API In Go](https://nordicapis.com/7-frameworks-to-build-a-rest-api-in-go/)

gin是一个小规模、高效的web框架，功能偏少但高并发时性能好

# curl 测试

命令

```
curl -v http://127.0.0.1:8080/hello/pfjhyyj
```

结果

```
StatusCode        : 200
StatusDescription : OK
Content           : Hello pfjhyyj
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 13
                    Content-Type: text/plain; charset=utf-8
                    Date: Sat, 11 Nov 2017 16:28:37 GMT
                    
                    Hello pfjhyyj
Forms             : {}
Headers           : {[Content-Length, 13], [Content-Type, text/plain; charset=utf-8], [Date, Sat, 11 Nov 2017 16:28:37 
                    GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 13
```

# Apache Web 测试

命令

```
ab -n 1000 -c 100 http://127.0.0.1:8080/hello/pfjhyyj
```

结果

```
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /hello/pfjhyyj
Document Length:        13 bytes

Concurrency Level:      100
Time taken for tests:   0.937 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      130000 bytes
HTML transferred:       13000 bytes
Requests per second:    1067.07 [#/sec] (mean)
Time per request:       93.715 [ms] (mean)
Time per request:       0.937 [ms] (mean, across all concurrent requests)
Transfer rate:          135.47 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       4
Processing:     2   90  38.4     85     296
Waiting:        1   82  37.2     79     286
Total:          2   90  38.4     85     296

Percentage of the requests served within a certain time (ms)
  50%     85
  66%     97
  75%    105
  80%    111
  90%    143
  95%    157
  98%    186
  99%    217
 100%    296 (longest request)
```