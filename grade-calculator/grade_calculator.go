package esepunittests

type GradeCalculator struct {
	scoreDetails []Grade
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
		scoreDetails: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

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
	gc.scoreDetails = append(gc.scoreDetails, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average, exam_average, essay_average := computeAverage(gc.scoreDetails)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) (int, int, int) {
	sumAssignments := 0
	sumExams := 0
	sumEssays := 0

	numberAssignment := 0
	numberExam := 0
	numberEssay := 0

	for _, value := range grades {
		if value.Type == Assignment {
			sumAssignments += value.Grade
			numberAssignment += 1
		} else if value.Type == Exam {
			sumExams += value.Grade
			numberExam += 1
		} else {
			sumEssays += value.Grade
			numberEssay += 1
		}

	}

	return sumAssignments / numberAssignment, sumExams / numberExam, sumEssays / numberEssay
}
