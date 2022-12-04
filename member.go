package member

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Member struct {
	id          uuid.UUID
	firstName   string
	lastName    string
	dateOfBirth time.Time
	joinedOn    time.Time
	active      bool
}

const (
	INT_SIZE    = 32
	DATE_LAYOUT = "01-02-2006"
	SHORT_FORM  = "2006-Jan-02"
	COLOR_RESET = "\033[0m"
	COLOR_RED   = "\033[31m"
)

func CreateMember(input *string) (*Member, error) {

	member := new(Member)
	data := strings.Split(*input, ",")
	var tmpDateOfBirth time.Time

	firstName := strings.TrimSpace(data[0])
	lastName := strings.TrimSpace(data[1])

	validateField("first name", &firstName)
	validateField("last name", &lastName)

	if len(data) < 3 {
		return nil, errors.New("you must provide 3 comma separated values")
	} else {
		dateOfBirth := strings.TrimSpace(data[2])
		validateField("date of birth", &dateOfBirth)
		dob, err := time.Parse(DATE_LAYOUT, dateOfBirth)
		if err != nil {
			return nil, err
		}
		tmpDateOfBirth = dob
	}

	(*member).id = uuid.New()
	(*member).firstName = strings.TrimSpace(data[0])
	(*member).lastName = strings.TrimSpace(data[1])
	(*member).dateOfBirth = tmpDateOfBirth
	(*member).joinedOn = time.Now()
	(*member).active = true

	return member, nil
}

func (member *Member) Print() {
	fmt.Println()
	fmt.Printf("%-20s%-38s\n", "ID:", (*member).id.String())
	fmt.Printf("%-20s%-38s\n", "Name:", (*member).lastName+", "+(*member).firstName)
	fmt.Printf("%-20s%-38s\n", "Date of Birth:", (*member).dateOfBirth.Format(SHORT_FORM))
	fmt.Printf("%-20s%-38s\n", "Joined On:", (*member).joinedOn.Format(SHORT_FORM))
	fmt.Printf("%-20s%-38t\n", "Active:", (*member).active)
	fmt.Println()
}

func validateField(fielName string, str *string) {
	if len(*str) == 0 {
		output := fmt.Sprintln(string(COLOR_RED), fielName+" cannot be empty", string(COLOR_RESET))
		panic(output)
	}
}
