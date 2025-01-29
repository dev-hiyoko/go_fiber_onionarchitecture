GO_AIR_VERSION=latest
GO_STATICCHECK_VERSION=latest
GO_OAPI_CODEGEN_VERSION=latest
GO_MOCKGEN_VERSION=latest

# go
go/install/tools:
	go install github.com/cosmtrek/air@$(GO_AIR_VERSION)
	go install honnef.co/go/tools/cmd/staticcheck@$(GO_STATICCHECK_VERSION)
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@$(GO_OAPI_CODEGEN_VERSION)
	go install github.com/golang/mock/mockgen@$(GO_MOCKGEN_VERSION)

# staticcheck
go/lint:
	staticcheck ./...

# ent
ent/gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob="./internal/pkg/ent/template/*.tmpl" ./internal/pkg/ent/schema

# oapi
oapi/gen/app:
	docker run --rm -v $(PWD)/docs:/spec redocly/cli:latest bundle api/app/oapi-spec/root.yml -o api/app/oapi-spec/root.gen.yml
oapi/validate/app:
	docker run --rm -v $(PWD)/docs/api/app/oapi-spec:/spec openapitools/openapi-generator-cli validate -i /spec/root.gen.yml
oapi/run/app:
	docker run -p 8081:8080 -v $(PWD)/docs/api/app/oapi-spec:/usr/share/nginx/html/api -e API_URL=api/root.gen.yml swaggerapi/swagger-ui
oapi/codegen/app:
	oapi-codegen  --config ./docs/api/app/oapi-spec/oapicodegen.yml ./docs/api/app/oapi-spec/root.gen.yml

# mockgen
mockgen:
	bash -x scripts/mockgen.sh

# docker
docker/up:
	docker compose --env-file ./cmd/app/.env up -d --build
docker/up/db:
	docker compose --env-file ./cmd/app/.env up -d db --build
docker/up/mailhog:
	docker compose --env-file ./cmd/app/.env up -d mailhog --build
docker/exec/go:
	docker compose --env-file ./cmd/app/.env exec go ash

# git
git/commit-template:
	git config commit.template ./.github/.gitmessage.txt &&\
	git config --add commit.cleanup strip

# other
sleep:
	sleep 20
