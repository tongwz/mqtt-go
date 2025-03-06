package consts

const (
	TimeYmdHis    = "2006-01-02 15:04:05"
	TimeYmdHi     = "2006-01-02 15:04"
	TimeYmdHisINT = "20060102150405"
	TimeYmd       = "2006-01-02"
	TimeYm        = "2006-01"
)

// 进行数据库操作的类型
var DaoExecType = map[string]bool{
	"create": true,
	"update": true,
	"delete": true,
}

// 富有支付类型
var FuiouPayType = map[string]string{
	"ALIPAY":  "支付宝",
	"LETPAY":  "微信小程序",
	"TAPPLET": "富有微信小程序",
}

// 游客id
const VisitorId = "-9999"

// 直播时长多久算直播一天
const EffectiveSecondsOneDay = 7200 // 秒
