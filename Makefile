.PHONY: report
report:
	go test -coverprofile=coverage.out ./...
	# Remove mock and stub files from coverage report
	more coverage.out | grep -v "mock" > coverage_filtered.out
	rm coverage.out
	mv coverage_filtered.out coverage.out

.PHONY: html_report
html_report: report
	go tool cover -html=coverage.out

.PHONY: func_report
func_report: report
	go tool cover -func=coverage.out