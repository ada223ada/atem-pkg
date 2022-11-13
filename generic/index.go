package generic

// NOTE This won't work since operator == not defined for T
// func Has[T any](a []T, v T) int {
//     for i, e := range a {
//         if  e == v {
//             return i
//         }
//     }
//     return -1
// }

func Index[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}
