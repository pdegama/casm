%bits 64

%text

        ;; print a string, with a size
        mov rcx, $hello
        mov rdx, 13
        call $print_string

        ;; print a string with ZERO size calculation
        mov rcx, $message
        call $print_asciiz_string

        mov rdx, $message
        call $print_asciiz_string_with_stars


        ;; print a string with an explicit size
        mov rcx, $goodbye
        mov rdx, 15
        call $print_string

        ;; exit this script
        mov rbx, 2
        call $_exit

        ;; Routine to print a string.
        ;;
        ;; Assumes string address is in RCX
        ;; Assumes string length is in RDX
        ;;
        ;; Traches: RAX, RBX, RCX, RDX
$print_string:
        mov rbx, 1         ;; output is STDOUT
        mov rax, 4         ;; sys_write
        int 0x80           ;; syscall

        ret

        ;; Routine to print a '0x00'-terminated string
        ;;
        ;; Assumes string address is in RCX
$print_asciiz_string:
        xor rdx, rdx            ; zero the length
        push rcx                ; save string
$len_loop:
        mov al, 0x00
        cmp  [rcx],  al
        je $len_loop_over
        inc rdx
        inc rcx
        jmp $len_loop
$len_loop_over:
        pop rcx                 ; restore string-pointer
                                ; rdx has the mesage
        call $print_string       ; call the print routine
        ret                     ; and return from here



        ;; Print a string, terminated by NULL, but change " " to "*"
        ;;
        ;; NOTE:This: destroys the string in the process.
$print_asciiz_string_with_stars:
        push rdx
$star_loop:
        mov al, 0x00
        cmp  [rdx],  al   ; end of string? we're done
        je $star_loop_over
        mov al, 0x20
        cmp  [rdx],  al  ; is this a space?
        jne $star_loop_cont         ; if not continue
        mov  [rdx],byte  42     ; so replace with "*"
$star_loop_cont:
        inc rdx                    ; increase our pointer
        jmp $star_loop              ; loop again
$star_loop_over:
        pop rcx
        call $print_asciiz_string
        ret


        ;; Exit
        ;;
        ;; Assumes RBX has exit-code
$_exit:
        mov rax, 1      ; SYS_exit
        int 0x80        ; syscall
        ret             ; Never reached

%data

$hello:
    str "Hello, world"
    db 10 00
$message: 
    str "This string has its size calculated dynamically!"
    db 10 00
$goodbye: 
    str "Goodbye, world"
    db 10 00
