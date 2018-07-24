package common

import (
    "bufio"
    "io"
    "os"
    "strings"
    "fmt"
)

var (
    All_port                   string
    UseCenterHost              string
    SQL_CONNECT                string
    THIRD_SQL_CONNECT          string
    DATABASE_TYPE              string
    GLOBAL_PAGESIZE            string
    GLOBAL_PAGENUM             string
    GLOBAL_TOP                 string
    CSV_CODING                 string
    TASK_DEFAULT_PRIORITY      string
    PromotionProductPriceRange string
    SAMPLE_LIB_DIR             string
    SAMPLE_LIB_DIR_DELETED     string
    CACHE_CLEAN_MINUTE         string
    CorpID                     string
    Corpsecret                 string
)

func InitConfig(configPath string) {
    configMap := make(map[string]string)
    configFile, err := os.Open(configPath)
    if err != nil {
        fmt.Println("configInit error")
        os.Exit(0)
    }
    reader := bufio.NewReader(configFile)
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
        line = strings.TrimSpace(strings.Split(line, "#")[0])
        equalIndex := strings.Index(line, "=")
        if equalIndex < 0 {
            continue
        }
        key := line[0:equalIndex]
        value := strings.Replace(strings.TrimSpace(line[equalIndex+1:]), "\"", "", -1)
        configMap[key] = value
    }
    All_port = configMap["All_port"]
    UseCenterHost = configMap["UseCenterHost"]
    SQL_CONNECT = configMap["SQL_CONNECT"]
    DATABASE_TYPE = configMap["DATABASE_TYPE"]
    GLOBAL_PAGESIZE = configMap["GLOBAL_PAGESIZE"]
    GLOBAL_PAGENUM = configMap["GLOBAL_PAGENUM"]
    GLOBAL_TOP = configMap["GLOBAL_TOP"]
    CSV_CODING = configMap["CSV_CODING"]
    TASK_DEFAULT_PRIORITY = configMap["TASK_DEFAULT_PRIORITY"]
    PromotionProductPriceRange = configMap["PromotionProductPriceRange"]
    THIRD_SQL_CONNECT = configMap["THIRD_SQL_CONNECT"]
    SAMPLE_LIB_DIR = configMap["SAMPLE_LIB_DIR"]
    SAMPLE_LIB_DIR_DELETED = configMap["SAMPLE_LIB_DIR_DELETED"]
    CACHE_CLEAN_MINUTE = configMap["CACHE_CLEAN_MINUTE"]
    CorpID = configMap["CorpID"]
    Corpsecret = configMap["Corpsecret"]
}
