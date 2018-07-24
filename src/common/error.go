package common

func ErrorDict() func(int) string {
	innerMap := map[int]string{
		//系统基本错误 10 000 端
		//数据库错误 20 000
		//用户权限错误 30 000
		//业务1 错误 60 000
		//业务2 错误 70 000
		//业务3 错误 80 000
		10001: "错误1",
		20001: "错误2",
		30001: "登录失败",
		30002: "文件上传错误",
	}
	return func(key int) string {
		return innerMap[key]
	}
}
