#include <stdio.h>
#include <stdlib.h>
#include "audiodb.h"

int main() {
  audiodb* res = audiodb_search_track_by_id(32771225);
  if (res->error != NULL) {
    printf("Something went wrong: %s\n", res->error);
    free(res->error);
    return 1;
  }
  printf("%s\n", ((audiodb_track*)res->buffer)->id_album);
printf("%s\n", ((audiodb_track*)res->buffer)->id_artist);
printf("%s\n", ((audiodb_track*)res->buffer)->id_imvdb);
printf("%s\n", ((audiodb_track*)res->buffer)->id_lyric);
printf("%s\n", ((audiodb_track*)res->buffer)->id_track);
printf("%s\n", ((audiodb_track*)res->buffer)->int_cd);
printf("%s\n", ((audiodb_track*)res->buffer)->int_duration);
printf("%s\n", ((audiodb_track*)res->buffer)->int_loved);
printf("%s\n", ((audiodb_track*)res->buffer)->int_music_vid_comments);
printf("%s\n", ((audiodb_track*)res->buffer)->int_music_vid_dislikes);
printf("%s\n", ((audiodb_track*)res->buffer)->int_music_vid_favorites);
printf("%s\n", ((audiodb_track*)res->buffer)->int_music_vid_likes);
printf("%s\n", ((audiodb_track*)res->buffer)->int_music_vid_views);
printf("%s\n", ((audiodb_track*)res->buffer)->int_score);
printf("%s\n", ((audiodb_track*)res->buffer)->int_scorevotes);
printf("%s\n", ((audiodb_track*)res->buffer)->int_total_listeners);
printf("%s\n", ((audiodb_track*)res->buffer)->int_totalplays);
printf("%s\n", ((audiodb_track*)res->buffer)->int_tracknumber);
printf("%s\n", ((audiodb_track*)res->buffer)->str_album);
printf("%s\n", ((audiodb_track*)res->buffer)->str_artist);
printf("%s\n", ((audiodb_track*)res->buffer)->str_artist_alternate);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptioncn);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionde);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionen);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptiones);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionfr);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionhu);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionil);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionit);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionjp);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionnl);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionno);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionpl);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionpt);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionru);
printf("%s\n", ((audiodb_track*)res->buffer)->str_descriptionse);
printf("%s\n", ((audiodb_track*)res->buffer)->str_genre);
printf("%s\n", ((audiodb_track*)res->buffer)->str_locked);
printf("%s\n", ((audiodb_track*)res->buffer)->str_mood);
printf("%s\n", ((audiodb_track*)res->buffer)->str_music_brainz_album_id);
printf("%s\n", ((audiodb_track*)res->buffer)->str_music_brainz_artist_id);
printf("%s\n", ((audiodb_track*)res->buffer)->str_music_brainz_id);
printf("%s\n", ((audiodb_track*)res->buffer)->str_musicvid);
printf("%s\n", ((audiodb_track*)res->buffer)->str_music_vid_company);
printf("%s\n", ((audiodb_track*)res->buffer)->str_music_vid_director);
printf("%s\n", ((audiodb_track*)res->buffer)->str_music_vid_screen1);
printf("%s\n", ((audiodb_track*)res->buffer)->str_music_vid_screen2);
printf("%s\n", ((audiodb_track*)res->buffer)->str_music_vid_screen3);
printf("%s\n", ((audiodb_track*)res->buffer)->str_style);
printf("%s\n", ((audiodb_track*)res->buffer)->str_theme);
printf("%s\n", ((audiodb_track*)res->buffer)->str_track);
printf("%s\n", ((audiodb_track*)res->buffer)->str_track3dcase);
printf("%s\n", ((audiodb_track*)res->buffer)->str_tracklyrics);
printf("%s\n", ((audiodb_track*)res->buffer)->str_trackthumb);

  audiodb_track_clean((audiodb_track*)res->buffer);
  return 0;
}