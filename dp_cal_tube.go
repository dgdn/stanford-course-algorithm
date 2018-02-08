package main

import (
	"fmt"
	"math"
	"time"
)

var tubeCapacities = map[string]int{
	"1": 5,
	"2": 5,
	"3": 3,
	"4": 5,
	"5": 3,
}

var tubeTypes = []string{"1", "2", "3"}

var cache = make(map[string]map[string]int)

type Tube struct {
	Type     string
	Capacity int
}

func main() {

	//计算检测项目所需试管数量
	//每种试管都有一个总容量
	//每个检测项目都有两种试管选择，每种试管分别需要对应的血量
	//目标是选出试管组合，使得检测这些项目所需试管数量最少

	//项目和可选的两种试管
	datas := [][2]Tube{
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"2", 2}, {"3", 1}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 3}},
		{{"1", 3}, {"2", 2}},
		{{"1", 2}, {"2", 2}},
		{{"1", 3}, {"2", 4}},
		{{"1", 3}, {"2", 2}},
		{{"1", 2}, {"3", 2}},
		{{"1", 1}, {"3", 3}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"1", 3}, {"2", 2}},
		{{"1", 2}, {"2", 2}},
		{{"1", 4}, {"2", 2}},
	}

	condition := make(map[string]int)

	ts := time.Now()

	result := dp(condition, datas)
	fmt.Println("tube combination:", result)
	fmt.Println("need tubes count:", calculateTubeCount(result))

	fmt.Println("calculate duration:", time.Now().Sub(ts))
}

func dp(condition map[string]int, datas [][2]Tube) map[string]int {

	if len(datas) == 1 {
		return calculateBaseCaseResult(condition, datas[0])
	}

	last := datas[len(datas)-1]
	//tube 1 case
	base1, cond1 := decomposeCondition(condition, last[0])
	//calculate subproblem
	result1 := merge(cacheDP(cond1, datas[:len(datas)-1]), base1)

	//tube 2 case
	base2, cond2 := decomposeCondition(condition, last[1])
	//calculate subproblem
	result2 := merge(cacheDP(cond2, datas[:len(datas)-1]), base2)

	//make desicion base on two subproblem
	return calculateOptimalResult(result1, result2)
}

func decomposeCondition(condition map[string]int, tube Tube) (base, cond map[string]int) {
	condition = appendCondition(condition, tube)

	base = make(map[string]int)
	cond = make(map[string]int)
	for typ, val := range condition {
		if r := val % tubeCapacities[typ]; r == 0 {
			base[typ] = val
		} else {
			base[typ] = val - r
			cond[typ] = r
		}
	}
	return
}

func cacheDP(condition map[string]int, datas [][2]Tube) map[string]int {
	var key = fmt.Sprintf("step%d-", len(datas))
	for i, typ := range tubeTypes {
		if i != 0 {
			key += "-"
		}
		key += fmt.Sprintf("%s:%d", typ, condition[typ])
	}
	if v, ok := cache[key]; ok {
		return newMapFrom(v)
	}
	v := dp(condition, datas)
	cache[key] = newMapFrom(v)
	return v
}

func merge(n1, n2 map[string]int) map[string]int {
	m := make(map[string]int)
	for k, v := range n1 {
		m[k] += v
	}
	for k, v := range n2 {
		m[k] += v
	}
	return m
}

func calculateBaseCaseResult(condition map[string]int, tubes [2]Tube) map[string]int {
	result1 := appendCondition(condition, tubes[0])
	result2 := appendCondition(condition, tubes[1])
	return calculateOptimalResult(result1, result2)
}

func appendCondition(condition map[string]int, tube Tube) map[string]int {
	ncond := newMapFrom(condition)
	ncond[tube.Type] += tube.Capacity
	return ncond
}

func calculateOptimalResult(result1, result2 map[string]int) map[string]int {

	tubeCount1 := calculateTubeCount(result1)
	tubeCount2 := calculateTubeCount(result2)

	//试管相同 血量少的优先
	if tubeCount1 == tubeCount2 {
		if sumValue(result1) > sumValue(result2) {
			return newMapFrom(result2)
		}
		return newMapFrom(result1)
	}

	//试管优先
	if tubeCount1 > tubeCount2 {
		return newMapFrom(result2)
	}
	return newMapFrom(result1)
}

func newMapFrom(src map[string]int) map[string]int {
	m := make(map[string]int)
	for k, v := range src {
		m[k] = v
	}
	return m
}

func sumValue(result map[string]int) int {
	var val int
	for _, v := range result {
		val += v
	}
	return val
}

func calculateTubeCount(result map[string]int) int {
	var count int
	for typ, val := range result {
		count += int(math.Ceil(float64(val) / float64(tubeCapacities[typ])))
	}
	return count
}
