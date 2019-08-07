package football

import (
	"bot-git/bot/abstract"
	"bot-git/config"
	"bot-git/footballDatabase"
	"bot-git/messageBuilders"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type football struct {
	commands abstract.ReactForMsgs
}

func New() *football {
	return &football{[]string{"gramy", "play", "game", "football", "soccer", "piłkarzyki"}}
}

func (f *football) CanHandle(msg string) bool {
	return f.commands.ContainsMessage(msg)
}

func (f *football) Handle(msg string, sender abstract.MessageSender) {
	if strings.Contains(msg, "-h") {
		sender.Send(messageBuilders.Text(f.GetHelp()))
		return
	}
	bookingTime := time.Now()
	if strings.Contains(msg, "-l") {
		f.getReservations(bookingTime, sender)
		return
	}
	resp := tryBookTable(msg, bookingTime)
	sender.Send(messageBuilders.Text(resp))
}

func tryBookTable(msg string, bookingTime time.Time) string {
	msgSplit := strings.Split(msg, "@")
	var err string
	if len(msgSplit) >= 2 {
		bookingTime, err = setTime(msgSplit[len(msgSplit)-1])
		if err != "" {
			return err
		}
	}
	free := footballDatabase.FreeReservation(bookingTime)
	if free.IsZero() {
		return "Nie można zarezerwować. Spróbuj inną godzinę."
	}
	if footballDatabase.TimeToString(free) == footballDatabase.TimeToString(bookingTime) {
		user, _ := config.ConnectionCfg.Client.GetUser(abstract.GetUserId(), "")
		if footballDatabase.SetReservation(user.Username, bookingTime) {
			return "Zarezerwowano piłkarzyki. Miłej gry!"
		} else {
			return "Już dzisiaj rezerwowałeś."
		}
	}
	return fmt.Sprintf("Stół zajęty. Stół będzie wolny o: %v:%v", free.Hour(), free.Minute())
}

func (f *football) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Rezerwacja stołu do gry w piłkarzyki na 20 minut. Domyślna godzina rezerwacji to godzina wysłania wiadomości.\n")
	sb.WriteString("Limit rezerwacji na użytkownika = 1\n")
	sb.WriteString("Szablon: _<komenda>_ (@_<godzinarezerwacji>_) (domyślnie ustawiana jest aktualna godzina)\n")
	sb.WriteString("_<komenda>_ -l - wyświetla wszystkie rezerwacje na dany dzień.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_football, game, gramy, piłkarzyki, play, soccer_\n")
	return sb.String()
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

const ballImgUrl = "https://a.espncdn.com/combiner/i?img=/redesign/assets/img/icons/ESPN-icon-soccer.png&w=288&h=288&transparent=true"

func (f *football) getReservations(startTime time.Time, sender abstract.MessageSender) {
	reservations := footballDatabase.GetAllReservationByStartTime(startTime)
	var sb strings.Builder
	for _, reservation := range reservations {
		startTimes := strings.Split(reservation.StartTime, " ")
		endTimes := strings.Split(reservation.EndTime, " ")
		sb.WriteString(fmt.Sprintf("%v - %v : %v\n", startTimes[1], endTimes[1], reservation.UserName))
	}
	y, m, d := time.Now().Date()
	title := fmt.Sprintf("%v %v %v", d, m, y)
	sender.Send(messageBuilders.TitleThumbUrl(title, sb.String(), ballImgUrl))
}
