package jokes

import "bot-git/bot/blacklist"

func getFreshJoke(jokes []*string, alreadySent *blacklist.BlackList) (*string, bool) {
	for _, joke := range jokes {
		if alreadySent.IsFresh(joke) {
			return joke, true
		}
	}
	return nil, false
}

func getFreshForFetcher(fetch func(int) []*string, try int, alreadySent *blacklist.BlackList) (*string, bool) {
	to := try + 1
	for i := 1; i <= to; i++ {
		fresh, ok := getFreshJoke(fetch(i), alreadySent)
		if ok {
			return fresh, true
		}
	}
	return nil, false
}
