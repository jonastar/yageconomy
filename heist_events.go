package yageconomy

import (
	"math/rand"
)

var (
	heistEvtHostageHero = &HeistEventChance{
		Chance: 0.25,
		Inner: &HeistMemberEvent{
			Message:        "A hostage played hero and started shooting at your team.",
			DeadMembersMin: 1,
			DeadMembersMax: 3,
		},
	}
	
	heistEvtBetrayal = &HeistEventChance{
		Chance: 0.35,
		Inner: &HeistMemberEvent{
			Message:        "One of your heist members got greedy and tried to take you out for more money, he was also killed.",
			DeadMembersMin: 2,
			DeadMembersMax: 2,
		},
	}

	heistEvtTripSpillMoney = &HeistEventChance{
		Chance: 0.5,
		Inner: &HeistMoneyEvent{
			Message:             "One of your members tripped, spilling some money in the process.",
			MoneyLossPercentMin: 5,
			MoneyLossPercentMax: 69,
		},
	}

	heistEvtTripInjury = &HeistEventChance{
		Chance: 0.25,
		Inner: &HeistMemberEvent{
			Message:           "A hostage ran into your team with a knife and started slashing!",
			InjuredMembersMin: 1,
			InjuredMembersMax: 10,
		},
	}
)

var OrderedHeistEvents = map[HeistProgressState][]HeistEvent{
	HeistProgressStateStarting: []HeistEvent{
		&BasicHeistEvent{
			Message:              "So uh, you're charging into the Wonderbank alone huh? Well you better get ready because you're starting now!",
			MessagePluralMembers: "Alright guys, check your guns. We are storming into the bank through all entrances. Let's get the wondercoins and get out before the cops get here.",
		},
	},
	HeistProgressStateInvading: []HeistEvent{
		&BasicHeistEvent{
			Message: "You've entered the building and you're trying to get control and also get the location of the money.",
		},
		&HeistEventChance{
			Chance: 0.25,
			Inner: &HeistMoneyEvent{
				Message:             "One of your crew members shot a hostage trying to play hero, you're gonna have to bribe some people to get out of this...",
				MoneyLossPercentMin: 20,
				MoneyLossPercentMax: 80,
			},
		},
		heistEvtHostageHero,
		heistEvtBetrayal,
		heistEvtTripInjury,
	},
	HeistProgressStateCollecting: []HeistEvent{
		&BasicHeistEvent{
			Message: "You've found the stash of wondercoins and started collecting the dough.",
		},

		&HeistEventChance{
			Chance: 0.5,
			Inner: &HeistMoneyEvent{
				Message:             "One of your bags ripped open and there's now wondercoins everywhere.",
				MoneyLossPercentMin: 10,
				MoneyLossPercentMax: 100,
			},
		},

		&HeistEventChance{
			Chance: 0.15,
			Inner: &HeistMemberEvent{
				Message:        "Someone walked into tripwire and set off an explosion near your team!",
				DeadMembersMin: 1,
				DeadMembersMax: 4,
			},
		},

		&HeistEventChance{
			Chance: 0.5,
			Inner: &HeistMemberEvent{
				Message:           "A sniper starts shooting at your team through the windows, you're taking too long!",
				InjuredMembersMin: 1,
				InjuredMembersMax: 20,
			},
		},
		heistEvtHostageHero,
		heistEvtBetrayal,
	},
	HeistProgressStateLeaving: []HeistEvent{
		&BasicHeistEvent{
			Message: "Alright the cops are getting close, time to head out!",
		},
		heistEvtTripSpillMoney,
		&HeistEventChance{
			Chance: 0.25,
			Inner: &HeistEventIncreaseChance{
				Message: "As you're walking out, a hostage spots a tattoo on one of your crew members.",
				Amount:  60,
			},
		},
		heistEvtBetrayal,
		heistEvtHostageHero,
		heistEvtTripInjury,
	},
	HeistProgressStateGetaway: []HeistEvent{
		&BasicHeistEvent{
			Message: "You've found the getaway car, ITS TIME TO LEAVE!",
		},
		heistEvtTripSpillMoney,
		&HeistEventChance{
			Chance: 0.4,
			Inner: &HeistMemberEvent{
				Message:           "There's a blockade up ahead, along with snipers on the roof giving you all kinds of trouble!",
				DeadMembersMin:    1,
				DeadMembersMax:    5,
				InjuredMembersMin: 1,
				InjuredMembersMax: 10,
			},
		},
	},
}

type HeistEvent interface {
	Run(session *HeistSession) *HeistEventEffect
}

type HeistEventEffect struct {
	Dead     []*HeistUser
	Injured  []*HeistUser
	Captured []*HeistUser

	HostagesKilled  int
	HostagesInjured int

	MoneyLostPercentage int
	MoneyLostFixed      int

	TextResponse string

	IncreasedChanceOfEvents int
}

type SimpleHeistEvent struct {
	Description   string
	Chance        float64
	MemberLossMin int
	MemberLossMax int

	MinMembers int
	MaxMembers int

	MoneyLossPercentage int
}

type BasicHeistEvent struct {
	Message string

	// optional seperate message if there's plural number of members
	MessagePluralMembers string
}

func (b *BasicHeistEvent) Run(session *HeistSession) *HeistEventEffect {
	resp := b.Message

	if b.MessagePluralMembers != "" && len(session.aliveUsers()) > 1 {
		resp = b.MessagePluralMembers
	}
	return &HeistEventEffect{
		TextResponse: resp,
	}
}

type HeistMoneyEvent struct {
	Message string

	MoneyLossPercentMin int
	MoneyLossPercentMax int

	MoneyLossFixedMin int
	MoneyLossFixedMax int
}

func (b *HeistMoneyEvent) Run(session *HeistSession) *HeistEventEffect {

	return &HeistEventEffect{
		TextResponse:        b.Message,
		MoneyLostPercentage: rand.Intn(b.MoneyLossPercentMax+1-b.MoneyLossFixedMin) + b.MoneyLossPercentMin,
		MoneyLostFixed:      rand.Intn(b.MoneyLossFixedMax+1-b.MoneyLossFixedMin) + b.MoneyLossFixedMin,
	}
}

type HeistEventChance struct {
	Chance float64
	Inner  HeistEvent
}

func (b *HeistEventChance) Run(session *HeistSession) *HeistEventEffect {
	if session.calcEventChance(b.Chance) {
		return b.Inner.Run(session)
	}

	return nil
}

type HeistMemberEvent struct {
	Message string

	DeadMembersMin int
	DeadMembersMax int

	InjuredMembersMin int
	InjuredMembersMax int
}

func (b *HeistMemberEvent) Run(session *HeistSession) *HeistEventEffect {

	numDead := rand.Intn(b.DeadMembersMax+1-b.DeadMembersMin) + b.DeadMembersMin
	numInjured := rand.Intn(b.InjuredMembersMax+1-b.InjuredMembersMin) + b.InjuredMembersMin

	alivePool := session.aliveUsers()

	killed := make([]*HeistUser, 0, numDead)

	// start by killing
	for i := 0; i < numDead; i++ {
		index := rand.Intn(len(alivePool))
		killed = append(killed, alivePool[index])
		alivePool = append(alivePool[:index], alivePool[index+1:]...)

		if len(alivePool) < 1 {
			break
		}
	}

	injured := make([]*HeistUser, 0, numDead)

	// follow up by injuring
	for i := 0; i < numInjured; i++ {
		if len(alivePool) < 1 {
			break
		}

		index := rand.Intn(len(alivePool))
		injured = append(injured, alivePool[index])
		alivePool = append(alivePool[:index], alivePool[index+1:]...)
	}

	return &HeistEventEffect{
		TextResponse: b.Message,
		Dead:         killed,
		Injured:      injured,
	}
}

type HeistEventIncreaseChance struct {
	Message string
	Amount  int
}

func (b *HeistEventIncreaseChance) Run(session *HeistSession) *HeistEventEffect {

	return &HeistEventEffect{
		TextResponse:            b.Message,
		IncreasedChanceOfEvents: b.Amount,
	}
}
