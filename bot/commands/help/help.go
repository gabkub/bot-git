package help

import (
	"bot-git/bot/abstract"
	"bot-git/messageBuilders"
	"strings"
)

const helpFlag = "-h"

type help struct {
	commands abstract.ReactForMsgs
}

func New() *help {
	return &help{[]string{"help", "pomocy", "pomoc"}}
}

func (h *help) CanHandle(msg string) bool {
	return msg == helpFlag || h.commands.ContainsMessage(msg)
}

func (h *help) Handle(msg string, sender abstract.MessageSender) {
	var sb strings.Builder
	sb.WriteString("LISTA KOMEND:\n")
	sb.WriteString("- _hard_ (dostępne tylko w wiadomościach prywatnych z botem) - losowy żart w kategorii **hard** (na własną odpowiedzialność!)\n")
	sb.WriteString("- _help, pomocy_ - pomoc\n")
	sb.WriteString("- _joke, żart_ - losowy dowcip\n")
	sb.WriteString("- _meme, mem_ - losowy mem\n")
	sb.WriteString("- _news <kategoria>_ - losowy news z danej kategorii (brak kategorii wysyła newsa technologicznego)\n")
	sb.WriteString("- _piłkarzyki, gramy_ - rezerwacja stołu do gry w piłkarzyki\n")
	sb.WriteString("- _suchar, nie, ..._ - usuwa ostatni dowcip/mem\n")
	sb.WriteString("- _ver_ - wersja\n\n")
	sb.WriteString(" _<komenda> -h_ zwraca szczegółowe informacje o komendzie\n")
	sender.Send(messageBuilders.Text(sb.String()))
}

func (h *help) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Wyświetlenie ogólnej pomocy dla podstawowych komend\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_help, pomoc, pomocy_\n")
	return sb.String()
}
