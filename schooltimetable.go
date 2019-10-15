package ConvICS

import (
	"github.com/google/uuid"
	"time"
)

// Timetable 是一天的时间表，规定了每天每节课的上课下课时间
type Timetable []struct {
	Description string
	Start, End  time.Time
}

// Schedule 是整个学期的课程安排
type Schedule struct {
	SemesterStart time.Time // 学期开始的第一天，用于计算周数
	TotalWeeks    int       // 总周数
	Subjects      map[uuid.UUID]Subject
}

// Subject 是一门课程
type Subject struct {
	Name     string
	Teacher  string
	Arranges []Course
}

// Course 是一节课
type Course struct {
	Start, End int
	Location   string
}