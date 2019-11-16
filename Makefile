.PHONY: install test build serve clean pack ship
OWNER=gd
SERVICE=awssm2env
# it will take short hash (first 7 symbols) of last git commit.
# Then, we export this variable, so it’s available in commands run by make.
#TAG=$(git rev-list HEAD --max-count=1 --abbrev-commit)
TAG=latest
export TAG

install:
	go get .

test:
	go test ./...

build: install
	# build a binary
	go build -ldflags "-X main.version=$(TAG)" -o ${SERVICE} .

serve: build
	./${SERVICE}

clean:
	# remove binary file
	rm ./${SERVICE}

pack:
	# build docker image
	GOOS=linux make build
	docker build -f docker/alpine/Dockerfile -t ${OWNER}/${SERVICE}:$(TAG) .

ship: test pack clean

