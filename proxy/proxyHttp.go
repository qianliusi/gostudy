package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql包
	"github.com/jmoiron/sqlx"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

import (
	"log"
	//	"strings"
)

var WB selenium.WebDriver
var SVC *selenium.Service
var proxyWeb = "http://test.hrwork.com"
var sizeOfMyStruct = int(unsafe.Sizeof(Sms{}))

func SmsToBytes(s *Sms) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func statistic(w http.ResponseWriter, r *http.Request) {
	log.Println("request from " + r.Host + " " + r.Method + " " + r.RequestURI)
	path := proxyWeb + r.URL.Path
	/*if strings.HasSuffix(path, ".js") {
		proxy(w, r)
		return
	}
	body, cookies := html(path)
	if cookies != nil {
		for _, v := range cookies {
			w.Header().Set("set-cookie", convertCookie(v))
		}
	}*/
	if strings.Contains(path, "querySms") {
		env := "daily"
		if strings.Contains(path, "gray") {
			env = "gray"
		}
		for _, v := range query(env) {
			bytes := []byte(fmt.Sprintln(v))
			bytes = append(bytes, '\n')
			w.Write(bytes)
		}

	}
}

var db *sql.DB

func convertCookie(cookie selenium.Cookie) (cookieStr string) {
	cookieStr = cookie.Name + "=" + cookie.Value + ";"
	if cookie.Expiry > 0 {
		exp := time.Unix(int64(cookie.Expiry), 0).Format(http.TimeFormat)
		cookieStr += "expires=" + exp + ";"
	}
	cookieStr += "Path=" + cookie.Path + ";"
	if cookie.HTTPOnly {
		cookieStr += "HttpOnly"
	}
	return
}

func proxy(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", proxyWeb+r.URL.Path, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Write(body)
}

func main() {
	//open()
	mux := http.NewServeMux()
	mux.HandleFunc("/", statistic)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	//err := http.ListenAndServeTLS(":443", "C:/Users/zhuyu/go/src/gostudy/proxy/cert.pem", "C:/Users/zhuyu/go/src/gostudy/proxy/key.pem", mux)
	//if err != nil {
	//	log.Fatal("ListenAndServe:", err)
	//}
}

const (
	seleniumPath = `D:\chormdriver\chromedriver.exe`
	port         = 9515
)

func open() {
	//如果seleniumServer没有启动，就启动一个seleniumServer所需要的参数，可以为空，示例请参见https://github.com/tebeka/selenium/blob/master/example_test.go
	opts := []selenium.ServiceOption{}
	//opts := []selenium.ServiceOption{
	//    selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
	//    selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
	//}
	//selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		return
	}
	//链接本地的浏览器 chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)
	// 调起chrome浏览器
	wb, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	WB = wb
	SVC = service
}

func html(url string) (body []byte, cookies []selenium.Cookie) {
	//关闭一个webDriver会对应关闭一个chrome窗口
	err := WB.Get(url)
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return nil, nil
	}
	cookies, err = WB.GetCookies()
	if err != nil {
		fmt.Println("get cookie faild", err.Error())
	}
	if strings.Contains(url, "api/account") {
		json, err := WB.ExecuteScript(`return document.body.innerText`, nil)
		if err != nil {
			fmt.Println("json faild", err.Error())
		}
		bJson := json.(string)
		fmt.Println(bJson)
		body = []byte(bJson)
		return
	}
	_, err = WB.ExecuteScript(`var wxCodeDiv = document.createElement('div');wxCodeDiv.id = 'wx_code';document.body.append(wxCodeDiv)`, nil)
	if err != nil {
		fmt.Println("wx_code faild", err.Error())
	}
	html, err := WB.ExecuteScript("return document.documentElement.outerHTML", nil)
	body = []byte(html.(string))
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return nil, nil
	}
	return
}

type Sms struct {
	Phone    string `db:"phone"`
	Content  string `db:"content"`
	SendTime string `db:"send_time"`
}

func query(env string) []Sms {
	//conn, err := goSqlHelper.MysqlOpen("hayswang:hayswang@3365@tcp(rm-wz9f4x465k399e7fvlo.mysql.rds.aliyuncs.com:3306)/platform-sms-daily?charset=utf8&parseTime=True")
	db, err := sqlx.Connect("mysql", "hayswang:hayswang@3365@tcp(rm-wz9f4x465k399e7fvlo.mysql.rds.aliyuncs.com:3306)/platform-sms-"+env+"?charset=utf8&parseTime=True")
	defer db.Close()
	checkErr(err)
	res := make([]Sms, 0)
	sms := Sms{}
	rows, err := db.Queryx("select phone, content,send_time from sms_record  order by sms_id desc limit 10")
	for rows.Next() {
		err := rows.StructScan(&sms)
		if err != nil {
			log.Fatalln(err)
		}
		res = append(res, sms)
	}
	//fmt.Println(res)
	return res
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
