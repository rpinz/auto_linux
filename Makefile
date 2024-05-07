#
# drupal_auto_linux
# ohrm_auto_linux
#
# Makefile
#

GOLANG := go

GOCYCLO := gocyclo
GOFMT := gofmt
GOLANGCI_LINT := golangci-lint
GOLINT := golint
GOSEC := gosec
GOVULNCHECK := govulncheck
INEFFASSIGN := ineffassign
STATICCHECK := staticcheck

.PHONY: all install build iterate format vet tidy verify golangci-lint lint staticcheck cyclo ineffassign gosec vulncheck clean

all: build

iterate: install format tidy verify lint golangci-lint vet staticcheck cyclo ineffassign gosec vulncheck build

install:
	@echo " 📦 staticcheck"
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@echo " 📦 cyclo"
	@go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@echo " 📦 ineffassign"
	@go install github.com/gordonklaus/ineffassign@latest
	@echo " 📦 vulncheck"
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo " 📦 golint"
	@go install golang.org/x/lint/golint@latest
	@echo " 📦 gosec"
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@echo " 📦 golangci-lint"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

drupal_auto:
	@echo " 🚧 building drupal_auto installer"
	@$(GOLANG) build -v ./cmd/drupal_auto/...

drupal_auto_linux: drupal_auto
	@echo " 💧 built drupal_auto installer"
	@mv drupal_auto drupal_auto_linux

ohrm_auto:
	@echo " 🚧 building ohrm_auto installer"
	@$(GOLANG) build -v ./cmd/ohrm_auto/...

ohrm_auto_linux: ohrm_auto
	@echo " 🍊 built ohrm_auto installer"
	@mv ohrm_auto ohrm_auto_linux

build: drupal_auto_linux ohrm_auto_linux
	@echo " 🚧💧🍊 built installers"

format:
	@echo " 🔘 Go Format"
	@$(GOFMT) -s -w .

vet:
	@echo " 🔘 Go Vet drupal_auto"
	@$(GOLANG) vet ./cmd/drupal_auto/...
	@echo " 🔘 Go Vet ohrm_auto"
	@$(GOLANG) vet ./cmd/ohrm_auto/...

tidy:
	@echo " 🔘 Go Mod Tidy"
	@$(GOLANG) mod tidy

verify:
	@echo " 🔘 Go Mod Verify"
	@$(GOLANG) mod verify

golangci-lint:
	@echo " 🔘 Run golangci-lint"
	@$(GOLANGCI_LINT) run ./...

lint:
	@echo " 🔘 Run golint"
	@$(GOLINT) ./...

staticcheck:
	@echo " 🔘 Run staticcheck"
	@$(STATICCHECK) ./...

cyclo:
	@echo " 🔘 Run gocyclo"
	@echo ' 🌀 ----- files -----'
	@$(GOCYCLO) . || exit 0
	@echo ' 🌀 ----- top 10 -----'
	@$(GOCYCLO) -top 10 . || exit 0
	@echo ' 🌀 ----- over 5 -----'
	@$(GOCYCLO) -over 5 . || exit 0
	@echo ' 🌀 ----- average -----'
	@$(GOCYCLO) -avg . || exit 0
	@echo ' 🌀 ----- over 2 average -----'
	@$(GOCYCLO) -over 2 -avg . || exit 0

ineffassign:
	@echo " 🔘 Run ineffassign"
	@$(INEFFASSIGN) ./...

gosec:
	@echo " 🔘 Run gosec"
	@$(GOSEC) ./...

vulncheck:
	@echo " 🔘 Run govulncheck"
	@$(GOVULNCHECK) -test -show verbose ./...

clean:
	@echo " 🗑️  cleaning ..."
	@$(GOLANG) clean
	@rm -f drupal_auto
	@rm -f drupal_auto_linux
	@rm -f ohrm_auto
	@rm -f ohrm_auto_linux
