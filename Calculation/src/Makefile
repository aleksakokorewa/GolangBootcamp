FILE = main.go
MEAN = -mean
MEDIAN = -median
SD = -sd
MODE = -mode
EXECUTABLE = src

build:
	go build
run: build
	./$(EXECUTABLE)
mean:
	./$(EXECUTABLE) $(MEAN)
mode:
	./$(EXECUTABLE) $(MODE)
median:
	./$(EXECUTABLE) $(MEDIAN)
sd:
	./$(EXECUTABLE) $(SD)
clean:
	@rm -rf $(EXECUTABLE)