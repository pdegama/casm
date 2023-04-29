% bits 32

$start:
  mov rax, 1        ; write(
  add rax, 0x2323121121        ; write(
  mov rdi, 0xffffff        ;   STDOUT_FILENO,
  mov rsi, $msg      ;   "Hello, world!\n",
  mov rdx, $msglen   ;   sizeof("Hello, world!\n")
  syscall           ; );

    mov rdx, 0xffffff        ;   STDOUT_FILENO,

  hello:

  mov rax, 0x03c      ; exit(
  mov rdi, 0x2          ;   EXIT_SUCCESS
  syscall           ; );

  mov rdi, qword 0xfffff0      ;   STDOUT_FILENO,

  mov rdi, 90

  mov rcx, rdx

  ;mov al, cx

$startto:
  add rax, rdx
  add dx,  0x74
  add cx, byte  0x74

  mov rcx, [rax]
  mov [ecx], rax

  mov [rax], word 0x74
  mov [dword 0x74], eax
  mov [qword 0x70], 13

  add rdx, 0x121231212

  add [rcx], 0x121231231
  add [rcx], $qwe ; label shorting

  $msg:
    db "I ❤️  Computers!", 10
  $myname: 
    db "Parth Degama", 10
  msglen: 
    db 0x17