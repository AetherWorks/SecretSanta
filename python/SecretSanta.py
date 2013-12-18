'''
Created on 17.12.2013

@author: Isabel R L Peters
'''
import random

# function that accepts a list of names, shuffles it and rotates it once, assigning a secret santa to the person next in the list
def secretSanta(names):
    random.shuffle(names)
    rotate = list(names)
    rotate.append(rotate.pop(0))
    
    printResults(names, rotate)

#print the results and a Christmas tree
def printResults(names, secret):
    print('*-------------------------------------------------------*')
    print('*------------------- Secret Santa ----------------------*')
    print('*-------------------------------------------------------*\n')
    for i in range(10):
        print(' '*(24-i)+'*'*(2*i+1))
    print(' '*23+'***\n')
    
        
    for i in range(len(names)):
        print("%10s == is buying a present for ==> %s"%(names[i], secret[i]))

  
if __name__ == '__main__':
    names = ['Isabel','Angus','Greg', 'Mike', 'Shannon', 'Lewis', 'Ally', 'Rob']
    secretSanta(names)