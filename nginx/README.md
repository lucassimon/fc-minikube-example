# Custom Nginx

[Docker Hub](https://hub.docker.com/r/lucassimon/nginx)

## Docker

Build image:

`$ dk build -t lucassimon/nginx -f Dockerfile .`

Run image

`$ dk run -itd --name=custom-nginx -p 8000:80 lucassimon/nginx:latest`

Push to docker hub

`$ dk push lucassimon/nginx:latest`

## K8s

All config files are in `k8s` folder

`$ kubectl apply -f nginx/k8s/deployment.yaml`

`$ kubectl apply -f nginx/k8s/service.yaml`

`$ minikube service nginx-service`
