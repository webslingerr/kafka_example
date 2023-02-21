package helper

import (
	"math/rand"
	"time"
)

func GetRandomColor() string {
	colors := []string{
		"#9f1616",
		"#278d39",
		"#e1c13b",
		"#00d3ff",
		"#4703a6",
		"#f205cb",
		"#f2463c",
		"#f3cfcf",
		"#761954",
		"#87ca9d",
	}

	rand.Seed(time.Now().UnixNano())

	return colors[rand.Intn(10)]
}
