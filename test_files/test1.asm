global _start

section .text

_start:


        mov rbx, qword a1

        mov [rbx], dword 0x09

        mov rax, [rbx]
        mov rdx, 0x04
        mul rdx

        mov [rbx], rax

        mov rdi, [rbx]      
        mov rax, 60
        syscall

        a1:
        mov rdi, [rbx]      
        mov rax, 60
        syscall

