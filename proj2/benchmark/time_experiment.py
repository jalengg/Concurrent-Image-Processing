import sys
import os

# runs experiment 5 times for each type of implementation and data directory for each of [2, 4, 6, 8, 12] cores
# USAGE: python3 time_experiment.py


# BSP
# os.system("echo small > output_bsp")
# os.system("echo ================ >> output_bsp")
# os.system("echo 1 >> output_bsp")
# for i in range(5):
    # os.system("(time go run ../editor/editor.go small) 2>> output_bsp")
# for core in range(2, 13, 2):
    # os.system("echo " + str(core) + " >> output_bsp")
    # for i in range(5):
        # instr = "time go run ../editor/editor.go small bsp " + str(core)
        # formatted = "(" + instr + ") 2>> output_bsp"
        # os.system(formatted)
# 
# os.system("echo mixture >> output_bsp")
# os.system("echo ================ >> output_bsp")
# os.system("echo 1 >> output_bsp")
# for i in range(5):
    # os.system("(time go run ../editor/editor.go mixture) 2>> output_bsp")
# for core in range(2, 13, 2):
    # os.system("echo " + str(core) + " >> output_bsp")
    # for i in range(5):
        # instr = "time go run ../editor/editor.go mixture bsp " + str(core)
        # formatted = "(" + instr + ") 2>> output_bsp"
        # os.system(formatted)
# 
# os.system("echo big >> output_bsp")
# os.system("echo ================ >> output_bsp")
# os.system("echo 1 >> output_bsp")
# for i in range(5):
    # os.system("(time go run ../editor/editor.go big) 2>> output_bsp")
# for core in range(2, 13, 2):
    # os.system("echo " + str(core) + " >> output_bsp")
    # for i in range(5):
        # instr = "time go run ../editor/editor.go big bsp " + str(core)
        # formatted = "(" + instr + ") 2>> output_bsp"
        # os.system(formatted)

## Pipeline

os.system("echo small > output_pipeline")
os.system("echo ================ >> output_pipeline")
os.system("echo 1 >> output_pipeline")
for i in range(5):
    os.system("(time go run ../editor/editor.go small) 2>> output_pipeline")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_pipeline")
    for i in range(5):
        instr = "time go run ../editor/editor.go small pipeline " + str(core)
        formatted = "(" + instr + ") 2>> output_pipeline"
        os.system(formatted)

os.system("echo mixture >> output_pipeline")
os.system("echo ================ >> output_pipeline")
os.system("echo 1 >> output_pipeline")
for i in range(5):
    os.system("(time go run ../editor/editor.go mixture) 2>> output_pipeline")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_pipeline")
    for i in range(5):
        instr = "time go run ../editor/editor.go mixture pipeline " + str(core)
        formatted = "(" + instr + ") 2>> output_pipeline"
        os.system(formatted)

os.system("echo big >> output_pipeline")
os.system("echo ================ >> output_pipeline")
os.system("echo 1 >> output_pipeline")
for i in range(5):
    os.system("(time go run ../editor/editor.go big) 2>> output_pipeline")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_pipeline")
    for i in range(5):
        instr = "time go run ../editor/editor.go big pipeline " + str(core)
        formatted = "(" + instr + ") 2>> output_pipeline"
        os.system(formatted)

#Script generated in collaboration with Maggie Zhao