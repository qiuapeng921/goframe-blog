package boot

import (
	"blog/boot/cronfunc"
	"github.com/gogf/gf/os/gcron"
)

func init() {
	_, _ = gcron.Add("1 * * * * *", cronfunc.CronTest, "second-cron")
}