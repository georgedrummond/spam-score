package message

import (
	"fmt"
)

type MessageStats struct {
	Processed string `json:"processed"`
	Requested string `json:"requested"`
	Failed    string `json:"failed"`
}

func Stats() MessageStats {
	client.Incr("hello")
	processed, _ := client.Get("hello")

	fmt.Println(processed)

	stats := MessageStats{
		Processed: string(processed),
		Requested: string(processed),
		Failed:    string(processed),
	}

	return stats
}
