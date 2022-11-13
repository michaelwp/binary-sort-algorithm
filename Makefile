.PHONY: run
run:
	@echo "> run application ..."
	@go run . --race
	@echo "> done"

.PHONY: test
test:
	@echo "> run testing ..."
	@go test
	@echo "> testing done"

.PHONY: benchmark
benchmark:
	@echo "> run benchmark ..."
	@go test -bench=. -count 5
	@echo "> benchmark done"
