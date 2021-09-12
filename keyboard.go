package keyboard

import "github.com/PaulSonOfLars/gotgbot/v2"

type keyboard [][]gotgbot.KeyboardButton

type Keyboard struct {
	keyboard
}

// Adds a button.
func (k *Keyboard) Add(buttons ...gotgbot.KeyboardButton) *Keyboard {
	if len(k.keyboard) == 0 {
		k.keyboard = append(k.keyboard, []gotgbot.KeyboardButton{})
	}

	k.keyboard[len(k.keyboard)-1] = append(k.keyboard[len(k.keyboard)-1], buttons...)
	return k
}

// Creates a row for the upcoming buttons.
func (k *Keyboard) Row() *Keyboard {
	k.keyboard = append(k.keyboard, []gotgbot.KeyboardButton{})
	return k
}

// Adds a text button.
func (k *Keyboard) Text(text string) *Keyboard {
	k.Add(gotgbot.KeyboardButton{Text: text})
	return k
}

// Adds a request contact button.
func (k *Keyboard) RequestContact(text string) *Keyboard {
	k.Add(gotgbot.KeyboardButton{Text: text, RequestContact: true})
	return k
}

// Adds a request location button.
func (k *Keyboard) RequestLocation(text string) *Keyboard {
	k.Add(gotgbot.KeyboardButton{Text: text, RequestLocation: true})
	return k
}

// Adds a request poll button.
func (k *Keyboard) RequestPoll(text string, pollType string) *Keyboard {
	k.Add(gotgbot.KeyboardButton{Text: text, RequestPoll: &gotgbot.KeyboardButtonPollType{Type: pollType}})
	return k
}

// Builds the reply keyboard markup. This should be called in the end only.
func (k *Keyboard) Build() *gotgbot.ReplyKeyboardMarkup {
	return &gotgbot.ReplyKeyboardMarkup{
		Keyboard: k.keyboard,
	}
}
