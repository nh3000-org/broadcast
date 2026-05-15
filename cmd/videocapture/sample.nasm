; x86-64 NASM sample: spawn ffmpeg to grab one frame from /dev/video0
; Builds with: nasm -felf64 sample.nasm && ld sample.o -o sample

global _start

section .data
	sh_path:    db "/bin/sh",0
	cmd:        db "ffmpeg -y -f video4linux2 -i /dev/video0 -frames:v 1 /tmp/frame.jpg",0
	arg0:       db "sh",0
	arg1:       db "-c",0

section .bss

section .text
_start:
	; Build argv array: ["sh","-c", cmd, NULL]
	mov rax, 0
	; push pointers on stack for argv (we'll allocate on stack)
	; Reserve stack space
	sub rsp, 40
	; store pointers
	lea rdi, [rel arg0]
	mov [rsp+0], rdi
	lea rdi, [rel arg1]
	mov [rsp+8], rdi
	lea rdi, [rel cmd]
	mov [rsp+16], rdi
	mov qword [rsp+24], 0

	; execve("/bin/sh", rsp, NULL)
	lea rdi, [rel sh_path]
	mov rsi, rsp
	xor rdx, rdx
	mov rax, 59         ; sys_execve
	syscall

	; if execve fails, exit with status 1
	mov rdi, 1
	mov rax, 60         ; sys_exit
	syscall
