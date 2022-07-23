echo "== Deploying app (go_user_microservice) to k8s... ==\n"
kubectl apply -f Deployment.yaml
sleep 2
echo "\n== Port-forwarding to svc/go-users, check localhost:8080==\n"
kubectl port-forward svc/go-users 8080:80