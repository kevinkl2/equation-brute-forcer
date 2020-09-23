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

def generateEquation(variables, operators):
    var = list(variables)
    random.shuffle(var)
    equation = []
    equationWithValues = []
    parenthesesCounter = len(variables)
    openParenthesesCounter = 0

    for i in var:
        if (parenthesesCounter > 0):
            openParens = random.randint(1,parenthesesCounter)
            equation += ["("]*openParens
            equationWithValues += ["("]*openParens
            openParenthesesCounter += openParens
            parenthesesCounter -= openParens

        equation.append(i)
        equationWithValues.append(variables[i])

        if (openParenthesesCounter > 0):
            closedParens = random.randint(0,openParenthesesCounter)
            equation += [")"]*closedParens
            equationWithValues += [")"]*closedParens
            openParenthesesCounter -= closedParens

        if (i != var[-1]):
            op = operators[random.randint(0,len(operators)-1)]
            equation.append(op)
            equationWithValues.append(op)
        else:
            equation += [")"]*openParenthesesCounter
            equationWithValues += [")"]*openParenthesesCounter
    
    return (equation, equationWithValues)

def validateEquation(equationWithValues, goal):
    return eval("".join(equationWithValues)) == goal

def findSolutions(variables, operators, goal, threadNumber):
    loopCounter = 0

    while loopCounter != 1000000000:
        if event.is_set():
            break

        loopCounter += 1
        iterations[threadNumber] += 1

        equation, equationWithValues = generateEquation(variables, operators)

        try:
            if validateEquation(equationWithValues, goal):
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

    threadCount = 1
    threads = []
    event = Event()

    iterations = defaultdict(int)
    potentialSolutions = []
    #[c*i + e*(a + b + d + g) + f + h, b*c + e*(a + d + f + i) + g + h, a + b + e*(c + d*(f + i)) + g + h, d*e*f + g + h*(a + i*(b + c)), a + b + c*i + e*(d + f + g) + h, a + b + e*(c*i + d*f) + g + h, d*(a + b + c*e + f + g) + h + i, e*(a + b*c + d + i) + f + g + h, e*(c + d*(a + b + i)) + f + g + h]
    variables = {"a":"10", "b":"5", "c":"3", "d":"1.5", "e":"1.75", "f":"15", "g":"20", "h":"1.02", "i":"2"}
    operators = ["+", "*"]
    goal = 85.895
    
    for threadNumber in range(1,threadCount+1):
        threads.append(Thread(target=findSolutions, args=(variables, operators, goal, threadNumber,)))
    
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