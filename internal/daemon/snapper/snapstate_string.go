// Code generated by "stringer -type=SnapState"; DO NOT EDIT.

package snapper

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SnapPending-1]
	_ = x[SnapStarted-2]
	_ = x[SnapDone-4]
	_ = x[SnapError-8]
}

const (
	_SnapState_name_0 = "SnapPendingSnapStarted"
	_SnapState_name_1 = "SnapDone"
	_SnapState_name_2 = "SnapError"
)

var (
	_SnapState_index_0 = [...]uint8{0, 11, 22}
)

func (i SnapState) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _SnapState_name_0[_SnapState_index_0[i]:_SnapState_index_0[i+1]]
	case i == 4:
		return _SnapState_name_1
	case i == 8:
		return _SnapState_name_2
	default:
		return "SnapState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}