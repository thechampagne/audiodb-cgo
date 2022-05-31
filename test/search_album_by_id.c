#include <stdio.h>
#include <stdlib.h>
#include "audiodb.h"

int main() {
  audiodb* res = audiodb_search_album_by_id(2114062);
  if (res->error != NULL) {
    printf("Something went wrong: %s\n", res->error);
    free(res->error);
    return 1;
  }
printf("%s\n", ((audiodb_album*)res->buffer)->id_album);
printf("%s\n", ((audiodb_album*)res->buffer)->id_artist);
printf("%s\n", ((audiodb_album*)res->buffer)->id_label);
printf("%s\n", ((audiodb_album*)res->buffer)->int_loved);
printf("%s\n", ((audiodb_album*)res->buffer)->int_sales);
printf("%s\n", ((audiodb_album*)res->buffer)->int_score);
printf("%s\n", ((audiodb_album*)res->buffer)->int_score_votes);
printf("%s\n", ((audiodb_album*)res->buffer)->int_year_released);
printf("%s\n", ((audiodb_album*)res->buffer)->str_album);
printf("%s\n", ((audiodb_album*)res->buffer)->str_album3dcase);
printf("%s\n", ((audiodb_album*)res->buffer)->str_album3dface);
printf("%s\n", ((audiodb_album*)res->buffer)->str_album3dflat);
printf("%s\n", ((audiodb_album*)res->buffer)->str_album3dthumb);
printf("%s\n", ((audiodb_album*)res->buffer)->str_albumcdart);
printf("%s\n", ((audiodb_album*)res->buffer)->str_albumspine);
printf("%s\n", ((audiodb_album*)res->buffer)->str_albumstripped);
printf("%s\n", ((audiodb_album*)res->buffer)->str_albumthumb);
printf("%s\n", ((audiodb_album*)res->buffer)->str_albumthumbback);
printf("%s\n", ((audiodb_album*)res->buffer)->str_albumthumbhq);
printf("%s\n", ((audiodb_album*)res->buffer)->str_allmusicid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_amazonid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_artist);
printf("%s\n", ((audiodb_album*)res->buffer)->str_artist_stripped);
printf("%s\n", ((audiodb_album*)res->buffer)->str_bbcreviewid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_description);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptioncn);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionde);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionen);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptiones);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionfr);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionhu);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionil);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionit);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionjp);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionnl);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionno);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionpl);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionpt);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionru);
printf("%s\n", ((audiodb_album*)res->buffer)->str_descriptionse);
printf("%s\n", ((audiodb_album*)res->buffer)->str_discogsid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_geniusid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_genre);
printf("%s\n", ((audiodb_album*)res->buffer)->str_itunesid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_label);
printf("%s\n", ((audiodb_album*)res->buffer)->str_location);
printf("%s\n", ((audiodb_album*)res->buffer)->str_locked);
printf("%s\n", ((audiodb_album*)res->buffer)->str_lyricwikiid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_mood);
printf("%s\n", ((audiodb_album*)res->buffer)->str_music_brainz_artist_id);
printf("%s\n", ((audiodb_album*)res->buffer)->str_music_brainz_id);
printf("%s\n", ((audiodb_album*)res->buffer)->str_musicmozid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_rate_your_music_id);
printf("%s\n", ((audiodb_album*)res->buffer)->str_release_format);
printf("%s\n", ((audiodb_album*)res->buffer)->str_review);
printf("%s\n", ((audiodb_album*)res->buffer)->str_speed);
printf("%s\n", ((audiodb_album*)res->buffer)->str_style);
printf("%s\n", ((audiodb_album*)res->buffer)->str_theme);
printf("%s\n", ((audiodb_album*)res->buffer)->str_wikidataid);
printf("%s\n", ((audiodb_album*)res->buffer)->str_wikipediaid);

  audiodb_album_clean((audiodb_album*)res->buffer);
  return 0;
}