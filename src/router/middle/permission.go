package middle

import (
	"../../common"
	"fmt"
	"../../model"
	"../../common/log"
)

func GetPermission(w common.ResponseWriter, r *common.Request, permission string) bool {
	r.ParseForm()

    permissionModel := &model.Permission{

    }

    //业务

    permissionModel.Userid="aaa"
    permissionModel.Datanote=1
    permissionModel.Adnote=1

    r.Permission= permissionModel
	var per = false;
	if permission == "login" {
		token := r.Form.Get("token")

		var sqlStr = fmt.Sprintf(`
        select id,username from admin where token = '%s'`, token)
		rows, _ := common.SqlReader(sqlStr)
		var id,username string;
		for rows.Next() {
			rows.Scan(&id, &username)
			per = true
		}
		log.Log(log.Info, "[%s][%s]", username, r.RequestURI)
	}else {
		log.Log(log.Info, "[%s][%s]", "notLogin", r.RequestURI)
		return true
	}

	//fmt.Println(token)
	//permissionModel := &model.Permission{
	//
	//}
	//
	////业务
	//
	//permissionModel.Userid="aaa"
	//permissionModel.Datanote=1
	//permissionModel.Adnote=1
	//
	//r.Permission= permissionModel

	return per

}
