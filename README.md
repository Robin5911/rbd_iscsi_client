# rbd-iscsi-client (golang)
参考 python rbd-iscsi-client 实现的go版本的 spdk iscsi target sdk， 用于请求操作基于 spdk iscsi target(bdev基于ceph rbd)
# Installation
```
go get github.com/Robin5911/rbd_iscsi_client
```
###### 客户端请求
![avatar](https://img-blog.csdnimg.cn/img_convert/5bab74f8ad589febdeb65eba48465029.png)
###### spdk iscsi target架构
https://spdk.io/doc/iscsi.html