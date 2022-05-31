#include <stdio.h>
#include <stdlib.h>
#include "audiodb.h"

int main() {
  audiodb* res = audiodb_search_tracks_by_album_id(2114062);
  if (res->error != NULL) {
    printf("Something went wrong: %s\n", res->error);
    free(res->error);
    return 1;
  }
  for (int i = 0; i < res->buffer_length; i++) {
  printf("%s\n", ((audiodb_track**)res->buffer)[i]->id_album);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->id_artist);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->id_imvdb);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->id_lyric);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->id_track);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_cd);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_duration);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_loved);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_music_vid_comments);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_music_vid_dislikes);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_music_vid_favorites);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_music_vid_likes);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_music_vid_views);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_score);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_scorevotes);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_total_listeners);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_totalplays);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->int_tracknumber);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_album);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_artist);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_artist_alternate);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptioncn);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionde);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionen);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptiones);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionfr);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionhu);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionil);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionit);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionjp);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionnl);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionno);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionpl);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionpt);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionru);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_descriptionse);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_genre);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_locked);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_mood);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_music_brainz_album_id);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_music_brainz_artist_id);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_music_brainz_id);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_musicvid);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_music_vid_company);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_music_vid_director);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_music_vid_screen1);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_music_vid_screen2);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_music_vid_screen3);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_style);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_theme);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_track);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_track3dcase);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_tracklyrics);
printf("%s\n", ((audiodb_track**)res->buffer)[i]->str_trackthumb);
}

  audiodb_tracks_clean((audiodb_track**)res->buffer, res->buffer_length);
  return 0;
}