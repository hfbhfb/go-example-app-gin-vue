
helmAppName ?= mysql1
NS ?= hwx1166232-a


template:
	rm -rf outdir-${helmAppName}
	helm template signal-mysql/ --namespace ${NS} --values ./values.yaml --name-template ${helmAppName} --output-dir outdir-${helmAppName}

install:
	helm install ${helmAppName} ./signal-mysql --values ./values.yaml --namespace ${NS} 

uninstall:
	helm uninstall ${helmAppName} --namespace ${NS} 



installold:
	kubectl apply -f mysql.yaml

uninstallold:
	kubectl delete -f mysql.yaml


