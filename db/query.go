package db

import (
	"fmt"
	"time"

	"github.com/sebarray/wiselink/model"
	"github.com/sebarray/wiselink/service"
)

func QueryCreateEvent(event model.Event) string {

	query := " INSERT INTO event (id, description ,descriptionshort ,state , organizer , place , date , title ) VALUES "
	query += fmt.Sprintf(" ('%s', '%s','%s',  '%s','%s','%s','%s' , '%s'); ",
		event.Id, event.Description,
		event.DescriptionShort,
		event.State, event.Organizer,
		event.Place, event.Date.Format("2006-01-02 15:04:05"), event.Title)
	return query
}

func QueryReadEvents(filter model.Filter) string {
	query := " SELECT * FROM event where id!='' "

	if filter.State != "" {
		query += fmt.Sprintf("and state='%s'", filter.State)
	}
	if filter.Title != "" {
		query += fmt.Sprintf(" and title='%s'", filter.Title)

	}
	if filter.Date != "" {
		query += fmt.Sprintf(" and date='%s'", filter.Date)
	}
	return query + ";"
}

func QueryUpdateEvents(event model.Event) string {
	query := "UPDATE event SET "
	query += fmt.Sprintf("description='%s',descriptionShort='%s',state='%s',organizer='%s', place='%s',date='%s', title='%s' ",
		event.Description, event.DescriptionShort, event.State,
		event.Organizer, event.Place, event.Date.Format("2006-01-02 15:04:05"), event.Title)
	query += fmt.Sprintf("where id='%s' ;", event.Id)
	return query
}

func QueryCreateUser(user model.User) string {
	user.Password, _ = service.Encryptpswd(user.Password)
	query := "INSERT INTO user (id, name, email, password, admin) VALUES "
	query += fmt.Sprintf("('%s', '%s','%s',  '%s',%v); ",
		user.Id, user.Name, user.Email, user.Password, user.Admin)
	return query

}

func QueryCreateSuscribe(id, iduser, idevent string) string {
	query := "  INSERT INTO suscribe (`id`, `iduser`, `idevent`) VALUES "
	query += fmt.Sprintf("('%s', '%s', '%s');", id, iduser, idevent)
	return query

}

func QueryReadSuscribe(id, filter string) string {
	if filter != "" {
		filter = fmt.Sprintf("AND  EVENT.date%s'%s';", filter, time.Now().Format("2006-01-02"))
	} else {
		filter = ";"
	}
	query := " SELECT EVENT.id, EVENT.description, EVENT.descriptionshort, EVENT.state, EVENT.organizer, EVENT.place, EVENT.date, EVENT.title FROM  event JOIN suscribe ON event.id=suscribe.idevent JOIN user ON user.id=suscribe.iduser WHERE "
	query += fmt.Sprintf("user.id='%s' ", id)
	query += filter
	return query
}
