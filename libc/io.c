#include "io.h"
#include "defines.h"

u8 inportb(u16 port) {
  u8 ret;
  asmv("inb %1, %0" : "=a"(ret) : "dN"(port));
  return ret;
}

void outportb(u16 port, u8 data) {
  asmv("outb %1, %0" ::"dN"(port), "a"(data));
}
