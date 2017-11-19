## Benchmark

**测试环境**
* CPU:    Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz, 4 cores
* Memory: 16G
* OS:     Linux Server-3 2.6.32-358.el6.x86_64, CentOS 6.4
* Go:     1.7

测试代码client是通过protobuf编解码和server通讯的。
请求发送给server, server解码、更新两个字段、编码再发送给client，所以整个测试会包含客户端的编解码和服务器端的编解码。
消息的内容大约为581 byte, 在传输的过程中会增加少许的头信息，所以完整的消息大小在600字节左右。

测试用的proto文件如下：

```proto
syntax = "proto2";

package main;

option optimize_for = SPEED;


message BenchmarkMessage {
  required string field1 = 1;
  optional string field9 = 9;
  optional string field18 = 18;
  optional bool field80 = 80 [default=false];
  optional bool field81 = 81 [default=true];
  required int32 field2 = 2;
  required int32 field3 = 3;
  optional int32 field280 = 280;
  optional int32 field6 = 6 [default=0];
  optional int64 field22 = 22;
  optional string field4 = 4;
  repeated fixed64 field5 = 5;
  optional bool field59 = 59 [default=false];
  optional string field7 = 7;
  optional int32 field16 = 16;
  optional int32 field130 = 130 [default=0];
  optional bool field12 = 12 [default=true];
  optional bool field17 = 17 [default=true];
  optional bool field13 = 13 [default=true];
  optional bool field14 = 14 [default=true];
  optional int32 field104 = 104 [default=0];
  optional int32 field100 = 100 [default=0];
  optional int32 field101 = 101 [default=0];
  optional string field102 = 102;
  optional string field103 = 103;
  optional int32 field29 = 29 [default=0];
  optional bool field30 = 30 [default=false];
  optional int32 field60 = 60 [default=-1];
  optional int32 field271 = 271 [default=-1];
  optional int32 field272 = 272 [default=-1];
  optional int32 field150 = 150;
  optional int32 field23 = 23 [default=0];
  optional bool field24 = 24 [default=false];
  optional int32 field25 = 25 [default=0];
  optional bool field78 = 78;
  optional int32 field67 = 67 [default=0];
  optional int32 field68 = 68;
  optional int32 field128 = 128 [default=0];
  optional string field129 = 129 [default="xxxxxxxxxxxxxxxxxxxxx"];
  optional int32 field131 = 131 [default=0];
}
```



测试的并发client是 100, 1000,2000 and 5000。总请求数一百万。

**测试结果**

### 一个服务器和一个客户端，在同一台机器上

- teleport

并发client|平均值(ms)|中位数(ms)|最大值(ms)|最小值(ms)|吞吐率(TPS)
-------------|-------------|-------------|-------------|-------------|-------------
100|3|3|32|0|26164
500|19|18|130|0|24923
1000|43|41|298|0|22211
2000|89|90|657|0|21543
5000|230|156|1961|0|19683

- teleport-socket

并发client|平均值(ms)|中位数(ms)|最大值(ms)|最小值(ms)|吞吐率(TPS)
-------------|-------------|-------------|-------------|-------------|-------------
100|2|2|25|0|35489
500|14|13|69|0|34689
1000|27|26|107|0|34663
2000|67|66|187|0|28476
5000|171|166|401|0|27669

- rpcx

并发client|平均值(ms)|中位数(ms)|最大值(ms)|最小值(ms)|吞吐率(TPS)
-------------|-------------|-------------|-------------|-------------|-------------
100|3|3|25|0|31477
500|18|16|108|0|27526
1000|37|34|230|0|26517
2000|77|74|327|0|24865
5000|193|197|844|0|23247
