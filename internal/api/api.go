// internal/api/api.go

//go:generate oapi-codegen -package=api -generate=types,chi-server -o api.gen.go api.yaml

package api

/*
Explanation:

The //go:generate directive tells Go to run the oapi-codegen command to generate code based on api.yaml.
We specify the package name as api and generate types and a Chi server (we'll adapt for Gorilla Mux later).
*/
