%bits 64

%text

mov cl, 250
mov dl, 5

add cl, dl

jnc $printHello

$exit:
  mov rax, 60   
  mov rdi, 0     
  syscall    

$printHello:
  mov rax, 1 
  mov rdi, 1       
  mov rsi, $mass     
  mov rdx, [$masslen]  
  syscall
  jmp $exit


%data

$mass:
    str "pePH1"
$masslen:
    db 5 0 0 0 0 0 0 0

