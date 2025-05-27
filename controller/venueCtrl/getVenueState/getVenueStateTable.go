package getVenueState

import (
	"Gym_booking_WeChat_mini_program/controller/getDBFromContext"
	"Gym_booking_WeChat_mini_program/model/venueModel"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

func GetStateTable(c *gin.Context) {
	mysql, err := getDBFromContext.GetGormDBFromContext(c, "readMySQL")
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	venueType := c.Query("venueType")
	//if !checkIfVenueTypeLegal(venueType) {
	//	c.String(http.StatusAccepted, "非法场地类型！")
	//	return
	//}
	vs := venueModel.VenueState{}
	var vsSlice1, vsSlice2 []venueModel.VenueState
	var wg sync.WaitGroup
	wg.Add(2)
	var todayVSTable, tomorrowVSTable [][]bool
	//由于读取出来的数据是结构体数组，且状态都是用0和1来表示，因此要把0转化为false，1转化为true。
	//rawState是经过StructToSlice处理后得到的场地状态二维切片数据，但是还包含id和date字段
	//vsTable是转化成二维切片的状态表，GetVenueStateSlice用来去除id和date字段
	go func() {
		defer wg.Done()
		mysql.Where("venue_type = ? AND date = ?", venueType, "today").Find(&vsSlice1)
		rawState1 := vs.VenueStateStructToSlice(vsSlice1)
		todayVSTable = RemoveVenueTypeAndIdAndDate(rawState1, 13)
	}()
	go func() {
		defer wg.Done()
		mysql.Where("venue_type = ? AND date = ?", venueType, "tomorrow").Find(&vsSlice2)
		rawState2 := vs.VenueStateStructToSlice(vsSlice2)
		tomorrowVSTable = RemoveVenueTypeAndIdAndDate(rawState2, 13)
	}()
	wg.Wait()
	VSTables := map[string][][]bool{
		"today":    todayVSTable,
		"tomorrow": tomorrowVSTable,
	}
	c.JSON(http.StatusOK, VSTables)
}
