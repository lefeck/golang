# Go 语言内存管理（一）：系统内存管理

要搞明白 Go 语言的内存管理，就必须先理解操作系统以及机器硬件是如何管理内存的。因为 Go 语言的内部机制是建立在这个基础之上的，它的设计，本质上就是尽可能的会发挥操作系统层面的优势，而避开导致低效情况。

## 操作系统内存管理

其实现在计算机内存管理的方式都是一步步演变来的，最开始是非常简单的，后来为了满足各种需求而增加了各种各样的机制，越来越复杂。这里我们只介绍和开发者息息相关的几个机制。

### 最原始的方式

我们可以把内存看成一个数组，每个数组元素的大小是 `1B`，也就是 8 位(bit)。CPU 通过内存地址来获取内存中的数据，内存地址可以看做成数组的游标（index）。

![img](https://upload-images.jianshu.io/upload_images/11662994-213100b52eabbefe.png?imageMogr2/auto-orient/strip|imageView2/2/w/391/format/webp)

CPU 在执行指令的时候，就是通过内存地址，将物理内存上的数据载入到寄存器，然后执行机器指令。但随着发展，出现了多任务的需求，也就是希望多个任务能同时在系统上运行。这就出现了一些问题：

1. **内存访问冲突：**程序很容易出现 bug，就是 2 或更多的程序使用了同一块内存空间，导致数据读写错乱，程序崩溃。更有一些黑客利用这个缺陷来制作病毒。
2. **内存不够用：**因为每个程序都需要自己单独使用的一块内存，内存的大小就成了任务数量的瓶颈。
3. **程序开发成本高：**你的程序要使用多少内存，内存地址是多少，这些都不能搞错，对于人来说，开发正确的程序很费脑子。

举个例子，假设有一个程序，当代码运行到某处时，需要使用 `100M` 内存，其他时候 `1M` 内存就够；为了避免和其他程序冲突，程序初始化时，就必须申请独立 `100M` 内存以保证正常运行，这就是一种很大的浪费，因为这 `100M` 它大多数时候用不上，其他程序还不能用。

### 虚拟内存

虚拟内存的出现，很好的为了解决上述的一些列问题。用户程序只能使用虚拟的内存地址来获取数据，系统会将这个虚拟地址翻译成实际的物理地址。

所有程序统一使用一套连续虚拟地址，比如 `0x0000 ~ 0xffff`。从程序的角度来看，它觉得自己独享了一整块内存。不用考虑访问冲突的问题。系统会将虚拟地址翻译成物理地址，从内存上加载数据。

对于内存不够用的问题，虚拟内存本质上是将磁盘当成最终存储，而主存作为了一个 cache。程序可以从虚拟内存上申请很大的空间使用，比如 `1G`；但操作系统不会真的在物理内存上开辟 `1G` 的空间，它只是开辟了很小一块，比如 `1M` 给程序使用。
这样程序在访问内存时，操作系统看访问的地址是否能转换成物理内存地址。能则正常访问，不能则再开辟。这使得内存得到了更高效的利用。

如下图所示，每个进程所使用的虚拟地址空间都是一样的，但他们的虚拟地址会被映射到主存上的不同区域，甚至映射到磁盘上（当内存不够用时）。

![img](https://upload-images.jianshu.io/upload_images/11662994-c5cf238298084324.png?imageMogr2/auto-orient/strip|imageView2/2/w/664/format/webp)

其实本质上很简单，就是操作系统将程序常用的数据放到内存里加速访问，不常用的数据放在磁盘上。这一切对用户程序来说完全是透明的，用户程序可以假装所有数据都在内存里，然后通过虚拟内存地址去访问数据。在这背后，操作系统会自动将数据在主存和磁盘之间进行交换。

#### 虚拟地址翻译

虚拟内存的实现方式，大多数都是通过**页表**来实现的。操作系统虚拟内存空间分成一页一页的来管理，每页的大小为 `4K`（当然这是可以配置的，不同操作系统不一样）。磁盘和主内存之间的置换也是以**页**为单位来操作的。`4K` 算是通过实践折中出来的通用值，太小了会出现频繁的置换，太大了又浪费内存。

`虚拟地址 -> 物理地址` 的映射关系由**页表（Page Table）**记录，它其实就是一个数组，数组中每个元素叫做**页表条目（Page Table Entry，简称 PTE）**，PTE 由一个有效位和 n 位地址字段构成，有效位标识这个虚拟地址是否分配了物理内存。

页表被操作系统放在物理内存的指定位置，CPU 上有个 Memory Management Unit（MMU） 单元，CPU 把虚拟地址给 MMU，MMU 去物理内存中查询页表，得到实际的物理地址。当然 MMU 不会每次都去查的，它自己也有一份缓存叫Translation Lookaside Buffer (TLB)，是为了加速地址翻译。

![img](https://upload-images.jianshu.io/upload_images/11662994-0be170c393fd296d.png?imageMogr2/auto-orient/strip|imageView2/2/w/795/format/webp)

虚拟地址翻译

> 你慢慢会发现整个计算机体系里面，缓存是无处不在的，整个计算机体系就是建立在一级级的缓存之上的，无论软硬件。

让我们来看一下 CPU 内存访问的完整过程：

1. CPU 使用虚拟地址访问数据，比如执行了 MOV 指令加载数据到寄存器，把地址传递给 MMU。
2. MMU 生成 PTE 地址，并从主存（或自己的 Cache）中得到它。
3. 如果 MMU 根据 PTE 得到真实的物理地址，正常读取数据。流程到此结束。
4. 如果 PTE 信息表示没有关联的物理地址，MMU 则触发一个缺页异常。
5. 操作系统捕获到这个异常，开始执行异常处理程序。在物理内存上创建一页内存，并更新页表。
6. 缺页处理程序在物理内存中确定一个**牺牲页**，如果这个牺牲页上有数据，则把数据保存到磁盘上。
7. 缺页处理程序更新 PTE。
8. 缺页处理程序结束，再回去执行上一条指令（导致缺页异常的那个指令，也就是 MOV 指令）。这次肯定命中了。

#### 内存命中率

你可能已经发现，上述的访问步骤中，从第 4 步开始都是些很繁琐的操作，频繁的执行对性能影响很大。毕竟访问磁盘是非常慢的，它会引发程序性能的急剧下降。如果内存访问到第 3 步成功结束了，我们就说**页命中**了；反之就是**未命中**，或者说**缺页**，表示它开始执行第 4 步了。

假设在 n 次内存访问中，出现命中的次数是 m，那么 `m / n * 100%` 就表示命中率，这是衡量内存管理程序好坏的一个很重要的指标。

如果物理内存不足了，数据会在主存和磁盘之间频繁交换，命中率很低，性能出现急剧下降，我们称这种现象叫**内存颠簸**。这时你会发现系统的 swap 空间利用率开始增高， CPU 利用率中 `iowait` 占比开始增高。

大多数情况下，只要物理内存够用，页命中率不会非常低，不会出现内存颠簸的情况。因为大多数程序都有一个特点，就是**局部性**。

**局部性就是说被引用过一次的存储器位置，很可能在后续再被引用多次；而且在该位置附近的其他位置，也很可能会在后续一段时间内被引用。**

前面说过计算机到处使用一级级的缓存来提升性能，归根结底就是利用了**局部性**的特征，如果没有这个特性，一级级的缓存不会有那么大的作用。所以一个局部性很好的程序运行速度会更快。

### CPU Cache

随着技术发展，CPU 的运算速度越来越快，但内存访问的速度却一直没什么突破。最终导致了 CPU 访问主存就成了整个机器的性能瓶颈。CPU Cache 的出现就是为了解决这个问题，在 CPU 和 主存之间再加了 Cache，用来缓存一块内存中的数据，而且还不只一个，现代计算机一般都有 3 级 Cache，其中 L1 Cache 的访问速度和寄存器差不多。

现在访问数据的大致的顺序是 `CPU --> L1 Cache --> L2 Cache --> L3 Cache --> 主存 --> 磁盘`。从左到右，访问速度越来越慢，空间越来越大，单位空间（比如每字节）的价格越来越低。

现在存储器的整体层次结构大致如下图：

![img](https://upload-images.jianshu.io/upload_images/11662994-5c6eecbc31233544.png?imageMogr2/auto-orient/strip|imageView2/2/w/700/format/webp)

存储器层次结构

在这种架构下，缓存的命中率就更加重要了，因为系统会假定所有程序都是有局部性特征的。如果某一级出现了未命中，他就会将该级存储的数据更新成最近使用的数据。

*主存与存储器之间以 page（通常是 4K） 为单位进行交换，cache 与 主存之间是以 cache line（通常 64 byte） 为单位交换的。*

## 举个例子

让我们通过一个例子来验证下命中率的问题，下面的函数是循环一个数组为每个元素赋值。



```go
func Loop(nums []int, step int) {
    l := len(nums)
    for i := 0; i < step; i++ {
        for j := i; j < l; j += step {
            nums[j] = 4
        }
    }
}
```

参数 step 为 1 时，和普通一层循环一样。假设 step 为 2 ，则效果就是跳跃式遍历数组，如 `1,3,5,7,9,2,4,6,8,10` 这样，step 越大，访问跨度也就越大，程序的局部性也就越不好。

下面是 nums 长度为 `10000`， `step = 1` 和 `step = 16` 时的压测结果：



```undefined
goos: darwin
goarch: amd64
BenchmarkLoopStep1-4              300000              5241 ns/op
BenchmarkLoopStep16-4             100000             22670 ns/op
```

可以看出，2 种遍历方式会出现 3 倍的性能差距。这种问题最容易出现在多维数组的处理上，比如遍历一个二维数组很容易就写出局部性很差的代码。

## 程序的内存布局

最后看一下程序的内存布局。现在我们知道了每个程序都有自己一套独立的地址空间可以使用，比如 `0x0000 ~ 0xffff`，但我们在用高级语言，无论是 C 还是 Go 写程序的时候，很少直接使用这些地址。我们都是通过变量名来访问数据的，编译器会自动将我们的变量名转换成真正的虚拟地址。

那最终编译出来的二进制文件，是如何被操作系统加载到内存中并执行的呢？

其实，操作系统已经将一整块内存划分好了区域，每个区域用来做不同的事情。如图：

![img](https://upload-images.jianshu.io/upload_images/11662994-5b4e90c15b38e1b8.png?imageMogr2/auto-orient/strip|imageView2/2/w/245/format/webp)

内存布局

- **text 段：**存储程序的二进制指令，及其他的一些静态内容
- **data 段：**用来存储已被初始化的全局变量。比如常量（`const`）。
- **bss 段：**用来存放未被初始化的全局变量。和 .data 段一样都属于静态分配，在这里面的变量数据在编译就确定了大小，不释放。
- **stack 段：**栈空间，主要用于函数调用时存储临时变量的。这部分的内存是自动分配自动释放的。
- **heap 段：**堆空间，用于动态分配，C 语言中 `malloc` 和 `free` 操作的内存就在这里；Go 语言主要靠 GC 自动管理这部分。

其实现在的操作系统，进程内部的内存区域没这么简单，要比这复杂多了，比如内核区域，共享库区域。因为我们不是要真的开发一套操作系统，细节可以忽略。这里只需要记住**堆空间**和**栈空间**即可。

- **栈空间**是通过压栈出栈方式自动分配释放的，由系统管理，使用起来高效无感知。
- **堆空间**是用以动态分配的，由程序自己管理分配和释放。Go 语言虽然可以帮我们自动管理分配和释放，但是代价也是很高的。

## 结论

局部性好的程序，可以提高缓存命中率，这对底层系统的内存管理是很友好的，可以提高程序的性能。CPU Cache 层面的低命中率导致的是程序运行缓慢，内存层面的低命中率会出现内存颠簸，出现这种现象时你的服务基本上已经瘫痪了。Go 语言的内存管理是参考 tcmalloc 实现的，它其实就是利用好了 OS 管理内存的这些特点，来最大化内存分配性能的。



# Go 语言内存管理（二）Go 内存管理

## 介绍

了解操作系统对内存的管理机制后，现在可以去看下 Go 语言是如何利用底层的这些特性来优化内存的。Go 的内存管理基本上参考 `tcmalloc` 来实现的，只是细节上根据自身的需要做了一些小的优化调整。

Go 的内存是自动管理的，我们可以随意定义变量直接使用，不需要考虑变量背后的内存申请和释放的问题。本文意在搞清楚 Go 在方面帮我们做了什么，使我们不用关心那些复杂内存的问题，还依旧能写出较为高效的程序。

## 池

程序动态申请内存空间，是要使用系统调用的，比如 Linux 系统上是调用 `mmap` 方法实现的。但对于大型系统服务来说，直接调用 `mmap` 申请内存，会有一定的代价。比如：

1. **系统调用会导致程序进入内核态，内核分配完内存后（也就是上篇所讲的，对虚拟地址和物理地址进行映射等操作），再返回到用户态。**
2. **频繁申请很小的内存空间，容易出现大量内存碎片，增大操作系统整理碎片的压力。**
3. **为了保证内存访问具有良好的局部性，开发者需要投入大量的精力去做优化，这是一个很重的负担。**

如何解决上面的问题呢？有经验的人，可能很快就想到解决方案，那就是我们常说的**对象池**（也可以说是缓存）。

假设系统需要频繁动态申请内存来存放一个数据结构，比如 `[10]int`。那么我们完全可以在程序启动之初，一次性申请几百甚至上千个 `[10]int`。这样完美的解决了上面遇到的问题：

1. 不需要频繁申请内存了，而是从对象池里拿，程序不会频繁进入内核态
2. 因为一次性申请一个连续的大空间，对象池会被重复利用，不会出现碎片。
3. 程序频繁访问的就是对象池背后的同一块内存空间，局部性良好。

这样做会造成一定的内存浪费，我们可以定时检测对象池的大小，保证可用对象的数量在一个合理的范围，少了就提前申请，多了就自动释放。

如果某种资源的申请和回收是昂贵的，我们都可以通过建立**资源池**的方式来解决，其他比如**连接池**，**内存池**等等，都是一个思路。

## Golang 内存管理

**Golang 的内存管理本质上就是一个内存池，只不过内部做了很多的优化。比如自动伸缩内存池大小，合理的切割内存块等等。**

### 内存池 mheap

Golang 的程序在启动之初，会一次性从操作系统那里申请一大块内存作为内存池。这块内存空间会放在一个叫 `mheap` 的 `struct` 中管理，mheap 负责将这一整块内存切割成不同的区域，并将其中一部分的内存切割成合适的大小，分配给用户使用。

我们需要先知道几个重要的概念：

- **`page`**: 内存页，一块 `8K` 大小的内存空间。Go 与操作系统之间的内存申请和释放，都是以 `page` 为单位的。
- **`span`**: 内存块，**一个或多个连续的** `page` 组成一个 `span`。如果把 `page` 比喻成工人，`span` 可看成是小队，工人被分成若干个队伍，不同的队伍干不同的活。
- **`sizeclass`**: 空间规格，每个 `span` 都带有一个 `sizeclass`，标记着该 `span` 中的 `page` 应该如何使用。使用上面的比喻，就是 `sizeclass` 标志着 `span` 是一个什么样的队伍。
- **`object`**: 对象，用来存储一个变量数据内存空间，一个 `span` 在初始化时，会被切割成一堆**等大**的 `object`。假设 `object` 的大小是 `16B`，`span` 大小是 `8K`，那么就会把 `span` 中的 `page` 就会被初始化 `8K / 16B = 512` 个 `object`。所谓内存分配，就是分配一个 `object` 出去。

示意图：

![img](https://upload-images.jianshu.io/upload_images/11662994-8361f3be115cf456.png?imageMogr2/auto-orient/strip|imageView2/2/w/404/format/webp)

上图中，不同颜色代表不同的 `span`，不同 `span` 的 `sizeclass` 不同，表示里面的 `page` 将会按照不同的规格切割成一个个等大的 `object` 用作分配。

使用 Go1.11.5 版本测试了下初始堆内存应该是 `64M` 左右，低版本会少点。

测试代码：

```go
package main
import "runtime"
var stat runtime.MemStats
func main() {
    runtime.ReadMemStats(&stat)
    println(stat.HeapSys)
}
```

内部的整体内存布局如下图所示：

![img](https://upload-images.jianshu.io/upload_images/11662994-356f568da2987e54.png?imageMogr2/auto-orient/strip|imageView2/2/w/817/format/webp)

- `mheap.spans`：用来存储 `page` 和 `span` 信息，比如一个 span 的起始地址是多少，有几个 page，已使用了多大等等。
- `mheap.bitmap` 存储着各个 `span` 中对象的标记信息，比如对象是否可回收等等。
- `mheap.arena_start`: 将要分配给应用程序使用的空间。

再说明下，图中的空间大小，是 Go 向操作系统申请的虚拟内存地址空间，操作系统会将该段地址空间预留出来不做它用；而不是真的创建出这么大的虚拟内存，在页表中创建出这么大的映射关系。



### mcentral



**用途相同**的 `span` 会以链表的形式组织在一起。 这里的用途用 `sizeclass` 来表示，就是指该 `span` 用来存储哪种大小的对象。比如当分配一块大小为 `n` 的内存时，系统计算 `n` 应该使用哪种 `sizeclass`，然后根据 `sizeclass` 的值去找到一个可用的 `span` 来用作分配。其中 `sizeclass` 一共有 67 种（Go1.5 版本，后续版本可能会不会改变不好说），如图所示：

![img](https:////upload-images.jianshu.io/upload_images/11662994-730fc9b0a604aea1.png?imageMogr2/auto-orient/strip|imageView2/2/w/551/format/webp)

找到合适的 `span` 后，会从中取一个 `object` 返回给上层使用。这些 `span` 被放在一个叫做 mcentral 的结构中管理。

mheap 将从 OS 那里申请过来的内存初始化成一个大 `span`(sizeclass=0)。然后根据需要从这个大 `span` 中切出小 `span`，放在 mcentral 中来管理。大 `span` 由 `mheap.freelarge` 和 `mheap.busylarge` 等管理。如果 mcentral 中的 `span` 不够用了，会从 `mheap.freelarge` 上再切一块，如果 `mheap.freelarge` 空间不够，会再次从 OS 那里申请内存重复上述步骤。下面是 mheap 和 mcentral 的数据结构：



```go
type mheap struct {
    // other fields
    lock      mutex
    free      [_MaxMHeapList]mspan // free lists of given length， 1M 以下
    freelarge mspan                // free lists length >= _MaxMHeapList, >= 1M
    busy      [_MaxMHeapList]mspan // busy lists of large objects of given length
    busylarge mspan                // busy lists of large objects length >= _MaxMHeapList

    central [_NumSizeClasses]struct { // _NumSizeClasses = 67
        mcentral mcentral
        // other fields
    }
    // other fields
}

// Central list of free objects of a given size.
type mcentral struct {
    lock      mutex // 分配时需要加锁
    sizeclass int32 // 哪种 sizeclass
    nonempty  mspan // 还有可用的空间的 span 链表
    empty     mspan // 没有可用的空间的 span 列表
}
```

这种方式可以避免出现外部碎片*（文章最后面有外部碎片的介绍）*，因为同一个 span 是按照固定大小分配和回收的，不会出现不可利用的一小块内存把内存分割掉。这个设计方式与现代操作系统中的伙伴系统有点类似。

### mcache

如果你阅读的比较仔细，会发现上面的 mcentral 结构中有一个 lock 字段；因为并发情况下，很有可能多个线程同时从 mcentral 那里申请内存的，必须要用锁来避免冲突。

**但锁是低效的，在高并发的服务中，它会使内存申请成为整个系统的瓶颈；所以在 mcentral 的前面又增加了一层 mcache。**

每一个 mcache 和每一个处理器(P) 是一一对应的，也就是说每一个 P 都有一个 mcache 成员。 Goroutine 申请内存时，首先从其所在的 P 的 mcache 中分配，如果 mcache 没有可用 `span`，再从 mcentral 中获取，并填充到 mcache 中。

**从 mcache 上分配内存空间是不需要加锁的，因为在同一时间里，一个 P 只有一个线程在其上面运行，不可能出现竞争。没有了锁的限制，大大加速了内存分配。**

所以整体的内存分配模型大致如下图所示：

![img](https:////upload-images.jianshu.io/upload_images/11662994-e6d7200368ec06b6.png?)

### 其他优化

#### zero size

有一些对象所需的内存大小是0，比如 `[0]int`, `struct{}`，这种类型的数据根本就不需要内存，所以没必要走上面那么复杂的逻辑。

系统会直接返回一个固定的内存地址。源码如下：

```go
func mallocgc(size uintptr, typ *_type, flags uint32) unsafe.Pointer {
    // 申请的 0 大小空间的内存
    if size == 0 {
        return unsafe.Pointer(&zerobase)
    }
    //.....
}
```

测试代码：

```go
package main

import (
    "fmt"
)

func main() {
    var (
        a struct{}
        b [0]int
        c [100]struct{}
        d = make([]struct{}, 1024)
    )
    fmt.Printf("%p\n", &a)
    fmt.Printf("%p\n", &b)
    fmt.Printf("%p\n", &c)
    fmt.Printf("%p\n", &(d[0]))
    fmt.Printf("%p\n", &(d[1]))
    fmt.Printf("%p\n", &(d[1000]))
}
// 运行结果，6 个变量的内存地址是相同的:
0x1180f88
0x1180f88
0x1180f88
0x1180f88
0x1180f88
0x1180f88
```

#### Tiny对象

上面提到的 `sizeclass=1` 的 span，用来给 `<= 8B` 的对象使用，所以像 `int32`, `byte`, `bool` 以及小字符串等常用的微小对象，都会使用 `sizeclass=1` 的 span，但分配给他们 `8B` 的空间，大部分是用不上的。并且这些类型使用频率非常高，就会导致出现大量的内部碎片。

所以 Go 尽量不使用 `sizeclass=1` 的 span， 而是将 `< 16B` 的对象为统一视为 tiny 对象(tinysize)。分配时，从 `sizeclass=2` 的 span 中获取一个 `16B` 的 object 用以分配。如果存储的对象小于 `16B`，这个空间会被暂时保存起来 (`mcache.tiny` 字段)，下次分配时会复用这个空间，直到这个 object 用完为止。

如图所示：

![img](https:////upload-images.jianshu.io/upload_images/11662994-d026190322b2c139.png?imageMogr2/auto-orient/strip|imageView2/2/w/479/format/webp)

以上图为例，这样的方式空间利用率是 `(1+2+8) / 16 * 100% = 68.75%`，而如果按照原始的管理方式，利用率是 `(1+2+8) / (8 * 3) = 45.83%`。
 源码中注释描述，说是对 tiny 对象的特殊处理，平均会节省 `20%` 左右的内存。

如果要存储的数据里有指针，即使 `<= 8B` 也不会作为 tiny 对象对待，而是正常使用 `sizeclass=1` 的 `span`。

#### 大对象

如上面所述，最大的 sizeclass 最大只能存放 `32K` 的对象。如果一次性申请超过 `32K` 的内存，系统会直接绕过 mcache 和 mcentral，直接从 mheap 上获取，mheap 中有一个 `freelarge` 字段管理着超大 span。

### 总结

内存的释放过程，没什么特别之处。就是分配的返过程，当 mcache 中存在较多空闲 span 时，会归还给 mcentral；而 mcentral 中存在较多空闲 span 时，会归还给 mheap；mheap 再归还给操作系统。这里就不详细介绍了。

总结一下，这种设计之所以快，主要有以下几个优势：

* 内存分配大多时候都是在用户态完成的，不需要频繁进入内核态。

* 每个 P 都有独立的 span cache，多个 CPU 不会并发读写同一块内存，进而减少 CPU L1 cache 的 cacheline 出现 dirty 情况，增大 cpu cache 命中率。

* 内存碎片的问题，Go 是自己在用户态管理的，在 OS 层面看是没有碎片的，使得操作系统层面对碎片的管理压力也会降低。

* mcache 的存在使得内存分配不需要加锁。

当然这不是没有代价的，Go 需要预申请大块内存，这必然会出现一定的浪费，不过好在现在内存比较廉价，不用太在意。

总体上来看，Go 内存管理也是一个金字塔结构：

![img](https:////upload-images.jianshu.io/upload_images/11662994-6d4f174886374a83.png?imageMogr2/auto-orient/strip|imageView2/2/w/484/format/webp)

这种设计比较通用，比如现在常见的 web 服务设计，为提升系统性能，一般都会设计成  `客户端 cache -> 服务端 cache -> 服务端 db` 这几层（当然也可能会加入更多层），也是金字塔结构。



**将有限的计算资源布局成金字塔结构，再将数据从热到冷分为几个层级，放置在金字塔结构上。调度器不断做调整，将热数据放在金字塔顶层，冷数据放在金字塔底层。**

这种设计利用了计算的局部性特征，认为冷热数据的交替是缓慢的。所以最怕的就是，数据访问出现冷热骤变。在操作系统上我们称这种现象为*内存颠簸*，系统架构上通常被说成是*缓存穿透*。其实都是一个意思，就是过度的使用了金字塔低端的资源。

这套内部机制，使得开发高性能服务容易很多，通俗来讲就是坑少了。一般情况下你随便写写性能都不会太差。我遇到过的导致内存分配出现压力的主要有 2 中情况：

1. 频繁申请大对象，常见于文本处理，比如写一个海量日志分析的服务，很多日志内容都很长。这种情况建议自己维护一个对象([]byte)池，避免每次都要去 mheap 上分配。
2. 滥用指针，指针的存在不仅容易造成内存浪费，对 GC 也会造成额外的压力，所以尽量不要使用指针。

### 附

#### 内存碎片

内存碎片是系统在内存管理过程中，会不可避免的出现一块块无法被使用的内存空间，这是内存管理的产物。

##### 内部碎片

一般都是因为字节对齐，如上面介绍 Tiny 对象分配的部分；为了字节对齐，会导致一部分内存空间直接被放弃掉，不做分配使用。
 再比如申请 28B 大小的内存空间，系统会分配 32B 的空间给它，这也导致了其中 4B 空间是被浪费掉的。这就是内部碎片。

##### 外部碎片

一般是因为内存的不断分配释放，导致一些释放的小内存块分散在内存各处，无法被用以分配。如图：

![img](https:////upload-images.jianshu.io/upload_images/11662994-29b706da7d0274f2.png?imageMogr2/auto-orient/strip|imageView2/2/w/531/format/webp)

上面的 8B 和 16B 的小空间，很难再被利用起来。不过 Go 的内存管理机制不会引起大量外部碎片。

#### 源代码调用流程图

针对 Go1.5 源码

![img](https:////upload-images.jianshu.io/upload_images/11662994-a418863ab8ea2190.png?imageMogr2/auto-orient/strip|imageView2/2/w/701/format/webp)

#### runtime.MemStats 部分注释

```go
type MemStats struct {
        // heap 分配出去的字节总数，和 HeapAlloc 值相同
        Alloc uint64

        // TotalAlloc 是 heap 累计分配出去字节数，每次分配
        // 都会累加这个值，但是释放时候不会减少
        TotalAlloc uint64

        // Sys 是指程序从 OS 那里一共申请了多少内存
        // 因为除了 heap，程序栈及其他内部结构都使用着从 OS 申请过来的内存
        Sys uint64

        // Mallocs heap 累积分配出去的对象数
        // 活动中的对象总数，即是 Mallocs - Frees
        Mallocs uint64
       
        // Frees 值 heap 累积释放掉的对象总数
        Frees uint64

        // HeapAlloc 是分配出去的堆对象总和大小，单位字节
        // object 的声明周期是 待分配 -> 分配使用 -> 待回收 -> 待分配
        // 只要不是待分配的状态，都会加到 HeapAlloc 中
        // 它和 HeapInuse 不同，HeapInuse 算的是使用中的 span，
        // 使用中的 span 里面可能还有很多 object 闲置
        HeapAlloc uint64

        // HeapSys 是 heap 从 OS 那里申请来的堆内存大小，单位字节
        // 指的是虚拟内存的大小，不是物理内存，物理内存大小 Go 语言层面是看不到的。
        // 等于 HeapIdle + HeapInuse
        HeapSys uint64

        // HeapIdle 表示所有 span 中还有多少内存是没使用的
        // 这些 span 上面没有 object，也就是完全闲置的，可以随时归还给 OS
        // 也可以用于堆栈分配
        HeapIdle uint64

        // HeapInuse 是处在使用中的所有 span 中的总字节数
        // 只要一个 span 中有至少一个对象，那么就表示它被使用了
        // HeapInuse - HeapAlloc 就表示已经被切割成固定 sizeclass 的 span 里
        HeapInuse uint64

        // HeapReleased 是返回给操作系统的物理内存总数
        HeapReleased uint64

        // HeapObjects 是分配出去的对象总数
        // 如同 HeapAlloc 一样，分配时增加，被清理或被释放时减少
        HeapObjects uint64

        // NextGC is the target heap size of the next GC cycle.
        // NextGC 表示当 HeapAlloc 增长到这个值时，会执行一次 GC
        // 垃圾回收的目标是保持 HeapAlloc ≤ NextGC，每次 GC 结束
        // 下次 GC 的目标，是根据当前可达数据和 GOGC 参数计算得来的
        NextGC uint64

        // LastGC 是最近一次垃圾回收结束的时间 (the UNIX epoch).
        LastGC uint64

        // PauseTotalNs 是自程序启动起， GC 造成 STW 暂停的累积纳秒值
        // STW 期间，所有的 goroutine 都会被暂停，只有 GC 的 goroutine 可以运行
        PauseTotalNs uint64

        // PauseNs 是循环队列，记录着 GC 引起的 STW 总时间
        //
        // 一次 GC 循环，可能会出现多次暂停，这里每项记录的是一次 GC 循环里多次暂停的综合。
        // 最近一次 GC 的数据所在的位置是 PauseNs[NumGC%256]
        PauseNs [256]uint64

        // PauseEnd 是一个循环队列，记录着最近 256 次 GC 结束的时间戳，单位是纳秒。
        //
        // 它和 PauseNs 的存储方式一样。一次 GC 可能会引发多次暂停，这里只记录一次 GC 最后一次暂停的时间
        PauseEnd [256]uint64

        // NumGC 指完成 GC 的次数
        NumGC uint32

        // NumForcedGC 是指应用调用了 runtime.GC() 进行强制 GC 的次数
        NumForcedGC uint32

        // BySize 统计各个 sizeclass 分配和释放的对象的个数
        //
        // BySize[N] 统计的是对象大小 S，满足 BySize[N-1].Size < S ≤ BySize[N].Size 的对象
        // 这里不记录大于 BySize[60].Size 的对象分配
        BySize [61]struct {
                // Size 表示该 sizeclass 的每个对象的空间大小
                // size class.
                Size uint32

                // Mallocs 是该 sizeclass 分配出去的对象的累积总数
                // Size * Mallocs 就是累积分配出去的字节总数
                // Mallocs - Frees 就是当前正在使用中的对象总数
                Mallocs uint64

                // Frees 是该 sizeclass 累积释放对象总数
                Frees uint64
        }
}
```



# Go 语言内存管理（三）：逃逸分析

## 介绍

Go 语言较之 C 语言一个很大的优势就是自带 GC 功能，可 GC 并不是没有代价的。写 C 语言的时候，在一个函数内声明的变量，在函数退出后会自动释放掉，因为这些变量分配在栈上。如果你想要变量的数据能在函数退出后还能访问，就需要调用 `malloc` 方法在堆上申请内存，如果程序不再需要这块内存了，再调用 `free` 方法释放掉。Go 语言不需要你主动调用 `malloc` 来分配堆空间，编译器会自动分析，找出需要 `malloc` 的变量，使用堆内存。编译器的这个分析过程就叫做逃逸分析。

所以你在一个函数中通过 `dict := make(map[string]int)` 创建一个 map 变量，其背后的数据是放在栈空间上还是堆空间上，是不一定的。这要看编译器分析的结果。

可逃逸分析并不是百分百准确的，它有缺陷。有的时候你会发现有些变量其实在栈空间上分配完全没问题的，但编译后程序还是把这些数据放在了堆上。如果你了解 Go 语言编译器逃逸分析的机制，在写代码的时候就可以有意识的绕开这些缺陷，使你的程序更高效。

## 关于堆栈

Go 语言虽然在内存管理方面降低了编程门槛，即使你不了解堆栈也能正常开发，但如果你要在性能上较真的话，还是要掌握这些基础知识。

这里不对堆内存和栈内存的区别做太多阐述。简单来说就是，**栈分配廉价，堆分配昂贵。**栈空间会随着一个函数的结束自动释放，堆空间需要 GC 模块不断的跟踪扫描回收。如果对这两个概念有些迷糊，建议阅读下面 ２ 个文章：

- [Language Mechanics On Stacks And Pointers](https://links.jianshu.com/go?to=https%3A%2F%2Fwww.ardanlabs.com%2Fblog%2F2017%2F05%2Flanguage-mechanics-on-stacks-and-pointers.html)
- [Language Mechanics On Escape Analysis](https://links.jianshu.com/go?to=https%3A%2F%2Fwww.ardanlabs.com%2Fblog%2F2017%2F05%2Flanguage-mechanics-on-escape-analysis.html)

这里举一个小例子，来对比下堆栈的差别：

```go
func stack() int { 
    // 变量 i 会在栈上分配
     i := 10
     return i
}
func heap() *int {
    // 变量 j 会在堆上分配
    j := 10
    return &j
}
```

`stack` 函数中的变量 `i` 在函数退出会自动释放；而 `heap` 函数返回的是对变量`i`的引用，也就是说 `heap()`退出后，表示变量 `i` 还要能被访问，它会自动被分配到堆空间上。

他们编译出来的代码如下：

```asm
// go build --gcflags '-l' test.go
// go tool objdump ./test


TEXT main.stack(SB) /tmp/test.go
  test.go:7     0x487240        48c74424080a000000  MOVQ $0xa, 0x8(SP)  
  test.go:7     0x487249        c3          RET         

TEXT main.heap(SB) /tmp/test.go
  test.go:9     0x487250        64488b0c25f8ffffff  MOVQ FS:0xfffffff8, CX          
  test.go:9     0x487259        483b6110        CMPQ 0x10(CX), SP           
  test.go:9     0x48725d        7639            JBE 0x487298                
  test.go:9     0x48725f        4883ec18        SUBQ $0x18, SP              
  test.go:9     0x487263        48896c2410      MOVQ BP, 0x10(SP)           
  test.go:9     0x487268        488d6c2410      LEAQ 0x10(SP), BP           
  test.go:10        0x48726d        488d05ac090100      LEAQ 0x109ac(IP), AX            
  test.go:10        0x487274        48890424        MOVQ AX, 0(SP)              
  test.go:10        0x487278        e8f33df8ff      CALL runtime.newobject(SB)      
  test.go:10        0x48727d        488b442408      MOVQ 0x8(SP), AX            
  test.go:10        0x487282        48c7000a000000      MOVQ $0xa, 0(AX)            
  test.go:11        0x487289        4889442420      MOVQ AX, 0x20(SP)           
  test.go:11        0x48728e        488b6c2410      MOVQ 0x10(SP), BP           
  test.go:11        0x487293        4883c418        ADDQ $0x18, SP              
  test.go:11        0x487297        c3          RET                 
  test.go:9     0x487298        e8a380fcff      CALL runtime.morestack_noctxt(SB)   
  test.go:9     0x48729d        ebb1            JMP main.heap(SB)           
// ...

TEXT runtime.newobject(SB) /usr/share/go/src/runtime/malloc.go
  malloc.go:1067    0x40b070        64488b0c25f8ffffff  MOVQ FS:0xfffffff8, CX          
  malloc.go:1067    0x40b079        483b6110        CMPQ 0x10(CX), SP           
  malloc.go:1067    0x40b07d        763d            JBE 0x40b0bc                
  malloc.go:1067    0x40b07f        4883ec28        SUBQ $0x28, SP              
  malloc.go:1067    0x40b083        48896c2420      MOVQ BP, 0x20(SP)           
  malloc.go:1067    0x40b088        488d6c2420      LEAQ 0x20(SP), BP           
  malloc.go:1068    0x40b08d        488b442430      MOVQ 0x30(SP), AX           
  malloc.go:1068    0x40b092        488b08          MOVQ 0(AX), CX              
  malloc.go:1068    0x40b095        48890c24        MOVQ CX, 0(SP)              
  malloc.go:1068    0x40b099        4889442408      MOVQ AX, 0x8(SP)            
  malloc.go:1068    0x40b09e        c644241001      MOVB $0x1, 0x10(SP)         
  malloc.go:1068    0x40b0a3        e888f4ffff      CALL runtime.mallocgc(SB)       
  malloc.go:1068    0x40b0a8        488b442418      MOVQ 0x18(SP), AX           
  malloc.go:1068    0x40b0ad        4889442438      MOVQ AX, 0x38(SP)           
  malloc.go:1068    0x40b0b2        488b6c2420      MOVQ 0x20(SP), BP           
  malloc.go:1068    0x40b0b7        4883c428        ADDQ $0x28, SP              
  malloc.go:1068    0x40b0bb        c3          RET                 
  malloc.go:1067    0x40b0bc        e87f420400      CALL runtime.morestack_noctxt(SB)   
  malloc.go:1067    0x40b0c1        ebad            JMP runtime.newobject(SB)       
```

逻辑的复杂度不言而喻，上面的汇编中可看到， `heap()` 函数调用了 `runtime.newobject()` 方法，它会调用 `mallocgc` 方法从 `mcache` 上申请内存，申请的内部逻辑前面文章已经讲述过。堆内存分配不仅分配上逻辑比栈空间分配复杂，它最致命的是会带来很大的管理成本，Go 语言要消耗很多的计算资源对其进行标记回收（也就是 GC 成本）。

> 不要以为使用了堆内存就一定会导致性能低下，使用栈内存会带来性能优势。因为实际项目中，系统的性能瓶颈一般都不会出现在内存分配上。千万不要盲目优化，找到系统瓶颈，用数据驱动优化。

## 逃逸分析

Go 编辑器会自动帮我们找出需要进行动态分配的变量，它是在编译时追踪一个变量的生命周期，如果能确认一个数据只在函数空间内访问，不会被外部使用，则使用栈空间，否则就要使用堆空间。

我们在 `go build` 编译代码时，可使用 `-gcflags '-m'` 参数来查看逃逸分析日志。

```sh
go build -gcflags '-m -m' test.go
```

以上面的两个函数为例，编译的日志输出是：

```go
/tmp/test.go:11:9: &i escapes to heap
/tmp/test.go:11:9:  from ~r0 (return) at /tmp/test.go:11:2
/tmp/test.go:10:2: moved to heap: i
/tmp/test.go:16:18: heap() escapes to heap
/tmp/test.go:16:18:     from ... argument (arg to ...) at /tmp/test.go:16:13
/tmp/test.go:16:18:     from *(... argument) (indirection) at /tmp/test.go:16:13
/tmp/test.go:16:18:     from ... argument (passed to call[argument content escapes]) at /tmp/test.go:16:13
/tmp/test.go:16:13: main ... argument does not escape
```

日志中的 `&i escapes to heap` 表示该变量数据逃逸到了堆上。

### 逃逸分析的缺陷

需要使用堆空间则逃逸，这没什么可争议的。但编译器有时会将**不需要**使用堆空间的变量，也逃逸掉。这里是容易出现性能问题的大坑。网上有很多相关文章，列举了一些导致逃逸情况，其实总结起来就一句话：

**多级间接赋值容易导致逃逸**。

这里的多级间接指的是，对某个引用类对象中的引用类成员进行赋值。Go 语言中的引用类数据类型有 `func`, `interface`, `slice`, `map`, `chan`, `*Type(指针)`。

记住公式 `Data.Field = Value`，如果 `Data`, `Field` 都是引用类的数据类型，则会导致 `Value` 逃逸。这里的等号 `=` 不单单只赋值，也表示参数传递。

根据公式，我们假设一个变量 `data` 是以下几种类型，相应的可得出结论：

- `[]interface{}`: `data[0] = 100` 会导致 `100` 逃逸
- `map[string]interface{}`: `data["key"] = "value"` 会导致 `"value"` 逃逸
- `map[interface{}]interface{}`: `data["key"] = "value"` 会导致 `key` 和 `value` 都逃逸
- `map[string][]string`: `data["key"] = []string{"hello"}` 会导致切片逃逸
- `map[string]*int`: 赋值时 `*int` 会 逃逸
- `[]*int`: `data[0] = &i` 会使 `i` 逃逸
- `func(*int)`: `data(&i)` 会使 `i` 逃逸
- `func([]string)`: `data([]{"hello"})` 会使 `[]string{"hello"}` 逃逸
- `chan []string`: `data <- []string{"hello"}` 会使 `[]string{"hello"}` 逃逸
- 以此类推，不一一列举了

下面给出一些实际的例子：

#### 函数变量

如果变量值是一个函数，函数的参数又是引用类型，则传递给它的参数都会逃逸。

```go
func test(i int)        {}
func testEscape(i *int) {}

func main() {
    i, j, m, n := 0, 0, 0, 0
    t, te := test, testEscape // 函数变量

    // 直接调用
    test(m)        // 不逃逸
    testEscape(&n) // 不逃逸
    // 间接调用
    t(i)   // 不逃逸
    te(&j) // 逃逸
}
```



```go
./test.go:4:17: testEscape i does not escape
./test.go:11:5: &j escapes to heap
./test.go:11:5:     from te(&j) (parameter to indirect call) at ./test.go:11:4
./test.go:7:5: moved to heap: j
./test.go:14:13: main &n does not escape
```

上例中 `te` 的类型是 `func(*int)`，属于引用类型，参数 `*int` 也是引用类型，则调用 `te(&j)` 形成了为 `te` 的参数(成员) `*int` 赋值的现象，即 `te.i = &j` 会导致逃逸。代码中其他几种调用都没有形成**多级间接赋值**情况。
同理，如果函数的参数类型是 `slice`, `map` 或 `interface{}` 都会导致参数逃逸。

```go
func testSlice(slice []int)       {}
func testMap(m map[int]int)       {}
func testInterface(i interface{}) {}

func main() {
    x, y, z := make([]int, 1), make(map[int]int), 100
    ts, tm, ti := testSlice, testMap, testInterface
    ts(x) // ts.slice = x 导致 x 逃逸
    tm(y) // tm.m = y 导致 y 逃逸
    ti(z) // ti.i = z 导致 z 逃逸
}
```



```go
./test.go:3:16: testSlice slice does not escape
./test.go:4:14: testMap m does not escape
./test.go:5:20: testInterface i does not escape
./test.go:8:17: make([]int, 1) escapes to heap
./test.go:8:17:     from x (assign-pair) at ./test.go:8:10
./test.go:8:17:     from ts(x) (parameter to indirect call) at ./test.go:10:4
./test.go:8:33: make(map[int]int) escapes to heap
./test.go:8:33:     from y (assign-pair) at ./test.go:8:10
./test.go:8:33:     from tm(y) (parameter to indirect call) at ./test.go:11:4
./test.go:12:4: z escapes to heap
./test.go:12:4:     from ti(z) (parameter to indirect call) at ./test.go:12:4
```

匿名函数的调用也是一样的，它本质上也是一个函数变量。有兴趣的可以自己测试一下。

#### 间接赋值

```go
type Data struct {
    data  map[int]int
    slice []int
    ch    chan int
    inf   interface{}
    p     *int
}

func main() {
    d1 := Data{}
    d1.data = make(map[int]int) // GOOD: does not escape
    d1.slice = make([]int, 4)   // GOOD: does not escape
    d1.ch = make(chan int, 4)   // GOOD: does not escape
    d1.inf = 3                  // GOOD: does not escape
    d1.p = new(int)             //  GOOD: does not escape

    d2 := new(Data)             // d2 是指针变量， 下面为该指针变量中的指针成员赋值
    d2.data = make(map[int]int) // BAD: escape to heap
    d2.slice = make([]int, 4)   // BAD:  escape to heap
    d2.ch = make(chan int, 4)   // BAD:  escape to heap
    d2.inf = 3                  // BAD:  escape to heap
    d2.p = new(int)             // BAD:  escape to heap
}
```



```go
./test.go:20:16: make(map[int]int) escapes to heap
./test.go:20:16:    from d2.data (star-dot-equals) at ./test.go:20:10
./test.go:21:17: make([]int, 4) escapes to heap
./test.go:21:17:    from d2.slice (star-dot-equals) at ./test.go:21:11
./test.go:22:14: make(chan int, 4) escapes to heap
./test.go:22:14:    from d2.ch (star-dot-equals) at ./test.go:22:8
./test.go:23:9: 3 escapes to heap
./test.go:23:9:     from d2.inf (star-dot-equals) at ./test.go:23:9
./test.go:24:12: new(int) escapes to heap
./test.go:24:12:    from d2.p (star-dot-equals) at ./test.go:24:7
./test.go:13:16: main make(map[int]int) does not escape
./test.go:14:17: main make([]int, 4) does not escape
./test.go:15:14: main make(chan int, 4) does not escape
./test.go:16:9: main 3 does not escape
./test.go:17:12: main new(int) does not escape
./test.go:19:11: main new(Data) does not escape
```

#### Interface

只要使用了 `Interface` 类型(不是 `interafce{}`)，那么赋值给它的变量一定会逃逸。因为 `interfaceVariable.Method()` 先是间接的定位到它的实际值，再调用实际值的同名方法，执行时实际值作为参数传递给方法。相当于`interfaceVariable.Method.this = realValue`



```go
type Iface interface {
    Dummy()
}
type Integer int
func (i Integer) Dummy() {}

func main() {
    var (
        iface Iface
        i     Integer
    )
    iface = i
    iface.Dummy() //  make i escape to heap
    // 形成 iface.Dummy.i = i
}
```

#### 引用类型的 channel

向 channel 中发送数据，本质上就是为 channel 内部的成员赋值，就像给一个 slice 中的某一项赋值一样。所以 `chan *Type`, `chan map[Type]Type`, `chan []Type`, `chan interface{}` 类型都会导致发送到 channel 中的数据逃逸。

这本来也是情理之中的，发送给 channel 的数据是要与其他函数分享的，为了保证发送过去的指针依然可用，只能使用堆分配。

```go
func test() {
    var (
        chInteger   = make(chan *int)
        chMap       = make(chan map[int]int)
        chSlice     = make(chan []int)
        chInterface = make(chan interface{})
        a, b, c, d  = 0, map[int]int{}, []int{}, 32
    )
    chInteger <- &a  // 逃逸
    chMap <- b       // 逃逸
    chSlice <- c     // 逃逸
    chInterface <- d // 逃逸
}
```



```go
./escape.go:11:15: &a escapes to heap
./escape.go:11:15:  from chInteger <- &a (send) at ./escape.go:11:12
./escape.go:9:3: moved to heap: a
./escape.go:9:31: map[int]int literal escapes to heap
./escape.go:9:31:   from b (assigned) at ./escape.go:9:3
./escape.go:9:31:   from chMap <- b (send) at ./escape.go:12:8
./escape.go:9:40: []int literal escapes to heap
./escape.go:9:40:   from c (assigned) at ./escape.go:9:3
./escape.go:9:40:   from chSlice <- c (send) at ./escape.go:13:10
./escape.go:14:14: d escapes to heap
./escape.go:14:14:  from chInterface <- (interface {})(d) (send) at ./escape.go:14:14
./escape.go:5:21: test make(chan *int) does not escape
./escape.go:6:21: test make(chan map[int]int) does not escape
./escape.go:7:21: test make(chan []int) does not escape
./escape.go:8:21: test make(chan interface {}) does not escape
```

#### 可变参数

可变参数如 `func(arg ...string)` 实际与 `func(arg []string)` 是一样的，会增加一层访问路径。这也是 `fmt.Sprintf` 总是会使参数逃逸的原因。

例子非常多，这里不能一一列举，我们只需要记住分析方法就好，即，2 级或更多级的访问赋值会**容易**导致数据逃逸。这里加上**容易**二字是因为随着语言的发展，相信这些问题会被慢慢解决，但现阶段，这个可以作为我们分析逃逸现象的依据。

下面代码中包含 2 种很常规的写法，但他们却有着很大的性能差距，建议自己想下为什么。

```go
type User struct {
    roles []string
}

func (u *User) SetRoles(roles []string) {
    u.roles = roles
}

func SetRoles(u User, roles []string) User {
    u.roles = roles
    return u
}
```

Benchmark 和 pprof 给出的结果:



```go
BenchmarkUserSetRoles-8     50000000            22.3 ns/op        16 B/op          1 allocs/op
BenchmarkSetRoles-8         2000000000           0.51 ns/op        0 B/op          0 allocs/op

  768.01MB   768.01MB (flat, cum)   100% of Total
         .          .      3:import "testing"
         .          .      4:
         .          .      5:func BenchmarkUserSetRoles(b *testing.B) {
         .          .      6:   u := new(User)
         .          .      7:   for i := 0; i < b.N; i++ {
  768.01MB   768.01MB      8:       u.SetRoles([]string{"admin"}) <- 看这里
         .          .      9:   }
         .          .     10:}
         .          .     11:
         .          .     12:func BenchmarkSetRoles(b *testing.B) {
         .          .     13:   for i := 0; i < b.N; i++ {
ROUTINE ======================== testing.(*B).launch in /usr/share/go/src/testing/benchmark.go
......
```

## 结论

熟悉堆栈概念可以让我们更容易看透 Go 程序的性能问题，并进行优化。

多级间接赋值会导致 Go 编译器出现不必要的逃逸，在一些情况下可能我们只需要修改一下数据结构就会使性能有大幅提升。这也是很多人不推荐在 Go 中使用指针的原因，因为它会增加一级访问路径，而 `map`, `slice`, `interface{}`等类型是不可避免要用到的，为了减少不必要的逃逸，只能拿指针开刀了。

大多数情况下，性能优化都会为程序带来一定的复杂度。建议实际项目中还是怎么方便怎么写，功能完成后通过性能分析找到瓶颈所在，再对局部进行优化。
