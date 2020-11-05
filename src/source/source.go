package source

type Source interface {
	Connect(string) error                  //连接
	Ping() error                           //判活
	Get(query string) (interface{}, error) //返回原始数据
	Set(query string) (interface{}, error) //设置格式化后的数据
}
