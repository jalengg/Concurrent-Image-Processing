import sys
import os

# runs experiment 3 times for each type of implementation for 3 different graph sizes and for each of [2, 4, 6, 8, 12] cores
# USAGE: python3 mapReduce_time_experiment.py


# steal
os.system("echo small > mapReduce_output_steal")
os.system("echo ================ >> mapReduce_output_steal")
os.system("echo 1 >> mapReduce_output_steal")
for i in range(3):
    os.system("(time go run ../SSSP/runSSSP.go 1000 steal 1 no) 2>> mapReduce_output_steal")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> mapReduce_output_steal")
    for i in range(3):
        instr = "time go run ../SSSP/runSSSP.go 1000 steal " + str(core) + " no" 
        formatted = "(" + instr + ") 2>> mapReduce_output_steal"
        os.system(formatted)


os.system("echo big >> mapReduce_output_steal")
os.system("echo ================ >> mapReduce_output_steal")
os.system("echo 1 >> mapReduce_output_steal")
for i in range(3):
    os.system("(time go run ../SSSP/runSSSP.go 10000 steal 1 no) 2>> mapReduce_output_steal")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> mapReduce_output_steal")
    for i in range(3):
        instr = "time go run ../SSSP/runSSSP.go 10000 steal " + str(core) + " no" 
        formatted = "(" + instr + ") 2>> mapReduce_output_steal"
        os.system(formatted)



# balance
os.system("echo small > mapReduce_output_balance")
os.system("echo ================ >> mapReduce_output_balance")
os.system("echo 1 >> mapReduce_output_balance")
for i in range(3):
    os.system("(time go run ../SSSP/runSSSP.go 1000 balance 1 no) 2>> mapReduce_output_balance")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> mapReduce_output_balance")
    for i in range(3):
        instr = "time go run ../SSSP/runSSSP.go 1000 balance " + str(core) + " no" 
        formatted = "(" + instr + ") 2>> mapReduce_output_balance"
        os.system(formatted)


os.system("echo big >> mapReduce_output_balance")
os.system("echo ================ >> mapReduce_output_balance")
os.system("echo 1 >> mapReduce_output_balance")
for i in range(3):
    os.system("(time go run ../SSSP/runSSSP.go 10000 balance 1 no) 2>> mapReduce_output_balance")
for core in range(2, 13, 2):
    os.system("echo " + str(core) + " >> mapReduce_output_balance")
    for i in range(3):
        instr = "time go run ../SSSP/runSSSP.go 10000 balance " + str(core) + " no" 
        formatted = "(" + instr + ") 2>> mapReduce_output_balance"
        os.system(formatted)