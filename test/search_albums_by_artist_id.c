#include <stdio.h>
#include <stdlib.h>
#include "audiodb.h"

int main() {
  audiodb* res = audiodb_search_albums_by_artist_id(111674);
  if (res->error != NULL) {
    printf("Something went wrong: %s\n", res->error);
    free(res->error);
    return 1;
  }
  for (int i = 0; i < res->buffer_length; i++) {
printf("%s\n", ((audiodb_album**)res->buffer)[i]->id_album);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->id_artist);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->id_label);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->int_loved);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->int_sales);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->int_score);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->int_score_votes);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->int_year_released);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_album);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_album3dcase);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_album3dface);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_album3dflat);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_album3dthumb);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_albumcdart);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_albumspine);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_albumstripped);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_albumthumb);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_albumthumbback);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_albumthumbhq);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_allmusicid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_amazonid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_artist);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_artist_stripped);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_bbcreviewid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_description);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptioncn);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionde);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionen);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptiones);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionfr);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionhu);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionil);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionit);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionjp);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionnl);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionno);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionpl);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionpt);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionru);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_descriptionse);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_discogsid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_geniusid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_genre);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_itunesid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_label);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_location);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_locked);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_lyricwikiid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_mood);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_music_brainz_artist_id);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_music_brainz_id);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_musicmozid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_rate_your_music_id);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_release_format);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_review);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_speed);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_style);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_theme);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_wikidataid);
printf("%s\n", ((audiodb_album**)res->buffer)[i]->str_wikipediaid);
}

  audiodb_albums_clean((audiodb_album**)res->buffer, res->buffer_length);
  return 0;
}