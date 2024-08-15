

## 单节点redis在集群里安装        **585hsss**      依赖： 429vuwo   

make install

#### windows下redis服务开启搜索 proj33-redis-server 
                   服务端开启本机 8749ksdfjsrw



- 在命令行中使用Redis可以通过`redis-cli`命令。以下是一些基本的Redis命令行使用示例：

```bash

docker tag redis:7.0.9 swr.cn-north-4.myhuaweicloud.com/hfbbg4/redis:7.0.9
docker push swr.cn-north-4.myhuaweicloud.com/hfbbg4/redis:7.0.9

kubectl exec -it redis-0 -n hwx1166232 -- bash -c "redis-cli -h 127.0.0.1 -p 6379 "


#1. **连接到本地Redis服务器：**
redis-cli -h 127.0.0.1 -p 6379

# 1.1 查看当前数据库信息
info keyspace

# 1.2 删除数据库
flushdb
flushall

# 2. **设置键值对：**
set my_key "Hello, Redis!"

# 3. **获取键对应的值：**
get my_key

# 4. **增加计数器：**
incr my_counter

# 5. **查看所有键：**
keys *

   # 注意：在生产环境中，避免在生产服务器上使用`keys *`，因为它会阻塞服务器一段时间。

# 6. **使用列表数据结构：**
lpush my_list "item1"
lpush my_list "item2"
lrange my_list 0 -1

   # 这个例子中，我们使用`lpush`向列表中添加元素，然后使用`lrange`获取列表的所有元素。

# 7. **使用集合数据结构：**
sadd my_set "member1"
sadd my_set "member2"
smembers my_set

# 8. **哈希表**
HSET myhash field1 "value1"
HSET myhash field2 "value1"
HGET myhash field1
HGETALL myhash 




```


##   这个例子中，我们使用`sadd`向集合中添加成员，然后使用`smembers`获取集合的所有成员。

- 这只是一些基本的Redis命令行操作示例，你可以根据自己的需求使用更多的命令和选项。确保在生产环境中谨慎操作，以避免数据丢失或其他不良后果。
