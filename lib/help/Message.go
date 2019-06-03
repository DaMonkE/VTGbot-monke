package help

// Message returns the help message
func Message() (s string) {
	s = "Discord Bot for Vampire the Masquerade\n"
	s += "    Created by DaMonkE(Ryan)\n\n"
	s += "To use this bot, send a message to the channel beginning with:\n"
	s += "    /create [NAME] -- Creates a blank character sheet\n"
	s += "    /set [NAME] <ATTRIB/SKILL> N -- Sets the Attribute or Skill indicated to N # of dots\n"
	s += "    /get [NAME] <ATTRIB/SKILL> -- Shows the number of dots in the Attribute or Skill indicated\n"
	s += "    /roll NdS -- Rolls N # of dice with S # of sides, /roll by itself rolls a single d10\n"
	s += "    /roll [NAME] <ATTRIB> <SKILL> [MOD] -- Rolls the Attribute + Skill + optional Modifier\n"
	s += "        Rolls of 10 will explode\n"
	s += "\n [NAME] is optional and if omitted will default to message sender\n"
	s += "Thanks for trying it out!\n"
	return
}
