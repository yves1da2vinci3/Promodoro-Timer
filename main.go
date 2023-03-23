package main

import (
	"fmt"
	"time"
)

const (
	pomodoroTime   = 25 * time.Minute
	shortBreakTime = 5 * time.Minute
	longBreakTime  = 15 * time.Minute
	numPomodoros   = 4
)

func main() {
	pomodorosCompleted := 0

	fmt.Println("Welcome to the Pomodoro timer!")

	for pomodorosCompleted < numPomodoros {
		fmt.Printf("\nStarting pomodoro %d/%d...\n", pomodorosCompleted+1, numPomodoros)
		startTimer(pomodoroTime)

		if pomodorosCompleted < numPomodoros-1 {
			fmt.Println("\nTake a short break...")
			startTimer(shortBreakTime)
		} else {
			fmt.Println("\nTake a long break...")
			startTimer(longBreakTime)
		}

		pomodorosCompleted++
	}

	fmt.Println("\nCongratulations! You have completed all the pomodoros.")
}

func startTimer(duration time.Duration) {
	startTime := time.Now()
	endTime := startTime.Add(duration)

	for time.Now().Before(endTime) {
		timeLeft := endTime.Sub(time.Now())
		fmt.Printf("\rTime left: %v", timeLeft.Truncate(time.Second))
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\rTime's up!")
}
