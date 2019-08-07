package notNowMsg

import (
	"bot-git/bot/messages"
	"math/rand"
)

var limitMessages = []string{
	"Do roboty!", "Hej ho, hej ho, do pracy by się szło...", "Już się zmęczyłem.", "Zostaw mnie w spokoju.",
	"Koniec śmieszków...", "Foch.", "Nie.", "Zaraz wracam. Albo i nie...", "A może by tak popracować?", "~~żart~~",
	"Kolego, poszukaj w eDoku - może tam znajdziesz...",
	"Może lepiej @dadoczek ?",
	"@dadoczek kawał",
	"Jestem na obiedzie w Bistro :pizza:",
	"Jestem zajęty - teraz bujam się po mieście BMW",
	"Jadę na wdrożenie do Gorzowa :car:",
	"Później - teraz wykręcam alusy z szarego BMW, które stoi u Was na parkingu. Nie wiecie czyje to?",
	"Głodny nie jesteś sobą - zjedz coś w Bistro :pizza:",
	"Teraz czytam książkę od @dadoczek :book:",
	"Kolego, bo pójdę spać :sleeping_bed:",
	"A chcesz pojechać na wdrożenie do Gorzowa?",
	"Czekaj, czekaj... celuje w tarczę ",
	"Dacie zapalić cygaro to może coś wrzucę",
	"Lepiej może piłkarzyki?",
	"Weź przykład z Daniela i popracuj trochę.",
	"Nie bierz przykładu z Dyrektora i pokoduj trochę... :briefcase:",
	"Jeśli w eDoku nie ma to może Kokpit? Tam podobno jest wszystko.",
	"Jak ubierzesz kurczaka to wrzucę",
	"Teraz z kolegą palimy :smoking:",
	"No co Ty gadasz!?",
	"Teraz nie mogem",
	"Idem do domu...",
	"Troche mnie zblokło",
	"niechcem",
	"muszem?",
	"A może burgerek? :hamburger:",
}

func Get() messages.Message {
	var msg messages.Message
	msg.New()
	msg.Text = limitMessages[rand.Intn(len(limitMessages))]
	return msg
}
