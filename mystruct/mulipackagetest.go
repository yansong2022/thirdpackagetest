package mystruct

import "fmt"

// 自定义的服务结构
type Myservice struct {
	SERVICE_NAME  string
	SERVICE_OWNER string
	SERVICE_COUNT int
}

/*
实例化一个服务
参数: 无
返回值： 生成的服务器实例
*/
func NewService(serviceName string, ownerName string, serviceCount int) *Myservice {
	retData := new(Myservice)
	retData.SERVICE_NAME = serviceName
	retData.SERVICE_OWNER = ownerName
	retData.SERVICE_COUNT = serviceCount

	return retData
}

/*
打印实例内容
参数: 无
返回值: 返回打印的实例内容
*/
func (myservice Myservice) PrintService() string {
	fmt.Println(myservice.SERVICE_NAME)
	return myservice.SERVICE_NAME
}
