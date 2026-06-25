package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	values := strings.Split(datastring, ",")

	if len(values) != 2 {
		return errors.New("There are not enough values in the row")
	}

	steps, err := strconv.Atoi(values[0])

	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("The steps must be greater than 0")
	}

	duration, err := time.ParseDuration(values[1])

	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("The duration must be greater than 0")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Height <= 0 {
		return "", errors.New("The height must be greater than 0")
	}

	if ds.Weight <= 0 {
		return "", errors.New("The weight must be greater than 0")
	}

	if ds.Steps <= 0 {
		return "", errors.New("The steps must be greater than 0")
	}

	if ds.Duration <= 0 {
		return "", errors.New("The duration must be greater than 0")
	}

	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	currDistance := spentenergy.Distance(ds.Steps, ds.Height)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, currDistance, spentCalories), nil

}
