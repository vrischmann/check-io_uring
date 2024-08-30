docker:
	docker build -t foobar -f Dockerfile --platform linux/arm64 .
	docker run --rm -ti --platform linux/arm64 foobar /run.sh

dagger:
	dagger call foo  --source=. --platform=linux/arm64 stdout
