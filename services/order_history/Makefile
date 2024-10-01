gen:
	@go generate ./...

test: gen
	@go test ./...

build-image:
	@docker build -t indigowar/food_out_order_history .

push-image:
	@docker push indigowar/food_out_order_history
