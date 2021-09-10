.PHONY = types,client

types:
	oapi-codegen -generate=types -package=skysql openapi.json > types.go

client:
	oapi-codegen -generate=client -package=skysql openapi.json > client.go

all: types client
