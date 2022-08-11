.PHONY = types,client

fetch-spec:
	rm openapi.json
	curl https://api.dev.gcp.mariadb.net/dps/docs/openapi.json | jq . > ./dps/openapi.json
	curl https://api.dev.gcp.mariadb.net/ous/docs/openapi.json | jq . > ./ous/openapi.json
	curl https://api.dev.gcp.mariadb.net/nos/docs/openapi.json | jq . > ./nos/openapi.json

all: fetch-spec
