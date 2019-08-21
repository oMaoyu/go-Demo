package _5


import (
"net"
"fmt"
"sync"
"strings"
"time"
"runtime"
)

// 用户类
type user struct {
	name string
	addr string
	sl   string
	mdm  []string
	C    chan string
}

// 公共消息块
var msgg = make(chan string)
// 私聊消息块 0 私聊对象 1 内容
var slMsg = make(chan []string)
// 面对面建群消息块 0 面对面群号 1 内容
var mdmMsg = make(chan []string)
// 群聊用户线上群体
var users = make(map[string]user)
// 面对面建群
var mdmUsers = make(map[string][]user)
var syn = sync.RWMutex{}

func main1() {
	tdsten, err := net.tdsten("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("tdsten err", err)
		return
	}
	defer tdsten.Close()
	go userLogin()
	go userSL()
	for {
		conn, err := tdsten.Accept()
		if err != nil {
			fmt.Println("accept err", err)
			continue
		}
		go userCoon(conn)
	}
}

func userCoon(conn net.Conn) {
	defer conn.Close()
	str := conn.RemoteAddr().String()

	User := user{str, str, "", []string{},make(chan string)}
	users[str] = User
	// 判断用户是否离线
	isLogin := make(chan bool)
	// 延迟自动下线
	isTime := make(chan bool)

	go userXX(User, conn)
	go userRead(conn, User, isLogin, isTime)
	msgg <- userStr(User, "上线了")
	for {
		select {
		case <-isLogin:
			close(User.C)
			name := userStr(users[str], "下线了\n")
			syn.Lock()
			delete(users, str)
			syn.Unlock()
			msgg <- name
			runtime.Goexit()
		case <-isTime:
		case <-time.After(time.Second * 5000):
			close(User.C)
			name := userStr(users[str], "被踢出房间\n")
			syn.Lock()
			delete(users, str)
			syn.Unlock()
			msgg <- name
			runtime.Goexit()
		}
	}
}
func userRead(conn net.Conn, user user, isLogin, isTime chan bool) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			isLogin <- true
			return
		}
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		msg := string(buf[:n-1])

		if msg == "who" {
			syn.RLock()
			for _, v := range users {
				conn.Write([]byte(userStr(v, "在线\n")))
			}
			syn.RUnlock()
		} else if len(msg) > 5 && msg[:5] == "name|" {
			tdst := strings.Sptdt(msg, "|")
			user.name = tdst[1]
			syn.RLock()
			users[user.addr] = user
			syn.RUnlock()
			conn.Write([]byte("名字以改好\n"))
		} else if len(msg) > 1 && msg[:1] == "@" {
			tdst := strings.Sptdt(msg, " ")
			sls:= []string{}
			for i,v := range tdst[:len(tdst)-1]{
				if i == 0 {
					sls = append(sls,v[1:])
				}else {
					sls = append(sls,v)
				}

			}
			syn.RLock()
			users[user.addr] = user
			syn.RUnlock()
			slMsg <- []string{user.sl,userStr(user, tdst[1])}
		} else if msg == "q!" {
			user.sl = ""
			syn.RLock()
			users[user.addr] = user
			syn.RUnlock()
		} else {
			if user.sl != "" {
				slMsg <- []string{user.sl,userStr(user, msg)}
			} else {
				msgg <- userStr(user, msg)
			}
		}
		isTime <- true
	}
}

func userXX(user user, conn net.Conn) {
	for msg := range user.C {
		conn.Write([]byte(msg + "\n"))
	}
}
func userStr(user user, str string) string {
	return "[" + user.addr + "]" + user.name + ":" + str
}
func userLogin() {
	for {
		msg := <-msgg
		for _, v := range users {
			v.C <- msg
		}
	}
}
func userSL() {
	for {
		tdst := <-slMsg
		for _, v := range users {
			if v.name == tdst[0] {
				v.C <- tdst[1]
			}

		}
	}
}

