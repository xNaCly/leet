run:
	gcc $(file) -o $(file).out
	./$(file).out
debug:
	gcc $(file) -fsanitize=address -g3 -o $(file).out
	gdb -ex 'run' ./$(file).out 
