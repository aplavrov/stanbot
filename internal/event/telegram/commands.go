package telegram

import (
	"fmt"
	"log"
	"net/url"
	"stanBot/internal/ads"
	"strings"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
	LastCmd  = "/last"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	p.saveChat(chatID)

	switch text {
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID, username)
	case LastCmd:
		return p.sendLast(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor) saveChat(chatID int) (err error) {
	return p.storage.Save(chatID)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int, username string) error {
	return p.tg.SendMessage(chatID, msgStart+username+msgHello)
}

func (p *Processor) sendLast(chatID int) error {
	res := p.ads.GetLast()
	resText := fmt.Sprintf("Цена %v евро", res.Price)
	picName := []byte(res.CoverPhoto)
	for i, _ := range picName {
		if picName[i] == ' ' {
			picName[i] = '_'
		}
	}

	photoURL := fmt.Sprintf("https://img.cityexpert.rs/properties/720x/%v000/%v/slike/%v@png", res.UniqueID[:2], res.UniqueID[:5], string(picName))

	log.Print(photoURL)
	return p.tg.SendPhoto(chatID, photoURL, resText)
}

func (p *Processor) sendProperty(chatID int, res ads.Property) error {
	log.Printf("sent property %v to chat %v", res.UniqueID, chatID)
	resText := fmt.Sprintf("Цена %v евро", res.Price)
	picName := []byte(res.CoverPhoto)
	for i, _ := range picName {
		if picName[i] == ' ' {
			picName[i] = '_'
		}
	}

	photoURL := fmt.Sprintf("https://img.cityexpert.rs/properties/720x/%v000/%v/slike/%v@png", res.UniqueID[:2], res.UniqueID[:5], string(picName))

	log.Print(photoURL)
	return p.tg.SendPhoto(chatID, photoURL, resText)
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
