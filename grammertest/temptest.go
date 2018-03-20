package main

import "fmt"

/*

  /* 用sprintf函数将x和message输出到字符数组buffer中
	sprintf(buffer, "%s%d", message, x);
 */
type DeviceReportDataRequest struct {
	ProductCode string       `json:"product_code" description:"产品代码, 如 A3161, Z6111"`
	SN          string       `json:"sn" description:"设备SN"`
	Events      []*DataEvent `json:"events" description:"设备端增量累计的所有事件数据"`
}

type DataEvent struct {
	Type      string `json:"type" required:"true" description:"数据事件类型，比如PLAY、PAUSE等，所有支持的类型参见文档 - 埋点追踪事件详解"`
	Value     int    `json:"value" required:"true" description:"对应事件类型发生的次数或者其他具体值"`
	LocalTime string `json:"local_time,omitempty" required:"false" description:"当地时间Time, 格式为 hh:mm，时间为设备本地的时间"`
}

func getString() string {

	const (
		a=1
		b=2
	)

	return ""
}
func main()  {
	 var getMassage string = getString()
	dr := DeviceReportDataRequest{
		ProductCode: "A3161",
		SN:          "",
		Events: nil,
	}


if len(dr.Events) ==0{
	fmt.Println(true)
}


	 fmt.Println(getMassage)
}
