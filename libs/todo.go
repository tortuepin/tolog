package tolog

import "bufio"
import "strings"

func TodoReader(scanner *bufio.Scanner) []TodoItem { //{{{
	replaceKey := []string{" ", "", "x", ""}
	items := []TodoItem{}
	i := -1
	tag := ""
	for scanner.Scan() {
		current_text := scanner.Text()
		// todoゾーンが終わったらreturn
		if strings.HasPrefix(current_text, "## ") {
			break
		}

		// todo要素の処理開始
		r := strings.NewReplacer(replaceKey...)
		if strings.HasPrefix(r.Replace(current_text), "-[]") {
			i = i + 1
			item := TodoItem{}
			items = append(items, item)

			// section substituiton
			title_i := strings.Index(current_text, "] ") + 1
			todo_i := strings.Index(current_text, "- [") + 3

			// タイトル、Doneの処理
			items[i].Title = strings.Trim(current_text[title_i:], " ")
			if current_text[todo_i:todo_i+1] == "x" {
				items[i].Done = true
			} else {
				items[i].Done = false
			}

			// タグの処理
			text_list := strings.Split(current_text, " ")
			if strings.HasPrefix(text_list[len(text_list)-1], "@") {
				items[i].Tag = strings.Trim(text_list[len(text_list)-1], " ")
			} else if len(tag) > 0 {
				items[i].Tag = tag
			}

			// 深さの処理
			start_i := strings.Index(current_text, "-")
			items[i].Depth = start_i / TabSetting

		} else {
			// todoのとこにtodo以外が入ってたときの処理

			// 先頭が@だったときはtagを設定する
			if strings.HasPrefix(current_text, "@") {
				tag = current_text
			}
			if current_text == "" && tag != "" {
				tag = ""
			}
		}

	}
	/* DEBUG
	for _, k := range items {
		fmt.Println(k.Title)
		fmt.Println(k.Done)
		fmt.Println(k.Tag)
	}
	*/
	return items
} //}}}

// TodoGetActiveはTodoItemの中からxのついてないやつを抜き出すやつ
//
// TodoItemは日付順に並んでいることを仮定する
func TodoGetActive(items []TodoItem) []TodoItem { // {{{
	activeItems := []TodoItem{}
	uniq := make(map[string]int)
	for _, v := range items {
		if v.Done == true {
			continue
		}
		_, exist := uniq[v.Title]
		if exist {
			activeItems[uniq[v.Title]].Done = v.Done
			activeItems[uniq[v.Title]].Tag = v.Tag
		} else {
			activeItems = append(activeItems, v)
			uniq[v.Title] = len(activeItems) - 1
		}
	}
	return activeItems
} //}}}

// TodoGetTagMapはTodoItemをタグをkeyにしたTodoItemのmapを返すやつ
func TodoGetTagMap(items []TodoItem) ([]string, map[string][]TodoItem) {
	items_tag := map[string][]TodoItem{}
	tags := []string{}
	for _, v := range items {
		if _, ok := items_tag[v.Tag]; !ok {
			tags = append(tags, v.Tag)
		}
		items_tag[v.Tag] = append(items_tag[v.Tag], v)

	}
	//fmt.Println(len(tags))
	//for _, v := range tags {
	//	fmt.Println(v)
	//	fmt.Println(items_tag[v])
	//}
	return tags, items_tag
}

// TodoMap2Stringsはタグをkeyにしたmapをtologに書く文字列に変換するやつ
func TodoMap2Strings(items map[string][]TodoItem, keys []string) []string {
	s := []string{}
	for _, k := range keys {
		s = append(s, k)
		for _, i := range items[k] {
			todo := "- ["
			if i.Done {
				todo = todo + "x"
			} else {
				todo = todo + " "
			}
			todo = todo + "] "
			todo = todo + i.Title
			todo = strings.Repeat(" ", TabSetting*i.Depth) + todo
			s = append(s, todo)
		}
		s = append(s, "")
	}
	return s
}
