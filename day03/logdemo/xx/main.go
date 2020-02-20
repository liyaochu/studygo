package xx

import "studygo/day03/logdemo/mylog"

func main() {
     fl := mylog.NewFileLogger(mylog.DEBUG, "./", "test.log")
     userID := 10
     fl.Debug("id是%d的用户一直在尝试登陆", userID)}
