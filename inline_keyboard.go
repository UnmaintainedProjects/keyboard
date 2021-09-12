package keyboard

import "github.com/PaulSonOfLars/gotgbot/v2"

type inlineKeyboard [][]gotgbot.InlineKeyboardButton

type InlineKeyboard struct {
	inlineKeyboard
}

func NewInlineKeyboard() *InlineKeyboard {
	return &InlineKeyboard{}
}

func (ik *InlineKeyboard) Add(buttons ...gotgbot.InlineKeyboardButton) *InlineKeyboard {
	if len(ik.inlineKeyboard) == 0 {
		ik.inlineKeyboard = append(ik.inlineKeyboard, []gotgbot.InlineKeyboardButton{})
	}

	ik.inlineKeyboard[len(ik.inlineKeyboard)-1] = append(ik.inlineKeyboard[len(ik.inlineKeyboard)-1], buttons...)
	return ik
}

func (ik *InlineKeyboard) Row() *InlineKeyboard {
	ik.inlineKeyboard = append(ik.inlineKeyboard, []gotgbot.InlineKeyboardButton{})
	return ik
}

func (ik *InlineKeyboard) Url(text string, url string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, Url: url})
	return ik
}

func (ik *InlineKeyboard) Login(text string, url string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, LoginUrl: &gotgbot.LoginUrl{Url: url}})
	return ik
}

func (ik *InlineKeyboard) Text(text string, data string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, CallbackData: data})
	return ik
}

func (ik *InlineKeyboard) SwitchInline(text string, query *string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, SwitchInlineQuery: query})
	return ik
}

func (ik *InlineKeyboard) SwitchInlineCurrent(text string, query *string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, SwitchInlineQueryCurrentChat: query})
	return ik
}

func (ik *InlineKeyboard) Game(text string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, CallbackGame: &gotgbot.CallbackGame{}})
	return ik
}

func (ik *InlineKeyboard) Pay(text string) *InlineKeyboard {
	ik.Add(gotgbot.InlineKeyboardButton{Text: text, Pay: true})
	return ik
}

func (ik *InlineKeyboard) Build() *gotgbot.InlineKeyboardMarkup {
	return &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: ik.inlineKeyboard,
	}
}
