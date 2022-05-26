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

func restMailingToDb(c *rest.Mailing) db.Mailing {
	mailing := db.Mailing{
		StartingTime:  c.StartingTime,
		Text:          c.Text,
		SendindFilter: db.SendindFilter(c.SendindFilter),
		EndingTime:    c.EndingTime,
		FilterValue:   c.FilterValue,
	}
	mailing.ID = c.ID
	return mailing
}

func dbMailingToRest(c *db.Mailing) rest.Mailing {
	return rest.Mailing{
		ID:            c.ID,
		StartingTime:  c.StartingTime,
		Text:          c.Text,
		SendindFilter: string(c.SendindFilter),
		EndingTime:    c.EndingTime,
		FilterValue:   c.FilterValue,
	}
}
