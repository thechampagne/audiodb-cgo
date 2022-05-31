package main

/*
#include <stdlib.h>

typedef struct {
  size_t buffer_length;
  char* error;
  void* buffer;
} audiodb;

typedef struct {
  char* id_artist;
  char* id_label;
  char* int_bornyear;
  char* int_charted;
  char* int_diedyear;
  char* int_formedyear;
  char* int_members;
  char* str_artist;
  char* str_artist_alternate;
  char* str_artistbanner;
  char* str_artistclearart;
  char* str_artistcutout;
  char* str_artistfanart;
  char* str_artistfanart2;
  char* str_artistfanart3;
  char* str_artistfanart4;
  char* str_artistlogo;
  char* str_artist_stripped;
  char* str_artist_thumb;
  char* str_artist_wide_thumb;
  char* str_biographycn;
  char* str_biographyde;
  char* str_biographyen;
  char* str_biographyes;
  char* str_biographyfr;
  char* str_biographyhu;
  char* str_biographyil;
  char* str_biographyit;
  char* str_biographyjp;
  char* str_biographynl;
  char* str_biographyno;
  char* str_biographypl;
  char* str_biographypt;
  char* str_biographyru;
  char* str_biographyse;
  char* str_country;
  char* str_countrycode;
  char* str_disbanded;
  char* str_facebook;
  char* str_gender;
  char* str_genre;
  char* str_isnicode;
  char* str_label;
  char* str_lastfmchart;
  char* str_locked;
  char* str_mood;
  char* str_music_brainz_id;
  char* str_style;
  char* str_twitter;
  char* str_website;
} audiodb_artist;

typedef struct {
  char* int_year_released;
  char* str_album;
} audiodb_discography_t;

typedef struct {
  char* id_album;
  char* id_artist;
  char* id_label;
  char* int_loved;
  char* int_sales;
  char* int_score;
  char* int_score_votes;
  char* int_year_released;
  char* str_album;
  char* str_album3dcase;
  char* str_album3dface;
  char* str_album3dflat;
  char* str_album3dthumb;
  char* str_albumcdart;
  char* str_albumspine;
  char* str_albumstripped;
  char* str_albumthumb;
  char* str_albumthumbback;
  char* str_albumthumbhq;
  char* str_allmusicid;
  char* str_amazonid;
  char* str_artist;
  char* str_artist_stripped;
  char* str_bbcreviewid;
  char* str_description;
  char* str_descriptioncn;
  char* str_descriptionde;
  char* str_descriptionen;
  char* str_descriptiones;
  char* str_descriptionfr;
  char* str_descriptionhu;
  char* str_descriptionil;
  char* str_descriptionit;
  char* str_descriptionjp;
  char* str_descriptionnl;
  char* str_descriptionno;
  char* str_descriptionpl;
  char* str_descriptionpt;
  char* str_descriptionru;
  char* str_descriptionse;
  char* str_discogsid;
  char* str_geniusid;
  char* str_genre;
  char* str_itunesid;
  char* str_label;
  char* str_location;
  char* str_locked;
  char* str_lyricwikiid;
  char* str_mood;
  char* str_music_brainz_artist_id;
  char* str_music_brainz_id;
  char* str_musicmozid;
  char* str_rate_your_music_id;
  char* str_release_format;
  char* str_review;
  char* str_speed;
  char* str_style;
  char* str_theme;
  char* str_wikidataid;
  char* str_wikipediaid;
} audiodb_album;

typedef struct {
  char* id_album;
  char* id_artist;
  char* id_imvdb;
  char* id_lyric;
  char* id_track;
  char* int_cd;
  char* int_duration;
  char* int_loved;
  char* int_music_vid_comments;
  char* int_music_vid_dislikes;
  char* int_music_vid_favorites;
  char* int_music_vid_likes;
  char* int_music_vid_views;
  char* int_score;
  char* int_scorevotes;
  char* int_total_listeners;
  char* int_totalplays;
  char* int_tracknumber;
  char* str_album;
  char* str_artist;
  char* str_artist_alternate;
  char* str_descriptioncn;
  char* str_descriptionde;
  char* str_descriptionen;
  char* str_descriptiones;
  char* str_descriptionfr;
  char* str_descriptionhu;
  char* str_descriptionil;
  char* str_descriptionit;
  char* str_descriptionjp;
  char* str_descriptionnl;
  char* str_descriptionno;
  char* str_descriptionpl;
  char* str_descriptionpt;
  char* str_descriptionru;
  char* str_descriptionse;
  char* str_genre;
  char* str_locked;
  char* str_mood;
  char* str_music_brainz_album_id;
  char* str_music_brainz_artist_id;
  char* str_music_brainz_id;
  char* str_musicvid;
  char* str_music_vid_company;
  char* str_music_vid_director;
  char* str_music_vid_screen1;
  char* str_music_vid_screen2;
  char* str_music_vid_screen3;
  char* str_style;
  char* str_theme;
  char* str_track;
  char* str_track3dcase;
  char* str_tracklyrics;
  char* str_trackthumb;
} audiodb_track;

typedef struct {
  char* id_album;
  char* id_artist;
  char* id_track;
  char* str_descriptionen;
  char* str_musicvid;
  char* str_track;
  char* str_trackthumb;
} audiodb_music_video;
*/
import "C"
import (
	"unsafe"
	"github.com/thechampagne/audiodb-go/audiodb"
)

//export audiodb_search_artist
func audiodb_search_artist(s *C.char) *C.audiodb {
  self := (*C.audiodb) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb{}))))
  res, err := audiodb.SearchArtist(C.GoString(s))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
  data := (*C.audiodb_artist) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb_artist{}))))
  data.id_artist = C.CString(res.IDArtist)
data.id_label = C.CString(res.IDLabel)
data.int_bornyear = C.CString(res.IntBornYear)
data.int_charted = C.CString(res.IntCharted)
data.int_diedyear = C.CString(res.IntDiedYear)
data.int_formedyear = C.CString(res.IntFormedYear)
data.int_members = C.CString(res.IntMembers)
data.str_artist = C.CString(res.StrArtist)
data.str_artist_alternate = C.CString(res.StrArtistAlternate)
data.str_artistbanner = C.CString(res.StrArtistBanner)
data.str_artistclearart = C.CString(res.StrArtistClearart)
data.str_artistcutout = C.CString(res.StrArtistCutout)
data.str_artistfanart = C.CString(res.StrArtistFanart)
data.str_artistfanart2 = C.CString(res.StrArtistFanart2)
data.str_artistfanart3 = C.CString(res.StrArtistFanart3)
data.str_artistfanart4 = C.CString(res.StrArtistFanart4)
data.str_artistlogo = C.CString(res.StrArtistLogo)
data.str_artist_stripped = C.CString(res.StrArtistStripped)
data.str_artist_thumb = C.CString(res.StrArtistThumb)
data.str_artist_wide_thumb = C.CString(res.StrArtistWideThumb)
data.str_biographycn = C.CString(res.StrBiographyCN)
data.str_biographyde = C.CString(res.StrBiographyDE)
data.str_biographyen = C.CString(res.StrBiographyEN)
data.str_biographyes = C.CString(res.StrBiographyES)
data.str_biographyfr = C.CString(res.StrBiographyFR)
data.str_biographyhu = C.CString(res.StrBiographyHU)
data.str_biographyil = C.CString(res.StrBiographyIL)
data.str_biographyit = C.CString(res.StrBiographyIT)
data.str_biographyjp = C.CString(res.StrBiographyJP)
data.str_biographynl = C.CString(res.StrBiographyNL)
data.str_biographyno = C.CString(res.StrBiographyNO)
data.str_biographypl = C.CString(res.StrBiographyPL)
data.str_biographypt = C.CString(res.StrBiographyPT)
data.str_biographyru = C.CString(res.StrBiographyRU)
data.str_biographyse = C.CString(res.StrBiographySE)
data.str_country = C.CString(res.StrCountry)
data.str_countrycode = C.CString(res.StrCountryCode)
data.str_disbanded = C.CString(res.StrDisbanded)
data.str_facebook = C.CString(res.StrFacebook)
data.str_gender = C.CString(res.StrGender)
data.str_genre = C.CString(res.StrGenre)
data.str_isnicode = C.CString(res.StrISNIcode)
data.str_label = C.CString(res.StrLabel)
data.str_lastfmchart = C.CString(res.StrLastFMChart)
data.str_locked = C.CString(res.StrLocked)
data.str_mood = C.CString(res.StrMood)
data.str_music_brainz_id = C.CString(res.StrMusicBrainzID)
data.str_style = C.CString(res.StrStyle)
data.str_twitter = C.CString(res.StrTwitter)
data.str_website = C.CString(res.StrWebsite)

self.buffer = unsafe.Pointer(data)
self.error = nil
return self
}

//export audiodb_artist_clean
func audiodb_artist_clean(self *C.audiodb_artist) {
  if self != nil {
 if self.id_artist != nil { C.free(unsafe.Pointer(self.id_artist)) }
if self.id_label != nil { C.free(unsafe.Pointer(self.id_label)) }
if self.int_bornyear != nil { C.free(unsafe.Pointer(self.int_bornyear)) }
if self.int_charted != nil { C.free(unsafe.Pointer(self.int_charted)) }
if self.int_diedyear != nil { C.free(unsafe.Pointer(self.int_diedyear)) }
if self.int_formedyear != nil { C.free(unsafe.Pointer(self.int_formedyear)) }
if self.int_members != nil { C.free(unsafe.Pointer(self.int_members)) }
if self.str_artist != nil { C.free(unsafe.Pointer(self.str_artist)) }
if self.str_artist_alternate != nil { C.free(unsafe.Pointer(self.str_artist_alternate)) }
if self.str_artistbanner != nil { C.free(unsafe.Pointer(self.str_artistbanner)) }
if self.str_artistclearart != nil { C.free(unsafe.Pointer(self.str_artistclearart)) }
if self.str_artistcutout != nil { C.free(unsafe.Pointer(self.str_artistcutout)) }
if self.str_artistfanart != nil { C.free(unsafe.Pointer(self.str_artistfanart)) }
if self.str_artistfanart2 != nil { C.free(unsafe.Pointer(self.str_artistfanart2)) }
if self.str_artistfanart3 != nil { C.free(unsafe.Pointer(self.str_artistfanart3)) }
if self.str_artistfanart4 != nil { C.free(unsafe.Pointer(self.str_artistfanart4)) }
if self.str_artistlogo != nil { C.free(unsafe.Pointer(self.str_artistlogo)) }
if self.str_artist_stripped != nil { C.free(unsafe.Pointer(self.str_artist_stripped)) }
if self.str_artist_thumb != nil { C.free(unsafe.Pointer(self.str_artist_thumb)) }
if self.str_artist_wide_thumb != nil { C.free(unsafe.Pointer(self.str_artist_wide_thumb)) }
if self.str_biographycn != nil { C.free(unsafe.Pointer(self.str_biographycn)) }
if self.str_biographyde != nil { C.free(unsafe.Pointer(self.str_biographyde)) }
if self.str_biographyen != nil { C.free(unsafe.Pointer(self.str_biographyen)) }
if self.str_biographyes != nil { C.free(unsafe.Pointer(self.str_biographyes)) }
if self.str_biographyfr != nil { C.free(unsafe.Pointer(self.str_biographyfr)) }
if self.str_biographyhu != nil { C.free(unsafe.Pointer(self.str_biographyhu)) }
if self.str_biographyil != nil { C.free(unsafe.Pointer(self.str_biographyil)) }
if self.str_biographyit != nil { C.free(unsafe.Pointer(self.str_biographyit)) }
if self.str_biographyjp != nil { C.free(unsafe.Pointer(self.str_biographyjp)) }
if self.str_biographynl != nil { C.free(unsafe.Pointer(self.str_biographynl)) }
if self.str_biographyno != nil { C.free(unsafe.Pointer(self.str_biographyno)) }
if self.str_biographypl != nil { C.free(unsafe.Pointer(self.str_biographypl)) }
if self.str_biographypt != nil { C.free(unsafe.Pointer(self.str_biographypt)) }
if self.str_biographyru != nil { C.free(unsafe.Pointer(self.str_biographyru)) }
if self.str_biographyse != nil { C.free(unsafe.Pointer(self.str_biographyse)) }
if self.str_country != nil { C.free(unsafe.Pointer(self.str_country)) }
if self.str_countrycode != nil { C.free(unsafe.Pointer(self.str_countrycode)) }
if self.str_disbanded != nil { C.free(unsafe.Pointer(self.str_disbanded)) }
if self.str_facebook != nil { C.free(unsafe.Pointer(self.str_facebook)) }
if self.str_gender != nil { C.free(unsafe.Pointer(self.str_gender)) }
if self.str_genre != nil { C.free(unsafe.Pointer(self.str_genre)) }
if self.str_isnicode != nil { C.free(unsafe.Pointer(self.str_isnicode)) }
if self.str_label != nil { C.free(unsafe.Pointer(self.str_label)) }
if self.str_lastfmchart != nil { C.free(unsafe.Pointer(self.str_lastfmchart)) }
if self.str_locked != nil { C.free(unsafe.Pointer(self.str_locked)) }
if self.str_mood != nil { C.free(unsafe.Pointer(self.str_mood)) }
if self.str_music_brainz_id != nil { C.free(unsafe.Pointer(self.str_music_brainz_id)) }
if self.str_style != nil { C.free(unsafe.Pointer(self.str_style)) }
if self.str_twitter != nil { C.free(unsafe.Pointer(self.str_twitter)) }
if self.str_website != nil { C.free(unsafe.Pointer(self.str_website)) }
    C.free(unsafe.Pointer(self))
  }
}

//export audiodb_discography
func audiodb_discography(s *C.char) *C.audiodb {
  self := (*C.audiodb) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb{}))))
  res, err := audiodb.Discography(C.GoString(s))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
  array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))
  
  slice := (*[1<<30 - 1]*C.audiodb_discography_t)(array)

  for i, val := range res {
        data := (*C.audiodb_discography_t) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb_discography_t{}))))
        data.int_year_released = C.CString(val.IntYearReleased)
        data.str_album = C.CString(val.StrAlbum)
        slice[i] = data
  }
  self.buffer = unsafe.Pointer((**C.audiodb_discography_t) (array))
  self.buffer_length = C.size_t(len(res))
  self.error = nil
  return self
  
}

//export audiodb_discography_clean
func audiodb_discography_clean(self **C.audiodb_discography_t, length C.int) {
  if self != nil {
    slice := (*[1<<30 - 1]*C.audiodb_discography_t)(unsafe.Pointer(self))[:length:length]
    for i := 0 ; i < len(slice); i++ {
      if slice[i] != nil {
        if slice[i].int_year_released != nil { C.free(unsafe.Pointer(slice[i].int_year_released)) }
        if slice[i].str_album != nil { C.free(unsafe.Pointer(slice[i].str_album)) }
        C.free(unsafe.Pointer(slice[i]))
      }
    }
    C.free(unsafe.Pointer(self))
  }
}

//export audiodb_search_artist_by_id
func audiodb_search_artist_by_id(i C.int) *C.audiodb {
  self := (*C.audiodb) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb{}))))
  res, err := audiodb.SearchArtistById(int64(i))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
  data := (*C.audiodb_artist) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb_artist{}))))
  data.id_artist = C.CString(res.IDArtist)
data.id_label = C.CString(res.IDLabel)
data.int_bornyear = C.CString(res.IntBornYear)
data.int_charted = C.CString(res.IntCharted)
data.int_diedyear = C.CString(res.IntDiedYear)
data.int_formedyear = C.CString(res.IntFormedYear)
data.int_members = C.CString(res.IntMembers)
data.str_artist = C.CString(res.StrArtist)
data.str_artist_alternate = C.CString(res.StrArtistAlternate)
data.str_artistbanner = C.CString(res.StrArtistBanner)
data.str_artistclearart = C.CString(res.StrArtistClearart)
data.str_artistcutout = C.CString(res.StrArtistCutout)
data.str_artistfanart = C.CString(res.StrArtistFanart)
data.str_artistfanart2 = C.CString(res.StrArtistFanart2)
data.str_artistfanart3 = C.CString(res.StrArtistFanart3)
data.str_artistfanart4 = C.CString(res.StrArtistFanart4)
data.str_artistlogo = C.CString(res.StrArtistLogo)
data.str_artist_stripped = C.CString(res.StrArtistStripped)
data.str_artist_thumb = C.CString(res.StrArtistThumb)
data.str_artist_wide_thumb = C.CString(res.StrArtistWideThumb)
data.str_biographycn = C.CString(res.StrBiographyCN)
data.str_biographyde = C.CString(res.StrBiographyDE)
data.str_biographyen = C.CString(res.StrBiographyEN)
data.str_biographyes = C.CString(res.StrBiographyES)
data.str_biographyfr = C.CString(res.StrBiographyFR)
data.str_biographyhu = C.CString(res.StrBiographyHU)
data.str_biographyil = C.CString(res.StrBiographyIL)
data.str_biographyit = C.CString(res.StrBiographyIT)
data.str_biographyjp = C.CString(res.StrBiographyJP)
data.str_biographynl = C.CString(res.StrBiographyNL)
data.str_biographyno = C.CString(res.StrBiographyNO)
data.str_biographypl = C.CString(res.StrBiographyPL)
data.str_biographypt = C.CString(res.StrBiographyPT)
data.str_biographyru = C.CString(res.StrBiographyRU)
data.str_biographyse = C.CString(res.StrBiographySE)
data.str_country = C.CString(res.StrCountry)
data.str_countrycode = C.CString(res.StrCountryCode)
data.str_disbanded = C.CString(res.StrDisbanded)
data.str_facebook = C.CString(res.StrFacebook)
data.str_gender = C.CString(res.StrGender)
data.str_genre = C.CString(res.StrGenre)
data.str_isnicode = C.CString(res.StrISNIcode)
data.str_label = C.CString(res.StrLabel)
data.str_lastfmchart = C.CString(res.StrLastFMChart)
data.str_locked = C.CString(res.StrLocked)
data.str_mood = C.CString(res.StrMood)
data.str_music_brainz_id = C.CString(res.StrMusicBrainzID)
data.str_style = C.CString(res.StrStyle)
data.str_twitter = C.CString(res.StrTwitter)
data.str_website = C.CString(res.StrWebsite)

self.buffer = unsafe.Pointer(data)
self.error = nil
return self
}

//export audiodb_search_album_by_id
func audiodb_search_album_by_id(i C.int) *C.audiodb {
  self := (*C.audiodb) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb{}))))
  res, err := audiodb.SearchAlbumById(int64(i))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
  data := (*C.audiodb_album) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb_album{}))))

  data.id_album = C.CString(res.IDAlbum)
data.id_artist = C.CString(res.IDArtist)
data.id_label = C.CString(res.IDLabel)
data.int_loved = C.CString(res.IntLoved)
data.int_sales = C.CString(res.IntSales)
data.int_score = C.CString(res.IntScore)
data.int_score_votes = C.CString(res.IntScoreVotes)
data.int_year_released = C.CString(res.IntYearReleased)
data.str_album = C.CString(res.StrAlbum)
data.str_album3dcase = C.CString(res.StrAlbum3DCase)
data.str_album3dface = C.CString(res.StrAlbum3DFace)
data.str_album3dflat = C.CString(res.StrAlbum3DFlat)
data.str_album3dthumb = C.CString(res.StrAlbum3DThumb)
data.str_albumcdart = C.CString(res.StrAlbumCDart)
data.str_albumspine = C.CString(res.StrAlbumSpine)
data.str_albumstripped = C.CString(res.StrAlbumStripped)
data.str_albumthumb = C.CString(res.StrAlbumThumb)
data.str_albumthumbback = C.CString(res.StrAlbumThumbBack)
data.str_albumthumbhq = C.CString(res.StrAlbumThumbHQ)
data.str_allmusicid = C.CString(res.StrAllMusicID)
data.str_amazonid = C.CString(res.StrAmazonID)
data.str_artist = C.CString(res.StrArtist)
data.str_artist_stripped = C.CString(res.StrArtistStripped)
data.str_bbcreviewid = C.CString(res.StrBBCReviewID)
data.str_description = C.CString(res.StrDescription)
data.str_descriptioncn = C.CString(res.StrDescriptionCN)
data.str_descriptionde = C.CString(res.StrDescriptionDE)
data.str_descriptionen = C.CString(res.StrDescriptionEN)
data.str_descriptiones = C.CString(res.StrDescriptionES)
data.str_descriptionfr = C.CString(res.StrDescriptionFR)
data.str_descriptionhu = C.CString(res.StrDescriptionHU)
data.str_descriptionil = C.CString(res.StrDescriptionIL)
data.str_descriptionit = C.CString(res.StrDescriptionIT)
data.str_descriptionjp = C.CString(res.StrDescriptionJP)
data.str_descriptionnl = C.CString(res.StrDescriptionNL)
data.str_descriptionno = C.CString(res.StrDescriptionNO)
data.str_descriptionpl = C.CString(res.StrDescriptionPL)
data.str_descriptionpt = C.CString(res.StrDescriptionPT)
data.str_descriptionru = C.CString(res.StrDescriptionRU)
data.str_descriptionse = C.CString(res.StrDescriptionSE)
data.str_discogsid = C.CString(res.StrDiscogsID)
data.str_geniusid = C.CString(res.StrGeniusID)
data.str_genre = C.CString(res.StrGenre)
data.str_itunesid = C.CString(res.StrItunesID)
data.str_label = C.CString(res.StrLabel)
data.str_location = C.CString(res.StrLocation)
data.str_locked = C.CString(res.StrLocked)
data.str_lyricwikiid = C.CString(res.StrLyricWikiID)
data.str_mood = C.CString(res.StrMood)
data.str_music_brainz_artist_id = C.CString(res.StrMusicBrainzArtistID)
data.str_music_brainz_id = C.CString(res.StrMusicBrainzID)
data.str_musicmozid = C.CString(res.StrMusicMozID)
data.str_rate_your_music_id = C.CString(res.StrRateYourMusicID)
data.str_release_format = C.CString(res.StrReleaseFormat)
data.str_review = C.CString(res.StrReview)
data.str_speed = C.CString(res.StrSpeed)
data.str_style = C.CString(res.StrStyle)
data.str_theme = C.CString(res.StrTheme)
data.str_wikidataid = C.CString(res.StrWikidataID)
data.str_wikipediaid = C.CString(res.StrWikipediaID)

  self.buffer = unsafe.Pointer(data)
self.error = nil
return self
}

//export audiodb_album_clean
func audiodb_album_clean(self *C.audiodb_album) {
  if self != nil {
    if self.id_album != nil { C.free(unsafe.Pointer(self.id_album)) }
if self.id_artist != nil { C.free(unsafe.Pointer(self.id_artist)) }
if self.id_label != nil { C.free(unsafe.Pointer(self.id_label)) }
if self.int_loved != nil { C.free(unsafe.Pointer(self.int_loved)) }
if self.int_sales != nil { C.free(unsafe.Pointer(self.int_sales)) }
if self.int_score != nil { C.free(unsafe.Pointer(self.int_score)) }
if self.int_score_votes != nil { C.free(unsafe.Pointer(self.int_score_votes)) }
if self.int_year_released != nil { C.free(unsafe.Pointer(self.int_year_released)) }
if self.str_album != nil { C.free(unsafe.Pointer(self.str_album)) }
if self.str_album3dcase != nil { C.free(unsafe.Pointer(self.str_album3dcase)) }
if self.str_album3dface != nil { C.free(unsafe.Pointer(self.str_album3dface)) }
if self.str_album3dflat != nil { C.free(unsafe.Pointer(self.str_album3dflat)) }
if self.str_album3dthumb != nil { C.free(unsafe.Pointer(self.str_album3dthumb)) }
if self.str_albumcdart != nil { C.free(unsafe.Pointer(self.str_albumcdart)) }
if self.str_albumspine != nil { C.free(unsafe.Pointer(self.str_albumspine)) }
if self.str_albumstripped != nil { C.free(unsafe.Pointer(self.str_albumstripped)) }
if self.str_albumthumb != nil { C.free(unsafe.Pointer(self.str_albumthumb)) }
if self.str_albumthumbback != nil { C.free(unsafe.Pointer(self.str_albumthumbback)) }
if self.str_albumthumbhq != nil { C.free(unsafe.Pointer(self.str_albumthumbhq)) }
if self.str_allmusicid != nil { C.free(unsafe.Pointer(self.str_allmusicid)) }
if self.str_amazonid != nil { C.free(unsafe.Pointer(self.str_amazonid)) }
if self.str_artist != nil { C.free(unsafe.Pointer(self.str_artist)) }
if self.str_artist_stripped != nil { C.free(unsafe.Pointer(self.str_artist_stripped)) }
if self.str_bbcreviewid != nil { C.free(unsafe.Pointer(self.str_bbcreviewid)) }
if self.str_description != nil { C.free(unsafe.Pointer(self.str_description)) }
if self.str_descriptioncn != nil { C.free(unsafe.Pointer(self.str_descriptioncn)) }
if self.str_descriptionde != nil { C.free(unsafe.Pointer(self.str_descriptionde)) }
if self.str_descriptionen != nil { C.free(unsafe.Pointer(self.str_descriptionen)) }
if self.str_descriptiones != nil { C.free(unsafe.Pointer(self.str_descriptiones)) }
if self.str_descriptionfr != nil { C.free(unsafe.Pointer(self.str_descriptionfr)) }
if self.str_descriptionhu != nil { C.free(unsafe.Pointer(self.str_descriptionhu)) }
if self.str_descriptionil != nil { C.free(unsafe.Pointer(self.str_descriptionil)) }
if self.str_descriptionit != nil { C.free(unsafe.Pointer(self.str_descriptionit)) }
if self.str_descriptionjp != nil { C.free(unsafe.Pointer(self.str_descriptionjp)) }
if self.str_descriptionnl != nil { C.free(unsafe.Pointer(self.str_descriptionnl)) }
if self.str_descriptionno != nil { C.free(unsafe.Pointer(self.str_descriptionno)) }
if self.str_descriptionpl != nil { C.free(unsafe.Pointer(self.str_descriptionpl)) }
if self.str_descriptionpt != nil { C.free(unsafe.Pointer(self.str_descriptionpt)) }
if self.str_descriptionru != nil { C.free(unsafe.Pointer(self.str_descriptionru)) }
if self.str_descriptionse != nil { C.free(unsafe.Pointer(self.str_descriptionse)) }
if self.str_discogsid != nil { C.free(unsafe.Pointer(self.str_discogsid)) }
if self.str_geniusid != nil { C.free(unsafe.Pointer(self.str_geniusid)) }
if self.str_genre != nil { C.free(unsafe.Pointer(self.str_genre)) }
if self.str_itunesid != nil { C.free(unsafe.Pointer(self.str_itunesid)) }
if self.str_label != nil { C.free(unsafe.Pointer(self.str_label)) }
if self.str_location != nil { C.free(unsafe.Pointer(self.str_location)) }
if self.str_locked != nil { C.free(unsafe.Pointer(self.str_locked)) }
if self.str_lyricwikiid != nil { C.free(unsafe.Pointer(self.str_lyricwikiid)) }
if self.str_mood != nil { C.free(unsafe.Pointer(self.str_mood)) }
if self.str_music_brainz_artist_id != nil { C.free(unsafe.Pointer(self.str_music_brainz_artist_id)) }
if self.str_music_brainz_id != nil { C.free(unsafe.Pointer(self.str_music_brainz_id)) }
if self.str_musicmozid != nil { C.free(unsafe.Pointer(self.str_musicmozid)) }
if self.str_rate_your_music_id != nil { C.free(unsafe.Pointer(self.str_rate_your_music_id)) }
if self.str_release_format != nil { C.free(unsafe.Pointer(self.str_release_format)) }
if self.str_review != nil { C.free(unsafe.Pointer(self.str_review)) }
if self.str_speed != nil { C.free(unsafe.Pointer(self.str_speed)) }
if self.str_style != nil { C.free(unsafe.Pointer(self.str_style)) }
if self.str_theme != nil { C.free(unsafe.Pointer(self.str_theme)) }
if self.str_wikidataid != nil { C.free(unsafe.Pointer(self.str_wikidataid)) }
if self.str_wikipediaid != nil { C.free(unsafe.Pointer(self.str_wikipediaid)) }
     C.free(unsafe.Pointer(self))
  }
}

//export audiodb_search_albums_by_artist_id
func audiodb_search_albums_by_artist_id(i C.int) *C.audiodb {
  self := (*C.audiodb) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb{}))))
  res, err := audiodb.SearchAlbumsByArtistId(int64(i))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
  array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))
  
  slice := (*[1<<30 - 1]*C.audiodb_album)(array)

  for i, val := range res {
    data := (*C.audiodb_album) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb_album{}))))

  data.id_album = C.CString(val.IDAlbum)
data.id_artist = C.CString(val.IDArtist)
data.id_label = C.CString(val.IDLabel)
data.int_loved = C.CString(val.IntLoved)
data.int_sales = C.CString(val.IntSales)
data.int_score = C.CString(val.IntScore)
data.int_score_votes = C.CString(val.IntScoreVotes)
data.int_year_released = C.CString(val.IntYearReleased)
data.str_album = C.CString(val.StrAlbum)
data.str_album3dcase = C.CString(val.StrAlbum3DCase)
data.str_album3dface = C.CString(val.StrAlbum3DFace)
data.str_album3dflat = C.CString(val.StrAlbum3DFlat)
data.str_album3dthumb = C.CString(val.StrAlbum3DThumb)
data.str_albumcdart = C.CString(val.StrAlbumCDart)
data.str_albumspine = C.CString(val.StrAlbumSpine)
data.str_albumstripped = C.CString(val.StrAlbumStripped)
data.str_albumthumb = C.CString(val.StrAlbumThumb)
data.str_albumthumbback = C.CString(val.StrAlbumThumbBack)
data.str_albumthumbhq = C.CString(val.StrAlbumThumbHQ)
data.str_allmusicid = C.CString(val.StrAllMusicID)
data.str_amazonid = C.CString(val.StrAmazonID)
data.str_artist = C.CString(val.StrArtist)
data.str_artist_stripped = C.CString(val.StrArtistStripped)
data.str_bbcreviewid = C.CString(val.StrBBCReviewID)
data.str_description = C.CString(val.StrDescription)
data.str_descriptioncn = C.CString(val.StrDescriptionCN)
data.str_descriptionde = C.CString(val.StrDescriptionDE)
data.str_descriptionen = C.CString(val.StrDescriptionEN)
data.str_descriptiones = C.CString(val.StrDescriptionES)
data.str_descriptionfr = C.CString(val.StrDescriptionFR)
data.str_descriptionhu = C.CString(val.StrDescriptionHU)
data.str_descriptionil = C.CString(val.StrDescriptionIL)
data.str_descriptionit = C.CString(val.StrDescriptionIT)
data.str_descriptionjp = C.CString(val.StrDescriptionJP)
data.str_descriptionnl = C.CString(val.StrDescriptionNL)
data.str_descriptionno = C.CString(val.StrDescriptionNO)
data.str_descriptionpl = C.CString(val.StrDescriptionPL)
data.str_descriptionpt = C.CString(val.StrDescriptionPT)
data.str_descriptionru = C.CString(val.StrDescriptionRU)
data.str_descriptionse = C.CString(val.StrDescriptionSE)
data.str_discogsid = C.CString(val.StrDiscogsID)
data.str_geniusid = C.CString(val.StrGeniusID)
data.str_genre = C.CString(val.StrGenre)
data.str_itunesid = C.CString(val.StrItunesID)
data.str_label = C.CString(val.StrLabel)
data.str_location = C.CString(val.StrLocation)
data.str_locked = C.CString(val.StrLocked)
data.str_lyricwikiid = C.CString(val.StrLyricWikiID)
data.str_mood = C.CString(val.StrMood)
data.str_music_brainz_artist_id = C.CString(val.StrMusicBrainzArtistID)
data.str_music_brainz_id = C.CString(val.StrMusicBrainzID)
data.str_musicmozid = C.CString(val.StrMusicMozID)
data.str_rate_your_music_id = C.CString(val.StrRateYourMusicID)
data.str_release_format = C.CString(val.StrReleaseFormat)
data.str_review = C.CString(val.StrReview)
data.str_speed = C.CString(val.StrSpeed)
data.str_style = C.CString(val.StrStyle)
data.str_theme = C.CString(val.StrTheme)
data.str_wikidataid = C.CString(val.StrWikidataID)
data.str_wikipediaid = C.CString(val.StrWikipediaID)
    slice[i] = data
  }

  self.buffer = unsafe.Pointer((**C.audiodb_album) (array))
  self.buffer_length = C.size_t(len(res))
  self.error = nil
  return self
}

//export audiodb_albums_clean
func audiodb_albums_clean(self **C.audiodb_album, length C.int) {
  if self != nil {
    slice := (*[1<<30 - 1]*C.audiodb_album)(unsafe.Pointer(self))[:length:length]
    for i := 0 ; i < len(slice); i++ {
      if slice[i] != nil {
        audiodb_album_clean(slice[i])
      }
    }
    C.free(unsafe.Pointer(self))
  }
}

//export audiodb_search_track_by_id
func audiodb_search_track_by_id(i C.int) *C.audiodb {
  self := (*C.audiodb) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb{}))))
  res, err := audiodb.SearchTrackById(int64(i))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
  data := (*C.audiodb_track) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb_track{}))))

  data.id_album = C.CString(res.IDAlbum)
data.id_artist = C.CString(res.IDArtist)
data.id_imvdb = C.CString(res.IDIMVDB)
data.id_lyric = C.CString(res.IDLyric)
data.id_track = C.CString(res.IDTrack)
data.int_cd = C.CString(res.IntCD)
data.int_duration = C.CString(res.IntDuration)
data.int_loved = C.CString(res.IntLoved)
data.int_music_vid_comments = C.CString(res.IntMusicVidComments)
data.int_music_vid_dislikes = C.CString(res.IntMusicVidDislikes)
data.int_music_vid_favorites = C.CString(res.IntMusicVidFavorites)
data.int_music_vid_likes = C.CString(res.IntMusicVidLikes)
data.int_music_vid_views = C.CString(res.IntMusicVidViews)
data.int_score = C.CString(res.IntScore)
data.int_scorevotes = C.CString(res.IntScoreVotes)
data.int_total_listeners = C.CString(res.IntTotalListeners)
data.int_totalplays = C.CString(res.IntTotalPlays)
data.int_tracknumber = C.CString(res.IntTrackNumber)
data.str_album = C.CString(res.StrAlbum)
data.str_artist = C.CString(res.StrArtist)
data.str_artist_alternate = C.CString(res.StrArtistAlternate)
data.str_descriptioncn = C.CString(res.StrDescriptionCN)
data.str_descriptionde = C.CString(res.StrDescriptionDE)
data.str_descriptionen = C.CString(res.StrDescriptionEN)
data.str_descriptiones = C.CString(res.StrDescriptionES)
data.str_descriptionfr = C.CString(res.StrDescriptionFR)
data.str_descriptionhu = C.CString(res.StrDescriptionHU)
data.str_descriptionil = C.CString(res.StrDescriptionIL)
data.str_descriptionit = C.CString(res.StrDescriptionIT)
data.str_descriptionjp = C.CString(res.StrDescriptionJP)
data.str_descriptionnl = C.CString(res.StrDescriptionNL)
data.str_descriptionno = C.CString(res.StrDescriptionNO)
data.str_descriptionpl = C.CString(res.StrDescriptionPL)
data.str_descriptionpt = C.CString(res.StrDescriptionPT)
data.str_descriptionru = C.CString(res.StrDescriptionRU)
data.str_descriptionse = C.CString(res.StrDescriptionSE)
data.str_genre = C.CString(res.StrGenre)
data.str_locked = C.CString(res.StrLocked)
data.str_mood = C.CString(res.StrMood)
data.str_music_brainz_album_id = C.CString(res.StrMusicBrainzAlbumID)
data.str_music_brainz_artist_id = C.CString(res.StrMusicBrainzArtistID)
data.str_music_brainz_id = C.CString(res.StrMusicBrainzID)
data.str_musicvid = C.CString(res.StrMusicVid)
data.str_music_vid_company = C.CString(res.StrMusicVidCompany)
data.str_music_vid_director = C.CString(res.StrMusicVidDirector)
data.str_music_vid_screen1 = C.CString(res.StrMusicVidScreen1)
data.str_music_vid_screen2 = C.CString(res.StrMusicVidScreen2)
data.str_music_vid_screen3 = C.CString(res.StrMusicVidScreen3)
data.str_style = C.CString(res.StrStyle)
data.str_theme = C.CString(res.StrTheme)
data.str_track = C.CString(res.StrTrack)
data.str_track3dcase = C.CString(res.StrTrack3DCase)
data.str_tracklyrics = C.CString(res.StrTrackLyrics)
data.str_trackthumb = C.CString(res.StrTrackThumb)

  self.buffer = unsafe.Pointer(data)
self.error = nil
return self
}

//export audiodb_track_clean
func audiodb_track_clean(self *C.audiodb_track) {
  if self != nil {
    if self.id_album != nil { C.free(unsafe.Pointer(self.id_album)) }
if self.id_artist != nil { C.free(unsafe.Pointer(self.id_artist)) }
if self.id_imvdb != nil { C.free(unsafe.Pointer(self.id_imvdb)) }
if self.id_lyric != nil { C.free(unsafe.Pointer(self.id_lyric)) }
if self.id_track != nil { C.free(unsafe.Pointer(self.id_track)) }
if self.int_cd != nil { C.free(unsafe.Pointer(self.int_cd)) }
if self.int_duration != nil { C.free(unsafe.Pointer(self.int_duration)) }
if self.int_loved != nil { C.free(unsafe.Pointer(self.int_loved)) }
if self.int_music_vid_comments != nil { C.free(unsafe.Pointer(self.int_music_vid_comments)) }
if self.int_music_vid_dislikes != nil { C.free(unsafe.Pointer(self.int_music_vid_dislikes)) }
if self.int_music_vid_favorites != nil { C.free(unsafe.Pointer(self.int_music_vid_favorites)) }
if self.int_music_vid_likes != nil { C.free(unsafe.Pointer(self.int_music_vid_likes)) }
if self.int_music_vid_views != nil { C.free(unsafe.Pointer(self.int_music_vid_views)) }
if self.int_score != nil { C.free(unsafe.Pointer(self.int_score)) }
if self.int_scorevotes != nil { C.free(unsafe.Pointer(self.int_scorevotes)) }
if self.int_total_listeners != nil { C.free(unsafe.Pointer(self.int_total_listeners)) }
if self.int_totalplays != nil { C.free(unsafe.Pointer(self.int_totalplays)) }
if self.int_tracknumber != nil { C.free(unsafe.Pointer(self.int_tracknumber)) }
if self.str_album != nil { C.free(unsafe.Pointer(self.str_album)) }
if self.str_artist != nil { C.free(unsafe.Pointer(self.str_artist)) }
if self.str_artist_alternate != nil { C.free(unsafe.Pointer(self.str_artist_alternate)) }
if self.str_descriptioncn != nil { C.free(unsafe.Pointer(self.str_descriptioncn)) }
if self.str_descriptionde != nil { C.free(unsafe.Pointer(self.str_descriptionde)) }
if self.str_descriptionen != nil { C.free(unsafe.Pointer(self.str_descriptionen)) }
if self.str_descriptiones != nil { C.free(unsafe.Pointer(self.str_descriptiones)) }
if self.str_descriptionfr != nil { C.free(unsafe.Pointer(self.str_descriptionfr)) }
if self.str_descriptionhu != nil { C.free(unsafe.Pointer(self.str_descriptionhu)) }
if self.str_descriptionil != nil { C.free(unsafe.Pointer(self.str_descriptionil)) }
if self.str_descriptionit != nil { C.free(unsafe.Pointer(self.str_descriptionit)) }
if self.str_descriptionjp != nil { C.free(unsafe.Pointer(self.str_descriptionjp)) }
if self.str_descriptionnl != nil { C.free(unsafe.Pointer(self.str_descriptionnl)) }
if self.str_descriptionno != nil { C.free(unsafe.Pointer(self.str_descriptionno)) }
if self.str_descriptionpl != nil { C.free(unsafe.Pointer(self.str_descriptionpl)) }
if self.str_descriptionpt != nil { C.free(unsafe.Pointer(self.str_descriptionpt)) }
if self.str_descriptionru != nil { C.free(unsafe.Pointer(self.str_descriptionru)) }
if self.str_descriptionse != nil { C.free(unsafe.Pointer(self.str_descriptionse)) }
if self.str_genre != nil { C.free(unsafe.Pointer(self.str_genre)) }
if self.str_locked != nil { C.free(unsafe.Pointer(self.str_locked)) }
if self.str_mood != nil { C.free(unsafe.Pointer(self.str_mood)) }
if self.str_music_brainz_album_id != nil { C.free(unsafe.Pointer(self.str_music_brainz_album_id)) }
if self.str_music_brainz_artist_id != nil { C.free(unsafe.Pointer(self.str_music_brainz_artist_id)) }
if self.str_music_brainz_id != nil { C.free(unsafe.Pointer(self.str_music_brainz_id)) }
if self.str_musicvid != nil { C.free(unsafe.Pointer(self.str_musicvid)) }
if self.str_music_vid_company != nil { C.free(unsafe.Pointer(self.str_music_vid_company)) }
if self.str_music_vid_director != nil { C.free(unsafe.Pointer(self.str_music_vid_director)) }
if self.str_music_vid_screen1 != nil { C.free(unsafe.Pointer(self.str_music_vid_screen1)) }
if self.str_music_vid_screen2 != nil { C.free(unsafe.Pointer(self.str_music_vid_screen2)) }
if self.str_music_vid_screen3 != nil { C.free(unsafe.Pointer(self.str_music_vid_screen3)) }
if self.str_style != nil { C.free(unsafe.Pointer(self.str_style)) }
if self.str_theme != nil { C.free(unsafe.Pointer(self.str_theme)) }
if self.str_track != nil { C.free(unsafe.Pointer(self.str_track)) }
if self.str_track3dcase != nil { C.free(unsafe.Pointer(self.str_track3dcase)) }
if self.str_tracklyrics != nil { C.free(unsafe.Pointer(self.str_tracklyrics)) }
if self.str_trackthumb != nil { C.free(unsafe.Pointer(self.str_trackthumb)) }
    C.free(unsafe.Pointer(self))
  }
}

//export audiodb_search_tracks_by_album_id
func audiodb_search_tracks_by_album_id(i C.int) *C.audiodb {
  self := (*C.audiodb) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb{}))))
  res, err := audiodb.SearchTracksByAlbumId(int64(i))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
  array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))
  
  slice := (*[1<<30 - 1]*C.audiodb_track)(array)

  for i, val := range res {
    data := (*C.audiodb_track) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb_track{}))))

  data.id_album = C.CString(val.IDAlbum)
data.id_artist = C.CString(val.IDArtist)
data.id_imvdb = C.CString(val.IDIMVDB)
data.id_lyric = C.CString(val.IDLyric)
data.id_track = C.CString(val.IDTrack)
data.int_cd = C.CString(val.IntCD)
data.int_duration = C.CString(val.IntDuration)
data.int_loved = C.CString(val.IntLoved)
data.int_music_vid_comments = C.CString(val.IntMusicVidComments)
data.int_music_vid_dislikes = C.CString(val.IntMusicVidDislikes)
data.int_music_vid_favorites = C.CString(val.IntMusicVidFavorites)
data.int_music_vid_likes = C.CString(val.IntMusicVidLikes)
data.int_music_vid_views = C.CString(val.IntMusicVidViews)
data.int_score = C.CString(val.IntScore)
data.int_scorevotes = C.CString(val.IntScoreVotes)
data.int_total_listeners = C.CString(val.IntTotalListeners)
data.int_totalplays = C.CString(val.IntTotalPlays)
data.int_tracknumber = C.CString(val.IntTrackNumber)
data.str_album = C.CString(val.StrAlbum)
data.str_artist = C.CString(val.StrArtist)
data.str_artist_alternate = C.CString(val.StrArtistAlternate)
data.str_descriptioncn = C.CString(val.StrDescriptionCN)
data.str_descriptionde = C.CString(val.StrDescriptionDE)
data.str_descriptionen = C.CString(val.StrDescriptionEN)
data.str_descriptiones = C.CString(val.StrDescriptionES)
data.str_descriptionfr = C.CString(val.StrDescriptionFR)
data.str_descriptionhu = C.CString(val.StrDescriptionHU)
data.str_descriptionil = C.CString(val.StrDescriptionIL)
data.str_descriptionit = C.CString(val.StrDescriptionIT)
data.str_descriptionjp = C.CString(val.StrDescriptionJP)
data.str_descriptionnl = C.CString(val.StrDescriptionNL)
data.str_descriptionno = C.CString(val.StrDescriptionNO)
data.str_descriptionpl = C.CString(val.StrDescriptionPL)
data.str_descriptionpt = C.CString(val.StrDescriptionPT)
data.str_descriptionru = C.CString(val.StrDescriptionRU)
data.str_descriptionse = C.CString(val.StrDescriptionSE)
data.str_genre = C.CString(val.StrGenre)
data.str_locked = C.CString(val.StrLocked)
data.str_mood = C.CString(val.StrMood)
data.str_music_brainz_album_id = C.CString(val.StrMusicBrainzAlbumID)
data.str_music_brainz_artist_id = C.CString(val.StrMusicBrainzArtistID)
data.str_music_brainz_id = C.CString(val.StrMusicBrainzID)
data.str_musicvid = C.CString(val.StrMusicVid)
data.str_music_vid_company = C.CString(val.StrMusicVidCompany)
data.str_music_vid_director = C.CString(val.StrMusicVidDirector)
data.str_music_vid_screen1 = C.CString(val.StrMusicVidScreen1)
data.str_music_vid_screen2 = C.CString(val.StrMusicVidScreen2)
data.str_music_vid_screen3 = C.CString(val.StrMusicVidScreen3)
data.str_style = C.CString(val.StrStyle)
data.str_theme = C.CString(val.StrTheme)
data.str_track = C.CString(val.StrTrack)
data.str_track3dcase = C.CString(val.StrTrack3DCase)
data.str_tracklyrics = C.CString(val.StrTrackLyrics)
data.str_trackthumb = C.CString(val.StrTrackThumb)
    slice[i] = data
  }

  self.buffer = unsafe.Pointer((**C.audiodb_track) (array))
  self.buffer_length = C.size_t(len(res))
  self.error = nil
  return self
}

//export audiodb_tracks_clean
func audiodb_tracks_clean(self **C.audiodb_track, length C.int) {
  if self != nil {
    slice := (*[1<<30 - 1]*C.audiodb_track)(unsafe.Pointer(self))[:length:length]
    for i := 0 ; i < len(slice); i++ {
      if slice[i] != nil {
        audiodb_track_clean(slice[i])
      }
    }
    C.free(unsafe.Pointer(self))
  }
}

//export audiodb_search_music_videos_by_artist_id
func audiodb_search_music_videos_by_artist_id(i C.int) *C.audiodb {
  self := (*C.audiodb) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb{}))))
  res, err := audiodb.SearchTracksByAlbumId(int64(i))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
  array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))
  
  slice := (*[1<<30 - 1]*C.audiodb_music_video)(array)

  for i, val := range res {
    data := (*C.audiodb_music_video) (C.malloc(C.size_t(unsafe.Sizeof(C.audiodb_music_video{}))))
    data.id_album = C.CString(val.IDAlbum)
data.id_artist = C.CString(val.IDArtist)
data.id_track = C.CString(val.IDTrack)
data.str_descriptionen = C.CString(val.StrDescriptionEN)
data.str_musicvid = C.CString(val.StrMusicVid)
data.str_track = C.CString(val.StrTrack)
data.str_trackthumb = C.CString(val.StrTrackThumb)
    slice[i] = data
  }

  self.buffer = unsafe.Pointer((**C.audiodb_music_video) (array))
  self.buffer_length = C.size_t(len(res))
  self.error = nil
  return self
}

//export audiodb_music_video_clean
func audiodb_music_video_clean(self **C.audiodb_music_video, length C.int) {
  if self != nil {
    slice := (*[1<<30 - 1]*C.audiodb_music_video)(unsafe.Pointer(self))[:length:length]
    for i := 0 ; i < len(slice); i++ {
      if slice[i] != nil {
        if slice[i].id_album != nil { C.free(unsafe.Pointer(slice[i].id_album)) }
if slice[i].id_artist != nil { C.free(unsafe.Pointer(slice[i].id_artist)) }
if slice[i].id_track != nil { C.free(unsafe.Pointer(slice[i].id_track)) }
if slice[i].str_descriptionen != nil { C.free(unsafe.Pointer(slice[i].str_descriptionen)) }
if slice[i].str_musicvid != nil { C.free(unsafe.Pointer(slice[i].str_musicvid)) }
if slice[i].str_track != nil { C.free(unsafe.Pointer(slice[i].str_track)) }
if slice[i].str_trackthumb != nil { C.free(unsafe.Pointer(slice[i].str_trackthumb)) }
        C.free(unsafe.Pointer(slice[i]))
      }
    }
    C.free(unsafe.Pointer(self))
  }
}

func main() {}