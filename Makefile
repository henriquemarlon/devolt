-include .env.develop

START_LOG = @echo "================================================= START OF LOG ==================================================="
END_LOG = @echo "================================================== END OF LOG ===================================================="

.PHONY: env
env: ./.env.develop
	$(START_LOG)
	@cp ./.env.develop.tmpl ./.env.develop
	@touch .cartesi.env
	@echo "Environment file created at ./.env.develop"
	$(END_LOG)

.PHONY: dev
dev:
	$(START_LOG)
	# @rm -r devolt.db
	@nonodo -- go run ./cmd/dapp/
	$(END_LOG)

.PHONY: build
build:
	$(START_LOG)
	@docker build \
		-t dapp:latest \
		-f ./build/Dockerfile.dapp .
	@cartesi build --from-image dapp:latest
	$(END_LOG)
	
.PHONY: generate
generate:
	$(START_LOG)
	@go run ./pkg/rollups_contracts/generate
	$(END_LOG)

.PHONY: test
test:
	@go test -p 1 ./... -coverprofile=./coverage.md -v

.PHONY: deploy
deploy:
	$(START_LOG)
	@cd contracts && forge script $(DEPLOY_NETWORK_ARGS)
	$(END_LOG)

.PHONY: coverage
coverage: test
	@go tool cover -html=./coverage.md

.PHONY: docs
docs:
	@cd docs && npm run dev


.PHONY: volt
volt:
	@cd contracts && forge script script/DeployVoltToken.s.sol --rpc-url $(RPC_URL) --broadcast --verify --etherscan-api-key $(API_KEY) -vvv