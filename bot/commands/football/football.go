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
	db       *footballDatabase.FootballDb
}

func New(db *footballDatabase.FootballDb) *football {
	return &football{[]string{"gramy", "play", "game", "football", "soccer", "piłkarzyki"}, db}
}

func (f *football) CanHandle(msg string) bool {
	return f.commands.ContainsMessage(msg)
}

func (f *football) Handle(msg string, sender abstract.MessageSender) {
	bookingTime := time.Now()
	if strings.Contains(msg, "-l") {
		f.getReservations(bookingTime, sender)
		return
	}
	resp := f.tryBookTable(sender.GetUserId(), msg, bookingTime)
	sender.Send(messageBuilders.Text(resp))
}

func (f *football) tryBookTable(userId abstract.UserId, msg string, bookingTime time.Time) string {
	msgSplit := strings.Split(msg, "@")
	var err string
	if len(msgSplit) >= 2 {
		bookingTime, err = setTime(msgSplit[len(msgSplit)-1])
		if err != "" {
			return err
		}
	}
	normDate := footballDatabase.NewNormalizeDate(bookingTime)
	isFree := f.db.IsFree(normDate)
	if isFree {
		user, _ := config.ConnectionCfg.Client.GetUser(string(userId), "")
		if f.db.IsBooked(user.Username, normDate) {
			return "Już dzisiaj rezerwowałeś."
		}
		f.db.SetReservation(user.Username, normDate)
		return "Zarezerwowano piłkarzyki. Miłej gry!"
	}
	freeTime := f.db.FirstFreeTimeFor(normDate)
	return fmt.Sprintf("Stół zajęty. Stół będzie wolny o: %v:%v", freeTime.Hour(), freeTime.Minute())
}

func setTime(toConvert string) (time.Time, string) {
	now := time.Now()
	if toConvert == "" {
		return now, ""
	}

	hourMinute := strings.Split(toConvert, ":")
	hour, e := strconv.Atoi(hourMinute[0])
	if e != nil || hour <= 6 || hour >= 20 {
		return time.Time{}, "Zły format godziny."
	}
	minute, e := strconv.Atoi(hourMinute[1])
	if e != nil || minute < 0 || minute >= 60 {
		return time.Time{}, "Zły format minut."
	}

	result := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.UTC)
	return result, ""
}

const ballImgUrl = "https://a.espncdn.com/combiner/i?img=/redesign/assets/img/icons/ESPN-icon-soccer.png&w=288&h=288&transparent=true"

func (f *football) getReservations(startTime time.Time, sender abstract.MessageSender) {
	reservations := f.db.GetAllReservationByStartTime(footballDatabase.NewNormalizeDate(startTime))
	var sb strings.Builder
	for _, reservation := range reservations {
		sb.WriteString(fmt.Sprintf("%v - %v : %v\n",
			reservation.StartTime.Format("15:04:05"), reservation.EndTime().Format("15:04:05"), reservation.UserName))
	}
	y, m, d := time.Now().Date()
	title := fmt.Sprintf("%v %v %v", d, m, y)
	sender.Send(messageBuilders.TitleThumbUrl(title, sb.String(), ballImgUrl))
}
