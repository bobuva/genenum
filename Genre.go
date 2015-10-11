package main


type Genre int8 
const (
  _Genre = iota
  Genre_Bluegrass
  Genre_Blues
  Genre_Choral
  Genre_Classical
  Genre_ClassicRock
  Genre_Country
  Genre_Electronica
  Genre_Folk
  Genre_HeavyMetal
  Genre_HipHop
  Genre_Jazz
  Genre_Lullabies
  Genre_Reggae
  Genre_RhythmAndBlues
  Genre_Soul
  Genre_SouthernRock
)


var GenreStrings = 
map[string]Genre {
  "Bluegrass":	Genre_Bluegrass,
  "Blues":	Genre_Blues,
  "Choral":	Genre_Choral,
  "Classical":	Genre_Classical,
  "ClassicRock":	Genre_ClassicRock,
  "Country":	Genre_Country,
  "Electronica":	Genre_Electronica,
  "Folk":	Genre_Folk,
  "HeavyMetal":	Genre_HeavyMetal,
  "HipHop":	Genre_HipHop,
  "Jazz":	Genre_Jazz,
  "Lullabies":	Genre_Lullabies,
  "Reggae":	Genre_Reggae,
  "RhythmAndBlues":	Genre_RhythmAndBlues,
  "Soul":	Genre_Soul,
  "SouthernRock":	Genre_SouthernRock,
}


func (l Genre) String() string {
  for s, v := range GenreStrings {
    if l == v {
      return s
    }
  }
  return "invalid"
}
