package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()
	//fmt.Println(numericalGrade)
	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	grade_average := computeWieghtedAverage(gc.grades)

	//fmt.Println(assignment_average)
	//weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return grade_average
}

func computeWieghtedAverage(grades []Grade) int {
	assignment_sum := 0.0
	assignment_count := 0
	exam_sum := 0.0
	exam_count := 0
	essay_sum := 0.0
	essay_count := 0

	for _, val := range grades {

		switch val.Type {
		case Assignment:
			assignment_sum += float64(val.Grade) * .5
			assignment_count += 1
		case Exam:
			exam_sum += float64(val.Grade) * .35
			exam_count += 1
		case Essay:
			essay_sum += float64(val.Grade) * .15
			essay_count += 1
		}

	}

	return int(
		assignment_sum/float64(assignment_count) +
			exam_sum/float64(exam_count) +
			essay_sum/float64(essay_count))
}
