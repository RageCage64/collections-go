.PHONY: install_tools
install_tools:
	go install github.com/google/addlicense@latest
	go install github.com/jordanlewis/gcassert/cmd/gcassert@latest

.PHONY: addlicense
addlicense:
	addlicense -c "Braydon Kains" -l mit .	
