package keyboard

import "github.com/PaulSonOfLars/gotgbot/v2"

type keyboard [][]gotgbot.KeyboardButton

type Keyboard struct {
	keyboard
}

func NewKeyboard() *Keyboard {
	return &Keyboard{}
}

func (k *Keyboard) Add(buttons ...gotgbot.KeyboardButton) *Keyboard {
	if len(k.keyboard) == 0 {
		k.keyboard = append(k.keyboard, []gotgbot.KeyboardButton{})
	}

	k.keyboard[len(k.keyboard)-1] = append(k.keyboard[len(k.keyboard)-1], buttons...)
	return k
}

func (k *Keyboard) Row() *Keyboard {
	k.keyboard = append(k.keyboard, []gotgbot.KeyboardButton{})
	return k
}

func (k *Keyboard) Text(text string) *Keyboard {
	k.Add(gotgbot.KeyboardButton{Text: text})
	return k
}

func (k *Keyboard) RequestContact(text string) *Keyboard {
	k.Add(gotgbot.KeyboardButton{Text: text, RequestContact: true})
	return k
}
func (k *Keyboard) RequestLocation(text string) *Keyboard {
	k.Add(gotgbot.KeyboardButton{Text: text, RequestLocation: true})
	return k
}

func (k *Keyboard) RequestPoll(text string, pollType string) *Keyboard {
	k.Add(gotgbot.KeyboardButton{Text: text, RequestPoll: &gotgbot.KeyboardButtonPollType{Type: pollType}})
	return k
}

func (k *Keyboard) Build() *gotgbot.ReplyKeyboardMarkup {
	return &gotgbot.ReplyKeyboardMarkup{
		Keyboard: k.keyboard,
	}
}
