package types

//game action ty
const (
	PBGameActionStart = iota + 1
	PBGameActionContinue
	PBGameActionQuit
	PBGameActionQuery
)

//包的名字可以通过配置文件来配置
//建议用github的组织名称，或者用户名字开头, 再加上自己的插件的名字
//如果发生重名，可以通过配置文件修改这些名字
var (
	PackageName = "chain33.pokerbull"
	JRPCName    = "PokerBull"
	PokerBullX  = "pokerbull"
	ExecerPokerBull  = []byte(PokerBullX)
)

const (
	//查询方法名
	FuncName_QueryGameListByIds           = "QueryGameListByIds"
	FuncName_QueryGameById                = "QueryGameById"
)
