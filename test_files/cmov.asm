%bits 64

%text

mov r11, 0x5
mov r15, [$masslen]

mov al, 250
mov bl, 6

add al, bl

CMOVC r15, r11
sub r15, 0x1

$printHello:
  mov rax, 1 
  mov rdi, 1       
  mov rsi, $mass     
  mov rdx, r15
  syscall
  jmp $exit

$exit:
  mov rax, 60   
  mov rdi, 0     
  syscall    



%data

$mass:
    str "pePH1"
$masslen:
    db 2 0 0 0 0 0 0 0

