if exists("g:loaded_tolog")
    finish
endif

let g:loaded_tolog = 1
let s:binDir = expand('<sfile>:h:h:gs!\\!/!') . "/bin/"

function! Tolog_todo_set_active(...)
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
