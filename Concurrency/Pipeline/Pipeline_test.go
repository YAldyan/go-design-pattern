package Pipeline

import "testing"

func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)
}

func LaunchPipelineX(amount int) int {
	return <-sum(power(generator(amount)))
}

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}
	// ...

	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, test[1])
	}
}
