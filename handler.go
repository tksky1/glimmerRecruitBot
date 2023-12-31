package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
)

func check(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func handlePost(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if check(err) {
		return
	}
	json, err := simplejson.NewJson(body)
	if check(err) {
		return
	}
	msgType, err := json.Get("message_type").String()
	if check(err) {
		return
	}
	if msgType != "group" {
		return
	}

	groupNum, err := json.Get("group_id").Int64()
	if check(err) {
		return
	}
	if groupNum != config.adminGroupId {
		return
	}

	msg, err := json.Get("message").String()
	if check(err) {
		return
	}
	parts := strings.Split(msg, " ")
	if len(parts) == 3 && parts[0] == "表扬" {
		println(msg)
		sendGroupMsg(parts[1], parts[2], "")
	}
	if len(parts) == 4 && parts[0] == "表扬" {
		println(msg)
		sendGroupMsg(parts[1], parts[2], parts[3])
	}
}
