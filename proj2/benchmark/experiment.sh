#!/bin/bash
#
#SBATCH --mail-user=zhiyuanxie@uchicago.edu
#SBATCH --mail-type=ALL
#SBATCH --job-name=proj2-experiment
#SBATCH --output=./slurm/out/%j.%N.stdout
#SBATCH --error=./slurm/out/%j.%N.stderr
#SBATCH --chdir=/home/zhiyuanxie/project-2-shiitavie/proj2/benchmark
#SBATCH --partition=debug 
#SBATCH --nodes=1
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=16
#SBATCH --mem-per-cpu=2000
#SBATCH --exclusive
#SBATCH --time=3:00:00

python3 time_experiment.py
python3 output_reader.py 
