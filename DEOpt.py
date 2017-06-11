#!/usr/bin/env python
# Written by Z-L DENG
# 11/06/2017

import numpy as np

# initialize the population
def init_pop(n_individuals, n_variables, x_lower, x_upper):
    x_all = np.zeros((n_individuals, n_variables))
    for i in range(n_individuals):
        x_all[i] = x_lower + np.random.random() * (x_upper - x_lower)
    return(x_all)

# mutation based on DE 
def mutation(g_generation, i_individual, f):
    g_without_i = np.delete(g_generation, i_individual, 0)
    np.random.shuffle(g_without_i)
    # randomly choose 3 individuals, new = old1 + f*(old2-old3)
    mutated_i = g_without_i[1] + f * (g_without_i[2] - g_without_i[3])
    return(mutated_i)

# check upper and lower bound
def check_bound(x, x_lower, x_upper):
    n = len(x)
    h_i = [x[item] if x[item] < x_upper[item] else x_upper[item]
            for item in range(n)]
    h_i = [h_i[item] if h_i[item] > x_lower[item] else x_lower[item]
            for item in range(n)]
    return(h_i)

# crossover to update the population
def crossover(x, cr, h_i):
    n = len(x)
    v_i = np.array(
        [x[j] if (np.random.random() > cr) else h_i[j] for j in range(n)])

    # evaluate the individuals above to decide keeping it or not
    if evaluate_func(x) > evaluate_func(v_i):
        x_new_generation = v_i
    else:
        x_new_generation = x
    return(x_new_generation)

# DE main interface 
def de(n_individuals=20,
       n_generations=100,
       f=0.5,
       cr=0.3,
       x_lower=np.array([0, 1, 0, 2]),
       x_upper=np.array([5, 6, 8, 4])):
    '''
    n_individuals: population size
    n_generations: number of generations
    f: scalling factor
    cr: crossover rate
    x_lower: lower bounds of variables
    x_upper: upper bonuds

    '''
    # n_variables: number of varibles to be optimized
    n_variables = len(x_lower)
    x_all = init_pop(n_individuals, n_variables, x_lower, x_upper)
    print("initialization of population finished")
    print("the variables to be optimized: ", n_variables,
          "\nthe bounds of variables are: ", x_lower, x_upper)

    for g in range(n_generations - 1):
        for i in range(n_individuals):
            # mutation
            mutated_i = mutation(x_all, i, f)
            print(mutated_i)
            # check bounds for mutations
            h_i = check_bound(mutated_i, x_lower, x_upper)
            # generate new population
            x_all[i] = crossover(x_all[i], cr, h_i)

    evaluate_result = [evaluate_func(
        x_all[i]) for i in range(n_individuals)]
    best_x_i = x_all[np.argmin(evaluate_result)]
    print("The resulting values of the objective function of final generation: ",
          evaluate_result)
    print("The final optimized value of variables: ", best_x_i)

# the object function for optimization
def evaluate_func(x):
    (a, b, c, d) = x
    return(5 * a**3 - 3 * b + 7 * c**3 - 2 * d)

def main():
    de()

if __name__ == '__main__':
    main()
