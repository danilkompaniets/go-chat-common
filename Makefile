OPENAPI_DIR = gen/openapi
SERVICES = auth chat user common

services-gen:
	protoc -I=api --go_out=gen --go-grpc_out=gen --grpc-gateway_out=gen ./api/auth.proto ./api/chat.proto ./api/common.proto ./api/user.proto

swagger-auth:
	swag init -g services/auth/cmd/main.go -o gen/openapi --outputTypes json
	mv gen/openapi/swagger.json gen/openapi/auth.swagger.json

swagger-gen:
	@for svc in $(SERVICES); do \
		echo "📄 Генерация swagger для $$svc.proto..."; \
		protoc \
			--proto_path=api \
			--openapiv2_out=$(OPENAPI_DIR) \
			--openapiv2_opt logtostderr=true \
			--openapiv2_opt allow_merge=false \
			--openapiv2_opt json_names_for_fields=false \
			api/$$svc.proto; \
	done

swagger-merge:
	@jq -s '{swagger:"2.0", info:{title:"Combined API", version:"v1"}, paths:(map(.paths // {}) | add), definitions:(map(.definitions // {}) | add)}' $(OPENAPI_DIR)/*.swagger.json > $(OPENAPI_DIR)/openapi.json

openapi: swagger-gen swagger-auth swagger-merge
	@echo "✅ OpenAPI spec сгенерирован: $(OPENAPI_DIR)/openapi.json"

.PHONY: swagger-gen swagger-merge openapi build
