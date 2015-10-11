.PHONY: .FORCE
GO=go
DOT=dot

PROGS = service

DOC_DIR = ./doc
SRCDIR = ./src

GRAPHS = $(DOC_DIR)/arch.png $(DOC_DIR)/fsm.png

all: fmt $(PROGS) $(GRAPHS)

$(PROGS):
	$(GO) install $@

$(DOC_DIR)/%.png: $(DOC_DIR)/%.dot
	$(DOT) -Tpng $< -o $@

clean:
	rm -rf $(addprefix bin/,$(PROGS))  pkg
 
fmt:
	$(GO) fmt $(SRCDIR)/...
