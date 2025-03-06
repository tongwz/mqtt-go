package consts

var CoinPayType = map[int]string{
	1:  "支付宝",
	2:  "微信",
	10: "富有支付宝",
	11: "杉德支付宝",
	12: "富有微信",
	13: "杉德微信",
	14: "拉卡拉支付宝",
	15: "拉卡拉微信",
}

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
const SNS_JSCODE = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
