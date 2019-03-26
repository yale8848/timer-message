// Create by Yale 2019/3/14 11:27
package main

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/go-vgo/robotgo"
	"github.com/robfig/cron"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const DURATION = "@hourly"

type kv struct {
	k string
	v string
}

func main() {

	CHARS := getCronValue()
	ch := make(chan bool)
	dur := DURATION

	name := filepath.Base(os.Args[0])
	if pos := strings.Index(name, "("); pos >= 0 {
		if pos1 := strings.LastIndex(name, ")"); pos1 >= 0 && pos1 != pos {
			dur = name[pos+1 : pos1]
			if strings.Index(dur, "@") >= 0 {
				dur = strings.Replace(dur, "_", " ", -1)
			} else {
				for _, v := range CHARS {
					dur = strings.Replace(dur, v.k, v.v, -1)
				}
			}
		}
	}
	sb := bytes.Buffer{}
	sb.WriteString("\r\n")
	for _, v := range CHARS {
		sb.WriteString(v.k)
		sb.WriteString(" --> ")
		sb.WriteString(v.v)
		sb.WriteString("\r\n")
	}
	fmt.Println(sb.String())
	go func() {
		alertCount := 0
		c := cron.New()
		err := c.AddFunc(dur, func() {
			if alertCount == 1 {
				return
			}
			alertCount++
			enc := mahonia.NewEncoder("gbk")
			tm:=time.Now().Format("2006-01-02 15:04:05")
			ret := robotgo.ShowAlert(enc.ConvertString("timer-message"), enc.ConvertString(fmt.Sprintf("执行时间:%s\r\n执行周期：%s\r\n时间映射：\r\n%s\r\n点击<确定>按钮程序退出", tm,dur, sb.String())))
			if ret == 0 {
				ret = robotgo.ShowAlert("", enc.ConvertString("确定要退出程序？"))
				if ret == 0 {
					ch <- true
				}
			}
			alertCount--
		})
		if err != nil {
			robotgo.ShowAlert("err", err.Error())
			ch <- true
		} else {
			c.Start()
		}
	}()

	<-ch
}

func getCronValue() []kv {
	return []kv{
		{
			k: "a",
			v: "*",
		},
		{
			k: "b",
			v: "/",
		},
		{
			k: "c",
			v: ",",
		},
		{
			k: "d",
			v: "-",
		},
		{
			k: "e",
			v: "?",
		},
		{
			k: "_",
			v: " ",
		},
	}
}
