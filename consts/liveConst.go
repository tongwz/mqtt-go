package consts

// 铂金火箭爱心火箭，热门卡id
const (
	PtRocket   int = 184 // 铂金火箭
	LoveRocket int = 183 // 爱心火箭
	HotCard    int = 217 // 热门卡 - 下架状态
)

// 关于websocket的状态码
const (
	SUCCESS = 0
	FAIL    = -1

	SYSTEM_ID_ERROR      = -1001
	ONLINE_MESSAGE_CODE  = 0
	OFFLINE_MESSAGE_CODE = 0
)

const (

	// 礼物
	GIFT_NOT_FOUND        int = 10002
	GIFT_NOT_GUARD_TYPE   int = 10003
	GIFT_NOT_GUARD_USER   int = 10004
	GIFT_BACKPACK_NUM_ERR int = 10005
	GIFT_BACKPACK_ERR     int = 10006
	// 基础异常
	MSG_ERR      int = 10100
	TOKEN_ERR    int = 10101
	PLATFORM_ERR int = 10102

	// 用户
	USER_NOT_FOUND   int = 11001
	USER_UPDATE_ERR  int = 11002
	USER_NOT_ONE_ERR int = 11003
	LIVE_USER_ERR    int = 11004

	// 金豆
	COIN_NOT_ENOUGH int = 12001
	COIN_RECORD_ERR int = 12002

	// 事务
	TRANSACTION_ERR int = 13001
)

var SocketRspCodeMap = map[int]string{

	10002: "礼物信息不存在",
	10003: "该礼物是年守护专属礼物",
	10004: "守护礼物只能赠送守护主播",
	10005: "背包中数量不足",
	10006: "背包扣除异常",

	10100: "入参异常",
	10101: "token异常",
	10102: "连接平台app,web,windows",

	11001: "未知用户",
	11002: "用户更新异常",
	11003: "登录与送礼物用户异常",
	11004: "主播已下播～",

	12001: "金豆数不足",
	12002: "金豆流水异常",

	13001: "事务异常",
}

type Method string

// 关于websocket中收到的消息方法类型
const (
	SendGift       Method = "SendGift"       // 直播间-送礼物 - 用户
	UsersPreview   Method = "UsersPreview"   // 直播间-用户概览(人数，前20人头像等) - 服务端
	EnterRoom      Method = "EnterRoom"      // 直播间-进入直播间
	LeaveRoom      Method = "LeaveRoom"      // 直播间-离开直播间
	SysTokenExpire Method = "SysTokenExpire" // 系统-token过期
	SysMultiLogin  Method = "SysMultiLogin"  // 系统-多端登录 - 挤掉线
	SysActivityRed Method = "SysActivityRed" // 系统 - 活动红包
	Online         Method = "Online"
	Offline        Method = "Offline"
	SendMsg        Method = "SendMsg"
	SysConn        Method = "SysConnect"    // 系统 - 连接
	SysDisConn     Method = "SysDisConnect" // 系统 - 断连
)
