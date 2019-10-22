default rel

section .data
black db "black", 0
brown db "brown", 0
red db "red", 0
orange db "orange", 0
yellow db "yellow", 0
green db "green", 0
blue db "blue", 0
violet db "violet", 0
grey db "grey", 0
white db "white", 0

color_list dq black, brown, red, orange, yellow, green, blue, violet, grey, white

section .text

global string_compare
string_compare:
    mov r10, rdi
char_compare:
    mov r8b, [r10]
    mov r9b, [rsi]
    cmp r8b, r9b
    jnz string_not_equal

    test r8b, r8b
    jz string_equal

    inc r10
    inc rsi
    jmp char_compare
string_equal:
    mov rax, 1
    ret
string_not_equal:
    xor rax, rax
    ret

global color_code
color_code:
    mov rcx, 0
    lea r11, [color_list]
check_color:
    mov rsi, [r11]
    call string_compare
    cmp rax, 1
    jz color_match

    add r11, 8
    inc rcx
    jmp check_color
color_match:
    mov rax, rcx
    ret

global colors
colors:
    lea rax, [color_list]
    ret
