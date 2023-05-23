%bits 64

        nop                     ; Nothing happens
        mov rdi, 26              ; first syscall argument: exit code
        mov rax,60              ; system call number (sys_exit)
        syscall


