package bt

// Strategy 定义了 策略接口
type Strategy interface {
	// 回调函数

	// OnInit 策略初始化
	OnInit()

	// OnStart 策略启动
	OnStart()

	// OnStop 策略结束
	OnStop()

	// OnTick 接收到 tick 数据
	// TODO: 添加参数
	// 自己合成 K 线的位置
	OnTick()

	// OnBar 接收到 K 线数据
	// TODO: 添加参数
	// 计算指标的位置
	OnBar()

	// OnTrade 成交的回报时候调用
	// TODO: 添加参数
	OnTrade()

	// OnOrder 委托更新时候调用
	// TODO: 添加参数
	OnOrder()

	// OnStopOrder 本地停止单状态变化
	// TODO: 添加参数
	OnStopOrder()
}

// 交易相关的方法
type trader interface {
	Buy()
	Sell()
	Short()
	Cover()

	// 撤销委托
	Cancel()
	CancelAll()
	SendOrder()
}

// 辅助接口
type helper interface {
	// 写日志
	log()
	// 区分实盘与回测
	getEngineType()
	// 加载历史 K 线数据
	loadBar()
	// 加载历史 tick 数据
	loadTick()
	// 通知 GUI 更新
	putEvent()

	// 通知接口
	Notifier() // sendEmail()

	// 保存与回复策略的状态
	save() //	syncData()
	load()
}
