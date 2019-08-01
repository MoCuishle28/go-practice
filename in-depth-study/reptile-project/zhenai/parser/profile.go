package parser

import(
	"regexp"
	"strconv"
	// "fmt"

	"Go-practice/in-depth-study/reptile-project/engine"
	"Go-practice/in-depth-study/reptile-project/model"
)

// [\d]就是所有数字 先编译好，不然每次解析都编译很耗时
var profileRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>(.+) \| ([\d]+)岁 \| (.+) \| (.+) \| ([\d]+)cm \| ([^<]*)</div>`)


func ParserProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	match := profileRe.FindSubmatch(contents)
	for i, m := range match[1:] {
		switch i {
			case 1:
				profile.Hokou = string(m)
			case 2:
				num, err := strconv.Atoi(string(m))
				if err != nil {
					num = -1
				}
				profile.Age = num
			case 3:
				profile.Education = string(m)
			case 4:
				profile.Marriage = string(m)
			case 5:
				num, err := strconv.Atoi(string(m))
				if err != nil {
					num = -1
				}
				profile.Height = num
			case 6:
				profile.Income = string(m)
		}
		// fmt.Printf("%s ", m)
	}
	profile.Name = name
	// fmt.Printf("%s ", name)
	// fmt.Println()

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}