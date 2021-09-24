.PHONY = types,client

fetch-spec:
	rm openapi.json
	curl https://api.dev.gcp.mariadb.net/openapi.json | jq . > openapi.json

types:
	oapi-codegen -generate=types -package=skysql openapi.json > types.go

client:
	oapi-codegen -generate=client -package=skysql openapi.json > client.go

all: fetch-spec types client
