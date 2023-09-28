package cqhttp

type MessageType string

const (
	MessageTypePrivate MessageType = "private"
	MessageTypeGroup   MessageType = "group"
)

// ref: https://docs.go-cqhttp.org/event/#%E6%B6%88%E6%81%AF%E4%B8%8A%E6%8A%A5
type Message struct {
	MessageType `json:"message_type"`
	MessageID   int32  `json:"message_id"`
	UserID      int64  `json:"user_id"`
	GroupID     int64  `json:"group_id"` // will not be provided if is private type
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
}

type MessageReply interface{}

type PrivateMessageReply struct {
	Reply      string `json:"reply,omitempty"`
	AutoEscape bool   `json:"auto_escape,omitemtpy"` // send message as plain-text
}

type GroupMessageReply struct {
	Reply      string `json:"reply,omitemtpy"`
	AutoEscape bool   `json:"auto_escape,omitemtpy"` // send message as plain-text
	AtSender   bool   `json:"as_sender,omitemtpy"`
	Delete     bool   `json:"delete,omitemtpy"`
}
