TAG=v1.$(shell date +%s)

build:
	go build -o http_server ./cmd/app/main.go

run: build
	./http_server

docker-build:
	@bash -c 'eval $$(minikube docker-env) && docker build -t http-server-app:$(TAG) .'
	sed -i "s|image: http-server-app:.*|image: http-server-app:$(TAG)|" ./k8s/deployment.yaml

commit:
	git commit -am "update image tag to $(TAG)"
	git push

argo-redeploy:
	argocd app sync http-server-app --force && kubectl rollout restart deployment http-server-app

redeploy:
	eval $(minikube docker-env) && docker build -t http-server-app:latest .
	kubectl rollout restart deployment http-server-app


argo-create:
	argocd app create http-server-app \
    --repo https://github.com/rreyfockandroid/app-web \
    --path k8s \
    --dest-namespace default \
    --dest-server "https://kubernetes.default.svc" \
    --sync-policy automated --self-heal --auto-prune

web:
	minikube service http-server-app
