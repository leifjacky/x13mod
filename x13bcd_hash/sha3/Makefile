
all: blake.o bmw.o groestl.o jh.o keccak.o skein.o cubehash.o echo.o luffa.o simd.o hamsi.o hamsi_helper.o fugue.o shavite.o sm3.o

%.o : %.c
	$(CC) -c $*.c -I .

clean:
	rm *.o
