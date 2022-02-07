package main

import (
	sql "database/sql"
	repository2 "github.com/sergioch2006/aprendendo-go/infra/repository"
	usecase2 "github.com/sergioch2006/aprendendo-go/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306/fullcycle")
	if err != nil {
		log.Fatalln(err)
	}
	repository := repository2.CourseMySQLRepository{Db: db}
	usecase := usecase2.CreateCourse{Repository: repository}

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap {
		"bootstrap.servers": "kafka:9094",
	}
	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	for msg := range msgChan {
		var input usecase2.CreateCourseInputDto
		json.Unmarshal(msg.Value, &input)
		output, err := usecase.Execute(input)
		if err != nil {
			fmt.Println("Error:", err)
		}
		else {
			fmt.Println(output)
		}
	}
}