package tolog

import "bufio"
import "strings"

func LogReader(scanner *bufio.Scanner) []LogItem { // {{{
	replaceKey := []string{"0", "?", "1", "?", "2", "?", "3", "?", "4", "?", "5", "?", "6", "?", "7", "?", "8", "?", "9", "?"}
	items := []LogItem{}
	i := -1
	for scanner.Scan() {
		current_text := scanner.Text()
		// logゾーンがおわったらreturn
		if strings.HasPrefix(current_text, HeaderPrefix) {
			return items
		}

		// [??:??] 形式だった場合ひとまとまりの開始
		r := strings.NewReplacer(replaceKey...)
		if strings.HasPrefix(r.Replace(current_text), "[??:??]") {
			i = i + 1
			item := LogItem{}
			items = append(items, item)

			// tag(@..)の認識
			header := strings.Split(strings.TrimRight(current_text, " "), " ")
			if len(header) > 1 {
				for k := 1; k < len(header); k++ {
					items[i].Tag = append(items[i].Tag, header[k])
				}
			}

			items[i].Name = current_text[1:6]
			continue
		}
		if i < 0 {
			continue
		}
		items[i].Contents = append(items[i].Contents, current_text)
	}

	/* DEBUG
	for _, k := range items {
		fmt.Println(k.Name)
		fmt.Println(k.Tag)
		for _, content := range k.Contents {
			fmt.Println(content)
		}
	}
	*/
	return items
} // }}}
