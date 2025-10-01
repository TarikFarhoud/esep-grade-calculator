package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 74, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 73, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 64, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 63, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 35, Assignment)
	gradeCalculator.AddGrade("exam 1", 34, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 33, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestStringMappingAssignment(t *testing.T) {
	expected_string := "assignment"

	actual_string := Assignment.String()

	if expected_string != actual_string {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_string, actual_string)
	}
}

func TestStringMappingExam(t *testing.T) {
	expected_string := "exam"

	actual_string := Exam.String()

	if expected_string != actual_string {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_string, actual_string)
	}
}

func TestStringMapping(t *testing.T) {
	expected_string := "essay"

	actual_string := Essay.String()

	if expected_string != actual_string {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_string, actual_string)
	}
}
