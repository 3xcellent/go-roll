package roll

import (
	"fmt"
)

func Simple(r Roll) string {
	return fmt.Sprintf("%d", r.CalculatedRoll)
}

func Verbose(r Roll) string {
	rollsStr := ""
	modifierStr := ""

	if len(r.rolls) > 1 {
		rollsStr = "Rolls: "
		for idx, amt := range r.rolls {
			if idx != 0 {
				rollsStr = rollsStr + ", "
			}
			rollsStr += fmt.Sprintf("%d", amt)

		}
		rollsStr += "\n"
	}
	if r.modifier != 0 {
		modifierDirection := ""
		if r.modifier < 0 {
			modifierDirection = "-"
		}
		if r.modifier > 0 {
			modifierDirection = "+"
		}
		modifierStr = fmt.Sprintf("Modifier: %s%d\n", modifierDirection, r.modifier)
	}

	if len(r.rolls) < 2 && r.modifier == 0 {
		return fmt.Sprintf("Roll: %d (min/max %d/%d)",
			r.CalculatedRoll,
			1*len(r.rolls)+r.modifier,
			r.maxScore*len(r.rolls)+r.modifier)
	} else {
		totalStr := fmt.Sprintf("Total: %d (min/max %d/%d)",
			r.CalculatedRoll,
			1*len(r.rolls)+r.modifier,
			r.maxScore*len(r.rolls)+r.modifier)
		return fmt.Sprintf("%s%s%s", rollsStr, modifierStr, totalStr)
	}
}
