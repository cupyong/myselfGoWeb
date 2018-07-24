package common

import (
    "./log"
    "fmt"
    "net/http"
    "encoding/csv"
    "github.com/360EntSecGroup-Skylar/excelize"
    //"golang.org/x/text/encoding/simplifiedchinese"
    //"golang.org/x/text/encoding/traditionalchinese"
    //"golang.org/x/text/transform"
    "io"
    "mime/multipart"
    //"bytes"
    //"io/ioutil"
    //"golang.org/x/text/transform"
    //"database/sql"
    //"time"
)

//
//func GbkToUtf8(s []byte) ([]byte, error) {
//	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
//	d, e := ioutil.ReadAll(reader)
//	if e != nil {
//		return nil, e
//	}
//	return d, nil
//}
//
//func Utf8ToGbk(s []byte) ([]byte, error) {
//	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
//	d, e := ioutil.ReadAll(reader)
//	if e != nil {
//		return nil, e
//	}
//	return d, nil
//}
//
//func Decode(s []byte) ([]byte, error) {
//	I := bytes.NewReader(s)
//	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
//	d, e := ioutil.ReadAll(O)
//	if e != nil {
//		return nil, e
//	}
//	return d, nil
//}

func CheckErr(errMasg error) {
    if errMasg != nil {
        log.Log(log.Err, "[%s][%d]", errMasg, log.File_line())
    }
}
func Service(port string) {
    bind := fmt.Sprintf("%s:%s", "0.0.0.0", port)
    http.ListenAndServe(bind, nil)
    fmt.Println(port)
}
func CSVtoList(file multipart.File) [][]string {
    defer file.Close()
    reader := csv.NewReader(file)
    var test = make([][]string, 0)
    for {
        record, err := reader.Read()

        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
        }
        //enc := mahonia.NewDecoder(CSV_CODING)
        //recordStr := strings.Join(record,",")
        //recordStrUTF8 := enc.ConvertString(recordStr)
        //recordListUTF8:= strings.Split(recordStrUTF8,",")

        test = append(test, record)
    }
    return test
    //fmt.Println(test)
    //fmt.Println("上传文件的大小为: %d", file.(Sizer).Size())
}

func ExcelToList(file multipart.File) [][]string {
    defer file.Close()
    xlsx, err := excelize.OpenReader(file)
    if err != nil {
        fmt.Println(err)
        return [][]string{}
    }
    rows := xlsx.GetRows(xlsx.GetSheetName(1))
    return rows
}
func Substr(str string, start int, length int) string {
    rs := []rune(str)
    rl := len(rs)
    end := 0

    if start < 0 {
        start = rl - 1 + start
    }
    end = start + length

    if start > end {
        start, end = end, start
    }

    if start < 0 {
        start = 0
    }
    if start > rl {
        start = rl
    }
    if end < 0 {
        end = 0
    }
    if end > rl {
        end = rl
    }

    return string(rs[start:end])
}

//func GetProgramDic() (map[string]string, error) {
//    if value, ok := Memory2.Load("programDic"); ok {
//        fmt.Println("use memory programDic")
//        return value.(MemoryValue).Result.(map[string]string), nil
//    }
//    fmt.Println("use sql programDic")
//    programDic := make(map[string]string)
//    sqlStr := `SELECT DISTINCT programId, programName FROM program_dic `
//    //sqlStr := `SELECT programId, programName FROM program_dic limit 10`
//    fmt.Println(sqlStr)
//    rows, err := SqlReader(sqlStr)
//    if err != nil {
//        fmt.Println(log.Err)
//        log.Log(log.Err, "[%s][%s]", err, log.File_line())
//        return map[string]string{}, err
//    }
//    var id, name sql.NullString
//    for rows.Next() {
//        rows.Scan(&id, &name)
//        programDic[id.String] = name.String
//    }
//    rows.Close()
//    Memory2.Store("programDic", MemoryValue{
//        Time:   time.Now().Unix() + (3600 * 24),
//        Result: programDic,
//    })
//    return programDic, nil
//}
