package main

import (
	ob "github.com/botuniverse/go-libonebot"
	log "github.com/sirupsen/logrus"
)

type OneBotDummy struct {
	*ob.OneBot
}

func main() {
	log.SetLevel(log.DebugLevel)

	msg := ob.Message{
		ob.TextSegment("1"),
		ob.TextSegment("2"),
		ob.MentionSegment("test"),
		ob.MentionSegment("test"),
		ob.TextSegment("3"),
		ob.TextSegment("4"),
		ob.TextSegment("5"),
		ob.MentionSegment("test"),
	}
	msg.Reduce()
	log.Debugf("%v", msg)
	return

	obdummy := &OneBotDummy{OneBot: ob.NewOneBot("dummy")}

	obdummy.HandleFunc(ob.ActionGetVersion, func(w ob.ResponseWriter, r *ob.Request) {
		w.WriteData(map[string]string{
			"version":         "1.0.0",
			"onebot_standard": "v12",
		})
	})

	obdummy.HandleFunc(ob.ActionSendMessage, func(w ob.ResponseWriter, r *ob.Request) {
		p := ob.NewParamGetter(r.Params, w)
		userID, ok := p.GetString("user_id")
		if !ok {
			return
		}
		msg, ok := p.GetMessage("message")
		if !ok {
			return
		}
		log.Debugf("Send message: %#v, to %v", msg, userID)
		w.WriteData(msg)
	})

	obdummy.HandleFuncExtended("do_something", func(w ob.ResponseWriter, r *ob.Request) {
	})

	// go func() {
	// 	for {
	// 		obdummy.Push(
	// 			&ob.MessageEvent{
	// 				Event: ob.Event{
	// 					SelfID:     "123",
	// 					Type:       ob.EventTypeMessage,
	// 					DetailType: "private",
	// 				},
	// 				UserID:  "234",
	// 				Message: ob.Message{ob.TextSegment("hello")},
	// 			},
	// 		)
	// 		time.Sleep(time.Duration(3) * time.Second)
	// 	}
	// }()

	// go func() {
	// 	time.Sleep(time.Duration(10) * time.Second)
	// 	obdummy.Shutdown()
	// }()

	obdummy.Run()
}
