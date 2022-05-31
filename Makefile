CC := gcc
CFLAGS := -Wall -Wextra -std=c99 -pthread
CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/audiodb.a build/audiodb.so

.PHONY: all
all: $(LIBS)

build/audiodb.a: audiodb.go
	$(CGO) $(STATIC) -o build/audiodb.a $<

build/audiodb.so: audiodb.go
	$(CGO) $(SHARED) -o build/audiodb.so $<

test: all
	for i in $$(find test -type f -name '*.c'); do \
		$(CC) $(CFLAGS) $$i -o test/main -Ibuild -Lbuild -l:audiodb.a; \
		echo $$i Test :; \
		test/main; \
	done

clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete