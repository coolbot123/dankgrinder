package responder

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dankgrinder/dankgrinder/discord"
	"github.com/dankgrinder/dankgrinder/scheduler"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const DMID = "270904126974590976"

var exp = struct {
	search,
	fh,
	hl,
	bal,
	gift,
	shop,
	event *regexp.Regexp
}{ //regex is amazing and terrable at the same time
	search: regexp.MustCompile(`Pick from the list below and type the name in chat\.\s\x60(.+)\x60,\s\x60(.+)\x60,\s\x60(.+)\x60`),
	fh:     regexp.MustCompile(`10\sseconds.*\s?([Tt]yping|[Tt]ype)\s\x60(.+)\x60`),
	hl:     regexp.MustCompile(`Your hint is \*\*([0-9]+)\*\*`),
	bal:    regexp.MustCompile(`\*\*Wallet\*\*: \x60?⏣?\s?([0-9,]+)\x60?`),
	event:  regexp.MustCompile(`^(Attack the boss by typing|Type) \x60(.+)\x60`),
	gift:   regexp.MustCompile(`[a-zA-Z\s]* \(([0-9]+) owned\)`),
	shop:   regexp.MustCompile(`pls shop ([a-zA-Z\s]+)`),
}

//what am I doing wahy am i doing wat amiwhat am idoiņ̭͆̍͒̏ġ͓̬̗͘ ̰͇͔̖̲̖̮̍̋̒͢ͅw̞͍͛̆͆ͧ͡h̛͕͕̒͋ͅÿ̨͎̻̜̘̮́̍͋ ̨͇͚̫̫͉̫̐͌a͖̜͓͖͚̙̦̭͌͢m̵͓̦̮̙ͫ̔͑í̶̜̣͉͚̺̞̮̐d̰̪̜͖͎ͥ̔̍ͩ͢ǫ̺̻̍́i̝̭̙̱̝͐ͣ̈́n̲̥̅̀ĝ̘͚̲͗ͬ̕how
var numFmt = message.NewPrinter(language.English)

func (r *Responder) fh(msg discord.Message) {
	res := exp.fh.FindStringSubmatch(msg.Content)[2]
	r.Sdlr.ResumeWithCommandOrPrioritySchedule(&scheduler.Command{
		Value: clean(res),
		Log:   "responding to fishing or hunting event",
	}) //i wish i could respond
}

func (r *Responder) pm(_ discord.Message) {
	res := r.PostmemeOpts[rand.Intn(len(r.PostmemeOpts))]
	r.Sdlr.ResumeWithCommandOrPrioritySchedule(&scheduler.Command{
		Value: res,
		Log:   "responding to postmeme",
	})
} //i̻͚̲̲͖ͬ͘ ͍̜̤̀ͯ͜ẅ̶̤̩̥͚̦͇̠́ͤ́͊i̗͕͕͈̜͐̇͘ş̪̲̞̻̯͉͕͌́ͬh̥̬͇̺ͯ͢ ̠̦͖̞̪̖̼̐̅͆͌͞i͕̻͎͈̺ͬͨ́͘ ̬͕̤̘͎ͦͮ̈͋͡c͓̲̙͋̚͞oͭ҉͕̯͙̮͎u̴̹̗ͤ̓ͣl͈̳͖̺ͦ͞d̶̲̖̳̰͇̘̆̾ͥͅ ̘̫͈ͩ͐ͫ͗̕ͅř͎̯̮̱͔̳̠͝eͧ̅͑͏̬s̢̰͇͖̹̭̟̍̚ͅp̶̝ͧ́̒̚o͗̿ͦ҉̦͈n̛͕̫̰̤͕̝̬͓ͦͮd̖̞͓̣̭͍͉͒ͯͩ͡

func (r *Responder) event(msg discord.Message) {
	res := exp.event.FindStringSubmatch(msg.Content)[2]
	r.Sdlr.PrioritySchedule(&scheduler.Command{
		Value: clean(res),
		Log:   "responding to global event",
	})
} //i̷̙̺̞͕̟̘͗ͩ̂͑ͅ ̗̭̬ͩ͑̐͝w̛̖̳͔͕̰̩̯ͭͭi̷̮̫̹͖̋ͮ͛͋ŝ̩͔ͨ́h̦͍͇̓̾̽̕ ̂̃̋̚͏̝͔͚̝̼̩̠i̓͌͏͉̺̦͙̘ ̼̥̯͓͑̐̎͂͡ͅc͉̻ͮͦ͘o̵͓̭͚̥̬̫͚ͤų̪̬̦͋̋̎l̙̐͞ͅḋ̃҉̬̲͉͈͙̙͈ ̧̬̻̫̼͎̮̋̏̊̆ͅr̹͔̜̤̦ͬ͝ͅe̞͇͈ͦͦͧ͠s͙̪͉͖̹͉͍ͫ̐̕p̸̝͎̬̗͓ͣ̂̈ơ̻͇͓̹ͨn̬͔̩̪̄ͥ͞ͅd͎̔͌̔̀͘ͅͅ

func (r *Responder) search(msg discord.Message) {
	choices := exp.search.FindStringSubmatch(msg.Content)[1:]
	for _, choice := range choices {
		for _, allowed := range r.AllowedSearches {
			if choice == allowed {
				r.Sdlr.ResumeWithCommandOrPrioritySchedule(&scheduler.Command{
					Value: choice,
					Log:   "responding to search",
				})
				return
			}
		}
	} //ị̡͚̞̩̬̫̮̫͊ ̨͕̙̺̭̰͒̆ͬw̲͍͓̼̥͙͙͉̃͟i̖͔͚͓ͦ̃͜s̜̹͉͌͜h̶͇̪̫͓̜̯̤̮̽̇̚ ̘̼̹̭̇̆̀ͅi̴̖̲̠̪̹̮̲̥͊ ̨̲̤͙̼̪̞͈̀̒c̵̝͕̰͍̮̉͗͒̎o͓͉̘̥̗̭ͣ͛̀ǔ̵͈̥̩̳͉̠͓ḻ̼͇͙̙ͭ̌͑͐͜ḓ̢̟͚̬̼̱ͧ͆ͩ ̷͈̯̫ͫ̈r̷̜̱̦̟͔̦̓ͅe̮̪͂͐͐͂͝s̴̗̠̯̖̣̰ͤp͕͍͙͖̩͉̬͍͐̏͐̌͝o̯̻̟̫͙̼͒͊́n̨̹͙ͨ͑̄͗ḑ̟̫̩͕̿ͅ

	r.Sdlr.ResumeWithCommandOrPrioritySchedule(&scheduler.Command{
		Value: r.SearchCancel[rand.Intn(len(r.SearchCancel))],
		Log:   "no allowed search options provided, responding",
	})
}

func (r *Responder) hl(msg discord.Message) {
	if !exp.hl.MatchString(msg.Embeds[0].Description) {
		return
	}
	nstr := strings.Replace(exp.hl.FindStringSubmatch(msg.Embeds[0].Description)[1], ",", "", -1)
	n, err := strconv.Atoi(nstr)
	if err != nil {
		r.Logger.Errorf("error while reading highlow hint: %v", err)
		return
	}
	res := "high"
	if n > 50 {
		res = "low"
	}
	r.Sdlr.ResumeWithCommandOrPrioritySchedule(&scheduler.Command{
		Value: res,
		Log:   "responding to highlow",
	})
} //i̛̝̬̮̮̼͙͙̹̿̒͐́ ̛̻͓̥͔͉̜ͨͣ̔w̛̫̞̉́̓ͤi͊ͩs̺̯̳͙̮͓͐͡ͅh̖͎̜͍̣̖͒ͧͩ͠ ̡̝͎̝ͬ͌͊̀i̭̺̻͓̽͘ ͙̰̖̹͖͉̋ͯ̒́c̴̬̤̠̗̩̞̳̙ͯ̐͌͋ô̵͇̗͇ͤͅu̯̝͙͉͉ͯ͆ͧ̀l̘̲͈̣̑͌͞d̪͎͚̙̘̗̠̾̈̌͠ ̴͙̫̰̀ͭṙ̛̫̞̯̥͚͇͇ͩ̄ę̮̭̗̬ͬ̈́s̹̠͔̙̠̊̕p̮͚̟̙̃̓͜ǒ̷̳͉̱͉̞͍̺͆́ͦn͚̞͈̙͎̲̻̖ͥ̋͡dͧͤ̅҉͇͎̫̫̥

func (r *Responder) balCheck(msg discord.Message) {
	if !r.BalanceCheck {
		return
	}
	if !strings.Contains(msg.Embeds[0].Title, r.Client.User.Username) {
		return
	}
	balstr := strings.Replace(exp.bal.FindStringSubmatch(msg.Embeds[0].Description)[1], ",", "", -1)
	bal, err := strconv.Atoi(balstr)
	if err != nil {
		r.Logger.Errorf("error while reading balance: %v", err)
		return
	}
	r.Logger.Infof(
		"current wallet balance: %v coins",
		numFmt.Sprintf("%d", bal),
	)
	if r.startingTime.IsZero() {
		r.startingBal = bal
		r.startingTime = time.Now()
		return
	}
	inc := bal - r.startingBal
	per := time.Now().Sub(r.startingTime)
	hourlyInc := int(math.Round(float64(inc) / per.Hours()))
	r.Logger.Infof(
		"average income: %v coins/h",
		numFmt.Sprintf("%d", hourlyInc),
	)
}

func (r *Responder) abLaptop(_ discord.Message) {
	if !r.AutoBuy.Laptop {
		return
	}
	r.Sdlr.PrioritySchedule(&scheduler.Command{
		Value: "pls buy laptop",
		Log:   "no laptop, buying a new one",
	})
}

func (r *Responder) abHuntingRifle(_ discord.Message) {
	if !r.AutoBuy.HuntingRifle {
		return
	}
	r.Sdlr.PrioritySchedule(&scheduler.Command{
		Value: "pls buy rifle",
		Log:   "no hunting rifle, buying a new one",
	})
}

func (r *Responder) abFishingPole(_ discord.Message) {
	if !r.AutoBuy.FishingPole {
		return
	}
	r.Sdlr.PrioritySchedule(&scheduler.Command{
		Value: "pls buy fishing pole",
		Log:   "no fishing pole, buying a new one",
	})
}

func (r *Responder) gift(msg discord.Message) {
	if !r.AutoGift.Enable {
		return
	}
	trigger := r.Sdlr.AwaitResumeTrigger()
	if !strings.Contains(trigger, "shop") {
		return
	}
	if r.Client.User.ID == r.AutoGift.To {
		r.Sdlr.Resume()
		return
	}
	if !exp.gift.Match([]byte(msg.Embeds[0].Title)) || !exp.shop.Match([]byte(trigger)) {
		r.Sdlr.Resume()
		return
	}
	amount := exp.gift.FindStringSubmatch(msg.Embeds[0].Title)[1]
	item := exp.shop.FindStringSubmatch(trigger)[1]
	r.Sdlr.ResumeWithCommandOrPrioritySchedule(&scheduler.Command{
		Value: fmt.Sprintf("pls gift %v %v <@%v>", amount, item, r.AutoGift.To),
		Log:   "gifting items",
	})
}

// clean removes all characters except for ASCII characters [32, 126] (basically
// all keys you would find on a US keyboard).
func clean(s string) string {
	var result string
	for _, char := range s {
		if regexp.MustCompile(`[\x20-\x7E]`).MatchString(string(char)) {
			result += string(char)
		}
	}
	return result
}

func (r *Responder) router() *discord.MessageRouter {
	rtr := &discord.MessageRouter{}

	// Fishing and hunting events.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		ContentMatchesExp(exp.fh).
		Mentions(r.Client.User.ID).
		Handler(r.fh)

	// Postmeme.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		ContentContains("What type of meme do you want to post").
		Mentions(r.Client.User.ID).
		Handler(r.pm)

	// Global events.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		ContentMatchesExp(exp.event).
		Handler(r.event)

	// Search.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		ContentMatchesExp(exp.search).
		Mentions(r.Client.User.ID).
		Handler(r.search)

	// Highlow.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		HasEmbeds(true).
		Mentions(r.Client.User.ID).
		Handler(r.hl)

	// Balance report.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		HasEmbeds(true).
		Handler(r.balCheck)

	// Auto-buy laptop.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		ContentContains("oi you need to buy a laptop in the shop to post memes").
		Mentions(r.Client.User.ID).
		Handler(r.abLaptop)

	// Auto-buy fishing pole.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		ContentContains("You don't have a fishing pole").
		Mentions(r.Client.User.ID).
		Handler(r.abFishingPole)

	// Auto-buy hunting rifle.
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		ContentContains("You don't have a hunting rifle").
		Mentions(r.Client.User.ID).
		Handler(r.abHuntingRifle)

	// Auto-gift
	rtr.NewRoute().
		Channel(r.ChannelID).
		Author(DMID).
		HasEmbeds(true).
		Handler(r.gift)

	return rtr
}

//Déjà vu
