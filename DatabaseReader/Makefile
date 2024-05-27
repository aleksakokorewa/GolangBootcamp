READER = dbReader
EXECUTABLE_READER = readDB

COMPARER = dbComparer
EXECUTABLE_COMPARER = compareDB

FS_COMPARER = fsComparer
EXECUTABLE_FS = compareFS

reader:
	@cd $(READER) && go build -o $(EXECUTABLE_READER)
	@mv $(READER)/$(EXECUTABLE_READER) ./

comparer:
	@cd $(COMPARER) && go build -o $(EXECUTABLE_COMPARER)
	@mv $(COMPARER)/$(EXECUTABLE_COMPARER) ./

fs_comparer:
	@cd $(FS_COMPARER) && go build -o $(EXECUTABLE_FS)
	@mv $(FS_COMPARER)/$(EXECUTABLE_FS) ./

clean:
	@rm -rf $(EXECUTABLE_READER)
	@rm -rf $(EXECUTABLE_COMPARER)