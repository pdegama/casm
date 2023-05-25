%bits 64

add [$msglen], byte 10

mov rax, 1        ; write(
mov rdi, 1        ;   STDOUT_FILENO,
mov rsi, $msg      ;   "Hello, world!\n",
mov rdx, [$msglen]   ;   sizeof("Hello, world!\n")
syscall           ; );


mov r13, $pre
jmp r13

$back:
mov r11, $msglen
mov [r11], r13
jmp [r11]

$exit:
mov rax, 60       ; exit(
mov rdi, 5        ;   EXIT_SUCCESS
syscall           ; );

$pre:
mov rax, 1        ; write(
mov rdi, 1        ;   STDOUT_FILENO,
mov rsi, $msg2      ;   "Hello, world!\n",
mov rdx, 3   ;   sizeof("Hello, world!\n")
syscall           ; );

mov rcx, $exit
mov r11, $msglen
mov [r11], rcx
jmp [r11]

$msg: 
  str "Hello, world!"
  db 10
$msglen: 
  db 14 0 0 0 0 0 0 0

$msg2: 
  str "H!"
  db 10
