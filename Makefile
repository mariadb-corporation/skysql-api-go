.PHONY = client, purge

purge:
	bash purge.sh

client: purge
	openapi-generator-cli generate -i openapi.json -g go --config .openapi-generator/config.yaml
	rm -rf api  # the openapi ignore still leaves the empty dir behind
