package main

import (
	"fmt"
	"time"

	"github.com/Advik-B/WhatsHappened/parsing"
)

func main() {
	const chat = `
12/11/25, 10:39 pm - Messages and calls are end-to-end encrypted. Only people in this chat can read, listen to, or share them. Learn more.
13/11/25, 12:41 am - Advik: He thinks it's AI 🤣🤣
13/11/25, 12:42 am - Person2: prank successful👏👏
a
13/11/25, 12:41 am - Advik: He thinks it's AI 🤣🤣
13/11/25, 12:42 am - Person2: prank successful👏👏
13/11/25, 12:42 am - Person2: and he still dont know its my number 🫡
13/11/25, 12:42 am - Advik: Exactly!
13/11/25, 12:43 am - Person2: I WILL TELL HIM TMRO
13/11/25, 12:43 am - Person2: dont say no
13/11/25, 12:43 am - Advik: Just don't say it's not AI
13/11/25, 12:43 am - Person2: okk..
13/11/25, 12:43 am - Person2: is he awake now?
`

	msgs, err := parsing.ParseChat(chat)
	if err != nil {
		panic(err)
	}

	for _, m := range msgs {
		fmt.Printf("[sender=%q] %s → %s\n",
			m.Sender,
			m.Time.Format(time.RFC3339),
			m.Content)
	}

}
