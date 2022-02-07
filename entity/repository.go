package entity

import "github.com/sergioch2006/aprendendo-go/entity"

type CourseRepository interface {
	Insert(course Course) error
}
