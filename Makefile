REPO=github.com/boynton
NAME=conf

install::
	go install $(REPO)/$(NAME)

test::
	go test $(REPO)/$(NAME)

proper::
	go fmt $(REPO)/$(NAME)
	go vet $(REPO)/$(NAME)
