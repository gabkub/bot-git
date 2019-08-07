package meme

import "bot-git/bot/abstract"

var short = "_meme, mem_ - losowy mem"
var long = `Wysyła losowy śmieszny obrazek. Odnośnik w tytule otwiera obrazek w nowej karcie.
Limity:
7:00-8:59 - 3 memy
9:00-14:59 - 1 mem na godzinę
15:00-6:59 - brak limitów

Pełna lista komend:
_meme, mem_`

var help = abstract.NewHelp(short, long)
