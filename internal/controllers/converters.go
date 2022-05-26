package controllers

import (
	"gitlab.com/alexalreadytaken/notify-service/internal/models/db"
	"gitlab.com/alexalreadytaken/notify-service/internal/models/rest"
)

func dbClientToRest(c *db.Client) rest.Client {
	return rest.Client{
		ID:                 c.ID,
		PhoneNumber:        c.PhoneNumber,
		MobileOperatorCode: c.MobileOperatorCode,
		Tag:                c.Tag,
		Timezone:           c.Timezone,
	}
}

func restClientToDb(c *rest.Client) db.Client {
	client := db.Client{
		PhoneNumber:        c.PhoneNumber,
		MobileOperatorCode: c.MobileOperatorCode,
		Tag:                c.Tag,
		Timezone:           c.Timezone,
	}
	client.ID = c.ID
	return client
}

func restMailingToDb(m *rest.Mailing) db.Mailing {
	mailing := db.Mailing{
		StartingTime:  m.StartingTime,
		Text:          m.Text,
		SendindFilter: db.SendindFilter(m.SendindFilter),
		EndingTime:    m.EndingTime,
		FilterValue:   m.FilterValue,
	}
	mailing.ID = m.ID
	return mailing
}

func dbMailingToRest(m *db.Mailing) rest.Mailing {
	return rest.Mailing{
		ID:            m.ID,
		StartingTime:  m.StartingTime,
		Text:          m.Text,
		SendindFilter: string(m.SendindFilter),
		EndingTime:    m.EndingTime,
		FilterValue:   m.FilterValue,
	}
}

func dbMessageToRest(m *db.Message) rest.Message {
	return rest.Message{
		ID:                       m.ID,
		ClientId:                 *m.ClientId,
		MailingId:                *m.MailingId,
		SendingStatus:            string(m.SendingStatus),
		SendingTime:              m.SendingTime,
		ConnectionTimeMillis:     m.ConnectionTime.Milliseconds(),
		ConnectionIdleTimeMillis: m.ConnectionIdleTime.Milliseconds(),
	}
}

func dbMessagesToRest(msgs []db.Message) []rest.Message {
	var restMessages []rest.Message
	for i := 0; i < len(msgs); i++ {
		restMessages = append(restMessages, dbMessageToRest(&msgs[i]))
	}
	return restMessages
}
