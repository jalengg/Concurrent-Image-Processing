#!/bin/bash
#
#SBATCH --mail-user=zhiyuanxie@uchicago.edu
#SBATCH --mail-type=ALL
#SBATCH --job-name=proj3-experiment
#SBATCH --output=./slurm/out/%j.%N.stdout
#SBATCH --error=./slurm/out/%j.%N.stderr
#SBATCH --chdir=/home/zhiyuanxie/proj3-shiitavie/benchmark
#SBATCH --partition=debug 
#SBATCH --nodes=1
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=16
#SBATCH --mem-per-cpu=2000
#SBATCH --exclusive
#SBATCH --time=4:00:00

python3 image_filtering_time_experiment.py
python3 output_reader.py 

python3 mapReduce_time_experiment.py
python3 mapReduce_output_reader.py

