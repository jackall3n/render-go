patch:
	git apply ./patches/add-env-vars-patch.patch

generate:
	oapi-codegen -package render openapi.yaml > render.go

build: patch generate