package generate

import "regexp"

func MatchBookInfoWithTitle(name string) (map[string]string, error) {
	result := make(map[string]string, 0)

	// pattern 1
	xp, err := regexp.Compile(`^\((.*?)\)\s*\[(.*?)]\s*(.*?)\s*\((.*?)\)$`)
	if err != nil {
		return nil, err
	}
	if xp.MatchString(name) {
		matchResult := xp.FindAllStringSubmatch(name, 1)
		if matchResult == nil {
			return nil, nil
		}
		result["series"] = matchResult[0][1]
		result["artist"] = matchResult[0][2]
		result["title"] = matchResult[0][3]
		result["theme"] = matchResult[0][4]
		return result, nil
	}

	// pattern2
	xp, err = regexp.Compile(`^\((.*?)\)\s*\[(.*?)]\s*(.*?)\s*$`)
	if err != nil {
		return nil, err
	}
	if xp.MatchString(name) {
		matchResult := xp.FindAllStringSubmatch(name, 1)
		if matchResult == nil {
			return nil, nil
		}
		result["series"] = matchResult[0][1]
		result["artist"] = matchResult[0][2]
		result["title"] = matchResult[0][3]
		return result, nil
	}
	return nil, nil
}
