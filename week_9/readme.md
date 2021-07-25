# week_9

总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用

- fix length: 定长结构，每一个数据包的长度固定，发送方按照固定长度发送数据。接收方在接受到数据后将数据保存到缓冲区，当缓冲区可读取的长度不足时继续等到，直到可读取的数据到达固定长度，从其中读取出来。如果缓冲区中数据大于固定长度，从其中读取出固定长度的数据，剩余的数据继续等待。
- delimiter based: 基于分隔符。在协议中定义一个头部标识符和尾部标识符。程序从缓冲区中逐个读取出数据，当读取到头部标识符后将后面的数据全部读入到缓存中，直到读取到尾部标识符。缓存中的数据即为完整的数据包。http协议就是使用的该方式。http协议中，使用`\n\r`来区分请求头，头部，空行和body。
- length field based frame decoder: 通过长度来标识数据包的大小。首先会读取到整个数据包的长度，之后读取出对应长度的内容，作为数据包。http POST请求时使用了这种方式，POST时通过ContentType头来指定请求中POSTbody的长度。