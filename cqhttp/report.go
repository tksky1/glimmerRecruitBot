package cqhttp

type PostType string

const (
	PostTypeMessage     PostType = "message"
	PostTypeMessageSent PostType = "message_sent"
	PostTypeRequest     PostType = "request"
	PostTypeNotice      PostType = "notice"
	PostTypeMetaEvent   PostType = "meta_type"
)

type GenericReport struct {
	TimeStamp int64 `json:"time"`
	SelfID    int64 `json:"self_id"` // bot's qq
	PostType  `json:"post_type"`
}
