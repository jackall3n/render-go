package render

//go:generate-old oapi-codegen --package=render -generate=types,client -o ./render.go ./openapi.yaml
//go:generate openapi-generator generate -i ./openapi.yaml -g go -o ./client -p "packageName=render"
