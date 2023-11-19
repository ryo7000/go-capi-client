#include "libclient.h"
#include <stdio.h>
#include <stdlib.h>

int main(void) {
  printf("rand: %llu\n", genrand());
  printf("error: %d\n", geterrors());
  printf("error msg: %s\n", (char *)errtostr(geterrors()));
  printf("\n");

  const auto res =
      get((char *)"35.68141046761117", (char *)"139.76716771217266");
  printf("%s\n", res);
  gofree(res);

  panic_test();
}
