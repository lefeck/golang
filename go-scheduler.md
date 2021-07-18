## 设计原理 [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#651-设计原理)

今天的 Go 语言调度器有着优异的性能，但是如果我们回头看 Go 语言的 0.x 版本的调度器会发现最初的调度器不仅实现非常简陋，也无法支撑高并发的服务。调度器经过几个大版本的迭代才有今天的优异性能，历史上几个不同版本的调度器引入了不同的改进，也存在着不同的缺陷:

- 单线程调度器 v0.x
  - 只包含 40 多行代码；
  - 程序中只能存在一个活跃线程，由 G-M 模型组成；
- 多线程调度器 v1.0
  - 允许运行多线程的程序；
  - 全局锁导致竞争严重；
- 任务窃取调度器 v1.1
  - 引入了处理器 P，构成了目前的 **G-M-P** 模型；
  - 在处理器 P 的基础上实现了基于**工作窃取**的调度器；
  - 在某些情况下，Goroutine 不会让出线程，进而造成饥饿问题；
  - 时间过长的垃圾回收（Stop-the-world，STW）会导致程序长时间无法工作；
- 抢占式调度器 v1.2 ~ 至今
  - 基于协作的抢占式调度器 - v1.2 ~ v1.13
    - 通过编译器在函数调用时插入**抢占检查**指令，在函数调用时检查当前 Goroutine 是否发起了抢占请求，实现基于协作的抢占式调度；
    - Goroutine 可能会因为垃圾回收和循环长时间占用资源导致程序暂停；
  - 基于信号的抢占式调度器 - v1.14 ~ 至今
    - 实现**基于信号的真抢占式调度**；
    - 垃圾回收在扫描栈时会触发抢占调度；
    - 抢占的时间点不够多，还不能覆盖全部的边缘情况；
- 非均匀存储访问调度器 · 提案
  - 对运行时的各种资源进行分区；
  - 实现非常复杂，到今天还没有提上日程；

### 多线程调度器 [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#多线程调度器)

多线程调度器的主要问题是调度时的锁竞争会严重浪费资源，[Scalable Go Scheduler Design Doc](http://golang.org/s/go11sched) 中对调度器做的性能测试发现 14% 的时间都花费在 [`runtime.futex:go1.0.1`](https://draveness.me/golang/tree/runtime.futex:go1.0.1) 上[3](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#fn:3)，该调度器有以下问题需要解决：

1. 调度器和锁是全局资源，所有的调度状态都是中心化存储的，锁竞争问题严重；
2. 线程需要经常互相传递可运行的 Goroutine，引入了大量的延迟；
3. 每个线程都需要处理内存缓存，导致大量的内存占用并影响数据局部性；
4. 系统调用频繁阻塞和解除阻塞正在运行的线程，增加了额外开销；

这里的全局锁问题和 Linux 操作系统调度器在早期遇到的问题比较相似，解决的方案也都大同小异。

### 任务窃取调度器 [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#任务窃取调度器)

2012 年 Google 的工程师 Dmitry Vyukov 在 [Scalable Go Scheduler Design Doc](http://golang.org/s/go11sched) 中指出了现有多线程调度器的问题并在多线程调度器上提出了两个改进的手段：

1. 在当前的 G-M 模型中引入了处理器 P，增加中间层；
2. 在处理器 P 的基础上实现基于工作窃取的调度器；

基于任务窃取的 Go 语言调度器使用了沿用至今的 G-M-P 模型，我们能在 [runtime: improved scheduler](https://github.com/golang/go/commit/779c45a50700bda0f6ec98429720802e6c1624e8) 提交中找到任务窃取调度器刚被实现时的[源代码](https://github.com/golang/go/blob/779c45a50700bda0f6ec98429720802e6c1624e8/src/pkg/runtime/proc.c)，调度器的 [`runtime.schedule:779c45a`](https://draveness.me/golang/tree/runtime.schedule:779c45a) 在这个版本的调度器中反而更简单了：

```go
static void schedule(void) {
    G *gp;
 top:
    if(runtime·gcwaiting) {
        gcstopm();
        goto top;
    }

    gp = runqget(m->p);
    if(gp == nil)
        gp = findrunnable();

    ...

    execute(gp);
}
```

1. 如果当前运行时在等待垃圾回收，调用 [`runtime.gcstopm:779c45a`](https://draveness.me/golang/tree/runtime.gcstopm:779c45a) 函数；
2. 调用 [`runtime.runqget:779c45a`](https://draveness.me/golang/tree/runtime.runqget:779c45a) 和 [`runtime.findrunnable:779c45a`](https://draveness.me/golang/tree/runtime.findrunnable:779c45a) 从本地或者全局的运行队列中获取待执行的 Goroutine；
3. 调用 [`runtime.execute:779c45a`](https://draveness.me/golang/tree/runtime.execute:779c45a) 在当前线程 M 上运行 Goroutine；

当前处理器本地的运行队列中不包含 Goroutine 时，调用 [`runtime.findrunnable:779c45a`](https://draveness.me/golang/tree/runtime.findrunnable:779c45a) 会触发工作窃取，从其它的处理器的队列中随机获取一些 Goroutine。

运行时 G-M-P 模型中引入的处理器 P 是线程和 Goroutine 的中间层，我们从它的结构体中就能看到处理器与 M 和 G 的关系：

```c
struct P {
	Lock;

	uint32	status;
	P*	link;
	uint32	tick;
	M*	m;
	MCache*	mcache;

	G**	runq;
	int32	runqhead;
	int32	runqtail;
	int32	runqsize;

	G*	gfree;
	int32	gfreecnt;
};
```

处理器持有一个由可运行的 Goroutine 组成的环形的运行队列 `runq`，还反向持有一个线程。调度器在调度时会从处理器的队列中选择队列头的 Goroutine 放到线程 M 上执行。如下所示的图片展示了 Go 语言中的线程 M、处理器 P 和 Goroutine 的关系。

![golang-gmp](https://img.draveness.me/2020-02-02-15805792666151-golang-gmp.png)

**图 6-27 - G-M-P 模型**

基于工作窃取的多线程调度器将每一个线程绑定到了独立的 CPU 上，这些线程会被不同处理器管理，不同的处理器通过工作窃取对任务进行再分配实现任务的平衡，也能提升调度器和 Go 语言程序的整体性能，今天所有的 Go 语言服务都受益于这一改动。

### 抢占式调度器 [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#抢占式调度器)

对 Go 语言并发模型的修改提升了调度器的性能，但是 1.1 版本中的调度器仍然不支持抢占式调度，程序只能依靠 Goroutine 主动让出 CPU 资源才能触发调度。Go 语言的调度器在 1.2 版本[4](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#fn:4)中引入基于协作的抢占式调度解决下面的问题[5](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#fn:5)：

- 某些 Goroutine 可以长时间占用线程，造成其它 Goroutine 的饥饿；
- 垃圾回收需要暂停整个程序（Stop-the-world，STW），最长可能需要几分钟的时间[6](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#fn:6)，导致整个程序无法工作；

1.2 版本的抢占式调度虽然能够缓解这个问题，但是它实现的抢占式调度是基于协作的，在之后很长的一段时间里 Go 语言的调度器都有一些无法被抢占的边缘情况，例如：for 循环或者垃圾回收长时间占用线程，这些问题中的一部分直到 1.14 才被基于信号的抢占式调度解决。

#### 基于协作的抢占式调度 [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#基于协作的抢占式调度)

基于协作的抢占式调度的工作原理：

1. 编译器会在调用函数前插入 [`runtime.morestack`](https://draveness.me/golang/tree/runtime.morestack)；
2. Go 语言运行时会在垃圾回收暂停程序、系统监控发现 Goroutine 运行超过 10ms 时发出抢占请求 `StackPreempt`；
3. 当发生函数调用时，可能会执行编译器插入的 [`runtime.morestack`](https://draveness.me/golang/tree/runtime.morestack)，它调用的 [`runtime.newstack`](https://draveness.me/golang/tree/runtime.newstack) 会检查 Goroutine 的 `stackguard0` 字段是否为 `StackPreempt`；
4. 如果 `stackguard0` 是 `StackPreempt`，就会触发抢占让出当前线程；

这种实现方式虽然增加了运行时的复杂度，但是实现相对简单，也没有带来过多的额外开销，总体来看还是比较成功的实现，也在 Go 语言中使用了 10 几个版本。因为这里的抢占是通过编译器插入函数实现的，还是需要函数调用作为入口才能触发抢占，所以这是一种**协作式的抢占式调度**。

#### 基于信号的抢占式调度 [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#基于信号的抢占式调度)

Go 语言在 1.14 版本中实现了非协作的抢占式调度，在实现的过程中我们重构已有的逻辑并为 Goroutine 增加新的状态和字段来支持抢占。Go 团队通过下面的一系列提交实现了这一功能，我们可以按时间顺序分析相关提交理解它的工作原理：

- runtime: add general suspendG/resumeG
  - 挂起 Goroutine 的过程是在垃圾回收的栈扫描时完成的，我们通过 [`runtime.suspendG`](https://draveness.me/golang/tree/runtime.suspendG) 和 [`runtime.resumeG`](https://draveness.me/golang/tree/runtime.resumeG) 两个函数重构栈扫描这一过程；
  - 调用 [`runtime.suspendG`](https://draveness.me/golang/tree/runtime.suspendG) 时会将处于运行状态的 Goroutine 的 `preemptStop` 标记成 `true`；
  - 调用 [`runtime.preemptPark`](https://draveness.me/golang/tree/runtime.preemptPark) 可以挂起当前 Goroutine、将其状态更新成 `_Gpreempted` 并触发调度器的重新调度，该函数能够交出线程控制权；
- runtime: asynchronous preemption function for x86
  - 在 x86 架构上增加异步抢占的函数 [`runtime.asyncPreempt`](https://draveness.me/golang/tree/runtime.asyncPreempt) 和 [`runtime.asyncPreempt2`](https://draveness.me/golang/tree/runtime.asyncPreempt2)；
- runtime: use signals to preempt Gs for suspendG
  - 支持通过向线程发送信号的方式暂停运行的 Goroutine；
  - 在 [`runtime.sighandler`](https://draveness.me/golang/tree/runtime.sighandler) 函数中注册 `SIGURG` 信号的处理函数 [`runtime.doSigPreempt`](https://draveness.me/golang/tree/runtime.doSigPreempt)；
  - 实现 [`runtime.preemptM`](https://draveness.me/golang/tree/runtime.preemptM)，它可以通过 `SIGURG` 信号向线程发送抢占请求；
- runtime: implement async scheduler preemption
  - 修改 [`runtime.preemptone`](https://draveness.me/golang/tree/runtime.preemptone) 函数的实现，加入异步抢占的逻辑；

目前的抢占式调度也只会在垃圾回收扫描任务时触发，我们可以梳理一下上述代码实现的抢占式调度过程：

1. 程序启动时，在 [`runtime.sighandler`](https://draveness.me/golang/tree/runtime.sighandler) 中注册 `SIGURG` 信号的处理函数 [`runtime.doSigPreempt`](https://draveness.me/golang/tree/runtime.doSigPreempt)；
2. 在触发垃圾回收的栈扫描时会调用`runtime.suspendG`挂起 Goroutine，该函数会执行下面的逻辑：
   1. 将 `_Grunning` 状态的 Goroutine 标记成可以被抢占，即将 `preemptStop` 设置成 `true`；
   2. 调用 [`runtime.preemptM`](https://draveness.me/golang/tree/runtime.preemptM) 触发抢占；
3. [`runtime.preemptM`](https://draveness.me/golang/tree/runtime.preemptM) 会调用 [`runtime.signalM`](https://draveness.me/golang/tree/runtime.signalM) 向线程发送信号 `SIGURG`；
4. 操作系统会中断正在运行的线程并执行预先注册的信号处理函数 [`runtime.doSigPreempt`](https://draveness.me/golang/tree/runtime.doSigPreempt)；
5. [`runtime.doSigPreempt`](https://draveness.me/golang/tree/runtime.doSigPreempt) 函数会处理抢占信号，获取当前的 SP 和 PC 寄存器并调用 [`runtime.sigctxt.pushCall`](https://draveness.me/golang/tree/runtime.sigctxt.pushCall)；
6. [`runtime.sigctxt.pushCall`](https://draveness.me/golang/tree/runtime.sigctxt.pushCall) 会修改寄存器并在程序回到用户态时执行 [`runtime.asyncPreempt`](https://draveness.me/golang/tree/runtime.asyncPreempt)；
7. 汇编指令 [`runtime.asyncPreempt`](https://draveness.me/golang/tree/runtime.asyncPreempt) 会调用运行时函数 [`runtime.asyncPreempt2`](https://draveness.me/golang/tree/runtime.asyncPreempt2)；
8. [`runtime.asyncPreempt2`](https://draveness.me/golang/tree/runtime.asyncPreempt2) 会调用 [`runtime.preemptPark`](https://draveness.me/golang/tree/runtime.preemptPark)；
9. [`runtime.preemptPark`](https://draveness.me/golang/tree/runtime.preemptPark) 会修改当前 Goroutine 的状态到 `_Gpreempted` 并调用 [`runtime.schedule`](https://draveness.me/golang/tree/runtime.schedule) 让当前函数陷入休眠并让出线程，调度器会选择其它的 Goroutine 继续执行；

上述 9 个步骤展示了基于信号的抢占式调度的执行过程。除了分析抢占的过程之外，我们还需要讨论一下抢占信号的选择，提案根据以下的四个原因选择 `SIGURG` 作为触发异步抢占的信号[7](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#fn:7)；

1. 该信号需要被调试器透传；
2. 该信号不会被内部的 libc 库使用并拦截；
3. 该信号可以随意出现并且不触发任何后果；
4. 我们需要处理多个平台上的不同信号；