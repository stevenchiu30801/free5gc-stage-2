.PHONY: install uninstall clean

GOFILENAME = pfcpMessage.go
GOFILEDIR = pfcpType

all:
		chmod +x pfcpDocxParser.py && ./pfcpDocxParser.py
		gofmt -w pfcpType
		gofmt -w pfcpMessage.go
install: $(GOFILENAME) $(GOFILEDIR)
		cp -r $^ ../
		
clean: $(GOFILENAME) $(GOFILEDIR)
		rm -rf $^

uninstall: $(GOFILENAME) $(GOFILEDIR)
		cd .. && rm -rf $^
