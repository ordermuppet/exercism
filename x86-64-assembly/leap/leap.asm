section .text
global leap_year
divide:
    xor rdx, rdx
    mov rax, rdi
    mov rcx, r8
    div rcx
    test rdx, rdx
    ret

leap_year:
    mov r8, 4
    call divide
    jnz is_not_leap_year

    mov r8, 100
    call divide
    jnz is_leap_year

    mov r8, 400
    call divide
    jz is_leap_year

is_not_leap_year:
    xor rax, rax
    ret
is_leap_year:
    mov rax, 1
    ret

