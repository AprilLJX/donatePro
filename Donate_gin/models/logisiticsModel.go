package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetLogisiticsModel(logiId string) (logisticList []map[string]interface{},state interface{}) {
	//logisticID, _ := strconv.Atoi(logiId)

	//todo 获取物流api返回的结果
	// 读取json文件内容 返回字节切片
	bytes,_ := ioutil.ReadFile("./entity/data.json")

	// 将字节切片映射到指定map上
	// key：string类型，value：interface{}  类型能存任何数据类型
	var jsonObj map[string]interface{}
	json.Unmarshal(bytes,&jsonObj)

	state = jsonObj["State"]

	// 将key：info 转换为[]interface{}类型（必须先转换成该类型才能实现range操作）
	traceList := jsonObj["Traces"].([]interface{})

	// 打印对象结构
	fmt.Println(jsonObj)

	// 打印所有学生信息
	for _,item := range traceList {
		oneLogistic := make(map[string]interface{})

		// 同样，遍历的时候也需要做转换
		x := item.(map[string]interface{})

		oneLogistic["AcceptTime"] = x["AcceptTime"]
		oneLogistic["AcceptStation"] = x["AcceptStation"]

		logisticList = append(logisticList, oneLogistic)


		// 更改名字
		//x["name"] = string("xxx")
		//fmt.Println(idx,x["AcceptTime"])
		//fmt.Println(idx,x["AcceptStation"])
		//fmt.Println(idx,x["Remark"])

	}
	//dao.GetLogisiticsDao(logisticID,state)

	return

}

