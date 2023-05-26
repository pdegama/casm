%bits 64

%text

$start:
mov rax, 1
mov rdi, 1
mov rsi, $msg
mov rdx, [$msglen]
syscall

mov rdi, $add1
mov [rdi], $start2
jmp [rdi]

$exit:
  mov rax, 60       ; exit(
  mov rdi, 0        ;   EXIT_SUCCESS
  syscall           ; );

$start2:
  mov r12, $add1
  mov [r12], $msg2
  mov rax, 1
  mov rdi, 1
  mov rsi, [r12]
  mov rdx, [$msglen]
  syscall

mov rdi, $add1
mov [rdi], $exit
jmp [rdi]

%data

$msg: 
  str "Hello, world!"
  db 10
$msglen: 
  db 14 0 0 0 0 0 0 0
$msg2: 
  str "Hiiii, world!"
  db 10
$add1: 
  db 0 0 0 0 0 0 0 0

