swagger:
	GO111MODULE=off swagger generate spec -o ./docs/swagger.yml  --scan-models

markdown:
	swagger generate markdown -f ./docs/swagger.yml --output=./docs/docs.md