


```sh



kubectl get po -n hwx1166232

docker images|grep mysql

docker mysq:5.7.41 swr.cn-north-4.myhuaweicloud.com/hfbbg4/mysq:5.7.41
docker push swr.cn-north-4.myhuaweicloud.com/hfbbg4/mysq:5.7.41

docker login -u cn-north-4@05EXSOCEPRNZ3HZJUML9 -p c8b6799ed788a276405b565070c4a1db64abc3ec6918c8dcd1a7af8398d8f163 swr.cn-north-4.myhuaweicloud.com




kubectl exec -it mysql-mysql1-0 -n hwx1166232 -- bash
kubectl exec -it mysql-mysql1-0 -n hwx1166232 -- bash -c "export LANG=C.UTF-8;export LC_ALL=C.UTF-8;export LC_CTYPE=C.UTF-8;mysql -p123456 mydb1"
kubectl exec -it mysql-mysql1-0 -n hwx1166232 -- bash -c "export LANG=C.UTF-8;export LC_ALL=C.UTF-8;export LC_CTYPE=C.UTF-8;mysql -p123456 "


show databases;
create database if not exists mydb1;
create database if not exists demo;
use mydb1;
-- 创建表
CREATE TABLE users (
    id INT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL
);
-- 插入数据
INSERT INTO users (id, username, email)
VALUES
    (1, 'john_doe', 'john@example.com'),
    (2, 'jane_smith', 'jane@example.com'),
    (3, 'bob_jones', 'bob@example.com');

use mydb1;
show tables;
select * from users;


use demo;
show tables;

```

