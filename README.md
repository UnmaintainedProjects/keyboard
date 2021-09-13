# gotgbot keyboard

Build reply markup keyboards easier than ever.

## Installation

```bash
go get github.com/gotgbot/keyboard
```

## Docs

[go.dev](https://pkg.go.dev/github.com/gotgbot/keyboard)

## Example

### Inline

```go
import "github.com/gotgbot/keyboard"

...

query := "query"

ctx.Message.Reply(
    b,
    "text",
    &gotgbot.SendMessageOpts{
        ReplyMarkup: new(
            keyboard.InlineKeyboard,
        ).Text(
            "text",
            "callback data",
        ).Url(
            "text",
            "https://github.com/gotgbot/keyboard",
        ).Row(
        ).SwitchInline(
            "text",
            &query,
        ).Build()
    }
)
```

### Reply

```go
import "github.com/gotgbot/keyboard"

...

ctx.Message.Reply(
    b,
    "text",
    &gotgbot.SendMessageOpts{
        ReplyMarkup: new(
            keyboard.Keyboard,
        ).Text(
            "text",
        ).RequestContact(
            "text",
        ).Row(
        ).RequestPoll(
            "text",
            "type",
        ).Build()
    }
)
```

## Credits

- [grammY framework](https://github.com/grammyjs/grammY)
