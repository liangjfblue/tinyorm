/**
 *
 * @author liangjf
 * @create on 2020/8/24
 * @version 1.0
 */
package utils

func DiffSlice(old, new []string) (diff []string) {
	m := make(map[string]string, len(old))
	for _, s := range old {
		m[s] = s
	}

	for _, s := range new {
		if _, ok := m[s]; !ok {
			diff = append(diff, s)
		}
	}
	return diff
}
