package main

import (
	"net/http"
	"strconv"
	"strings"
)

func sendGroupMsg(name string, problemName string, comment string) {
	var err error
	if comment == "" {
		_, err = http.Post("http://127.0.0.1:"+strconv.Itoa(config.postTo)+"/send_group_msg",
			"application/x-www-form-urlencoded",
			strings.NewReader("group_id="+strconv.FormatInt(config.recruitGroupId, 10)+
				"&message="+"管理员表扬了"+name+"同学提交的题目"+problemName+"，再接再厉哦~"))
	} else {
		_, err = http.Post("http://127.0.0.1:"+strconv.Itoa(config.postTo)+"/send_group_msg",
			"application/x-www-form-urlencoded",
			strings.NewReader("group_id="+strconv.FormatInt(config.recruitGroupId, 10)+
				"&message="+"管理员表扬了"+name+"同学提交的题目"+problemName+"，评语："+comment+" 再接再厉哦~"))
	}
	if err != nil {
		println("post err:", err.Error())
	}
}
