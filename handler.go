package main

import (
	tele "github.com/3JoB/telebot"
	tb "github.com/3JoB/ulib/telebot/utils"
)

var (
	btn = &tele.ReplyMarkup{}

	alertbtn = btn.Data("Show Alert", "tele_show")
	alert2btn = btn.Data("Show Next Alert", "tele_show2")
)

func showBtn() *tele.ReplyMarkup {
	btn.Row(alertbtn, alert2btn)
	return btn
}

func Handler(b *tele.Bot) {
	b.Handle("/start", Start)
	b.Handle("/autodelete", AutoStart)
	b.Handle(&alertbtn, Alert1)
	b.Handle(&alert2btn, Alert2)
}

func Start(c tele.Context) error {
	c.Send("Hello World!")
	return nil
}

func AutoStart(c tele.Context) error {
	use := tb.New().SetContext(c)
	//Messages will be automatically deleted after ten seconds.
	use.SetAutoDelete(10).Send("Hello! Telebot!")
	return nil
}

func SendBtn(c tele.Context) error {
	use := tb.New().SetContext(c)
	use.SetBtn(showBtn()).Send("choose a button")
	return nil
}

func Alert1(c tele.Context) error {
	use := tb.New().SetContext(c)
	use.Alert("Wow! Alert1!")
	return nil
}

func Alert2(c tele.Context) error {
	use := tb.New().SetContext(c)
	use.SetShowAlert().Alert("Woooooooooooooooooooooww! Alert2!")
	return nil
}