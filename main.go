package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
	"time"
)

// Hardcoded API key - SECURITY ISSUE
const API_KEY = "sk-1234567890abcdef"

var db *sql.DB

func main() {
	// Demo 1: SQL Injection vulnerability
	userID := getUserInput()
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", userID)
	fmt.Printf("Query: %s\n", query)

	// Demo 2: Inefficient loop
	data := make([]int, 1000000)
	result := processData(data)
	fmt.Printf("Processed: %d\n", result)

	// Demo 3: Missing error handling
	hashPassword("admin123")

	// Demo 4: Race condition
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter: %d\n", counter)

	// Demo 5: Division by zero
	num := divide(10, 0)
	fmt.Printf("Division: %d\n", num)
}

func getUserInput() string {
	return "1 OR 1=1"
}

func processData(data []int) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}

func hashPassword(password string) string {
	// MD5 is deprecated for password hashing - SECURITY ISSUE
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func divide(a, b int) int {
	// No error handling for division by zero
	return a / b
}

func unsafeSliceAccess(arr []int, index int) int {
	// No bounds checking
	return arr[index]
}

func initDB() {
	// Missing error handling
	db, _ = sql.Open("mysql", "user:password@/dbname")
}

func processWithTimeout() {
	// No timeout handling
	ch := make(chan string)
	result := <-ch
	fmt.Println(result)
}

func logSensitiveData() {
	// Logging sensitive data
	log.Printf("Processing payment with card: %s", "4111-1111-1111-1111")
}

func inefficientStringConcat() []string {
	// Inefficient string concatenation in loop
	result := ""
	for i := 0; i < 1000; i++ {
		result += "data"
	}
	return []string{result}
}

func missingNilCheck(data *string) {
	// No nil check before dereferencing
	fmt.Println(*data)
}

func sleepWithoutContext() {
	// Sleep without context cancellation
	time.Sleep(10 * time.Second)
}
