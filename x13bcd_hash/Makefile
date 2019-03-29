CC=gcc
FLAGS=-std=gnu99 -I. -I./sha3
CFLAGS=-fPIC -pedantic -Wall -Wextra -march=native
LDFLAGS=-shared

TARGET=libx13bcd.so
SOURCES = x13.c $(shell echo sha3/*.c)
HEADERS = x13.h $(shell echo sha3/*.h)
OBJECTS = $(SOURCES:.c=.o)

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $(LDFLAGS) -o $(TARGET) $(OBJECTS)

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

clean:
	rm *.o sha3/*.o *.so
