// Code generated by "stringer -type=Code"; DO NOT EDIT.

package acres

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DONE - -1]
	_ = x[ALREADY - -2]
	_ = x[NO_JOB - -3]
	_ = x[NONE-0]
	_ = x[FAILURE-1]
}

const _Code_name = "NO_JOBALREADYDONENONEFAILURE"

var _Code_index = [...]uint8{0, 6, 13, 17, 21, 28}

func (i Code) String() string {
	i -= -3
	if i < 0 || i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i+-3), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}
