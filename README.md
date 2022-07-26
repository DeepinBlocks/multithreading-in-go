# multithreading-in-go
## Memory Sharing
Inter-process communication(IPC)
* message passing  
threads are passing messages to each other.
* shared memory  
threads are sharing a common space in memory, and one of the threads writes particular value and the other one reads it.
### message passing
### shared memory
Communication between two threads using memory sharing is easy and efficient because threads are sharing their memory space.
You can have multiple variables or a common space where both threads read and write from.