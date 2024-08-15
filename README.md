
## 实践学习
git config --global url."https://".insteadOf git://

- 镜像编译
  - make build-image

- 依赖：website编译
  - cd website;npm install;npm run build:prod;cp -rf dist ../server/

- 在k8s安装部署
 - cd yamls-in-k8s/server ; make install
 - 查看配置文件： docker run -it swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1  sh -c "ls -lh /root/ ; cat config.yaml"


# go-example-app

<!-- PROJECT SHIELDS -->


<!-- ABOUT THE PROJECT -->
## About The Project

This is a `gin + vue`  example app, include user manage , rbac permision, 
menu generator and kubernetes manage. 

### Using Component

This Demo Project Using These Component, Thanks for these project: 

- gin
- vue-element-admin
- mysql
- redis
- kubernetes
- swagger
- prometheus-metrics
- cors
- casbin
- gorm
- zap

### Project Structure

The project not batter structure, beacuse I'm a beginner, 😄.

```
├── LICENSE
├── Makefile # simple run commad
├── README.md 
├── docker-compose.yaml # run this project on local
├── go.mod
├── go.sum
├── server  # golang backend  api-server
└── website # vue frontend project
```

## Try the project

pre required need kubernetes, you can use kind create it.

### Local

after run this project. nedd docker-compose kubernetes cluster and kubeconfig.


change kubeconfig path:

```bash
kubernetes:
  # type in or out cluster type
  type: out
  # type out need kubeconfig path
  config: /root/.kube/config
```

run docker-compose

```bash
make run
```

run backend 

```bash
cd server
go run main.go --config configs/dev.yaml
```

run frontend

```bash
yarn dev
```

### SnapShot

<details>
<summary>Expand view</summary>
    <pre>
        <code>
        <img src="./img/menu.png" width="100%"/><br/><br/>
        <img src="./img/role.png" width="100%"/><br/><br/>
        <img src="./img/namespace.png" width="100%"/><br/><br/>
        <img src="./img/deployment.png" width="100%"/><br/><br/>
        <img src="./img/swagger.png" width="100%"/><br/><br/>
        </code>
    </pre>
</details>


