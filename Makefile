include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

server:
	@ echo
	@ echo "Spinning up server..."
	@ echo
	@ go run ./cmd/main.go

docker-image:
	@ echo
	@ echo "Building docker image..."
	@ echo
	@ docker build -t machado-br/aws-api:latest .

docker-tag:
	@ echo
	@ echo "Tagging docker image for AWS..."
	@ echo
	@ docker tag machado-br/aws-api:latest 774429751797.dkr.ecr.us-west-2.amazonaws.com/aws-api:latest

docker-push:
	@ echo
	@ echo "Pushing docker image to AWS ECR..."
	@ echo
	@ docker push 774429751797.dkr.ecr.us-west-2.amazonaws.com/aws-api:latest