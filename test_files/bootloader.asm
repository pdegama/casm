%bits 16                  

MOV SI, $HelloString
CALL $PrintString
JMP $PrintCharacter      

$PrintCharacter:

MOV AH, 0x0E         
MOV BH, 0x00              
MOV BL, 0x07              

INT 0x10
RET                       

$PrintString:
                            
$next_character:
MOV AL, [SI]
INC SI

cmp al, 0x00
Je $exit_function

CALL $PrintCharacter
JMP $next_character
$exit_function:
RET


$HelloString:
    str "Hiiii...."
    db 0

db 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
db 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

dw 0xAA55
