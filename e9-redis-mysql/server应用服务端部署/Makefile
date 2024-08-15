



helmAppName ?= mrk1 # mrk-helm 是指 mysql redis kubernetes
NS ?= hwx1166232-a


template:
	rm -rf outdir-${helmAppName}
	helm template mrk-helm/ --namespace ${NS} --values ./values.yaml --name-template ${helmAppName} --output-dir outdir-${helmAppName}

install:
	helm install ${helmAppName} ./mrk-helm --values ./values.yaml --namespace ${NS} 

uninstall:
	helm uninstall ${helmAppName} --namespace ${NS} 




