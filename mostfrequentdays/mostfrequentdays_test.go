package mostfrequentdays_test

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/mostfrequentdays"
)

func test(input int, expected ...string) {
	Expect(MostFrequentDays(input)).To(Equal(expected))
}

var _ = Describe("Sample Tests", func() {
	It("Sample Tests", func() {
		test(1770, "Monday")
		test(1785, "Saturday")
		test(1984, "Monday", "Sunday")
		test(2185, "Saturday")
		test(2427, "Friday")
		test(3076, "Saturday", "Sunday")
	})
})

// Years with edge case results
var mondaySunday = []int{1612, 1640, 1668, 1696, 1708, 1736, 1764, 1792, 1804, 1832, 1860, 1888, 1928, 1956, 1984, 2012, 2040, 2068, 2096, 2108, 2136, 2164, 2192, 2204, 2232, 2260, 2288, 2328, 2356, 2384, 2412, 2440, 2468, 2496, 2508, 2536, 2564, 2592, 2604, 2632, 2660, 2688, 2728, 2756, 2784, 2812, 2840, 2868, 2896, 2908, 2936, 2964, 2992, 3004, 3032, 3060, 3088, 3128, 3156, 3184, 3212, 3240, 3268, 3296, 3308, 3336, 3364, 3392, 3404, 3432, 3460, 3488, 3528, 3556, 3584, 3612, 3640, 3668, 3696, 3708, 3736, 3764, 3792, 3804, 3832, 3860, 3888, 3928, 3956, 3984}

var saturdaySunday = []int{1600, 1628, 1656, 1684, 1724, 1752, 1780, 1820, 1848, 1876, 1916, 1944, 1972, 2000, 2028, 2056, 2084, 2124, 2152, 2180, 2220, 2248, 2276, 2316, 2344, 2372, 2400, 2428, 2456, 2484, 2524, 2552, 2580, 2620, 2648, 2676, 2716, 2744, 2772, 2800, 2828, 2856, 2884, 2924, 2952, 2980, 3020, 3048, 3076, 3116, 3144, 3172, 3200, 3228, 3256, 3284, 3324, 3352, 3380, 3420, 3448, 3476, 3516, 3544, 3572, 3600, 3628, 3656, 3684, 3724, 3752, 3780, 3820, 3848, 3876, 3916, 3944, 3972, 4000}

var _ = Describe("Test Suite", func() {
	It("Fixed Tests", func() {
		test(1770, "Monday")
		test(1785, "Saturday")
		test(1794, "Wednesday")
		test(1901, "Tuesday")
		test(1910, "Saturday")
		test(1968, "Monday", "Tuesday")
		test(1984, "Monday", "Sunday")
		test(1986, "Wednesday")
		test(2001, "Monday")
		test(2016, "Friday", "Saturday")
		test(2135, "Saturday")
		test(3043, "Sunday")
		test(3150, "Sunday")
		test(3230, "Tuesday")
		test(3361, "Thursday")
	})

	It("Random Tests", func() {
		for i := 0; i < 90; i++ {
			ry := 1593 + rand.Intn(8408)
			Expect(MostFrequentDays(ry)).Should(Equal(reference(ry)))
		}

		rand.Shuffle(len(mondaySunday), func(i, j int) {
			mondaySunday[i], mondaySunday[j] = mondaySunday[j], mondaySunday[i]
		})

		rand.Shuffle(len(saturdaySunday), func(i, j int) {
			saturdaySunday[i], saturdaySunday[j] = saturdaySunday[j], saturdaySunday[i]
		})

		for _, y := range mondaySunday[:10] {
			Expect(MostFrequentDays(y)).Should(Equal(reference(y)))
		}

	})

})

var _ = Describe("Edge Cases", func() {
	It("Saturday, Sunday", func() {
		for _, y := range saturdaySunday {
			test(y, "Saturday", "Sunday")
		}
	})
	It("Monday, Sunday", func() {
		for _, y := range mondaySunday {
			test(y, "Monday", "Sunday")
		}
	})
})

func reference(year int) []string {
	beg := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).Weekday()
	end := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).Weekday()
	if beg == end {
		return []string{beg.String()}
	}
	if beg == time.Sunday {
		beg, end = end, beg
	}
	return []string{beg.String(), end.String()}
}
