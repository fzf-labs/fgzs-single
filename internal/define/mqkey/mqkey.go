package mqkey

import "github.com/fzf-labs/fpkg/mq"

var (
	UserCancellation = mq.NewBusiness("用户注销", "user_cancellation", "user_cancellation", "user_cancellation")
)
