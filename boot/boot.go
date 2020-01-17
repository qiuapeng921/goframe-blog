package boot

import (
	"blog/boot/cron"
	"github.com/gogf/gf/os/gcron"
)

func init() {
	_, _ = gcron.Add("* * * * * *", cron.CronTest, "second-cron")
}