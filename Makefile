# metweather
# Copyright (C) 2021 Dakota Walsh
# GPL3+ See LICENSE in this repo for details.
.POSIX:

include config.mk

all: clean build

build:
	go build
	scdoc < metweather.1.scd | sed "s/VERSION/$(VERSION)/g" > metweather.1

clean:
	rm -f metweather
	rm -f metweather.1

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f metweather $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/metweather
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f metweather.1 $(DESTDIR)$(MANPREFIX)/man1/metweather.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/metweather.1

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/metweather
	rm -f $(DESTDIR)$(MANPREFIX)/man1/metweather.1

.PHONY: all build clean install uninstall
