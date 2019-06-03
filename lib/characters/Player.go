package characters

// Player contains the player character data
type Player struct {
	Name          string `json:"name"`
	Strength      int    `json:"strength"`
	Dexterity     int    `json:"dexterity"`
	Stamina       int    `json:"stamina"`
	Charisma      int    `json:"charisman"`
	Manipulation  int    `json:"manipulation"`
	Composure     int    `json:"composure"`
	Intelligence  int    `json:"intelligence"`
	Wits          int    `json:"wits"`
	Resolve       int    `json:"resolve"`
	Athletics     int    `json:"athletics"`
	Brawl         int    `json:"brawl"`
	Craft         int    `json:"craft"`
	Drive         int    `json:"drive"`
	Firearms      int    `json:"firearms"`
	Larceny       int    `json:"larceny"`
	Melee         int    `json:"melee"`
	Stealth       int    `json:"stealth"`
	Survival      int    `json:"survival"`
	AnimalKen     int    `json:"animalken"`
	Etiquette     int    `json:"etiquette"`
	Insight       int    `json:"insight"`
	Intimidation  int    `json:"intimidation"`
	Leadership    int    `json:"leadership"`
	Performance   int    `json:"performance"`
	Persuasion    int    `json:"persuasion"`
	Streetwise    int    `json:"streetwise"`
	Subterfuge    int    `json:"subterfuge"`
	Academics     int    `json:"academics"`
	Awareness     int    `json:"awareness"`
	Finance       int    `json:"finance"`
	Investigation int    `json:"investigation"`
	Medicine      int    `json:"medicine"`
	Occult        int    `json:"occult"`
	Politics      int    `json:"politics"`
	Science       int    `json:"science"`
	Technology    int    `json:"technology"`
}
