package login

import (
	"encoding/json"
	"game_framework/src/eassy/login/accounts"
	"game_framework/src/eassy/util"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type LoginJson struct {
	Phone   string
	AccPwd  string
	ReqType int
	Flag    string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(403)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form) == 0 {
		w.WriteHeader(400)
		return
	}
	phone := r.Form["phone"][0]
	token := r.Form["token"][0]
	if util.CheckToken(phone, token) {
		util.ResponseWithJson(w, http.StatusOK, util.Response{
			Code: http.StatusOK,
			Msg:  "验证通过",
			Data: nil,
		})
	} else {
		util.ResponseWithJson(w, http.StatusOK, util.Response{
			Code: http.StatusOK,
			Msg:  "验证失败",
			Data: nil,
		})
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//验证key
	r.ParseForm()
	if len(r.Form) == 0 {
		w.WriteHeader(400)
		return
	}
	var jData LoginJson
	data := r.Form["data"][0]
	_ = json.Unmarshal([]byte(data), &jData)
	phone := jData.Phone     // r.Form["phone"][0]
	accPwd := jData.AccPwd   // r.Form["accPwd"][0]
	reqType := jData.ReqType // r.Form["reqType"][0]
	flag := jData.Flag       //r.Form["flag"][0]
	md5str := strings.ToUpper(util.MD5(Md5Key + phone + accPwd + strconv.Itoa(reqType)))
	if md5str == strings.ToUpper(flag) {
		//md5验证通过
		switch reqType {
		case 1:
			//register
			// todo 过滤字符
			if ok, _ := accounts.Acc.FindAccount(phone); ok {
				util.ResponseWithJson(w, http.StatusOK, util.Response{
					Code: 0,
					Msg:  "已注册！",
					Data: nil,
				})
				return
			}
			if accounts.Acc.Register(phone, accPwd) {
				util.ResponseWithJson(w, http.StatusOK, util.Response{
					Code: 1,
					Msg:  "注册成功！",
					Data: nil,
				})
				return
			}
			util.ResponseWithJson(w, http.StatusOK, util.Response{
				Code: 0,
				Msg:  "注册失败！",
				Data: nil,
			})
		case 2:
			//login
			if ok, acc := accounts.Acc.FindAccount(phone); ok {
				if util.MD5(accPwd) == acc.Pwd {
					//todo 如果已有token 则删除原有token
					//修改登录时间和登录IP
					acc.LastIP = int64(util.Ip2long(util.ClientIP(r)))
					acc.LastLoginTime = time.Now().Unix()
					go accounts.Acc.UpdateInfo(acc)
					//生成token
					var tokenStruct = &util.TokenStruct{Phone: phone, AccPwd: accPwd}
					token, _ := util.GenerateToken(tokenStruct)
					util.ResponseWithJson(w, http.StatusOK, util.Response{Code: 1, Msg: "登录成功！", Data: util.JwtToken{Token: token}})
					return
				}
				util.ResponseWithJson(w, http.StatusOK, util.Response{
					Code: 0,
					Msg:  "账号或密码错误！",
					Data: nil,
				})
				return
			}
			util.ResponseWithJson(w, http.StatusOK, util.Response{
				Code: 0,
				Msg:  "账号或密码错误！",
				Data: nil,
			})
		case 3: //修改密码 手机验证码

		default:
			log.Fatal("请求类型错误")
			w.WriteHeader(400)
		}
		return
	}
	w.WriteHeader(400)
}
