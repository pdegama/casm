% bits 32

$start:
  mov rax, 1        ; write(
  mov rdi, 0xfff        ;   STDOUT_FILENO,
  mov rsi, $msg      ;   "Hello, world!\n",
  mov rdx, $msglen   ;   sizeof("Hello, world!\n")
  syscall           ; );

  hello:

  mov rax, 0x03c      ; exit(
  mov rdi, 0x2          ;   EXIT_SUCCESS
  syscall           ; );

  mov al, cx

  $msg:
    db "I ❤️  Computers!", 10
  $myname: 
    db "Parth Degama", 10
  msglen: 
    db 0x17