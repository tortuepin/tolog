""" vim:set foldmethod=marker:
if exists("g:loaded_tolog")
    finish
endif

let g:loaded_tolog = 1
let s:binDir = expand('<sfile>:h:h:gs!\\!/!') . "/bin/"
let s:tag_file = ".tags.tolog"

function! Tolog_todo_set_active(...) "{{{
    " TODO g:tolog_dirセットしてって"
    let l:func_name = "[Tolog_todo_set_active]"
    let l:command =  s:binDir . "tolog_set_active_todo"
    let l:option = " -d " . g:tolog_dir
    let l:option = l:option . " -f " . fnamemodify(expand("%"), ":t")
    if exists("a:1")
        let l:option = l:option . " -n " . a:1
    endif

    echo l:func_name . " Start "
    echo l:func_name . " Do " . l:command . l:option
    echo system(l:command . l:option)
    echo l:func_name . " ReLoad : " . fnamemodify(expand("%"), ":t")
    echo execute("e")
    echo l:func_name . " Done " . l:command . l:option
endfunction "}}}
function! Tolog_tag_collect() "{{{
    let l:func_name = "[Tolog_tag_collect]"
    let l:command = s:binDir . "tolog_tag_collect"
    let l:option = " -d " . g:tolog_dir

    echo l:func_name . " Start "
    echo l:func_name . " Do " . l:command . l:option
    echo system(l:command . l:option)
    echo l:func_name . " Done " . l:command . l:option
endfunction "}}}
function! Tolog_Complete_tag(...) "{{{
    " タグのリストを返す
    let l:tag_file = g:tolog_dir . "/" . s:tag_file
    let l:tag_list = ""
    try
    for line in readfile(l:tag_file)
        let l:tag_list = line . "\n" . l:tag_list
    endfor
    catch
        " tag_listに何も入ってなかったらtag_collectしてもっかい
        echo l:tag_file . " is empty"
        echo "call tolog_tag_collect()"
        call Tolog_tag_collect()
        for line in readfile(l:tag_file)
            let l:tag_list = line . "\n" . l:tag_list
        endfor
    endtry

    return l:tag_list
endfunction "}}}
function! Tolog_log_search_bytag(...) "{{{
    let l:func_name = "[Tolog_log_search_bytag]"
    let l:command = s:binDir . "tolog_log_search_bytag"
    let l:option = " -d " . g:tolog_dir
    let l:args = " " . join(a:000)

    echo l:func_name . " Start "
    echo l:func_name . " Do " . l:command.l:option.l:args
    let l:ret = systemlist(l:command.l:option.l:args)
    let l:fname = tempname()
    echo l:fname
    vs
    let l:com = "edit +call\\ append(0,l:ret) " . l:fname
    echo l:com
    execute l:com
    nnoremap <buffer> q <C-w>c
    setlocal bufhidden=hide buftype=nofile noswapfile nobuflisted
endfunction "}}}
function! Tolog_read_template() " {{{
    " テンプレートの読み込み
    execute "0read " . g:tolog_template_dir
endfunction
"}}}
function! Tolog_add_log(...) "{{{
    let l:option = ""
    if a:0 > 0
        let l:option = " " . join(a:000)
    endif
    call append(line('$'), "")
    call append(line('$'), strftime("[%H:%M]") . l:option)
    call append(line('$'), "")
    call append(line('$'), "")
    call cursor(line('$'), 0)
endfunction "}}}
function! Tolog_open_log(filename) "{{{
    execute "e " . a:filename
endfunction 
" }}}
""" 前後の日を開く Tolog_get_*_filename(){{{
function! Tolog_get_today_filename()
    let a:next =  strftime("%y%m%d")
    return g:tolog_dir . a:next . ".md"
endfunction
function! Tolog_get_next_filename()
    " 次の日のファイル名を出力
    let a:t = expand("%") " 現在のファイル名
    let a:date = fnamemodify(a:t, ":t:r")
    let a:year = strpart(a:date, 0, 2)
    let a:month = strpart(a:date, 2, 2)
    let a:day = strpart(a:date, 4, 2)

    let a:d = s:localtime("20".a:year, a:month, a:day, 0, 0, 0)

    let day = (60 * 60 * 24)
    let a:next =  strftime("%y%m%d", a:d + day)
    return g:tolog_dir . a:next . ".md"
endfunction
function! Tolog_get_prev_filename()
    " 次の日のファイル名を出力
    let a:t = expand("%") " 現在のファイル名
    let a:date = fnamemodify(a:t, ":t:r")
    let a:year = strpart(a:date, 0, 2)
    let a:month = strpart(a:date, 2, 2)
    let a:day = strpart(a:date, 4, 2)

    let a:d = s:localtime("20".a:year, a:month, a:day, 0, 0, 0)

    let day = (60 * 60 * 24)
    let a:next =  strftime("%y%m%d", a:d - day)
    return g:tolog_dir . a:next . ".md"
endfunction 
"}}}
""" サイコロをふる Tolog_tumbling_dice(){{{
function! Tolog_tumbling_dice() range
    let a:linenum = a:lastline - a:firstline
    let a:n_dice = s:getRandom(a:linenum)
    let a:line = getline(a:firstline + a:n_dice)
    let a:ret = a:n_dice+1 . " : " . a:line
    echo a:ret
    call append(a:lastline, ['', a:ret])
endfunction
""" }}}


""" Utils {{{
function! s:localtime(year, month, day, hour, minute, second)
    " days from 0000/01/01
    let l:year  = a:month < 3 ? a:year - 1   : a:year
    let l:month = a:month < 3 ? 12 + a:month : a:month
    let l:days = 365*l:year + l:year/4 - l:year/100 + l:year/400 + 306*(l:month+1)/10 + a:day - 428

    " days from 0000/01/01 to 1970/01/01
    " 1970/01/01 == 1969/13/01
    let l:ybase = 1969
    let l:mbase = 13
    let l:dbase = 1
    let l:basedays = 365*l:ybase + l:ybase/4 - l:ybase/100 + l:ybase/400 + 306*(l:mbase+1)/10 + l:dbase - 428

    " seconds from 1970/01/01
    return (l:days-l:basedays)*86400 + (a:hour-9)*3600 + a:minute*60 + a:second
endfunction

function! s:getRandom(max)
    let a:match_end = matchend(reltimestr(reltime()), '\d\+\.') + 1
    let a:rand = reltimestr(reltime())[a:match_end : ] % (a:max + 1)
    return a:rand
endfunction
" }}}

