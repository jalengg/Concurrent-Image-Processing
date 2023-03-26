#!/bin/bash
#
#SBATCH --mail-user=zhiyuanxie@uchicago.edu
#SBATCH --mail-type=ALL
#SBATCH --job-name=proj1_benchmark 
#SBATCH --output=./slurm/out/%j.%N.stdout
#SBATCH --error=./slurm/out/%j.%N.stderr
#SBATCH --chdir=/home/zhiyuanxie/project-2-shiitavie/proj2/benchmark
#SBATCH --partition=debug 
#SBATCH --nodes=1
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=16
#SBATCH --mem-per-cpu=900
#SBATCH --exclusive
#SBATCH --time=10:00


module load golang/1.19
go run ../editor/editor.go small
go run ../editor/editor.go big pipeline 8
go run ../editor/editor.go mixture bsp 6
