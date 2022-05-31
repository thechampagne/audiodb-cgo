#include <stdio.h>
#include <stdlib.h>
#include "audiodb.h"

int main() {
  audiodb* res = audiodb_search_artist_by_id(111674);
  if (res->error != NULL) {
    printf("Something went wrong: %s\n", res->error);
    free(res->error);
    return 1;
  }
  printf("%s\n", ((audiodb_artist*)res->buffer)->id_artist);
  printf("%s\n", ((audiodb_artist*)res->buffer)->id_label);
  printf("%s\n", ((audiodb_artist*)res->buffer)->int_bornyear);
  printf("%s\n", ((audiodb_artist*)res->buffer)->int_charted);
  printf("%s\n", ((audiodb_artist*)res->buffer)->int_diedyear);
  printf("%s\n", ((audiodb_artist*)res->buffer)->int_formedyear);
  printf("%s\n", ((audiodb_artist*)res->buffer)->int_members);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artist);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artist_alternate);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artistbanner);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artistclearart);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artistcutout);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artistfanart);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artistfanart2);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artistfanart3);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artistfanart4);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artistlogo);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artist_stripped);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artist_thumb);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_artist_wide_thumb);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographycn);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyde);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyen);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyes);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyfr);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyhu);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyil);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyit);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyjp);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographynl);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyno);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographypl);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographypt);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyru);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_biographyse);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_country);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_countrycode);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_disbanded);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_facebook);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_gender);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_genre);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_isnicode);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_label);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_lastfmchart);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_locked);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_mood);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_music_brainz_id);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_style);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_twitter);
  printf("%s\n", ((audiodb_artist*)res->buffer)->str_website);

  audiodb_artist_clean((audiodb_artist*)res->buffer);
  return 0;
}