#include <stdio.h>
#include <stdlib.h>
#include "audiodb.h"

int main() {
  audiodb* res = audiodb_search_music_videos_by_artist_id(111674);
  if (res->error != NULL) {
    printf("Something went wrong: %s\n", res->error);
    free(res->error);
    return 1;
  }
  for (int i = 0; i < res->buffer_length; i++) {
    printf("%s\n", ((audiodb_music_video**)res->buffer)[i]->id_album);
printf("%s\n", ((audiodb_music_video**)res->buffer)[i]->id_artist);
printf("%s\n", ((audiodb_music_video**)res->buffer)[i]->id_track);
printf("%s\n", ((audiodb_music_video**)res->buffer)[i]->str_descriptionen);
printf("%s\n", ((audiodb_music_video**)res->buffer)[i]->str_musicvid);
printf("%s\n", ((audiodb_music_video**)res->buffer)[i]->str_track);
printf("%s\n", ((audiodb_music_video**)res->buffer)[i]->str_trackthumb);
  }

  audiodb_music_video_clean((music_video**)res->buffer, res->buffer_length);
  return 0;
}