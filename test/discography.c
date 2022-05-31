#include <stdio.h>
#include <stdlib.h>
#include "audiodb.h"

int main() {
  audiodb* res = audiodb_discography("Bruno mars");
  if (res->error != NULL) {
    printf("Something went wrong: %s\n", res->error);
    free(res->error);
    return 1;
  }
  for (int i = 0; i < res->buffer_length; i++) {
    printf("%s\n",
           ((audiodb_discography_t**)res->buffer)[i]->int_year_released);
    printf("%s\n", ((audiodb_discography_t**)res->buffer)[i]->str_album);
  }
  audiodb_discography_clean((audiodb_discography_t**)res->buffer,
                            res->buffer_length);
  return 0;
}