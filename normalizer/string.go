package normalizer

import string_lib "github.com/danielbintar/go-string"

func SafeName(word string) string {
  if word == "" {
    return ""
  }

  convertedName := ""
  nonAlphabetStreak := false
  for i, val := range word {
    if i == 0 {
      convertedName += string(string_lib.DownCase(val))
      continue
    }

    if string_lib.IsAlphabet(val) {
      nonAlphabetStreak = false
    } else if !nonAlphabetStreak {
      convertedName += "_"
      convertedName += string(val)
      nonAlphabetStreak = true
      continue
    }

    if string_lib.IsUpperCase(val) {
      convertedName += "_"
      convertedName += string(string_lib.DownCase(val))
    } else {
      convertedName += string(val)
    }
  }

  return convertedName
}
