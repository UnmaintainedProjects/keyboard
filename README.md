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

ctx.Message.Reply(
    b,
    &gotgbot.SendMessageOpts{
        ReplyMarkup: new(
            keyboard.InlineKeyboard,
        ).Text(
            "text",
            "callback data",
        ).Url(
            "text",
            "url",
        ).Row(
        ).SwitchInline(
            "text",
            "query",
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
    &gotgbot.SendMessageOpts{
        ReplyMarkup: new(
            keyboard.Keyboard,
        ).Text(
            "text",
        ).RequestContact(
            "text".
        ).Row(
        ).Game(
            "text",
        ).Build()
    }
)
```

## Credits

- [grammY framework](https://github.com/grammyjs/grammY)
