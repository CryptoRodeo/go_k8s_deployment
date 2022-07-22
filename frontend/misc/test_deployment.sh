echo "== Deploying app to k8s... ==\n"
kubectl apply -f Deployment.yaml
sleep 2
echo "\n== Port-forwarding to svc, check localhost:3000 ==\n"
kubectl port-forward svc/frontend 3000:80