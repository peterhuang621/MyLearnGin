package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("39.156.70.46:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// fmt.Printf("%d error at %s!!\n", p, err)
			results <- -p
			continue
		}
		conn.Close()
		fmt.Printf("%d opened!!\n", p)
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports, closedports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= 1024; i++ {
		port := <-results
		if port >= 0 {
			openports = append(openports, port)
		} else {
			closedports = append(closedports, port)
		}
	}

	defer close(ports)
	defer close(results)

	sort.Ints(openports)
	sort.Ints(closedports)

	for _, port := range closedports {
		fmt.Printf("%d closed\n", port)
	}

	for _, port := range openports {
		fmt.Printf("%d opened\n", port)
	}
}

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	for i := 21; i < 120; i++ {
// 		wg.Add(1)
// 		go func(j int) {
// 			defer wg.Done()
// 			address := fmt.Sprintf("39.156.70.46:%d", j)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				fmt.Printf("%s closed!\n", address)
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("%s opened!\n", address)
// 			elapsed := time.Since(start) / 1e6
// 			fmt.Printf("\n\n%d millseconds", elapsed)
// 		}(i)
// 	}
// 	wg.Wait()
// }
