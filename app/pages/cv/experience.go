package cv

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"time"
)

type Company struct {
	Name    string
	Url     string
	LogoURL string
}

type SkillTag struct {
	Name string
}

type Job struct {
	Company     Company
	Position    string
	DateFrom    time.Time
	DateTo      time.Time
	Tags        []string
	Description string
}

func NewJobs() []Job {
	return []Job{
		{
			Company: Company{
				Name:    "Sbermarket",
				Url:     "https://sbermarket.ru/",
				LogoURL: "https://getmatch.ru/uploads/companies_logos/0edb9116-1f9c-4eed-9d7d-e9c0669103f1.png",
			},
			DateFrom: time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC),
			DateTo:   time.Time{},
			Position: "Backend developer",
			Tags: []string{
				"Go",
				"Python",
				"Postgresql",
				"Kafka",
				"Redis",
				"OpenRTB",
			},
			Description: `Стандартный бэкенд с ETL, думами об архитектуре и иногда интересными задачами. Отдел онлайн, а позже — оффлайн рекламы.


**Codeowner of**:


* Админка рекламных инструментов (*python, django*).
* Спиннер для показа нативной рекламы по таргетингу (*go, OpenRTB*).
* Сервис аутентификации/авторизации для всего домена (*go, JWT, oauth*).
* Сервис оффлайн-сэмплинга (*go*).
* Telegram бот для двусторонней связи с работниками в "полях" (*go, Telegram API*).


В промежутках: *PostgreSQL, Redis, Kafka, Kubernetes, S3*. 

*Jaeger, Sentry, Prometheus, ELK*.`,
		},
		{
			Company: Company{
				Name:    "[NDA]",
				Url:     "",
				LogoURL: "",
			},
			DateFrom: time.Date(2018, 9, 1, 0, 0, 0, 0, time.UTC),
			DateTo:   time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC),
			Position: "Backend developer",
			Tags: []string{
				"Python",
				"C++",
				"AsyncIO",
				"Kafka",
				"RabbitMQ",
				"PostgreSQL",
				"Blockchain",
			},
			Description: `Обязанности/проекты:

* Разработка бирж высокочастотной торговли на основе блокчейна.
* Поддержка и разработка смарт-контрактов (C++/EOS CDT).
* Проведение технических собеседований кандидатов (Python).
* Разработка anti money laundering систем.
* Разработка системы реконсиляции счетов.

Стек: Python 3.7+; asyncio; Python multiprocessing/multithreading; Kafka, RabbitMQ; PostgreSQL; Blockchain (BTC, Ethereum, EOS).`,
		},
	}
}

func (r Job) DescriptionHtml() string {
	return string(markdown.ToHTML([]byte(r.Description), nil, nil))
}

func (r Job) DateFromHumanize() string {
	return r.DateFrom.Format("January 2006")
}

func (r Job) DateToHumanize() string {
	if r.DateTo.IsZero() {
		return "Present"
	}
	return r.DateTo.Format("January 2006")
}

func (r Job) Dates() string {
	return fmt.Sprintf("%s - %s", r.DateFromHumanize(), r.DateToHumanize())
}

func (r Job) Duration() string {
	dateTo := r.DateTo
	if dateTo.IsZero() {
		now := time.Now()
		dateTo = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	}
	dur := dateTo.Sub(r.DateFrom)
	seconds := int64(dur.Seconds())
	monthsTotal := seconds / 2600640
	years := monthsTotal / 12
	monthsResidue := monthsTotal % 12
	if monthsResidue == 0 {
		return fmt.Sprintf("%d y.", years)
	} else {
		return fmt.Sprintf("%d y. %d mo.", years, monthsResidue)
	}
}
