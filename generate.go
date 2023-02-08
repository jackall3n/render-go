package render

//go:generate-old oapi-codegen --package=render -generate=types,client -o ./render.go ./openapi.yaml
//go:generate openapi-generator generate -i ./api/openapi.yaml -g go -o . -p "packageName=render,isGoSubmodule=yes" --git-user-id jackall3n --git-repo-id=render-go
