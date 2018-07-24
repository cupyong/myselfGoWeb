package taskModel

// 任务详情
type TaskModel struct {
    TaskId     string  `json:"taskId"`
    TaskName string `json:"taskName"`
    StartDate  string `json:"startDate"`
    EndDate    string `json:"endDate"`
    Type       string `json:"type"`
    Status     string `json:"status"`
    Priority   int    `json:"priority"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
    RelevanceId string `json:"relevanceId"`
    CreatorId  string `json:"creatorId"`
    CreatorName  string `json:"creatorName"`
}

// 任务列表
type TaskModelList struct {
    List         []TaskModel `json:"list"`
    PageSize     string      `json:"page_size"`
    PageNum      string      `json:"page_num"`
    PageTotalNum string      `json:"page_total_num"`
}

func (taskModelList *TaskModelList) AddTaskModel(data TaskModel) {
    taskModelList.List = append(taskModelList.List, data)
}

// 执行状态
type StatusModel struct {
    Status int `json:"status"`
}

// 自定义任务
type CustomModel struct {
    TaskId string `json:"task_id"`
    TaskName string `json:"task_name"`
    Live string `json:"live"`
    Vod string `json:"vod"`
    Order string `json:"order"`
    New string `json:"new"`
    Output int `json:"output"`
    Token string `json:"token"`
}