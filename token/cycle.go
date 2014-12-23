package token

import(
	"time"
	"github.com/robfig/cron"
)

type Cycle struct {
	duration int64
	interval string
	quantity int64
	respawn Time
}

const (
	DurationTypes = [1,2,3,4,5,6,7,8,9,10]
	IntervalTypes = map[string]int64
		"a day" : time.Hour * 24,
		"a week": time.Hour * 24 * 7,
		"a month": time.Hour * 24 * (time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day())
	QuantityTypes = [1,2,3,4,5]
)


// Checks if a token is available based on last token time within quantity available
// I think.
func (t *Token) isTokenAvailable() bool {
	if(int64(len(t.tokens)) >=t.cycle.quantity) {
		timeUsed := time.Parse(t.tokens[t.cycle.quantity].timestamp)
		if(time.Now().after(timeUsed+t.token.cycle.intervalTypes)){
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}
