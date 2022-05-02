package main

import (
	"gotrading/app/controllers"
	"gotrading/config"
	"gotrading/utils"
	"log"
)

func main() {
	//df, _ := models.GetAllCandle(config.Config.ProductCode, time.Minute, 365)
	//fmt.Printf("%+v\n", df.OptimizeParams())
	utils.LoggingSettings(config.Config.LogFile)
	//fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}

//func main() {
//	s := models.NewSignalEvents()
//	df, _ := models.GetAllCandle("BTC_JPY", time.Minute, 10)
//	c1 := df.Candles[len(df.Candles)-2]
//	c2 := df.Candles[len(df.Candles)-1]
//	s.Buy("BTC_JPY", c1.Time.UTC(), c1.Close, 0.0001, true)
//	s.Sell("BTC_JPY", c2.Time.UTC(), c2.Close, 0.0001, true)
//	fmt.Println(models.GetSignalEventsByCount(1))
//	fmt.Println(models.GetSignalEventsAfterTime(c1.Time))
//	fmt.Println(s.CollectAfter(time.Now().UTC()))
//	fmt.Println(s.CollectAfter(c1.Time))
//}
