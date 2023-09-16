build:
	cd web && yarn build
	ENV=prod go build -buildvcs=false -o ./bin/go-vite ./main.go

dev:
	cd web && yarn start & air && fg

client:
	cd web && yarn start

server:
	nodemon --exec  "ENV=dev go run" main.go