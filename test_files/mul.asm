%bits 64

        mov rbx, $a1

        mov [rbx], dword 42

        mov rax, [rbx]
        mov r15b, 0x02
        mul r15b

        mov [rbx], rax

        mov rdi, [rbx]      
        mov rax, 60
        syscall


$a1:
    db 0 0 0 0 0 0 0 0