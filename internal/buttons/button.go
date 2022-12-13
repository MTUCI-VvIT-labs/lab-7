package buttons

import (
	"MTUCI-VvIT-labs/lab-7/internal/entities"
	"MTUCI-VvIT-labs/lab-7/internal/errors"
	"MTUCI-VvIT-labs/lab-7/pkg/pg"
	"context"
	"fmt"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v4"
	"log"
	"strconv"
)

var weekDays = map[int]string{
	0: "Понедельник",
	1: "Вторник",
	2: "Среда",
	3: "Четверг",
	4: "Пятница",
}

var timeSlots = map[int]string{
	0: "9:30 - 11:05",
	1: "11:20 - 12:55",
	2: "13:10 - 14:45",
	3: "15:25 - 17:00",
	4: "17:15 - 18:50",
}

func ButtonHandler(bot *botapi.BotAPI, msg botapi.MessageConfig, button string) {
	switch button {
	case "Понедельник":
		weekDay(bot, msg, 0)
	case "Вторник":
		weekDay(bot, msg, 1)
	case "Среда":
		weekDay(bot, msg, 2)
	case "Четверг":
		weekDay(bot, msg, 3)
	case "Пятница":
		weekDay(bot, msg, 4)
	case "Текущая неделя":
		currentWeek(bot, msg)
	case "Следующая неделя":
		nextWeek(bot, msg)
	default:
		errors.NotCommandMessage(bot, msg)
	}
}

func weekDay(bot *botapi.BotAPI, msg botapi.MessageConfig, day int) {
	lessons := dayLessons(entities.NewWeek().WeekType, day)

	msg.Text = weekDays[day] + ". Неделя №" + strconv.Itoa(entities.NewWeek().WeekNumber) + "." + "\n" + makeIndent()
	for i := 0; i < 5; i++ {
		isLesson := false
		for _, lesson := range lessons {
			if lesson.Number == i {
				fmt.Println("i = ", i)
				fmt.Println("lesson.Number = ", lesson.Number)
				isLesson = true
			}
		}

		fmt.Println("is lesson: ", isLesson)
		if isLesson {
			msg.Text += strconv.Itoa(i+1) + ". " + timeSlots[i] + "\n" +
				lessons[0].Subject + "\n" +
				lessons[0].Teacher + "\n" +
				lessons[0].Place + "\n"
			lessons = lessons[1:]
		} else {
			msg.Text += strconv.Itoa(i+1) + ". Пары нет.\n"
		}

		if i != 4 {
			msg.Text += makeIndent()
		}
	}

	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func dayLessons(weekType string, day int) (lessons []entities.Lesson) {
	var rows pgx.Rows
	var err error

	if weekType == "нечётная" {
		rows, err = pg.DB.Query(context.Background(), "SELECT * FROM schedule WHERE day = $1 AND week = $2", day, 0)
		if err != nil {
			log.Panic(err)
		}
	} else {
		rows, err = pg.DB.Query(context.Background(), "SELECT * FROM schedule WHERE day = $1 AND week = $2", day, 1)
		if err != nil {
			log.Panic(err)
		}
	}

	for rows.Next() {
		var lesson entities.Lesson
		err := rows.Scan(&lesson.Subject, &lesson.Teacher, &lesson.Number, &lesson.Place, &lesson.Day, &lesson.Week)
		if err != nil {
			log.Panic(err)
		}
		lessons = append(lessons, lesson)
	}

	return lessons
}

func currentWeek(bot *botapi.BotAPI, msg botapi.MessageConfig) {
	msg.Text = "Текущая неделя: " + entities.NewWeek().WeekType + "\n\n"

	lessons := weekLessons(entities.NewWeek().WeekType)

	makeWeekMessage(&msg, lessons)

	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func nextWeek(bot *botapi.BotAPI, msg botapi.MessageConfig) {
	var lessons []entities.Lesson
	if entities.NewWeek().WeekType == "нечётная" {
		lessons = weekLessons("чётная")
		msg.Text = "Следующая неделя: чётная\n\n"
	} else {
		lessons = weekLessons("нечётная")
		msg.Text = "Следующая неделя: нечётная\n\n"
	}

	makeWeekMessage(&msg, lessons)
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func weekLessons(weekType string) (lessons []entities.Lesson) {
	var rows pgx.Rows
	var err error

	if weekType == "нечётная" {
		rows, err = pg.DB.Query(context.Background(), "SELECT * FROM schedule WHERE week = $1", 0)
		if err != nil {
			log.Panic(err)
		}
	} else {
		rows, err = pg.DB.Query(context.Background(), "SELECT * FROM schedule WHERE week = $1", 1)
		if err != nil {
			log.Panic(err)
		}
	}

	for rows.Next() {
		var lesson entities.Lesson
		err := rows.Scan(&lesson.Subject, &lesson.Teacher, &lesson.Number, &lesson.Place, &lesson.Day, &lesson.Week)
		if err != nil {
			log.Panic(err)
		}
		lessons = append(lessons, lesson)
	}
	
	return lessons
}

func makeWeekMessage(msg *botapi.MessageConfig, lessons []entities.Lesson) {
	for i := 0; i < 5; i++ {
		msg.Text += weekDays[i] + "\n" + makeIndent()
		for k := 0; k < 5; k++ {
			isLesson := false
			for _, lesson := range lessons {
				if lesson.Number == k && lesson.Day == i {
					isLesson = true
				}
			}

			if isLesson {
				msg.Text += strconv.Itoa(k+1) + ". " + timeSlots[k] + "\n" +
					lessons[0].Subject + "\n" +
					lessons[0].Teacher + "\n" +
					lessons[0].Place + "\n"
				lessons = lessons[1:]
			} else {
				msg.Text += strconv.Itoa(k+1) + ". Пары нет.\n"
			}

			if k != 4 {
				msg.Text += makeIndent()
			} else {
				msg.Text += "\n\n"
			}
		}
	}
}

func makeIndent() string {
	return "-----------------------\n"
}
