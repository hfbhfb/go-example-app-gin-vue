
helmAppName ?= myredis1
NS ?= hwx1166232-a


template:
	rm -rf outdir-${helmAppName}
	helm template redishelm/ --namespace ${NS} --values ./values.yaml --name-template ${helmAppName} --output-dir outdir-${helmAppName}

install:
	helm install ${helmAppName} ./redishelm --values ./values.yaml --namespace ${NS} 

uninstall:
	helm uninstall ${helmAppName} --namespace ${NS} 



installold:
	kubectl apply -f redis.yaml

uninstallold:
	kubectl delete -f redis.yaml