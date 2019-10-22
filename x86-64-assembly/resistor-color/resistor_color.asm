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
    mov r8, rdi
char_compare:
    mov r9b, [r8]
    mov r10b, [rsi]
    cmp r9b, r10b
    jnz string_not_equal
    test r9b, r9b
    jz string_equal
    inc r8
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
