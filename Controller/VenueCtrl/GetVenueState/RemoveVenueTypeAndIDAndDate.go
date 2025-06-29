package GetVenueState

func RemoveVenueTypeAndIdAndDate(rawState [][]byte, timeNum int) [][]bool {
	n := len(rawState)
	vsTable := make([][]bool, n)
	for i := range vsTable {
		//vsTable[i]是每个场地不同时间的状态，也就是数据库里的一行，预约界面里的一列
		vsTable[i] = make([]bool, timeNum)
	}
	//vs := venueModel.VenueState{}
	//venueType、id和date会各占据16、8和16位,m=40
	//m := int(unsafe.Sizeof(vs.VenueType) + unsafe.Sizeof(vs.Date) + unsafe.Sizeof(vs.ID))
	//fmt.Println("m:", m)
	for i := 0; i < n; i++ {
		//因为venueType、id和date会各占据16、8和16位，后面的才是状态，因此j从40开始
		for j := 40; j < 40+timeNum; j++ {
			if rawState[i][j] == 1 {
				vsTable[i][j-40] = true
			} else {
				vsTable[i][j-40] = false
			}
		}
	}
	return vsTable
}
