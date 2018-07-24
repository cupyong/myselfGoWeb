package router

func Head(pathExp string, handlerFunc HandlerFunc, values ...interface{}) *Route {
	return HttpCommonMethod("HEAD", pathExp, handlerFunc, values)
}
func Get(pathExp string, handlerFunc HandlerFunc, values ...interface{}) *Route {
	return HttpCommonMethod("GET", pathExp, handlerFunc, values)
}

func Post(pathExp string, handlerFunc HandlerFunc, values ...interface{}) *Route {
	return HttpCommonMethod("POST", pathExp, handlerFunc, values)
}
func Put(pathExp string, handlerFunc HandlerFunc, values ...interface{}) *Route {
	return HttpCommonMethod("PUT", pathExp, handlerFunc, values)
}

func Patch(pathExp string, handlerFunc HandlerFunc, values ...interface{}) *Route {
	return HttpCommonMethod("PATCH", pathExp, handlerFunc, values)
}

func Delete(pathExp string, handlerFunc HandlerFunc, values ...interface{}) *Route {
	return HttpCommonMethod("DELETE", pathExp, handlerFunc, values)
}
func Options(pathExp string, handlerFunc HandlerFunc, values ...interface{}) *Route {
	return HttpCommonMethod("OPTIONS", pathExp, handlerFunc, values)
}

func HttpCommonMethod(httpMethod, pathExp string, handlerFunc HandlerFunc, values []interface{}) *Route {
	var option = HttpOption{}
	if len(values) > 0 {
		option = values[0].(HttpOption)
	}
	return &Route{
		HttpMethod: httpMethod,
		PathExp:    pathExp,
		httpOption: option,
		Func:       handlerFunc,
	}
}

func Use(pathExp string, childrenFunc ChildrenFunc) *Route {
	return &Route{
		HttpMethod: "Use",
		PathExp:    pathExp,
		Children:   childrenFunc,
	}
}
