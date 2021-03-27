gonce - mini package implementing singleton.

* **Vendor**: bavix
* **Package**: gonce

### Installation

```bash
go get github.com/bavix/gonce
```

### Get started

let's make a simple singleton object for telegrams

```go
package telegram

import (
	"github.com/bavix/gonce"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var instance *tgbotapi.BotAPI
var once gonce.Once

func GetInstance(token string) *tgbotapi.BotAPI {
	once.Do(func() {
		botApi, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			panic(err)
		}

		instance = botApi
	})

	return instance
}
```

PS, don't forget, singleton is an anti-pattern. It should be used extremely rarely.
