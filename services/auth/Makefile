gen:
	@go generate ./...

test:
	@go test ./...


build-image:
	@docker build -t indigowar/food_out_auth:latest .

push-image:
	@docker push indigowar/food_out_auth:latest
