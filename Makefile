build:
	go build -o http_server ./cmd/app/main.go

run: build
	./http_server

docker-build:
	eval $(minikube docker-env)
	docker build -t http-server-app:latest .

argo-redeploy:
	argocd app sync http-server-app --force

