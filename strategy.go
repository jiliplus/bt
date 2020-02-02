package bt

type account struct {
	// 持仓量
	pos float64
}

// demo 实现了 Strategy 接口
// demo 会更具两条均线的金叉与死叉来进行交易
type demo struct {
	account
	// 可调参数
	fastWindow, slowWindow int

	// 内部变量
	fastMA0, fastMA1 float64
	slowMA0, slowMA1 float64
}

// TODO: 完成这个方法
func (d *demo) log(str string) {
}

// TODO: 完成这个方法
func (d *demo) loadBar(day int) {
}

func (d *demo) OnInit() {
	d.log("策略初始化")
	d.loadBar(10)
}

func (d *demo) OnStart() {
	d.log("策略启动")
}

func (d *demo) OnStop() {
	d.log("策略停止")
}

func (d *demo) OnBar(line string) {
	// 利用 line 更新指标
	// if ! am.inited {
	//   return
	// }

	// 获取均线值
	d.fastMA0, d.fastMA1 = -1, -2
	d.slowMA0, d.slowMA1 = -1, -2

	// 金叉
	crossOver := d.fastMA1 < d.slowMA1 && d.fastMA0 >= d.slowMA0
	crossBelow := d.fastMA1 > d.slowMA1 && d.fastMA0 <= d.slowMA0

	if crossOver {
		d.log("发生了金叉，准备买入")
		d.pos = -1
		// price := bar.close + 5
		// if d.pos == 0 {
		//   d.buy("梭哈")
		// }else if d.pos < 0 {
		//   d.cover("全部")
		//   d.buy("梭哈")
		// }
	} else if crossBelow {
		d.log("发生了死叉，准备卖出")
		// price := bar.close - 5
		// if d.pos == 0 {
		//   d.short("梭哈")
		// }else if d.pos > 0 {
		//   d.sell("全部")
		//   d.short("梭哈")
		// }
	}
	d.putEvent()
}

// 通知界面更新
func (d *demo) putEvent() {
}
