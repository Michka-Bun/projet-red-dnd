package dnd

func HasSkill(p *Player, skill string) bool {
	if p == nil {
		return false
	}
	if p.Skills == nil {
		return false
	}
	return p.Skills[skill]
}

func AddSkill(p *Player, skill string) bool {
	if p == nil || skill == "" {
		return false
	}
	if p.Skills == nil {
		p.Skills = map[string]bool{}
	}
	if p.Skills[skill] {
		return false
	}
	p.Skills[skill] = true
	return true
}
