package attribute

// Attribute
type Attribute interface {
	Attribute() // Tag
	Id() int
	String() string
}

// static
var (
	attributes           = []Attribute{}
	fixedAttributes      = []*Fixed{}
	upgradableAttributes = []*Upgradable{}
	generatedAttributes  = []*Generated{}
)

// constructor
func newFixedAttribute(id int, name string, formatValue func(value interface{}) string) *Fixed {
	attr := &Fixed{
		id,
		name,
		formatValue,
	}
	attributes = append(attributes, attr)
	fixedAttributes = append(fixedAttributes, attr)
	return attr
}

// constructor
func newUpgradableAttribute(id int, name string, formatValues func(values interface{}) []string) *Upgradable {
	attr := &Upgradable{
		id,
		name,
		formatValues,
	}
	attributes = append(attributes, attr)
	upgradableAttributes = append(upgradableAttributes, attr)
	return attr
}

// constructor
func newGeneratedAttribute(id int, uattr *Upgradable, generateValues func(baseValue interface{}) []int) *Generated {
	attr := &Generated{
		id,
		uattr,
		generateValues,
	}
	attributes = append(attributes, attr)
	generatedAttributes = append(generatedAttributes, attr)
	return attr
}

func ForEachAttribute(f func(Attribute)) {
	for _, attr := range attributes {
		f(attr)
	}
}

func ForEachFixedAttribute(f func(*Fixed)) {
	for _, attr := range fixedAttributes {
		f(attr)
	}
}

func ForEachUpgradableAttribute(f func(*Upgradable)) {
	for _, attr := range upgradableAttributes {
		f(attr)
	}
}

func ForEachGeneratedAttribute(f func(*Generated)) {
	for _, attr := range generatedAttributes {
		f(attr)
	}
}

var (
	NAME        = newFixedAttribute(0, "Name", formatString)
	ARENA       = newFixedAttribute(1, "Arena", formatString)
	RARITY      = newFixedAttribute(2, "Rarity", formatString)
	TYPE        = newFixedAttribute(3, "Type", formatString)
	DESC        = newFixedAttribute(4, "Description", formatString)
	COST        = newFixedAttribute(5, "Elixir Cost", formatInt)
	HP          = newUpgradableAttribute(100, "Hitpoints", formatInts)
	SHP         = newUpgradableAttribute(110, "Shield Hitpoints", formatInts)
	DPS         = newUpgradableAttribute(200, "Damage per Second", formatInts)
	DPSL        = newUpgradableAttribute(210, "Damage per Second (L)", formatInts)
	DPSH        = newUpgradableAttribute(220, "Damage per Second (H)", formatInts)
	CTDPS       = newUpgradableAttribute(230, "Crown Tower Damage/sec", formatInts)
	DAM         = newUpgradableAttribute(300, "Damage", formatInts)
	DAML        = newUpgradableAttribute(310, "Damage (L)", formatInts)
	DAMH        = newUpgradableAttribute(320, "Damage (H)", formatInts)
	ADAM        = newUpgradableAttribute(330, "Area Damage", formatInts)
	DDAM        = newUpgradableAttribute(340, "Death Damage", formatInts)
	CTDAM       = newUpgradableAttribute(350, "Crown Tower Damage", formatInts)
	GOB_LV      = newUpgradableAttribute(400, "Goblin Level", formatInts)
	SGO_LV      = newUpgradableAttribute(410, "Spear Goblin Level", formatInts)
	SKE_LV      = newUpgradableAttribute(420, "Skeleton Level", formatInts)
	BAR_LV      = newUpgradableAttribute(430, "Barbarian Level", formatInts)
	SSPD        = newFixedAttribute(500, "Spawn Speed", formatTime)
	PSPD        = newFixedAttribute(510, "Production Speed", formatTime)
	HSPD        = newFixedAttribute(520, "Hit Speed", formatTime)
	TGTS        = newFixedAttribute(600, "Targets", formatString)
	SPD         = newFixedAttribute(700, "Speed", formatString)
	RNG         = newFixedAttribute(800, "Range", formatRange)
	DTIME       = newFixedAttribute(900, "Deploy Time", formatTime)
	LTIME       = newFixedAttribute(910, "Lifetime", formatTime)
	DUR_F       = newFixedAttribute(920, "Duration", formatTime)
	DUR_U       = newUpgradableAttribute(921, "Duration", formatTimes)
	RAD         = newFixedAttribute(1000, "Radius", formatFloat)
	COUNT       = newFixedAttribute(1100, "Count", formatCount)
	GOB_COUNT   = newFixedAttribute(1200, "Goblin Count", formatCount)
	MC_LV       = newUpgradableAttribute(1300, "Mirrored Common Level", formatInts)
	MR_LV       = newUpgradableAttribute(1310, "Mirrored Rare Level", formatInts)
	ME_LV       = newUpgradableAttribute(1320, "Mirrored Epic Level", formatInts)
	ML_LV       = newUpgradableAttribute(1330, "Mirrored Legendary Level", formatInts)
	CARDS_REQ   = newUpgradableAttribute(9990, "Cards Required", formatInts)
	GOLD_REQ    = newUpgradableAttribute(9991, "Gold Required", formatInts)
	EXP_GAIN    = newUpgradableAttribute(9992, "Experience Gained", formatInts)
	BASE_HP     = newGeneratedAttribute(10100, HP, generateHp)
	BASE_SHP    = newGeneratedAttribute(10110, SHP, generateHp)
	BASE_DAM    = newGeneratedAttribute(10300, DAM, generateDam)
	BASE_DAML   = newGeneratedAttribute(10310, DAML, generateDam)
	BASE_DAMH   = newGeneratedAttribute(10320, DAMH, generateDam)
	BASE_ADAM   = newGeneratedAttribute(10330, ADAM, generateDam)
	BASE_DDAM   = newGeneratedAttribute(10340, DDAM, generateDam)
	BASE_GOB_LV = newGeneratedAttribute(10400, GOB_LV, generateLv)
	BASE_SGO_LV = newGeneratedAttribute(10410, SGO_LV, generateLv)
	BASE_SKE_LV = newGeneratedAttribute(10420, SKE_LV, generateLv)
	BASE_BAR_LV = newGeneratedAttribute(10430, BAR_LV, generateLv)
	BASE_MC_LV  = newGeneratedAttribute(11300, MC_LV, generateLv)
	BASE_MR_LV  = newGeneratedAttribute(11310, MR_LV, generateLv)
	BASE_ME_LV  = newGeneratedAttribute(11320, ME_LV, generateLv)
)
