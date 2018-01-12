package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Job struct {
	Weight     int
	Length     int
	RatioScore float32 // equal to weight/length
	DiffScore  int     // equal to weight/length
}

func main() {

	f, err := os.Open("job.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var jobs []Job

	for line := 0; scanner.Scan(); line++ {
		//ignore the first line data(total number)
		if line == 0 {
			continue
		}
		lineData := strings.Fields(scanner.Text())

		weight, err := strconv.Atoi(lineData[0])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(lineData[1])
		if err != nil {
			panic(err)
		}

		var job Job
		job.Weight = weight
		job.Length = length
		job.RatioScore = float32(weight) / float32(length)
		job.DiffScore = weight - length

		jobs = append(jobs, job)
	}

	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].DiffScore == jobs[j].DiffScore {
			return jobs[i].Weight > jobs[i].Weight
		}
		return jobs[i].DiffScore > jobs[j].DiffScore
	})

	var accLength, totalCompleteTime int
	for _, job := range jobs {
		accLength += job.Length
		totalCompleteTime += job.Weight * accLength
	}
	fmt.Println("difference:", totalCompleteTime)

	sort.Slice(jobs, func(i, j int) bool {
		//no need to handle tie
		return jobs[i].RatioScore > jobs[j].RatioScore
	})

	accLength, totalCompleteTime = 0, 0
	for _, job := range jobs {
		accLength += job.Length
		totalCompleteTime += job.Weight * accLength
	}
	fmt.Println("ratio:", totalCompleteTime)

}
