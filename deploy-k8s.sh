#!/usr/bin/sh
kompose convert

kubectl create -f alertmanager-claim0-persistentvolumeclaim.yaml,alertmanager-service.yaml,alertmanager-deployment.yaml,db-claim0-persistentvolumeclaim.yaml,db-deployment.yaml,grafana-deployment.yaml,grafana-service.yaml,monitaapi-claim0-persistentvolumeclaim.yaml,monitaapi-deployment.yaml,monitaapi-service.yaml,monitanet-networkpolicy.yaml,mysqlexporter-deployment.yaml,prometheus-claim0-persistentvolumeclaim.yaml,prometheus-claim1-persistentvolumeclaim.yaml,prometheus-deployment.yaml,prometheus-service.yaml,prom-receiver-deployment.yaml

rm alertmanager-claim0-persistentvolumeclaim.yaml alertmanager-service.yaml alertmanager-deployment.yaml db-claim0-persistentvolumeclaim.yaml db-deployment.yaml grafana-deployment.yaml grafana-service.yaml monitaapi-claim0-persistentvolumeclaim.yaml monitaapi-deployment.yaml monitaapi-service.yaml monitanet-networkpolicy.yaml mysqlexporter-deployment.yaml prometheus-claim0-persistentvolumeclaim.yaml prometheus-claim1-persistentvolumeclaim.yaml prometheus-deployment.yaml prometheus-service.yaml prom-receiver-deployment.yaml 
