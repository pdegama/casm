bits 

start:
  mov rax, 1        ; write(
  mov rdi, 0x01        ;   STDOUT_FILENO,
  mov rsi, msg      ;   "Hello, world!\n",
  mov rdx, msglen   ;   sizeof("Hello, world!\n")
  syscall           ; );

  hello:

  mov rax, 60       ; exit(
  mov rdi, 0        ;   EXIT_SUCCESS
  syscall           ; );

  msg:
    db "I ❤️  Computers!", 10
  msglen: 
    db 0x17