package _022_07_24

import (
	"sort"
	"testing"
)

type food struct {
	foodName string
	cuisine  string
	rating   int
}

type FoodRatings struct {
	foodMp          map[string][]*food
	foodNameMp      map[string]*food
	cuisineNeedSort map[string]bool
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodMp := map[string][]*food{}
	foodNameMp := map[string]*food{}
	for i := 0; i < len(foods); i++ {
		f := &food{foodName: foods[i], cuisine: cuisines[i], rating: ratings[i]}
		foodMp[cuisines[i]] = append(foodMp[cuisines[i]], f)
		foodNameMp[foods[i]] = f
	}
	for _, v := range foodMp {
		v = sortFood(v)
	}
	return FoodRatings{foodMp: foodMp, foodNameMp: foodNameMp, cuisineNeedSort: map[string]bool{}}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	f := this.foodNameMp[food]
	f.rating = newRating
	this.cuisineNeedSort[f.cuisine] = true
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	if this.cuisineNeedSort[cuisine] {
		this.foodMp[cuisine] = sortFood(this.foodMp[cuisine])
		this.cuisineNeedSort[cuisine] = false
	}
	return this.foodMp[cuisine][0].foodName
}
func sortFood(list []*food) []*food {
	v := list
	sort.Slice(v, func(i, j int) bool {
		if v[i].rating > v[j].rating {
			return true
		} else if v[i].rating == v[j].rating {
			return v[i].foodName < v[j].foodName
		}
		return false
	})
	return v
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

func TestFoodRatings(t *testing.T) {

}
