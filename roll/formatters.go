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

	// rolls
	if len(r.rolls) > 1 {
		rollsStr = "Rolls: "
		if r.chooseLow || r.chooseHigh {
			var first int
			for idx, amt := range r.rolls {

				if idx%2 == 0 {
					first = amt
				} else {
					rollsStr += fmt.Sprintf("(%d, %d)", first, amt)
					if idx != 0 && idx != len(r.rolls)-1 {
						rollsStr = rollsStr + ", "
					}
				}
			}
		} else {
			for idx, amt := range r.rolls {
				if idx != 0 {
					rollsStr = rollsStr + ", "
				}
				rollsStr += fmt.Sprintf("%d", amt)

			}
		}
		rollsStr += "\n"
	}

	// modifier
	if r.modifier != 0 {
		modifierDirection := ""
		if r.modifier > 0 {
			modifierDirection = "+"
		}
		modifierStr = fmt.Sprintf("Modifier: %s%d\n", modifierDirection, r.modifier)
	}

	// total
	if !r.chooseHigh && !r.chooseLow && len(r.rolls) < 2 && r.modifier == 0 {
		return fmt.Sprintf("Roll: %d (min/max %d/%d)",
			r.CalculatedRoll,
			1*len(r.rolls)+r.modifier,
			r.maxScore*len(r.rolls)+r.modifier)
	} else {
		adjustedLength := len(r.rolls)

		if r.chooseLow || r.chooseHigh {
			adjustedLength = adjustedLength / 2
		}
		totalStr := fmt.Sprintf("Total: %d (min/max %d/%d)",
			r.CalculatedRoll,
			adjustedLength+r.modifier,
			r.maxScore*adjustedLength+r.modifier)
		return fmt.Sprintf("%s%s%s", rollsStr, modifierStr, totalStr)
	}
}
