# Go 语言汇编入门

## 为什么要学 Go 语言汇编

首先是要破除迷信，同一个问题网上的答案众说纷纭，比如到底是传值还是传引用争论不休，不如静下心看一下汇编来的踏实。

下面写的这些东西不一定都对，但是希望能与你分享一些方法和思路，授之以渔。学习的目的不是掌握这个知识，而是掌握学习知识的方法，举一反三，触类旁通，不管学什么都有自己的一套方法支撑，快速如何，快速解决问题，长远来看知识本身是没什么太大作用的。

学习 Go 语言汇编不是为了以后用汇编来做开发，只是可以用通过阅读汇编来深刻的理解 Go 语言背后的实现细节，真正的精通这门语言，在使用的过程中可以更加安心。

这篇文章将会首先介绍在 Linux 平台上用汇编输出 "Hello, World!"，通过这个例子顺带介绍汇编的一些基本的概念。后面我们介绍 Go 语言 Plan9 汇编打下基础。

## 用 C 语言写一个 Hello World 输出

之前看了不少的汇编的书，有一个感觉是，咋没有跟其它编程书籍一样，介绍如何输出 "Hello, World!" 呢？看得多以后就慢慢知道了，用汇编在控制台输出 "Hello, World!" 没有那么简单，不是三两行简单调用一个函数就完了。

为了搞清楚如何在终端中输出字符串，我们先来写一段 C 语言的实现：

```c
#include <stdio.h>

int main() {
    char *str = "Hello, World!\n";
    printf("%s", str);
}
```

更接近系统调用层的写法是：

```c
#include <unistd.h>

int main() {
    int stdout_fd = 1;
    char* str = "Hello, World!\n";
    int length =  14;
    write(stdout_fd, str, length);
}
```

Unix 的设计哲学，一切皆文件，一个程序运行以后都至少包含三个文件描述符（file descriptor，简称 fd）：

- 标准输入 stdin(0)
- 标准输出 stdout(1)
- 错误输出 stderr(2)

在终端执行程序输出字符串，实际上就是往标准输出 stdout 文件描述符写数据，stdout 的 fd 值等于 1。

write 是一个系统调用，把数据写入到文件，它的函数签名如下：

```c
ssize_t	 write(int fd, void * buffer, size_t count)
```

第一个参数 fd 表示要写入的文件描述符，第二个参数 buffer 表示要写入文件中数据的内存地址，第三个参数表示从 buffer 写入文件的数据字节数。因此在标准输出中输出"Hello, World!\n"实际上是调用 write 系统调用，往 fd 为 1 的文件描述符写入 14 个字节的字符串。

编译并执行上面的 C 代码，就可以看到输出了 "Hello, World!" 字符串

```c
gcc main.c -o main
./main 

Hello, World!
```

## CPU、内存与寄存器

汇编主要是跟 CPU 和内存打交道，CPU 本身只负责运算，不负责存储，数据存储一般都是放在内存中，我们知道 CPU 的运算速度远高于内存的读写速度，为了 CPU 不被内存读写拖后腿，CPU 内部引入一级缓存、二级缓存和寄存器的概念，这些资源都非常宝贵，至今都记得有一位老师说过：“二级缓存贵如黄金”。寄存器可以认为是在 CPU 内可以存储非常少量数据的超高速的存储单元。因为寄存器个数有限且非常重要，每个寄存器都有自己的名字，最常用的有下面这些，这些先混个眼熟，在后续的文章中再详细介绍。

```
%EAX %EBX %ECX %EDX %EDI %ESI %EBP %ESP
```

## 系统调用（System Call）：内核和应用程序之间的契约

下面我们来介绍**系统调用**概念。



![img](https://user-gold-cdn.xitu.io/2019/9/2/16cef38ca9027923?imageView2/0/w/1280/h/960/ignore-error/1)



内核对外暴露的接口被称为系统调用，应用程序可以调用对应的接口请求内核去完成某些动作，我们常见的创建新进程、IO 读写等都属于系统调用。

需要注意一下这些知识：

- 系统调用将处理器从「用户态」切换到「内核态」
- 应用程序都是按「名字」来执行系统调用，比如 exit、write，底层上每个系统调用都对应一个数字，比如 exit 对应 1，write 对应 4，这些数字编号需要被存储到寄存器 `%eax` 中
- 在调用系统调用时，参数值需要放置到规定好的寄存器中
- `int 0x80` 指令用来触发处理器从用户态切换到内核态，int 是 interrupt（中断）的缩写，不是整数的那个 int。内核收到 0x80 的中断请求以后，就会并根据前面准备好的寄存器的内容调用相应的系统调用。

执行一个 write 调用的流程如下图所示：

![go assembly.001](https://user-gold-cdn.xitu.io/2019/9/2/16cef38cab63622c?imageView2/0/w/1280/h/960/ignore-error/1)



## 汇编写 Hello World

有了上面的基础，再来看汇编的代码。文件名是 `helloworld.s`，下面是汇编的代码

```apl
.section .data

msg:
    .ascii "Hello, World!\n"

.section .text
.globl _start

_start:
    # write 的第 3个参数 count: 14
    movl $14,  %edx
    # write 的第 2 个参数 buffer: "Hello, World!\n"
    movl $msg, %ecx
    # write 的第 1 个参数 fd: 1
    movl $1,   %ebx
    #  write 系统调用本身的数字标识：4
    movl $4,   %eax
    #  执行系统调用: write(fd, buffer, count)
    int $0x80

    # status: 0
    movl $0,   %ebx
    # 函数: exit
    movl $1,   %eax
    # system call: exit(status)
    int $0x80
```

在汇编中，任何以点(.)开头的都不会被直接翻译为机器指令，`.section` 将汇编代码划分为多个段，`.section .data`是数据段的开始，数据段中存储后面程序需要用到的数据，相当于一个全局变量。在数据段中，我们定义了一个 msg，ascii 编码表示的内容是 "Hello, World!\n"，

接下来的 `.section .text` 表示是文本段的开始，文本段是存放程序指令的地方。

接下来的指令是 `.globl _start`，这里并没有拼错，不是 global，`_start` 是一个标签。接下来是真正的汇编指令部分了。

前面介绍过，执行 write 系统调用时，`%eax`寄存器存储 write 的系统调用号 4，`%ebx`存储标准输出的 fd，`%ecx`存储着输出buffer 的地址。`%edx`存储字节数。所以看到 `_start`便签后有四个 movl 指令，movl 指令的格式是：

```apl
movl src dst
```

比如`movl $4, %eax`指令是讲常量 4 存储到寄存器 `%eax` 中，数字 4 前面的 $ 表示「立即寻址」，汇编的其它寻址方式后面的文章还会详细介绍，这里先不展开，只需要知道立即寻址是本身就包含要访问的数据，比如要把数据初始化为 4，不用去哪个地址去读 4，在指令中直接给出数字 4。

接下来指令是 `int $0x80`，前面介绍过，这是一条中断触发指令，把执行流程交给内核继续处理，应用程序不用关心内核是如何处理的，内核处理完会把执行流程还给应用程序，同时根据执行成功与否设置全局变量 errno 的值。一般情况下，在 linux 上系统调用成功会返回非负值，发送错误时会返回负值。

接下来的指令实际上执行 exit(0) 退出程序，指令和逻辑与之前的一样，不再赘述。

下面来编译和执行上面的汇编代码。在 Linux 上，可以使用 as 和 ld 汇编和链接程序

```apl
as $helloworld.s -o helloworld.o
ld $helloworld.o -o helloworld

执行：
./helloworld
```

可以看到输出了

```
Hello, World!
```

## Go 语言汇编输出 Hello World 浅尝

刚开始接触 Go 语言汇编的时候一脸懵逼，这都是些啥，居然用的是一个从来没听说过的操作系统 plan9 所自带的汇编器语法，不过没有办法，技术选型永远是 leader 和 CTO 说了算。

文件结构如下：

```
.
├── helloworld
│   ├── helloworld.go
│   └── helloworld.s
├── main.go
```

main.go 的内容如下，调用了 helloworld.go 中的 PrintMe 方法：

```go
package main

import (
	"./helloworld"
)

func main() {
	helloworld.PrintMe()
}
```

helloworld.go 的内容只是声明了一个 PrintMe() 的空函数：

```go
package helloworld

func PrintMe()
```

具体的实现是在 helloworld.s 这个汇编文件中，内容如下：

```apl
#include "textflag.h"

DATA  msg<>+0x00(SB)/8, $"Hello, W"
DATA  msg<>+0x08(SB)/8, $"orld!\n"
GLOBL msg<>(SB),NOPTR,$16

TEXT ·PrintMe(SB), NOSPLIT, $0
	MOVL 	$(0x2000000+4), AX 	 // write 系统调用数字编号 4
	MOVQ 	$1, DI 			      // 第 1 个参数 fd
	LEAQ 	msg<>(SB), SI 		// 第 2 个参数 buffer 指针地址
	MOVL 	$16, DX 		        // 第 3 个参数 count
	SYSCALL
	RET
```

虽然指令不太一样，但是整体的汇编代码逻辑是一样的，同样是分了 Data 段、Text 段，同样是用 mov 等指令给寄存器赋值。下面简单介绍一下上面的汇编代码，后面的文章会有更详细的介绍。

### 关于寄存器

plan9 中使用寄存器不需要带 r 或 e 的前缀，例如 rax，只要写 AX 就可以了。

```
eax->AX
ebx->BX
ecx->CX
...
```

Go 汇编引入了四个伪寄存器，这四个伪寄存器非常重要：

- FP: Frame pointer，用来访问函数的参数
- PC: Program counter: 用于分支和跳转
- SB: Static base pointer: 一般用于声明函数或者全局变量
- SP: Stack pointer：指向当前栈帧的局部变量的开始位置，一般用来引用函数的局部变量

### 变量声明

Go 汇编语言中 DATA 命令用于初始化变量，语法如下：

```
DATA symbol+offset(SB)/width, value
```

比如声明 msg 这个变量：

```
DATA  msg<>+0x00(SB)/8, $"Hello, W"
```

下面来看 GLOBL 指令

```
GLOBL msg<>(SB),NOPTR,$16
```

GLOBL 指令将变量声明为 global，后面需要跟两个参数，flag 和变量的大小，这的 NOPTR 不影响后面的阅读，这里先不做介绍。

注意到 msg 后面有一个`<>`，这表示这个全局变量只在当前文件中可以被访问，类似于 C 语言中的 static。

### 函数定义

函数定义的语法如下：

```
TEXT symbol(SB), [flags,] $framesize[-argsize]
```

分为 5 个组成部分：TEXT 指令、函数名、可选的 flags 标志、函数帧大小和可选的函数参数大小

以例子中的汇编代码为例：

```
TEXT ·PrintMe(SB), NOSPLIT, $0
```

- TEXT 表示是汇编中的 .text 分段，
- 注意到 TEXT 和 PrintMe中间除了一个空格以外，还有一个反人类的「中点`·`」，不知道当初设计这个的人是有一种什么样的癖好，😁。这个中点在编译以后会被替换为`.`，同时也会加上包名，比如这里的 `helloworld.PrintMe`
- NOSPLIT 标志位这里先不介绍
- $0 表示栈帧大小为 0

接下来的就是具体的函数体的内容。

MOVL $(0x2000000+4)中的 0x2000000 是什么鬼？Mac 下的系统调用数字编号需要加 0x2000000，不要问为什么，问就是系统约定。Mac 下的系统调用编号可以在这里查：[opensource.apple.com/source/xnu/…](https://opensource.apple.com/source/xnu/xnu-1504.3.12/bsd/kern/syscalls.master)

与前面介绍的 Linux 下的汇编稍有不同，Mac 下的系统调用参数需要存储在 DI、SI、DX 等寄存器中，系统调用编号存储在 AX 中。