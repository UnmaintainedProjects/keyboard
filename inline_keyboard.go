package keyboard

import "github.com/PaulSonOfLars/gotgbot/v2"

type inlineKeyboard [][]gotgbot.InlineKeyboardButton

type InlineKeyboard struct {
	inlineKeyboard
}

// Adds a button.
func (ik *InlineKeyboard) Add(buttons ...gotgbot.InlineKeyboardButton) *InlineKeyboard {
	if len(ik.inlineKeyboard) == 0 {
		ik.inlineKeyboard = append(ik.inlineKeyboard, []gotgbot.InlineKeyboardButton{})
	}

	ik.inlineKeyboard[len(ik.inlineKeyboard)-1] = append(ik.inlineKeyboard[len(ik.inlineKeyboard)-1], buttons...)
	return ik
}

// Creates a row for the upcoming buttons.
func (ik *InlineKeyboard) Row() *InlineKeyboard {
	ik.inlineKeyboard = append(ik.inlineKeyboard, []gotgbot.InlineKeyboardButton{})
	return ik
}

// Adds a url button.
func (ik *InlineKeyboard) Url(text string, url string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, Url: url})
	return ik
}

// Adds a login button.
func (ik *InlineKeyboard) Login(text string, url string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, LoginUrl: &gotgbot.LoginUrl{Url: url}})
	return ik
}

// Adds a callback button.
func (ik *InlineKeyboard) Text(text string, data string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, CallbackData: data})
	return ik
}

// Adds a button which opens a chat selection dialog for the user to open inline mode in.
func (ik *InlineKeyboard) SwitchInline(text string, query *string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, SwitchInlineQuery: query})
	return ik
}

// Adds a button which opens inline button in the current chat.
func (ik *InlineKeyboard) SwitchInlineCurrent(text string, query *string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, SwitchInlineQueryCurrentChat: query})
	return ik
}

// Adds a game button.
func (ik *InlineKeyboard) Game(text string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, CallbackGame: &gotgbot.CallbackGame{}})
	return ik
}

// Adds a pay button.
func (ik *InlineKeyboard) Pay(text string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, Pay: true})
	return ik
}

// Builds the inline keyboard markup. This should be called in the end only.
func (ik *InlineKeyboard) Build() *gotgbot.InlineKeyboardMarkup {
	return &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: ik.inlineKeyboard,
	}
}
