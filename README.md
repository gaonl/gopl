## 内存模型  happens-before
1. 初始化  init()函数
2. 创建goroutine(注意：销毁goroutine并不提供保证)
3. 通道（区分带缓冲和不带缓冲）
4. 锁
5. 一次执行（once.Do(setup)）
[详细的，请请点击我](https://blog.csdn.net/zhailuxu/article/details/88085864)
