gen:
	@go generate ./...

test: gen
	@go test ./...

build-image:
	@docker build -t indigowar/food_out_orders .

push-image:
	@docker push indigowar/food_out_orders
