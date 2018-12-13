package cv

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"io/ioutil"
	"net/http"
)



type User struct{
	Name string `yaml:"name"`
	Age int64 `yaml:"age"`
	Introduce string `yaml:"introduce"`
	Skill string `yaml:"skill"`
	WorkExperence string `yaml:workexperence`


}


func (c *User) getConf() *User {
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("/Users/northstar/Documents/go/src/restful/demo02/conf/cv.yaml")
	fmt.Printf("yaml file: %v",string(yamlFile))
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)
	fmt.Printf("yaml file: %v",c)
	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}
func GetCVDetail() (string){
	var c User
	conf := c.getConf()
	data,err := json.Marshal(conf)
	if err != nil {
		fmt.Println("err:\t", err.Error())
		return err.Error()
	}

	//最终以json格式，输出
	fmt.Println("data:\t", string(data))
	return string(data)
}


func GetFullCV(c *gin.Context) {
	log.Info("get something")

	//c.Get("restful/demo02/front/creative-cv_free_1-0-2/index.html")
	//c.String(http.StatusOK,"\n"+aa.username+"\n"+string(aa.age))

	c.JSON(http.StatusOK,gin.H{"myv":GetCVDetail()})



}
//func HealthCheck(c *gin.Context) {
//	message := "OK"
//	c.String(http.StatusOK,"\n"+message)
//
//}
//
//func DiskCheck(c *gin.Context){
//	u, _:= disk.Usage("/")
//	usedMB := int(u.Used) / MB
//	usedGB := int(u.Used) / GB
//	totalMB := int(u.Total) / MB
//	totalGB := int(u.Total) / GB
//	usedPercent := int(u.UsedPercent)
//
//	status := http.StatusOK
//	text := "OK"
//
//	if usedPercent >= 95 {
//		status = http.StatusOK
//		text = "CRITICAL"
//	} else if usedPercent >= 90 {
//		status = http.StatusTooManyRequests
//		text = "WARNING"
//	}
//	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%",text,usedMB,usedGB,totalMB,totalGB,usedPercent)
//	c.String(status,"\n"+message)
//}
//
//func CPUCheck(c *gin.Context) {
//	cores, _ := cpu.Counts(false)
//	a,_ := load.Avg()
//	l1 := a.Load1
//	l5 := a.Load5
//	l15 := a.Load15
//
//	status := http.StatusOK
//	text := "OK"
//
//	if l5 >= float64(cores-1) {
//		status = http.StatusInternalServerError
//		text = "CRITICAL"
//	}else if l5 >= float64(cores-2) {
//		status = http.StatusTooManyRequests
//		text = "WARNING"
//	}
//
//	message := fmt.Sprintf("%s - Load average: %.2f,%.2f,%.2f | Cores: %d",text,l1,l5,l15,cores)
//	c.String(status,"\n"+message)
//
//}
//
//func RAMCheck(c *gin.Context) {
//	u, _ := mem.VirtualMemory()
//
//	usedMB := int(u.Used) / MB
//	usedGB := int(u.Used) / GB
//	totalMB := int(u.Total) / MB
//	totalGB := int(u.Total) / GB
//	usedPercent := int(u.UsedPercent)
//
//	status := http.StatusOK
//	text := "OK"
//
//	if usedPercent >= 95 {
//		status = http.StatusInternalServerError
//		text = "CRITICAL"
//	} else if usedPercent >= 90 {
//		status = http.StatusTooManyRequests
//		text = "WARNING"
//	}
//
//	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
//	c.String(status, "\n"+message)
//}
