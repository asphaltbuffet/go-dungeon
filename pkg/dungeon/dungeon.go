package dungeon

type Dungeon struct {
	Lang  string `json:"lang"`
	Title string `json:"title"`
	About string `json:"about"`
	Map   struct {
		Theme       []string   `json:"theme"`
		Type        []string   `json:"type"`
		Adjectives  [][]string `json:"adjectives"`
		Nouns       [][]string `json:"nouns"`
		Discoveries struct {
			SpecialRoom [][]string `json:"special_room"`
			Feature     []string   `json:"feature"`
			Item        []string   `json:"item"`
			Treasure    []string   `json:"treasure"`
			Passive     []string   `json:"passive"`
		} `json:"discoveries"`
		Dangers struct {
			Hazard    []string `json:"hazard"`
			Trap      []string `json:"trap"`
			Encounter []string `json:"encounter"`
		} `json:"dangers"`
		Dressing struct {
			Natural  []string `json:"natural"`
			ManMade  []string `json:"man-made"`
			Lighting []string `json:"lighting"`
			Odor     []string `json:"odor"`
			Odd      []string `json:"odd"`
			Mystical []string `json:"mystical"`
		} `json:"dressing"`
		Descriptions [][]string `json:"descriptions"`
		History      struct {
			Destruction []string `json:"destruction"`
			Builder     []string `json:"builder"`
			Purpose     []string `json:"purpose"`
		} `json:"history"`
		Denizens struct {
			Humanoid  []string `json:"humanoid"`
			Creature  []string `json:"creature"`
			Unnatural []string `json:"unnatural"`
		} `json:"denizens"`
		Loot      [][]string `json:"loot"`
		Explore   []string   `json:"explore"`
		Omen      []string   `json:"omen"`
		Discovery []string   `json:"discovery"`
		Danger    []string   `json:"danger"`
		Reaction  []string   `json:"reaction"`
		Door      []string   `json:"door"`
	} `json:"map"`
	Denizens struct {
		Tier1 struct {
			Humanoid  []string `json:"humanoid"`
			Creature  []string `json:"creature"`
			Unnatural []string `json:"unnatural"`
		} `json:"tier1"`
		Tier2 struct {
			Humanoid  []string `json:"humanoid"`
			Creature  []string `json:"creature"`
			Unnatural []string `json:"unnatural"`
		} `json:"tier2"`
	} `json:"denizens"`
	Boss []struct {
		Name        string `json:"name"`
		Stats       string `json:"stats"`
		Attack      string `json:"attack,omitempty"`
		Description string `json:"description"`
	} `json:"boss"`
	Encounters struct {
		Num2 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"2"`
		Num3 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"3"`
		Num4 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"4"`
		Num5 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"5"`
		Num6 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"6"`
		Num7 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"7"`
		Num8 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"8"`
		Num9 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"9"`
		Num10 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"10"`
		Num11 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"11"`
		Num12 struct {
			Name        string   `json:"name"`
			Description []string `json:"description"`
		} `json:"12"`
	} `json:"encounters"`
}
