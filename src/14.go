
import ("strings")

func longestCommonPrefix(strs []string) string {
	var common []rune = []rune(strs[0])
	for i := 1; i < len(strs); i++ {
		temp := make([]rune, len(common))
		for i, v := range strs[i] {
			if i >= len(common) {
				break;
			}
			if common[i] == v {
				temp[i] = v
			} else {
				break;
			}
		}
		common = temp;
	}
	var strBuilder strings.Builder;
	for _, v := range common {
		if v == 0 {break;}
		strBuilder.WriteRune(v)
	}
	return strBuilder.String()
}