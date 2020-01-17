package boot

import (
	"blog/boot/cron"
	"github.com/gogf/gf/os/gcron"
)

func init() {
	_, _ = gcron.Add("1 * * * * *", cron.CronTest, "second-cron")
}