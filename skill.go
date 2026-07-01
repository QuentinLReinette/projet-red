package main

type skillType int

const (
	punch skillType = iota
	fireBall
)

var skillNames = map[skillType]string{
	punch:    "Punch",
	fireBall: "Fireball",
}

type skill struct {
	skillType
	char   *character
	action func()
}

type book struct {
	skillType
	learn func()
}

func (s skill) String() string {
	return skillNames[s.skillType]
}

func (t skillType) String() string {
	return skillNames[t]
}

func newPunch(char *character) skill {
	newSkill := skill{punch, char, nil}
	newSkill.action = newSkill.skillPunch
	return newSkill
}

func newFireBall(char *character) skill {
	newSkill := skill{fireBall, char, nil}
	newSkill.action = newSkill.skillFireBall
	return newSkill
}

func (s skill) skillPunch() {}

func (s skill) skillFireBall() {}
