package hf_go_ctp

type BaseTrade interface {
	ReqConnect(addr string)
	ReqLogin(investor, pwd, broker string)
	Release()

	ReqOrderInsert(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string
	ReqOrderInsertMarket(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string
	ReqOrderInsertFOK(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string
	ReqOrderInsertFAK(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string

	RegOnFrontConnected(on OnFrontConnectedType)
	RegOnRspUserLogin(on OnRspUserLoginType)
	RegOnRtnOrder(on OnRtnOrderType)
	RegOnErrRtnOrder(on OnRtnErrOrderType)
	RegOnErrAction(on OnRtnOrderType)
	RegOnRtnCancel(on OnRtnOrderType)
	RegOnRtnTrade(on OnRtnTradeType)
}
