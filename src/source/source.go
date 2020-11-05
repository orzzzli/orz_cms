package source

type Source interface {
	Connect()         //连接
	Ping()            //判活
	Get() interface{} //返回格式化后的数据
	Set(interface{})  //设置格式化后的数据
	format()          //格式化数据
}
