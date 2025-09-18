package dnd

import "fmt"

// These functions return a message so callers can assign it to lastMsg.

func TakeHealPot(p *Player) string {
	if p == nil {
		return ""
	}
	if p.HP >= p.HPmax {
		return "You are already at full health."
	}
	if !RemoveItem(p, "Health potion", 1) {
		return "You don't have a Health potion."
	}
	heal := p.HPmax / 2
	if heal <= 0 {
		heal = 1
	}
	p.HP += heal
	if p.HP > p.HPmax {
		p.HP = p.HPmax
	}
	return fmt.Sprintf("You used a Health potion. Your health is now %d/%d.", p.HP, p.HPmax)
}

func PoisonPot(p *Player, m *Monster) string {
	if p == nil {
		return ""
	}
	if !RemoveItem(p, "Poison potion", 1) {
		return "You don't have a Poison potion."
	}
	msg := ""
	if p.PoisonEffect > 0 {
		p.HP -= 15
		p.PoisonEffect--
		msg = fmt.Sprintf("You took 15 HP damage from the poison, you have %d turn(s) left.", p.PoisonEffect)
		if p.HP < 1 {
			IsDead(p)
		}
	}
	if m != nil && m.PoisonEffect > 0 {
		m.HP -= 15
		m.PoisonEffect--
		if msg == "" {
			msg = "The monster took 15 HP damage from the poison."
		} else {
			msg += " The monster took 15 HP damage from the poison."
		}
	}
	if msg == "" {
		msg = "The poison has no effect right now."
	}
	return msg
}

func TakePoisonPot(p *Player) string {
	if p == nil {
		return ""
	}
	if !RemoveItem(p, "Poison potion", 1) {
		return "You don't have a Poison potion."
	}
	p.HP -= 30
	if p.HP < 1 {
		IsDead(p)
	}
	return fmt.Sprintf("You did a bit of self-sabotage and drank a Poison potion. You took \033[31m30 damage\033[0m. Your health is now %d/%d.", p.HP, p.HPmax)
}
