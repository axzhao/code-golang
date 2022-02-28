package set

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

func ExampleSet() {
	requiredClasses := mapset.NewSet()
	requiredClasses.Add("Cooking")
	requiredClasses.Add("English")
	requiredClasses.Add("Math")
	requiredClasses.Add("Biology")

	scienceSlice := []interface{}{"Biology", "Chemistry"}
	scienceClasses := mapset.NewSetFromSlice(scienceSlice)

	electiveClasses := mapset.NewSet()
	electiveClasses.Add("Welding")
	electiveClasses.Add("Music")
	electiveClasses.Add("Automotive")

	bonusClasses := mapset.NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")

	allClasses := requiredClasses.Union(scienceClasses).Union(electiveClasses).Union(bonusClasses)
	fmt.Println(allClasses)

	fmt.Println(scienceClasses.Contains("Cooking"))
	fmt.Println(allClasses.Difference(scienceClasses))

	// Output:
	//Set{Cooking, English, Math, Chemistry, Welding, Biology, Music, Automotive, Go Programming, Python Programming}
	//false
	//Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}
}
