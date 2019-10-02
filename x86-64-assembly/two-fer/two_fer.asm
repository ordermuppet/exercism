default rel

section .data
prefix db "One for ", 0
suffix db ", one for me.", 0
default_name db "you", 0

section .text
global two_fer
print:
    mov cl, [rax]
    mov byte [rsi], cl
    inc rax
    inc rsi
    mov cl, [rax]
    cmp cl, 0
    jnz print
    ret

two_fer:
    lea rax, [prefix]
    call print ; print "One for "

    cmp rdi, 0
    jnz print_name
    lea rdi, [default_name]
print_name:
    lea rax, [rdi]
    call print ; print "you" if a NULL pointer has been passed, otherwise print the passed name
    lea rax, [suffix]
    call print ; print ", one for me."
    mov byte [rsi], 0
    ret
