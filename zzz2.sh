



```sh


make install


git remote -v

# server-redis-mysql
helm create srmhelm


docker push swr.cn-north-4.myhuaweicloud.com/hfbbg4/myalpinetest:curl

make build-image

docker images |grep  hfbbg4/mr-server



docker history swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1
docker push swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1
docker history swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1 --no-trunc


docker run -it swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1  sh -c "pwd;ls -lh; "
docker run -it swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1  sh -c "pwd;ls dist; "
docker run -it swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1  sh -c "pwd;ls;./app --config config.yaml"
docker run -it swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1  sh -c "ls -lh /root/ ; cat config.yaml"



kubectl delete -f ./yamls-in-k8s/server/server.yaml -n hwx1166232
kubectl apply -f ./yamls-in-k8s/server/server.yaml -n hwx1166232

kubectl get serviceaccount/mr-server-redis-mysql  -n hwx1166232 -oyaml

docker tag centos:centos7.6.1810 swr.cn-north-4.myhuaweicloud.com/hfbbg4/centos:centos7.6.1810
docker push swr.cn-north-4.myhuaweicloud.com/hfbbg4/centos:centos7.6.1810


# 测试访问
http://192.168.113.249:30712
http://myelbip:30711



```



