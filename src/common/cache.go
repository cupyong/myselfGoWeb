package common

//import (
//    "time"
//    "strconv"
//    "sync"
//    "fmt"
//)
//
//type MemoryValue struct {
//    Time   int64
//    Result Result
//}
//
////var Memory = make(map[string]MemoryValue)
//var Memory2 = sync.Map{}
//var cleanInterval, _ = strconv.ParseInt(CACHE_CLEAN_MINUTE, 10, 64)
//
////func CleanCacheLoop() {
////    for {
////        now := time.Now().Unix()
////        for key, value := range Memory {
////            if now-value.Time > cleanInterval*60 {
////                delete(Memory, key)
////            }
////        }
////        time.Sleep(10 * 60 * time.Second)
////    }
////}
//
//func CleanCacheSafe() {
//    for {
//        fmt.Println("start cleaner")
//        Memory2.Range(eachMap)
//        time.Sleep(10 * 60 * time.Second)
//        hour, _, _ := time.Now().Clock()
//        if hour == 5 {
//            CleanCache("programDic")
//            GetProgramDic()
//        }
//    }
//}
//func eachMap(key, memory interface{}) bool {
//    now := time.Now().Unix()
//    if now-memory.(MemoryValue).Time > cleanInterval*60 {
//        Memory2.Delete(key)
//        fmt.Println("clean map: ", key)
//    }
//    return true
//}
//
//func CleanCache(key string) bool {
//    if key == "" {
//        Memory2 = sync.Map{}
//    } else {
//        Memory2.Delete(key)
//        //delete(Memory, key)
//    }
//    return true
//}
