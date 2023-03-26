<<<<<<< Hi!
# Concurrent Image Processing
This is a project written in C from my parallel programming class MPCS52060
A concurrent image processing system that parallelizes two dimensional image convolution with >2x speedup. In edition 1, you can find two designs of the concurrent program: Fan-In-Fan-Out Model and Bulk Synchronization Model. In edition 2. 
=======
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
>>>>>>> 2a79667
