import random
from os import system, name
from sympy import *
from threading import Thread, Event
from time import sleep
from collections import defaultdict

def clear():   
    if name == 'nt': 
        _ = system('cls') 
    else: 
        _ = system('clear')

def generateEquation(variables, values, operators):
    var = list(variables)
    random.shuffle(var)
    equation = []
    equationWithValues = []
    parenthesesCounter = len(variables)
    openParenthesesCounter = 0

    for valueSet in values:
        equationWithValues.append([])

    for i in var:
        if (parenthesesCounter > 0):
            openParens = random.randint(1,parenthesesCounter)
            equation += ["("]*openParens
            for a in equationWithValues:
                a += ["("]*openParens
            openParenthesesCounter += openParens
            parenthesesCounter -= openParens

        equation.append(i)
        for a in range(0,len(equationWithValues)):
            equationWithValues[a].append(values[a][i])

        if (openParenthesesCounter > 0):
            closedParens = random.randint(0,openParenthesesCounter)
            equation += [")"]*closedParens
            for a in equationWithValues:
                a += [")"]*closedParens
            openParenthesesCounter -= closedParens

        if (i != var[-1]):
            op = operators[random.randint(0,len(operators)-1)]
            equation.append(op)
            for a in equationWithValues:
                a.append(op)
        else:
            equation += [")"]*openParenthesesCounter
            for a in equationWithValues:
                a += [")"]*openParenthesesCounter
    
    return (equation, equationWithValues)

def validateEquation(equationWithValues, goals):
    for equationNum in range(0,len(equationWithValues)):
        if eval("".join(equationWithValues[equationNum])) != goals[equationNum]:
            return False
    return True

def findSolutions(variables, values, operators, goals, threadNumber):
    loopCounter = 0

    while loopCounter != 1000000000:
        if event.is_set():
            break

        loopCounter += 1
        iterations[threadNumber] += 1

        equation, equationWithValues = generateEquation(variables, values, operators)

        try:
            if validateEquation(equationWithValues, goals):
                simplified = simplify("".join(equation))
                if (simplified not in potentialSolutions):
                    potentialSolutions.append(simplified)
        except Exception as e:
            print(e)

def printSolutions(iterations, potentialSolutions):
    clear()
    print("{} = {:,}".format(iterations, sum(iterations.values())))
    print(len(potentialSolutions))
    print(potentialSolutions)

if __name__ == "__main__":
    a,b,c,d,e,f,g,h,i = symbols('a b c d e f g h i')
    global iterations
    global potentialSolutions 

    threadCount = 8
    threads = []
    event = Event()

    iterations = defaultdict(int)
    potentialSolutions = []
    variables = ['baseDamage', 'smite', 'vigor', 'agilityConst']
    values = [{"baseDamage":"3", "smite":"5", "vigor":"1.02", "agilityConst":"1"}]
    operators = ["+", "*"]
    goals = [10.076]
    
    for threadNumber in range(1,threadCount+1):
        threads.append(Thread(target=findSolutions, args=(variables, values, operators, goals, threadNumber,)))
    
    for thread in threads:
        thread.start()

    while True:
        try:
            printSolutions(iterations, potentialSolutions)
            sleep(1)
        except KeyboardInterrupt:
            event.set()
            break
    
    for thread in threads:
        thread.join()

    printSolutions(iterations, potentialSolutions)