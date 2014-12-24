package token

import(
	"time"
	"fmt"
)

// type Cycle struct {
// 	quantity int64 // How many times per interval
// 	interval string // Intervals
// 	start int64 // Timestamp for initial creation...
// }


var IntervalTypes = map[string]time.Duration{
"a day" : time.Hour * 24,
"a week": time.Hour * 24 * 7,
"a month": time.Hour * 24 * 30,// lazy for now //(time.Date(time.Now().year(), m+1, 0, 0, 0, 0, 0, time.UTC).Day()),
}


// func GetCycle(q int64, i string) *Cycle {
// 	return &Cycle{
// 		q,
// 		i,
// 		time.Now().Unix(),
// 	}
// }

// Checks if a token is available based on last token time within quantity available
// I think.
// Conceptually, this sets up the timing as "time since last used the last one," which may or may
// not be how people actually concieve of this. Curious to try a few implementations/mental models.
func (t *Token) IsTokenAvailable() bool {

	if(int64(len(t.tokens)) >= t.quantity) {
		lastToken := t.tokens[t.quantity-1]
		timeUsed := time.Unix(lastToken.timestamp,0)

		// Check if right now is AFTER the last token w/i bounds was used.
		fmt.Println("TimeUsed ",timeUsed.Add(IntervalTypes[t.interval]) )
		fmt.Println("Now ", time.Now())
		if(time.Now().Before(timeUsed.Add(IntervalTypes[t.interval])) ) {
			fmt.Println(t.title, "is NOT available")
			return false
		} else {
			fmt.Println(t.title, "is available")
			return true
		}
	} else {
		fmt.Println(t.title, "is available")
		return true
	}
}
