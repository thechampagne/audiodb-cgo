# audiodb-cgo

[![](https://img.shields.io/github/v/tag/thechampagne/audiodb-cgo?label=version)](https://github.com/thechampagne/audiodb-cgo/releases/latest) [![](https://img.shields.io/github/license/thechampagne/audiodb-cgo)](https://github.com/thechampagne/audiodb-cgo/blob/main/LICENSE)
### Installation & Setup

#### 1. Clone the repository
```
git clone https://github.com/thechampagne/audiodb-cgo.git
```
#### 2. Navigate to the root
```
cd audiodb-cgo
```
#### 3. Build the project
```
make
```
#### 4. Run test
```
make test
```

### Example

#### audiodb_discography
```c
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
```

#### audiodb_search_album_by_id
```c
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
```

#### audiodb_search_albums_by_artist_id
```c
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
```

#### audiodb_search_artist
```c
#include <stdio.h>
#include <stdlib.h>
#include "audiodb.h"

int main() {
  audiodb* res = audiodb_search_artist("Bruno mars");
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
```

#### audiodb_search_artist_by_id
```c
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
```

#### audiodb_search_music_videos_by_artist_id
```c
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

  audiodb_music_video_clean((audiodb_music_video**)res->buffer, res->buffer_length);
  return 0;
}
```

#### audiodb_search_track_by_id
```c
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
```

#### audiodb_search_tracks_by_album_id
```c
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
```

### References
 - [audiodb-go](https://github.com/thechampagne/audiodb-go)

### License

This repo is released under the [Apache License 2.0](https://github.com/thechampagne/audiodb-cgo/blob/main/LICENSE).
```
 Copyright 2022 XXIV

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
```