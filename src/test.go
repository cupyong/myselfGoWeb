// Auther: Mjy
// Create Date: 2018/2/26
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/json", jsonHandle)
	http.HandleFunc("/str", strHandle)
	http.ListenAndServe(":8080", nil)

}
func jsonHandle(writer http.ResponseWriter, request *http.Request) {
	s := `{"code":0,"data":
             {"list":[{"id":"12319","name":"CCTV-10","program_name":"特警力量(35)","live_share":0.000145929,"live_rating":0.00003,"duration":0,"uv":15},{"id":"12325","name":"CCTV-9","program_name":"寰宇视野","live_share":0.0465026,"live_rating":0.00956,"duration":0,"uv":2957},{"id":"12310","name":"CCTV-4","program_name":"中央电视台春节联欢晚会","live_share":0.000826929,"live_rating":0.00017,"duration":0,"uv":90},{"id":"3224","name":"CCTV-13","program_name":"新闻直播间","live_share":0.0571067,"live_rating":0.01174,"duration":0,"uv":3753},{"id":"12320","name":"CCTV-12","program_name":"父母爱情(44)","live_share":0.00141064,"live_rating":0.00029,"duration":0,"uv":121},{"id":"12311","name":"CCTV-7","program_name":"人与自然","live_share":0.0017025,"live_rating":0.00035,"duration":0,"uv":126},{"id":"3226","name":"CCTV-15","program_name":"《一起音乐吧》新春特别节目","live_share":0.0242728,"live_rating":0.00499,"duration":0,"uv":1396},{"id":"12308","name":"CCTV-5＋","program_name":"2018年平昌冬奥会","live_share":0.0003405,"live_rating":0.00007,"duration":0,"uv":87},{"id":"12303","name":"CCTV-14","program_name":"优秀儿童剧(42)","live_share":0.035412,"live_rating":0.00728,"duration":0,"uv":2299},{"id":"3222","name":"CCTV-11","program_name":"空中剧院大型系列戏曲演唱会","live_share":0.023154,"live_rating":0.00476,"duration":0,"uv":1281},{"id":"12309","name":"CCTV-2","program_name":"中国财经报道","live_share":0.0255861,"live_rating":0.00526,"duration":0,"uv":1706},{"id":"12304","name":"CCTV-1","program_name":"我们的爱(9)","live_share":0.00209164,"live_rating":0.00043,"duration":0,"uv":209},{"id":"3227","name":"CCTV-NEWS","program_name":"今日世界","live_share":0.00418329,"live_rating":0.00086,"duration":0,"uv":310}]},"message":""}`
	var data map[string]interface{}
	json.Unmarshal([]byte(s), &data)
	fmt.Println(data)
	fmt.Println(s)
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))
	fmt.Fprint(writer, string(jsonStr))
}
func strHandle(writer http.ResponseWriter, request *http.Request) {
	s := `{"code":0,"data":{"list":[{"id":"12319","name":"CCTV-10","program_name":"特警力量(35)","live_share":0.000145929,"live_rating":0.00003,"duration":0,"uv":15},{"id":"12325","name":"CCTV-9","program_name":"寰宇视野","live_share":0.0465026,"live_rating":0.00956,"duration":0,"uv":2957},{"id":"12310","name":"CCTV-4","program_name":"中央电视台春节联欢晚会","live_share":0.000826929,"live_rating":0.00017,"duration":0,"uv":90},{"id":"3224","name":"CCTV-13","program_name":"新闻直播间","live_share":0.0571067,"live_rating":0.01174,"duration":0,"uv":3753},{"id":"12320","name":"CCTV-12","program_name":"父母爱情(44)","live_share":0.00141064,"live_rating":0.00029,"duration":0,"uv":121},{"id":"12311","name":"CCTV-7","program_name":"人与自然","live_share":0.0017025,"live_rating":0.00035,"duration":0,"uv":126},{"id":"3226","name":"CCTV-15","program_name":"《一起音乐吧》新春特别节目","live_share":0.0242728,"live_rating":0.00499,"duration":0,"uv":1396},{"id":"12308","name":"CCTV-5＋","program_name":"2018年平昌冬奥会","live_share":0.0003405,"live_rating":0.00007,"duration":0,"uv":87},{"id":"12303","name":"CCTV-14","program_name":"优秀儿童剧(42)","live_share":0.035412,"live_rating":0.00728,"duration":0,"uv":2299},{"id":"3222","name":"CCTV-11","program_name":"空中剧院大型系列戏曲演唱会","live_share":0.023154,"live_rating":0.00476,"duration":0,"uv":1281},{"id":"12309","name":"CCTV-2","program_name":"中国财经报道","live_share":0.0255861,"live_rating":0.00526,"duration":0,"uv":1706},{"id":"12304","name":"CCTV-1","program_name":"我们的爱(9)","live_share":0.00209164,"live_rating":0.00043,"duration":0,"uv":209},{"id":"3227","name":"CCTV-NEWS","program_name":"今日世界","live_share":0.00418329,"live_rating":0.00086,"duration":0,"uv":310}]},"message":""}`
	fmt.Fprint(writer, s)
}