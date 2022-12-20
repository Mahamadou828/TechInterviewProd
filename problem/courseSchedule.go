package problem

//CheckCourseSchedule check if a given array of course schedules is valid.
//An invalid course schedule is a schedule where one the prerequisite form a loop.
//For example: [[1,0],[0,1]] is not valid because course 0 need course 1 and
//course 1 need course 0
//[[1,2,3], [0,5,6]] is valid because there's no loop
//Time complexity O(n^3)
//Space complexity O(n)
func CheckCourseSchedule(courses [][]int) bool {
	mapCoursesPrerequisites := make(map[int][]int)
	for _, course := range courses {
		courseIdx, course := course[0], course[1:]
		for _, prerequisite := range course {
			mapCoursesPrerequisites[courseIdx] = append(mapCoursesPrerequisites[courseIdx], prerequisite)
		}
	}

	for course, prerequisites := range mapCoursesPrerequisites {
		for _, prerequisiteNumber := range prerequisites {
			prerequisite := mapCoursesPrerequisites[prerequisiteNumber]
			for _, elt := range prerequisite {
				if elt == course {
					return false
				}
			}
		}
	}

	return true
}
