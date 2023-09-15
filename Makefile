build:
	cd web && yarn build
	ENV=prod go build -buildvcs=false -o ./bin/go-vite ./main.go

dev:
	cd web && yarn start & air && fg