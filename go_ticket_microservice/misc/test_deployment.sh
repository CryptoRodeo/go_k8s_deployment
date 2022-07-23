echo "== Deploying app (go_ticket_microservice) to k8s... ==\n"
kubectl apply -f Deployment.yaml
sleep 2
echo "\n== Port-forwarding to svc/go-tickets, check localhost:8080==\n"
kubectl port-forward svc/go-tickets 8081:80