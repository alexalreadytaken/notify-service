package rest

type MailingsDashboard struct {
	Dashboard []MailingCountsByStatus
}

type MailingCountsByStatus struct {
	Mailing Mailing
	Counts  []CountMessagesByStatus
}

type CountMessagesByStatus struct {
	Status string
	Count  uint
}
