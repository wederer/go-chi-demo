run:
	DB_URL=http://localhost:31529 DB_USER=root DB_PW=test go run github.com/wederer/go-chi-demo/cmd/web
build:
	go build -o bin/ github.com/wederer/go-chi-demo/cmd/web
build-docker:
	docker build -f "service/Dockerfile" . -t go-chi-demo:latest
kubectl-apply:
	kubectl apply -f service
