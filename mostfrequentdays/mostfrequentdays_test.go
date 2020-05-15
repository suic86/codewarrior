package mostfrequentdays_test

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/suic86/codewarrior/mostfrequentdays"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func test(input int, expected ...string) {
	Expect(MostFrequentDays(input)).To(Equal(expected))
}

var _ = Describe("Sample Tests", func() {
	It("Sample Tests", func() {
		test(2427, "Friday")
		test(2185, "Saturday")
		test(1167, "Sunday")
		test(1492, "Friday", "Saturday")
		test(1770, "Monday")
		test(1785, "Saturday")
		test(1984, "Monday", "Sunday")
		test(3076, "Saturday", "Sunday")
	})
})

var _ = Describe("Test Suite", func() {
	It("Fixed Tests", func() {
		test(1492, "Friday", "Saturday")
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
		test(5863, "Thursday")
		test(7548, "Thursday", "Friday")
		test(8116, "Wednesday", "Thursday")
		test(9137, "Friday")
	})

	mondaySunday := []int{1404, 1432, 1460, 1488, 1528, 1556, 1584, 1612, 1640, 1668, 1696, 1708, 1736, 1764, 1792, 1804, 1832, 1860, 1888, 1928, 1956, 1984, 2012, 2040, 2068, 2096, 2108, 2136, 2164, 2192, 2204, 2232, 2260, 2288, 2328, 2356, 2384, 2412, 2440, 2468, 2496, 2508, 2536, 2564, 2592, 2604, 2632, 2660, 2688, 2728, 2756, 2784, 2812, 2840, 2868, 2896, 2908, 2936, 2964, 2992, 3004, 3032, 3060, 3088, 3128, 3156, 3184, 3212, 3240, 3268, 3296, 3308, 3336, 3364, 3392, 3404, 3432, 3460, 3488, 3528, 3556, 3584, 3612, 3640, 3668, 3696, 3708, 3736, 3764, 3792, 3804, 3832, 3860, 3888, 3928, 3956, 3984, 4012, 4040, 4068, 4096, 4108, 4136, 4164, 4192, 4204, 4232, 4260, 4288, 4328, 4356, 4384, 4412, 4440, 4468, 4496, 4508, 4536, 4564, 4592, 4604, 4632, 4660, 4688, 4728, 4756, 4784, 4812, 4840, 4868, 4896, 4908, 4936, 4964, 4992, 5004, 5032, 5060, 5088, 5128, 5156, 5184, 5212, 5240, 5268, 5296, 5308, 5336, 5364, 5392, 5404, 5432, 5460, 5488, 5528, 5556, 5584, 5612, 5640, 5668, 5696, 5708, 5736, 5764, 5792, 5804, 5832, 5860, 5888, 5928, 5956, 5984, 6012, 6040, 6068, 6096, 6108, 6136, 6164, 6192, 6204, 6232, 6260, 6288, 6328, 6356, 6384, 6412, 6440, 6468, 6496, 6508, 6536, 6564, 6592, 6604, 6632, 6660, 6688, 6728, 6756, 6784, 6812, 6840, 6868, 6896, 6908, 6936, 6964, 6992, 7004, 7032, 7060, 7088, 7128, 7156, 7184, 7212, 7240, 7268, 7296, 7308, 7336, 7364, 7392, 7404, 7432, 7460, 7488, 7528, 7556, 7584, 7612, 7640, 7668, 7696, 7708, 7736, 7764, 7792, 7804, 7832, 7860, 7888, 7928, 7956, 7984, 8012, 8040, 8068, 8096, 8108, 8136, 8164, 8192, 8204, 8232, 8260, 8288, 8328, 8356, 8384, 8412, 8440, 8468, 8496, 8508, 8536, 8564, 8592, 8604, 8632, 8660, 8688, 8728, 8756, 8784, 8812, 8840, 8868, 8896, 8908, 8936, 8964, 8992, 9004, 9032, 9060, 9088, 9128, 9156, 9184, 9212, 9240, 9268, 9296, 9308, 9336, 9364, 9392, 9404, 9432, 9460, 9488, 9528, 9556, 9584, 9612, 9640, 9668, 9696, 9708, 9736, 9764, 9792, 9804, 9832, 9860, 9888, 9928, 9956, 9984}

	It("Random Tests", func() {
		for i := 0; i < 90; i++ {
			ry := 1593 + rand.Intn(8408)
			Expect(MostFrequentDays(ry)).Should(Equal(reference(ry)))
		}

		rand.Shuffle(len(mondaySunday), func(i, j int) {
			mondaySunday[i], mondaySunday[j] = mondaySunday[j], mondaySunday[i]
		})

		for _, y := range mondaySunday[:10] {
			Expect(MostFrequentDays(y)).Should(Equal(reference(y)))
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
