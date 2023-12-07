package days

import "fmt"

type Day6 struct{}

func (d Day6) Part1(_ []string) string {
	race1 := []int{42, 284}
	race2 := []int{68, 1005}
	race3 := []int{69, 1122}
	race4 := []int{85, 1341}

	calcQuotient := func(ms int, mm int) (quotient int) {
		earliestWinningTime := 0
		latestWinningTime := 0
		for i := 0; i < ms; i++ {
			speed := i
			remainingTime := ms - i
			distanceTravelled := speed * remainingTime
			if distanceTravelled > mm {
				if earliestWinningTime == 0 {
					earliestWinningTime = i
				}
				latestWinningTime = i
			}
		}
		fmt.Println(earliestWinningTime, latestWinningTime)
		return latestWinningTime - earliestWinningTime + 1
	}

	quo1 := calcQuotient(race1[0], race1[1])
	quo2 := calcQuotient(race2[0], race2[1])
	quo3 := calcQuotient(race3[0], race3[1])
	quo4 := calcQuotient(race4[0], race4[1])

	return fmt.Sprintf("%d", quo1*quo2*quo3*quo4)
}

func (d Day6) Part2(_ []string) string {
	raceLong := []int{42686985, 284100511221341}

	earliestWin := func(ms int, mm int) (earliest int) {
		for i := 0; i < ms; i++ {
			speed := i
			remainingTime := ms - i
			distanceTravelled := speed * remainingTime
			if distanceTravelled > mm {
				if earliest == 0 {
					earliest = i
				}
			}
		}
		return earliest
	}(raceLong[0], raceLong[1])
	fmt.Println(earliestWin)

	latestWin := func(ms int, mm int) (latest int) {
		for i := ms; i > 0; i-- {
			speed := i
			remainingTime := ms - i
			distanceTravelled := speed * remainingTime
			if distanceTravelled > mm {
				if latest == 0 {
					latest = i
				}
			}
		}
		return latest
	}(raceLong[0], raceLong[1])

	return fmt.Sprintf("%d", latestWin-earliestWin+1)
}

func NewDay6() Day6 {
	return Day6{}
}
