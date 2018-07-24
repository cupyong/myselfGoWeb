package router

import (
	"../common"
	"net/http"
	"fmt"
	"./middle"
	"strings"
	"time"
	"encoding/json"
	"bytes"
	"strconv"
	"github.com/tealeg/xlsx"
	"../common/log"
)

type HttpOption struct {
	Permission string //接口权限
	Cache      bool   //是否设置缓存
	ExportType string //导出类型 （execl csv等）
	ExportExeclModel ExportExeclModel
}
type ExportExeclModel struct {
	FileName string //导出文件名
	Title    string //第一行标题
}

func setRestApiHeader(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,DEL")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length, Authorization, Accept,X-Requested-With")
}

func setExeclHeader(res http.ResponseWriter, filename string) {
	res.Header().Set("Content-Type", "vnd.ms-excel")
	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", filename))
}

type OptionsModel struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type router struct {
	Routes []*Route
}

func MakeRouter(routes ...*Route) (*Trie) {
	time.Now().Unix()
	r := &router{
		Routes: routes,
	}
	trie := New()
	for _, route := range r.Routes {
		//var a = fmt.Sprintf("/:data%s", route.PathExp)
		//fmt.Println(route.HttpMethod, a)
		if route.HttpMethod == "Use" {
			var list = make([]*Route, 0)
			childRouter := MakeChildRouter(route.PathExp, route, list)
			for _, route1 := range childRouter {
				trie.AddRoute(route1.HttpMethod, route1.PathExp, route1)
			}

		} else {
			trie.AddRoute(route.HttpMethod, route.PathExp, route)
		}

	}
	trie.Compress()
	return trie;
}

func MakeDicRouter(routes ...*Route) []*Route {
	var list = make([]*Route, 0)

	r := &router{
		Routes: routes,
	}
	for _, route := range r.Routes {
		list = append(list, route)
	}
	return list;
}

func MakeChildRouter(path string, route *Route, list []*Route) []*Route {
	oldRouter := route.Children()
	for _, route := range oldRouter {
		route.PathExp = path + route.PathExp
		if route.HttpMethod == "Use" {
			list = MakeChildRouter(route.PathExp, route, list)
		} else {
			list = append(list, route)
		}
	}
	return list
}

func HttpService(tries *Trie) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		setRestApiHeader(res)
		s := strings.Split(req.RequestURI, "?")
		matches := tries.FindRoutes(req.Method, s[0])
		request := &common.Request{
			req,
			nil,
			map[string]interface{}{},
			nil,
			nil,
		}
		writer := &responseWriter{
			res,
			false,
		}
		if (req.Method == "OPTIONS") {
			writer.WriteJson(OptionsModel{
				Code: "0",
				Name: "ok",
			}, );
			return
		}
		if len(matches) != 1 {
			common.NotFound(writer, request)
			return
		}


		//异常处理
		defer func() {
			if e := recover(); e != nil {
				log.Log(log.Err, "Panicing %s\r\n", e)
				writer.WriteJsonPermission(OptionsModel{
					Code: "-1",
					Name:fmt.Sprintf("api接口错误%s", e),
				},);
			}
		}()

		httpOption := matches[0].Route.(*Route).httpOption
		ifPermission := middle.GetPermission(writer, request, httpOption.Permission)
		if (!ifPermission) {
			writer.WriteJsonPermission(OptionsModel{
				Code: "-100",
				Name: "ok",
			}, );
			return
			//common.NotPermission(writer, request)
			//return
		}

		switch httpOption.ExportType {
		case "execl":
			setExeclHeader(res,httpOption.ExportExeclModel.FileName)
			ExeclExport(request, writer, matches[0],httpOption.ExportExeclModel);
			break
		case "csv":
			CsvExport(request, writer, matches[0]);
			break
		default:
			RestApi(request, writer, matches[0])
			break;
		}
	})
}

//restApi
func RestApi(request *common.Request, writer *responseWriter, matche *Match) {
	//permission验证 todo
	httpOption := matche.Route.(*Route).httpOption
	//判断接口是否读取缓存
	if httpOption.Cache {
		cacheValue, error := common.GetCache(request.RequestURI)
		if error != nil {
			fmt.Println("redis set failed:", error)
		} else {
			cacheValue = fmt.Sprintf(`{"result":%s}`, cacheValue)
			var data map[string]interface{}
			json.Unmarshal([]byte(cacheValue), &data)
			fmt.Println(cacheValue, data)
			writer.WriteJson(data["result"])
			return
		}
	}
	request.PathParams = matche.Params
	thisRoute := matche.Route.(*Route)
	var result, _ = thisRoute.Func(writer, request)
	if httpOption.Cache {
		cacheString, _ := json.Marshal(result)
		common.SetCache(request.RequestURI, string(cacheString))
		fmt.Println(string(cacheString))
	}
	writer.WriteJson(result)
}

//execl
func ExeclExport(request *common.Request, writer *responseWriter, matche *Match,execlModel ExportExeclModel) {
	request.PathParams = matche.Params
	thisRoute := matche.Route.(*Route)
	var exportString, _ = thisRoute.Func(writer, request)

	buff := bytes.Buffer{}
	buff.WriteString(execlModel.Title + "\n")
	buff.WriteString(exportString.(string))

	xlsxFile := xlsx.NewFile()
	sheet, _ := xlsxFile.AddSheet("Sheet1")
	lines := strings.Split(buff.String(), "\n")
	for _, line := range lines {
		row := sheet.AddRow()
		values := strings.Split(line, ",")
		for _, value := range values {
			cell := row.AddCell()
			valueFloat, err := strconv.ParseFloat(value, 64)
			if err != nil || len(value) > 15 {
				cell.Value = value
			} else {
				cell.SetFloat(valueFloat)
			}
		}
	}
	xlsxFile.Write(writer)
}

//csv
func CsvExport(request *common.Request, writer *responseWriter, matche *Match) {

}
