package academy

import "math"

type Student struct {
	Name      string
	Grades    []int
	Project   int
	Attendace []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {

	if len(grades) == 0 {
		return 0
	}
	var sum float64 = 0.0

	for _, num := range grades {
		sum += float64(num)
	}

	return int(math.Round(sum / float64(len(grades))))
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from  0 to 1,
// with 2 digits of precision.
func AttendancePercentage(attendance []bool) float64 {
	var num int = 0

	for _, val := range attendance {
		if val == true {
			num += 1

		}

	}
	return float64(num) / float64(len(attendance))
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {

	switch {
	case AttendancePercentage(s.Attendace) < 0.6:
		return 1
	case AverageGrade(s.Grades) == 1:
		return 1
	case s.Project == 1:
		return 1
	}

	var grade float64 = float64(s.Project) + float64(AverageGrade(s.Grades))
	grade /= 2

	if AttendancePercentage(s.Attendace) < 0.8 {
		grade -= 1
	}

	return int(math.Round(grade))
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {

	m := make(map[string]uint8, len(students))

	for _, s := range students {
		m[s.Name] = uint8(FinalGrade(s))
	}
	return m
}
