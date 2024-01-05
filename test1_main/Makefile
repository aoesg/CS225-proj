exec_name = server_test1
exec = $(exec_name)

CC = /usr/local/go/bin/go

$(exec) : main.go
	$(CC) build -o $@ $^

.PHONY : clean
clean:
	-rm $(exec)

.PHONY : run
run:
	./$(exec)