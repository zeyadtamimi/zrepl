// Code generated by "enumer -type=PropertySource -trimprefix=Source"; DO NOT EDIT.

package zfs

import (
	"fmt"
)

const (
	_PropertySourceName_0 = "LocalDefault"
	_PropertySourceName_1 = "Inherited"
	_PropertySourceName_2 = "None"
	_PropertySourceName_3 = "Temporary"
	_PropertySourceName_4 = "Received"
	_PropertySourceName_5 = "Any"
)

var (
	_PropertySourceIndex_0 = [...]uint8{0, 5, 12}
	_PropertySourceIndex_1 = [...]uint8{0, 9}
	_PropertySourceIndex_2 = [...]uint8{0, 4}
	_PropertySourceIndex_3 = [...]uint8{0, 9}
	_PropertySourceIndex_4 = [...]uint8{0, 8}
	_PropertySourceIndex_5 = [...]uint8{0, 3}
)

func (i PropertySource) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _PropertySourceName_0[_PropertySourceIndex_0[i]:_PropertySourceIndex_0[i+1]]
	case i == 4:
		return _PropertySourceName_1
	case i == 8:
		return _PropertySourceName_2
	case i == 16:
		return _PropertySourceName_3
	case i == 32:
		return _PropertySourceName_4
	case i == 4294967295:
		return _PropertySourceName_5
	default:
		return fmt.Sprintf("PropertySource(%d)", i)
	}
}

var _PropertySourceValues = []PropertySource{1, 2, 4, 8, 16, 32, 4294967295}

var _PropertySourceNameToValueMap = map[string]PropertySource{
	_PropertySourceName_0[0:5]:  1,
	_PropertySourceName_0[5:12]: 2,
	_PropertySourceName_1[0:9]:  4,
	_PropertySourceName_2[0:4]:  8,
	_PropertySourceName_3[0:9]:  16,
	_PropertySourceName_4[0:8]:  32,
	_PropertySourceName_5[0:3]:  4294967295,
}

// PropertySourceString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PropertySourceString(s string) (PropertySource, error) {
	if val, ok := _PropertySourceNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PropertySource values", s)
}

// PropertySourceValues returns all values of the enum
func PropertySourceValues() []PropertySource {
	return _PropertySourceValues
}

// IsAPropertySource returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PropertySource) IsAPropertySource() bool {
	for _, v := range _PropertySourceValues {
		if i == v {
			return true
		}
	}
	return false
}
