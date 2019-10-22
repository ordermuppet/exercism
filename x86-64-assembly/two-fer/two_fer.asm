default rel

section .data
prefix db "One for ", 0
suffix db ", one for me.", 0
default_name db "you", 0

section .text
global two_fer
string_copy:
    mov cl, [rax]
    mov byte [rsi], cl
    inc rax
    inc rsi
    cmp byte [rax], 0
    jnz string_copy
    ret

two_fer:
    lea rax, [prefix]
    call string_copy ; print "One for "

    cmp rdi, 0
    jnz copy_name
    lea rdi, [default_name]
copy_name:
    mov rax, rdi
    call string_copy ; print "you" if a NULL pointer has been passed, otherwise print the passed name
    lea rax, [suffix]
    call string_copy ; print ", one for me."
    mov byte [rsi], 0
    ret
