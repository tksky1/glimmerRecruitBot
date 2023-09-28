package dice

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"glimmerBot/cqhttp"
)

type rp struct {
	userID       int64
	rp           int64
	generateTime time.Time
}

func (r *rp) hasExpired() bool {
	return time.Now().Day() != r.generateTime.Day()
}

var rpCache map[int64]rp

func getRP(userID int64) rp {
	if rpCache == nil {
		rpCache = make(map[int64]rp)
	}

	r, ok := rpCache[userID]
	if ok && !r.hasExpired() {
		return r
	}

	randomRP, _ := dice(5, 20)

	newRP := rp{
		userID:       userID,
		rp:           randomRP,
		generateTime: time.Now(),
	}

	rpCache[userID] = newRP

	return newRP
}

func generateMessageReply(answer string, message *cqhttp.Message) cqhttp.MessageReply {
	switch message.MessageType {
	case cqhttp.MessageTypeGroup:
		return cqhttp.GroupMessageReply{
			Reply:      answer,
			AutoEscape: false,
			AtSender:   true,
		}

	case cqhttp.MessageTypePrivate:
		return cqhttp.PrivateMessageReply{
			Reply:      answer,
			AutoEscape: false,
		}

	default:
		return nil
	}
}

func LuckyHandler(c *gin.Context) {
	message := &cqhttp.Message{}
	c.ShouldBindBodyWith(message, binding.JSON)

	answer := fmt.Sprintf("今日人品 %d", getRP(message.UserID).rp)

	reply := generateMessageReply(answer, message)
	if reply != nil {
		c.JSON(
			http.StatusOK,
			reply,
		)
	}
}
