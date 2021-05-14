package main

import (
	"./exit"
	"lesson10/input"
	"lesson10/print"
	"lesson10/config"
	"fmt"
)


func main() {
	var (
		isValid bool
		patternNum int
		lineCount int
		repeatCount int
	)

	sig := exit.SignalHandle()
	go exit.Interrupt(sig)
	step := config.InputPatternNumStep

	for {
		switch step{
		case config.InputPatternNumStep :
			fmt.Println(config.PatternNumGuide)
			patternNum, isValid = input.GetValidInput(config.PatternNumInputGuide,config.PatternNumStartRange,config.PatternNumEndRange)
		case config.InputPatternRepeatCountStep:
			repeatCount, isValid = input.GetPatternRepeatCount(patternNum)
		case config.InputLineCountStep:
			for {
				lineCount, isValid = input.GetValidInput(config.LineCountInputGuide,config.InputStartRange,config.InputEndRange)
				if patternNum == config.Diamond && !input.IsValidDiamondLineCount(lineCount) {
					continue
				}	else {
					break
				}
			}
		}

		if !isValid{
			return
		}
		step += config.NextStep

		switch step {
		case config.PrintPatternStep:
			print.Pattern(patternNum, lineCount, repeatCount)
			step = config.InputPatternNumStep
			continue
		}
	}
}



