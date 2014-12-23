package token

import(
	"time"
	"github.com/robfig/cron"
)

type Cycle struct {
	interval int64
	duration int64
	quantity int64
}