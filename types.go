package genshinapi

const (
	ArtifactsDType  = "artifacts"
	CharactersDType = "characters"
	DomainsDType    = "domains"
	ElementsDType   = "elements"
	NationsDType    = "nations"
	WeaponsDType    = "weapons"
)

type DataTypes struct {
	Types []string `json:"types"`
}

type Artifact struct {
	Name           string `json:"name"`
	MaxRarity      int    `json:"max_rarity"`
	TwoPieceBonus  string `json:"2-piece_bonus"`
	FourPieceBonus string `json:"4-piece_bonus"`
}

type Character struct {
	Name           string   `json:"name"`
	Title          string   `json:"title"`
	Vision         string   `json:"vision"`
	Weapon         string   `json:"weapon"`
	Gender         string   `json:"gender"`
	Nation         string   `json:"nation"`
	Affiliation    string   `json:"affiliation"`
	SpecialDish    string   `json:"specialDish"`
	Rarity         int      `json:"rarity"`
	Constellation  string   `json:"constellation"`
	Birthday       string   `json:"birthday"`
	Description    string   `json:"description"`
	SkillTalents   []Talent `json:"skillTalents"`
	PassiveTalents []Talent `json:"passiveTalents"`
	Constellations []Talent `json:"constellations"`
}

type Talent struct {
	Name        string         `json:"name"`
	Unlock      string         `json:"unlock"`
	Description string         `json:"description"`
	Upgrades    []UpgradeLevel `json:"upgrades,omitempty"`
}

type UpgradeLevel struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Domain struct {
	Name                string               `json:"name"`
	Type                string               `json:"type"`
	Description         string               `json:"description"`
	Location            string               `json:"location"`
	Nation              string               `json:"nation"`
	Requirements        []RequirementLevel   `json:"requirements"`
	RecommendedElements []RecommendedElement `json:"recommendedElements"`
	Rewards             []Reward             `json:"rewards"`
}

type RequirementLevel struct {
	Level            string `json:"level"`
	AdventureRank    string `json:"adventureRank"`
	RecommendedLevel string `json:"recommendedLevel"`
	LeyLineDisorder  string `json:"leyLineDisorder"`
}

type RecommendedElement struct {
	Element string `json:"element"`
}

type Reward struct {
	Day     string         `json:"day"`
	Details []RewardDetail `json:"details"`
}

type RewardDetail struct {
	Level                   string `json:"level"`
	AdventureExperience     string `json:"adventureExperience"`
	CompanionshipExperience string `json:"companionshipExperience"`
	Mora                    string `json:"mora"`
	Drops                   []Drop `json:"drops"`
}

type Drop struct {
	Name string `json:"name"`
	Rate string `json:"rate"`
}

type Element struct {
	Name string `json:"name"`
}

type Nation struct {
	Name              string `json:"name"`
	Element           string `json:"element"`
	Archon            string `json:"archon"`
	ControllingEntity string `json:"controllingEntity"`
}

type Weapon struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Rarity      int    `json:"rarity"`
	BaseAttack  string `json:"baseAttack"`
	SubStat     string `json:"subStat"`
	PassiveName string `json:"passiveName"`
	PassiveDesc string `json:"passiveDesc"`
	Location    string `json:"location"`
}