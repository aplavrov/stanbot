package telegram

import (
	"fmt"
	"log"
	"stanBot/internal/ads"
	"stanBot/internal/client/telegram"
	"stanBot/internal/event"
	"stanBot/internal/storage"
	"time"
)

type Processor struct {
	tg      *telegramClient.Client
	offset  int
	storage storage.Storage
	ads     *ads.Fetcher
}

type Meta struct {
	ChatID   int
	Username string
}

func New(tg *telegramClient.Client, storage storage.Storage, fetcher *ads.Fetcher) *Processor {
	return &Processor{
		tg:      tg,
		storage: storage,
		ads:     fetcher,
	}
}

func (p *Processor) StartGetter() {
	go p.GetAds()
}

const (
	timeBeforeNextUpdate = 5 * time.Minute
	timeNewAds           = 310 * time.Second
)

func (p *Processor) GetAds() {
	log.Print("started ads getter")
	for {
		log.Print("ads getter is getting")
		res := p.ads.GetAll()
		for _, val := range res {
			log.Printf("current time: %v", time.Now().UTC())
			log.Printf("check ads time: %v", val.FirstPublished.Add(2*time.Hour))
			if val.FirstPublished.Add(2 * time.Hour).Before(time.Now().UTC().Add(-timeNewAds)) {
				break
			}
			ids, _ := p.storage.GetAll()
			for _, id := range ids {
				p.sendProperty(id, val)
			}
		}

		time.Sleep(timeBeforeNextUpdate)
	}
}

func (p *Processor) Fetch(limit int) ([]event.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]event.Event, 0, len(updates))
	for _, update := range updates {
		res = append(res, convertToEvent(update))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *Processor) Process(currEvent event.Event) error {
	switch currEvent.Type {
	case event.Message:
		return p.processMessage(currEvent)
	default:
		return fmt.Errorf("can't process message")
	}
}

func (p *Processor) processMessage(event event.Event) error {
	meta, err := meta(event)
	if err != nil {
		return err
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
		return err
	}

	return nil
}

func meta(event event.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, fmt.Errorf("can't get meta")
	}

	return res, nil
}

func convertToEvent(update telegramClient.Update) event.Event {
	updateType := fetchType(update)

	res := event.Event{
		Text: fetchText(update),
		Type: updateType,
	}

	if updateType == event.Message {
		res.Meta = Meta{
			ChatID:   update.Message.Chat.ID,
			Username: update.Message.From.Username,
		}
	}

	return res
}

func fetchType(update telegramClient.Update) event.Type {
	if update.Message == nil {
		return event.Unknown
	}
	return event.Message
}

func fetchText(update telegramClient.Update) string {
	if update.Message == nil {
		return ""
	}
	return update.Message.Text
}
