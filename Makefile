
helmAppName ?= srm1 #server-redis-mysql  srmhelm
NS ?= hwx1166232

template:
	rm -rf outdir-${helmAppName}
	helm template srmhelm/ --namespace ${NS} --values ./values.yaml --name-template ${helmAppName} --output-dir outdir-${helmAppName}

install:
	helm install ${helmAppName} ./srmhelm --values ./values.yaml --namespace ${NS} 

uninstall:
	helm uninstall ${helmAppName} --namespace ${NS} 


# mysql redis app 
ImgInfoGo=swr.cn-north-4.myhuaweicloud.com/hfbbg4/mr-server:v0.1

aaa:
	@echo ${ImgInfo}

build-image:
	docker build -f ./docker/Dockerfile-golang -t ${ImgInfoGo} .

build-go-layels:
	docker build -f ./docker/Dockerfile-golang-layels -t ${ImgInfoGo}-layels .


run:
	@docker-compose up -d
