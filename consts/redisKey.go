package consts

const (
	// 礼物详情缓存key
	RedisGiftInfoPrefix = "gift_info:%d"
	// 守护关系 key 客户id 主播id按顺序替换 PHP
	RedisGuardRelationPrefix = "getUserGuard_%d_%d"
	// 直播间送礼物人集合PHP
	RedisSendGiftUsersSet = "sendgift_users_%d"
	// option私有配置 PHP
	RedisConfigPri = "getConfigPri"
	// option私有配置 PHP
	RedisThreePay = "getThreePay"
	// 主播所在家族
	RedisUidFromFamily = "uid_from_family:%d"
	// 家族所在族长
	RedisFamilyFromLeader = "family_from_leader:%d"
	// 直播PK 双方uid 类型hashMap key 发起方 value 接受pk方 PHP
	RedisLivePK = "LivePK"
	// 直播PK 双方礼物金豆数 类型hashMap key uid value 金豆数 PHP
	RedisLivePKGift = "LivePK_gift"
	// 单个直播间送礼物排行榜 zSet user_stream 其中stream 是 uid_时间戳（秒） key uid value 金豆数 PHP
	RedisLiveStreamRankingList = "user_%s"
	// 小黑屋 - 送礼物排行榜 zSet %d = liveuid key uid value 金豆数
	RedisOfflineLiveStreamRankingList = "user_offline_live_stream_%d"
	// gift token前缀
	RedisGiftTokenKey = "gift_token_for:%s"

	// 中奖次数缓存 - %d = poolId ;%d = uId
	RedisLotterySizePool = "lottery_size_pool_%d_%d"
	// user token PHP %d = uid
	RedisUserTokenKey = "token_%d"
	// 用户信息缓存 %d = uid PHP - 个人建议缓存时间10分钟
	RedisUserInfoKey = "userinfo_%d"
	// 主播当日贡献榜 - %s = 日期2024-11-08 - %d = liveuid
	RedisDailyShell = "daily_shell:%s:%d"
	// 主播单场 - 收入珍珠数量 - %s stream 值 例如 10008_1732170968
	RedisStreamIncome = "live_stream_income:%s"
	// token订阅消息Key - token变化监听
	RedisTokenMonitor = "token_change"
	// 活动红包key %d = 红包id
	RedisActivityRedListKey = "activity_red:%d"
	// 活动红包 - 用户抢红包已经抢了几次 key %s 日期=例子 2024-12-20 %d = uid
	RedisActivityRedUserGrabCountKey = "activity_red:user_grab_count:%s:%d"
	// 热门首页 - 缓存直播主播列表 - json串
	RedisHotListKey = "live_hot_list"
	// 红包列表 %s = stream PHP
	RedisLiveRoomRedListKey = "red_list_%s"

	// 坐骑列表 所有列表json - PHP
	RedisCarInfoKey = "car_info"
	// 徽章
	RedisSpecialBadgeListKey = "{2025}special_badge_list"
	// 进场特效
	RedisSpecialEntryListKey = "{2025}special_entry_list"
	// 空间特效
	RedisSpecialHomepageListKey = "{2025}special_homepage_list"
	// 头像框
	RedisSpecialAvatarListKey = "{2025}special_avatar_list"
)
