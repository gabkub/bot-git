package help

import (
	"bot-git/bot/abstract"
	"bot-git/messageBuilders"
	"strings"
)

const helpFlag = "-h"

type help struct {
	handlers []abstract.Handler
}

var commands abstract.ReactForMsgs = []string{"help", "pomocy", "pomoc"}

func New() *help {
	return &help{}
}

func (h *help) Init(handlers []abstract.Handler) {
	h.handlers = handlers
}

func (h *help) CanHandle(msg string) bool {
	return msg == helpFlag || commands.ContainsMessage(msg)
}

func (h *help) Handle(msg string, sender abstract.MessageSender) {
	sender.Send(messageBuilders.Text(h.GetHelp().Long))
}

var helpMsg *abstract.Help

func (h *help) GetHelp() *abstract.Help {
	if helpMsg == nil {
		helpMsg = h.buildHelp()
	}
	return helpMsg
}

func (h *help) buildHelp() *abstract.Help {
	var sb strings.Builder
	sb.WriteString("LISTA KOMEND:\n")
	for _, hnd := range h.handlers {
		if hnd == h {
			continue
		}
		sb.WriteString("- ")
		shortHelp := hnd.GetHelp().Short
		sb.WriteString(shortHelp)
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	sb.WriteString(" _<komenda> -h_ zwraca szczegółowe informacje o komendzie\n")
	return abstract.NewHelp("", sb.String())
}
