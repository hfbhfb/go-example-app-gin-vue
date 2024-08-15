

## k8s集群中部署mysql  **id329shqi**  依赖： 429vuwo


make uninstall
make install


#### 测试连接
kubectl exec -it mysql-0 -- bash -c "export LANG=C.UTF-8;export LC_ALL=C.UTF-8;export LC_CTYPE=C.UTF-8;mysql -p123456 mydb1"

drop database mydb1;
create database mydb1;

## 把负本数调整为0
exit
kubectl get po
kubectl scale statefulset mysql --replicas=0 -n default
kubectl scale statefulset mysql --replicas=1 -n default




#### 使用 mysqlworkbench 测试连接 创建一些测试数据
render.tpddns.cn:30306
```sql


```
