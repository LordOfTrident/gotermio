SRC   = $(wildcard *.go)
TESTS = $(basename $(wildcard tests/*.go))

GO = go

tests: $(TESTS)

$(TESTS): tests/% : tests/%.go $(SRC)
	$(GO) build -o $@ $<

clean:
	rm $(TESTS)

all:
	@echo tests, clean
