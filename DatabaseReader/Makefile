READER = dbReader
EXECUTABLE_READER = readDB

COMPARER = dbComparer
EXECUTABLE_COMPARER = compareDB

reader:
	@cd $(READER) && go build -o $(EXECUTABLE_READER)
	@mv $(READER)/$(EXECUTABLE_READER) ./

comparer:
	@cd $(COMPARER) && go build -o $(EXECUTABLE_COMPARER)
	@mv $(COMPARER)/$(EXECUTABLE_COMPARER) ./

clean:
	@rm -rf $(EXECUTABLE_READER)
	@rm -rf $(EXECUTABLE_COMPARER)