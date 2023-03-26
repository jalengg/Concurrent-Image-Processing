Hi!

# Concurrent Image Processing

=======

This is a project written in G for a parallel programming class MPCS 52060.
It features concurrent image processing system that parallelizes two dimensional image convolution with >2x speedup. In edition 1, you can find two designs of the concurrent program: Fan-In-Fan-Out Model and Bulk Synchronization Model. In edition 2. 

# Edition 1 
Please make sure you are in the following directory to run the program: 
```
/edition_1/editor/
```
Usage Specification:
```
Usage: editor data_dir [mode] [number_of_threads]
data_dir = The data directories to use to load the images.
mode     = (bsp) run the BSP mode, (pipeline) run the pipeline mode
number_of_threads = Runs the parallel version of the program with the specified number of threads (i.e., goroutines).
```

For example, running the program as follows:
```
$: go run editor.go big bsp 4
```
Here’s an example of a combination run:
```
$: go run editor.go big+small pipeline 2
```
will produce inside the out directory the following files:
```
big_IMG_2020_Out.png
big_IMG_2724_Out.png
big_IMG_3695_Out.png
big_IMG_3696_Out.png
big_IMG_3996_Out.png
big_IMG_4061_Out.png
big_IMG_4065_Out.png
big_IMG_4066_Out.png
big_IMG_4067_Out.png
big_IMG_4069_Out.png
small_IMG_2020_Out.png
small_IMG_2724_Out.png
small_IMG_3695_Out.png
small_IMG_3696_Out.png
small_IMG_3996_Out.png
small_IMG_4061_Out.png
small_IMG_4065_Out.png
small_IMG_4066_Out.png
small_IMG_4067_Out.png
small_IMG_4069_Out.png
```
The structure of the Fan-In-Fan-Out Model is illustrated in the following scheme: 
<img width="784" alt="image" src="https://user-images.githubusercontent.com/66903483/227790549-befe3ab4-a609-4324-ac75-6e8943b60b1a.png">

The structure of the Bulk Synchronization Model is illustrated in the following scheme: 

<img width="552" alt="image" src="https://user-images.githubusercontent.com/66903483/227790632-25602670-76a1-483c-8e91-f91882927327.png">

In ./edition_1/benchmark, you can find all relevant benchmarking pythong and bash scripts. Additionally, there is a detailed report for the system there. 

> BSP speedup increases to around 2 as thread number increases to 6, after which the speedup plateaus around 2 regardless of picture size. This eventual plateau is expected as BSP has a global synchronization step which create a bottleneck: no matter how fast each super-step is, the global synchronization step dictates the fastest speed-up as the global synchronization step takes constant time, can’t be parallelized, and has to happen when all goroutines are paused.
Pipeline speedup, on the other hand, seems uniform across thread numbers. The speedup for big pictures with 12 threads is about the same as the speedup with 2 threads, which is around 2.5. This also make sense because the pipeline implementation creates a total of N*N goroutines for each run, where N is the command line input thread number. In other words, when you put 2 in the command line, the program actually spawns a total of 4 goroutines. This quickly exhausts physical cores and the advantage of physical concurrency, and the context-switching has to be coordinated by the scheduler to handle application concurrency. This scheduling and context switching overhead cannot be parallelized and become the upper bound of the speedup graph.

<img width="588" alt="image" src="https://user-images.githubusercontent.com/66903483/227790826-11956bbf-9106-4e62-9773-f7c9cda620f2.png">
<img width="595" alt="image" src="https://user-images.githubusercontent.com/66903483/227790868-ba5ffcaa-819f-467e-a143-cbf33f08579a.png">

# Edition 2

Edition 2 is set up in a similar structure as in edition 1, but used different synchronization techniques: it utilizes WorkStealing/WorkBalancing Algorithms and MapReduce structure. For a detailed report, please go to [Edition2 Report](/edition2_new/proj3%20report.pdf)

Scheme for MapReduce: 
<img width="528" alt="image" src="https://user-images.githubusercontent.com/66903483/227794897-c7010eb5-986c-4152-8054-9b854ebc3871.png">


