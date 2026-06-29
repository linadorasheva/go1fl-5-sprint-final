package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	values := strings.Split(datastring, ",")

	if len(values) != 3 {
		return errors.New("There are not enough values in the row")
	}

	steps, err := strconv.Atoi(values[0])

	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("The steps must be greater than 0")
	}

	duration, err := time.ParseDuration(values[2])

	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("The duration must be greater than 0")
	}

	trainingType := values[1]

	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Height <= 0 {
		return "", errors.New("The height must be greater than 0")
	}
	if t.Weight <= 0 {
		return "", errors.New("The weight must be greater than 0")
	}
	if t.Steps <= 0 {
		return "", errors.New("The steps must be greater than 0")
	}

	var spentCalories float64 = 0
	var err error

	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		if err != nil {
			return "", err
		}
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		if err != nil {
			return "", err
		}

	default:
		return "", errors.New("Unknown type of training")
	}

	currDistance := spentenergy.Distance(t.Steps, t.Height)
	currMeanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), currDistance, currMeanSpeed, spentCalories), nil

}
