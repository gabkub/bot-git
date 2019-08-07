package commands

import (
	"bot-git/bot/abstract"
	"bot-git/bot/messages"
	"bot-git/config"
	"bot-git/footballDatabase"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type football struct {
	commands []string
}

var FootballHandler football

func (f *football) New() abstract.Handler {
	f.commands = []string{"gramy", "play", "game", "football", "soccer", "piłkarzyki"}
	return f
}

func (f *football) CanHandle(msg string) bool {
	return abstract.FindCommand(f.commands, msg)
}

func (f *football) Handle(msg string) messages.Message {
	if strings.Contains(msg, "-h") {
		return f.GetHelp()
	}
	bookingTime := time.Now()
	if strings.Contains(msg, "-l") {
		return f.getReservations(bookingTime)
	}
	msgSplit := strings.Split(msg, "@")
	var err string
	if len(msgSplit) >= 2 {
		bookingTime, err = setTime(msgSplit[len(msgSplit)-1])
		if err != "" {
			messages.Response.Text = err
			return messages.Response
		}
	}

	free := footballDatabase.FreeReservation(bookingTime)
	if free.IsZero() {
		messages.Response.Text = "Nie można zarezerwować. Spróbuj inną godzinę."
		return messages.Response
	}
	if footballDatabase.TimeToString(free) == footballDatabase.TimeToString(bookingTime) {
		user, _ := config.ConnectionCfg.Client.GetUser(abstract.GetUserId(), "")
		if footballDatabase.SetReservation(user.Username, bookingTime) {
			messages.Response.Text = "Zarezerwowano piłkarzyki. Miłej gry!"
		} else {
			messages.Response.Text = "Już dzisiaj rezerwowałeś."
		}
		return messages.Response
	}

	messages.Response.Text = fmt.Sprintf("Stół zajęty. Stół będzie wolny o: %v:%v", free.Hour(), free.Minute())
	return messages.Response
}

func (f *football) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Rezerwacja stołu do gry w piłkarzyki na 20 minut. Domyślna godzina rezerwacji to godzina wysłania wiadomości.\n")
	sb.WriteString("Limit rezerwacji na użytkownika = 1\n")
	sb.WriteString("Szablon: _<komenda>_ (@_<godzinarezerwacji>_) (domyślnie ustawiana jest aktualna godzina)\n")
	sb.WriteString("_<komenda>_ -l - wyświetla wszystkie rezerwacje na dany dzień.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_football, game, gramy, piłkarzyki, play, soccer_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}

func setTime(toConvert string) (time.Time, string) {
	now := time.Now()
	if toConvert == "" {
		return now, ""
	}

	hour_minute := strings.Split(toConvert, ":")
	hour, e := strconv.Atoi(hour_minute[0])
	if e != nil || hour <= 6 || hour >= 20 {
		return time.Time{}, "Zły format godziny."
	}
	minute, e := strconv.Atoi(hour_minute[1])
	if e != nil || minute < 0 || minute >= 60 {
		return time.Time{}, "Zły format minut."
	}

	result := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.UTC)
	return result, ""
}

func (f *football) getReservations(startTime time.Time) messages.Message {
	reservations := footballDatabase.GetAllReservationByStartTime(startTime)
	var sb strings.Builder
	for _, reservation := range reservations {
		startTimes := strings.Split(reservation.StartTime, " ")
		endTimes := strings.Split(reservation.EndTime, " ")
		sb.WriteString(fmt.Sprintf("%v - %v : %v\n", startTimes[1], endTimes[1], reservation.UserName))
	}
	messages.Response.Text = sb.String()
	y, m, d := time.Now().Date()
	messages.Response.Title = fmt.Sprintf("%v %v %v", d, m, y)
	messages.Response.ThumbUrl = "https://a.espncdn.com/combiner/i?img=/redesign/assets/img/icons/ESPN-icon-soccer.png&w=288&h=288&transparent=true"
	return messages.Response
}
