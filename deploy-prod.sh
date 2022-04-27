#!/bin/sh

echo "Deploying services on kubernetes.."


kubectl apply -f deployments/k8s/. -R

sleep 30

kubectl -n csapi get all -o wide

echo "\n\n\nPlease use EXTERNAL-IP from LoadBalancer to request the API "
echo "kubectl port-forward -n csapi deployment.apps/grafana 3000:3000"
echo ".. and browse localhost:3000 to connect grafana dashboard, use admin/admin credentials\n\n\n\n"

kubectl -n csapi get svc

