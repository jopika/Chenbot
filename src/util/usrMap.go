package util

var PersonNameIDMap map[string]string

func IntializeMapping() map[string]string {
	if PersonNameIDMap == nil {
		PersonNameIDMap = make(map[string]string)

		PersonNameIDMap["Jenna"] = "UKQA8VBHR"
		PersonNameIDMap["Sam"] = "UL5194VE2"
		PersonNameIDMap["Chen"] = "UL1MWS8D6"
		PersonNameIDMap["Shyam"] = "UL3HFQPHD"
		PersonNameIDMap["Derek"] = "UKQBVGZGT"
		PersonNameIDMap["Wenting"] = "UKQ9P39B4"
		PersonNameIDMap["Trevor"] = "ULJTDFB4H"
		PersonNameIDMap["Jonathan"] = "UL1MY9DLY"
	}

	return PersonNameIDMap
}







