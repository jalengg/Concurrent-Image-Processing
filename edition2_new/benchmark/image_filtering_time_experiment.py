import sys
import os

# run experiment 5 times for each type of implementation and data directory for each of [2, 4, 6, 8, 12] cores
# USAGE: python3 time_experiment.py


# steal
os.system("echo small > output_steal")
os.system("echo ================ >> output_steal")
os.system("echo 1 >> output_steal")
for i in range(3):
    os.system("(time go run ../editor/editor.go small) 2>> output_steal")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_steal")
    for i in range(3):
        instr = "time go run ../editor/editor.go small steal " + str(core)
        formatted = "(" + instr + ") 2>> output_steal"
        os.system(formatted)

os.system("echo mixture >> output_steal")
os.system("echo ================ >> output_steal")
os.system("echo 1 >> output_steal")
for i in range(3):
    os.system("(time go run ../editor/editor.go mixture) 2>> output_steal")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_steal")
    for i in range(3):
        instr = "time go run ../editor/editor.go mixture steal " + str(core)
        formatted = "(" + instr + ") 2>> output_steal"
        os.system(formatted)

os.system("echo big >> output_steal")
os.system("echo ================ >> output_steal")
os.system("echo 1 >> output_steal")
for i in range(3):
    os.system("(time go run ../editor/editor.go big) 2>> output_steal")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_steal")
    for i in range(3):
        instr = "time go run ../editor/editor.go big steal " + str(core)
        formatted = "(" + instr + ") 2>> output_steal"
        os.system(formatted)


# Balance
os.system("echo small > output_balance")
os.system("echo ================ >> output_balance")
os.system("echo 1 >> output_balance")
for i in range(3):
    os.system("(time go run ../editor/editor.go small) 2>> output_balance")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_balance")
    for i in range(3):
        instr = "time go run ../editor/editor.go small balance " + str(core)
        formatted = "(" + instr + ") 2>> output_balance"
        os.system(formatted)

os.system("echo mixture >> output_balance")
os.system("echo ================ >> output_balance")
os.system("echo 1 >> output_balance")
for i in range(3):
    os.system("(time go run ../editor/editor.go mixture) 2>> output_balance")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_balance")
    for i in range(3):
        instr = "time go run ../editor/editor.go mixture balance " + str(core)
        formatted = "(" + instr + ") 2>> output_balance"
        os.system(formatted)

os.system("echo big >> output_balance")
os.system("echo ================ >> output_balance")
os.system("echo 1 >> output_balance")
for i in range(3):
    os.system("(time go run ../editor/editor.go big) 2>> output_balance")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> output_balance")
    for i in range(3):
        instr = "time go run ../editor/editor.go big balance " + str(core)
        formatted = "(" + instr + ") 2>> output_balance"
        os.system(formatted)

#Script generated in collaboration with Maggie Zhao