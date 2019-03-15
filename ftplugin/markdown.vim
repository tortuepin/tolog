if exists("g:loaded_tolog")
    finish
endif

let g:loaded_tolog = 1
let s:binDir = expand('<sfile>:h:h:gs!\\!/!') . "/bin/"
let s:tag_file = ".tags.tolog"

function! Tolog_todo_set_active(...)
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
endfunction


function! Tolog_tag_collect()
    let l:func_name = "[Tolog_tag_collect]"
    let l:command = s:binDir . "tolog_tag_collect"
    let l:option = " -d " . g:tolog_dir

    echo l:func_name . " Start "
    echo l:func_name . " Do " . l:command . l:option
    echo system(l:command . l:option)
    echo l:func_name . " Done " . l:command . l:option
endfunction


function! Tolog_Complete_tag(...)
    " タグのリストを返す
    let l:tag_file = g:tolog_dir . "/" . s:tag_file
    let l:tag_list = ""
    try
    for line in readfile(l:tag_file)
        let l:tag_list = l:taglist . "\n"
    endfor
    catch
        " tag_listに何も入ってなかったらtag_collectしてもっかい
        echo l:tag_file . " is empty"
        echo "call tolog_tag_collect()"
        call Tolog_tag_collect()
        for line in readfile(l:tag_file)
            let l:tag_list = l:taglist . "\n"
        endfor
    endtry

    return l:tag_list
endfunction
