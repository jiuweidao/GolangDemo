package main

import (
	"regexp"
	"fmt"
)

func main()  {
	EVENT_TYPE_SOUNDCORE :=
		"PLAY,PAUSE,VOLUMN_UP,VOLUMN_DOWN,POWER_BTN,BT_SHORT," +
			"BT_LONG,MULTI_BTN_SHORT,MULTI_BTN_LONG,MULTI_BTN_DOUBLE,BASS_UP_ON,BASS_UP_OFF," +
			"MULTI_BTN_LONG,MULTI_BTN_SHORT,AUX_IN_PLAY,PLAY_DURATION,PLUG_TIMES,PLUG_ONE_FIFTH,PLUG_TWO_FIFTH," +
			"PLUG_THREE_FIFTH,PLUG_FOUR_FIFTH,PLUG_FIVE_FIFTH,BT_CONNECT_TIMES,"

	EVENT_TYPE_S2   :=
		"PLAY,PAUSE,VOLUMN_UP,VOLUMN_DOWN,VOLUMN_VALUE,BTN_ACT_GOOGLE_VOICE," +
			"POWER_ON,POWER_OFF,POWER_VALUE,BT_SHORT,BT_LONG,DOLBY_ON," +
			"DOLBY_OFF,MUTEFACTORY_RESET,RESET,AUX_IN,BT_PLAY,NETWORK_PLAY," +
			"AUX_IN_PLAY,VO_ACT_GOOGLE_VOICE,PLAY_DURATION,PLUG_TIMES,PLUG_ONE_FIFTH,PLUG_TWO_FIFTH," +
			"PLUG_THREE_FIFTH,PLUG_FOUR_FIFTH,PLUG_FIVE_FIFTH,UPGRADE_TIME,SUPGRADE_SUCCESS_TIMES,"
			types :="PLAY"
	match, _ := regexp.MatchString(","+types+",", EVENT_TYPE_SOUNDCORE+EVENT_TYPE_S2)
	if match {
		println("true")
	}else {
		println("wrong")
	}


		//这个测试一个字符串是否符合一个表达式。
		match1, _ := regexp.MatchString("\\d{1,2}:\\d{1,2}", "12:34")
		fmt.Println(match1)
	
}