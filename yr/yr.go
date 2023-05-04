package yr

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mathhaug/funtemps/conv"
)

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

func CelsiusToFahrenheitLine(line string) (string, error) {
	dividedString := strings.Split(line, ";")
	var err error

	if len(dividedString) == 4 {
		if strings.HasPrefix(dividedString[0], "Data er gyldig") {
			return "Data er basert paa gyldig data (per 18.03.2023) (CCBY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Mathias Haugen", err
		}

		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil
}

func CountLines(input string) int {
	var fileName = input

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func AverageTemp(sum float64, count float64) float64 {
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		if count == 0 {
			count++
		}
		dividedString := strings.Split(scanner.Text(), ";")

		if len(dividedString) < 4 {
			continue // Skip lines that do not have expected format
		}

		if dividedString[3] == "Lufttemperatur" || strings.HasPrefix(dividedString[0], "Data er gyldig") {
			continue // Skip unwanted lines
		}

		num, err := strconv.ParseFloat(dividedString[3], 64)
		if err != nil {
			log.Fatalln(err)
		}
		sum += num
		count++
	}

	avg := sum / (count - 2)
	return avg
}

func AverageTempFahrenheit(sum float64, count float64) float64 {
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		if count == 0 {
			count++
		}
		dividedString := strings.Split(scanner.Text(), ";")

		if len(dividedString) < 4 {
			continue // Skip lines that do not have expected format
		}

		if dividedString[3] == "Lufttemperatur" || strings.HasPrefix(dividedString[0], "Data er gyldig") {
			continue // Skip unwanted lines
		}

		num, err := strconv.ParseFloat(dividedString[3], 64)
		if err != nil {
			log.Fatalln(err)
		}
		sum += num
		count++
	}

	avg := sum / (count - 2)
	return conv.CelsiusToFahrenheit(avg)
}
